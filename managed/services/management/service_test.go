// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package management

import (
	"context"
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm/api/inventorypb"
	"github.com/percona/pmm/api/managementpb"
	servicev1beta1 "github.com/percona/pmm/api/managementpb/service"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/utils/logger"
	"github.com/percona/pmm/managed/utils/testdb"
	"github.com/percona/pmm/managed/utils/tests"
)

func TestServiceService(t *testing.T) {
	t.Run("Remove", func(t *testing.T) {
		setup := func(t *testing.T) (context.Context, *ServiceService, func(t *testing.T), *mockPrometheusService) { //nolint:unparam
			t.Helper()

			ctx := logger.Set(context.Background(), t.Name())
			uuid.SetRand(&tests.IDReader{})

			sqlDB := testdb.Open(t, models.SetupFixtures, nil)
			db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))

			vmdb := &mockPrometheusService{}
			vmdb.Test(t)

			state := &mockAgentsStateUpdater{}
			state.Test(t)

			ar := &mockAgentsRegistry{}
			ar.Test(t)

			teardown := func(t *testing.T) {
				uuid.SetRand(nil)

				require.NoError(t, sqlDB.Close())
				vmdb.AssertExpectations(t)
				state.AssertExpectations(t)
				ar.AssertExpectations(t)
			}
			s := NewServiceService(db, ar, state, vmdb)

			return ctx, s, teardown, vmdb
		}
		t.Run("No params", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			response, err := s.RemoveService(ctx, &managementpb.RemoveServiceRequest{})
			assert.Nil(t, response)
			tests.AssertGRPCError(t, status.New(codes.InvalidArgument, `service_id or service_name expected`), err)
		})

		t.Run("Both params", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			response, err := s.RemoveService(ctx, &managementpb.RemoveServiceRequest{ServiceId: "some-id", ServiceName: "some-service-name"})
			assert.Nil(t, response)
			tests.AssertGRPCError(t, status.New(codes.InvalidArgument, `service_id or service_name expected; not both`), err)
		})

		t.Run("Not found", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			response, err := s.RemoveService(ctx, &managementpb.RemoveServiceRequest{ServiceName: "some-service-name"})
			assert.Nil(t, response)
			tests.AssertGRPCError(t, status.New(codes.NotFound, `Service with name "some-service-name" not found.`), err)
		})

		t.Run("Wrong service type", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			service, err := models.AddNewService(s.db.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
				ServiceName: "test-mysql",
				NodeID:      models.PMMServerNodeID,
				Address:     pointer.ToString("127.0.0.1"),
				Port:        pointer.ToUint16(3306),
			})
			require.NoError(t, err)

			response, err := s.RemoveService(ctx, &managementpb.RemoveServiceRequest{ServiceId: service.ServiceID, ServiceType: inventorypb.ServiceType_POSTGRESQL_SERVICE})
			assert.Nil(t, response)
			tests.AssertGRPCError(t, status.New(codes.InvalidArgument, `wrong service type`), err)
		})

		t.Run("Basic", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			service, err := models.AddNewService(s.db.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
				ServiceName: "test-mysql",
				NodeID:      models.PMMServerNodeID,
				Address:     pointer.ToString("127.0.0.1"),
				Port:        pointer.ToUint16(3306),
			})
			require.NoError(t, err)

			pmmAgent, err := models.CreatePMMAgent(s.db.Querier, models.PMMServerNodeID, nil)
			require.NoError(t, err)

			mysqldExporter, err := models.CreateAgent(s.db.Querier, models.MySQLdExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
				Password:   "password",
				Username:   "username",
				// TODO TLS
			})
			require.NoError(t, err)

			s.state.(*mockAgentsStateUpdater).On("RequestStateUpdate", ctx, pmmAgent.AgentID)
			response, err := s.RemoveService(ctx, &managementpb.RemoveServiceRequest{ServiceName: service.ServiceName, ServiceType: inventorypb.ServiceType_MYSQL_SERVICE})
			assert.NotNil(t, response)
			assert.NoError(t, err)

			agent, err := models.FindAgentByID(s.db.Querier, mysqldExporter.AgentID)
			assert.Nil(t, agent)
			tests.AssertGRPCError(t, status.New(codes.NotFound, `Agent with ID "/agent_id/00000000-0000-4000-8000-000000000007" not found.`), err)

			service, err = models.FindServiceByID(s.db.Querier, service.ServiceID)
			assert.Nil(t, service)
			tests.AssertGRPCError(t, status.New(codes.NotFound, `Service with ID "/service_id/00000000-0000-4000-8000-000000000005" not found.`), err)
		})

		t.Run("RDS", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			node, err := models.CreateNode(s.db.Querier, models.RemoteRDSNodeType, &models.CreateNodeParams{
				NodeName: "test",
				Address:  "test-address",
				Region:   pointer.ToString("test-region"),
			})
			require.NoError(t, err)

			service, err := models.AddNewService(s.db.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
				ServiceName: "test-mysql",
				NodeID:      node.NodeID,
				Address:     pointer.ToString("127.0.0.1"),
				Port:        pointer.ToUint16(3306),
			})
			require.NoError(t, err)

			pmmAgent, err := models.CreatePMMAgent(s.db.Querier, models.PMMServerNodeID, nil)
			require.NoError(t, err)

			mysqldExporter, err := models.CreateAgent(s.db.Querier, models.MySQLdExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
				Password:   "password",
				Username:   "username",
				// TODO TLS
			})
			require.NoError(t, err)

			rdsExporter, err := models.CreateAgent(s.db.Querier, models.RDSExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				NodeID:     node.NodeID,
			})
			require.NoError(t, err)

			s.state.(*mockAgentsStateUpdater).On("RequestStateUpdate", ctx, pmmAgent.AgentID)
			_, err = s.RemoveService(ctx, &managementpb.RemoveServiceRequest{ServiceName: service.ServiceName, ServiceType: inventorypb.ServiceType_MYSQL_SERVICE})
			assert.NoError(t, err)

			_, err = models.FindServiceByID(s.db.Querier, service.ServiceID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Service with ID "%s" not found.`, service.ServiceID)), err)

			_, err = models.FindAgentByID(s.db.Querier, mysqldExporter.AgentID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Agent with ID "%s" not found.`, mysqldExporter.AgentID)), err)

			_, err = models.FindAgentByID(s.db.Querier, rdsExporter.AgentID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Agent with ID "%s" not found.`, rdsExporter.AgentID)), err)

			_, err = models.FindNodeByID(s.db.Querier, node.NodeID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Node with ID "%s" not found.`, node.NodeID)), err)
		})

		t.Run("Azure", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			node, err := models.CreateNode(s.db.Querier, models.RemoteAzureDatabaseNodeType, &models.CreateNodeParams{
				NodeName: "test",
				Address:  "test-address",
				Region:   pointer.ToString("test-region"),
			})
			require.NoError(t, err)

			service, err := models.AddNewService(s.db.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
				ServiceName: "test-mysql",
				NodeID:      node.NodeID,
				Address:     pointer.ToString("127.0.0.1"),
				Port:        pointer.ToUint16(3306),
			})
			require.NoError(t, err)

			pmmAgent, err := models.CreatePMMAgent(s.db.Querier, models.PMMServerNodeID, nil)
			require.NoError(t, err)

			mysqldExporter, err := models.CreateAgent(s.db.Querier, models.MySQLdExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
				Password:   "password",
				Username:   "username",
				// TODO TLS
			})
			require.NoError(t, err)

			azureExporter, err := models.CreateAgent(s.db.Querier, models.AzureDatabaseExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				NodeID:     node.NodeID,
			})
			require.NoError(t, err)

			s.state.(*mockAgentsStateUpdater).On("RequestStateUpdate", ctx, pmmAgent.AgentID)
			_, err = s.RemoveService(ctx, &managementpb.RemoveServiceRequest{ServiceName: service.ServiceName, ServiceType: inventorypb.ServiceType_MYSQL_SERVICE})
			assert.NoError(t, err)

			_, err = models.FindServiceByID(s.db.Querier, service.ServiceID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Service with ID "%s" not found.`, service.ServiceID)), err)

			_, err = models.FindAgentByID(s.db.Querier, mysqldExporter.AgentID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Agent with ID "%s" not found.`, mysqldExporter.AgentID)), err)

			_, err = models.FindAgentByID(s.db.Querier, azureExporter.AgentID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Agent with ID "%s" not found.`, azureExporter.AgentID)), err)

			_, err = models.FindNodeByID(s.db.Querier, node.NodeID)
			tests.AssertGRPCError(t, status.New(codes.NotFound, fmt.Sprintf(`Node with ID "%s" not found.`, node.NodeID)), err)
		})
	})

	t.Run("List", func(t *testing.T) {
		setup := func(t *testing.T) (context.Context, *MgmtServiceService, func(t *testing.T), *mockPrometheusService) { //nolint:unparam
			t.Helper()

			ctx := logger.Set(context.Background(), t.Name())
			uuid.SetRand(&tests.IDReader{})

			sqlDB := testdb.Open(t, models.SetupFixtures, nil)
			db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))

			vmdb := &mockPrometheusService{}
			vmdb.Test(t)

			state := &mockAgentsStateUpdater{}
			state.Test(t)

			ar := &mockAgentsRegistry{}
			ar.Test(t)

			vmClient := &mockVictoriaMetricsClient{}

			teardown := func(t *testing.T) {
				uuid.SetRand(nil)

				require.NoError(t, sqlDB.Close())
				vmdb.AssertExpectations(t)
				state.AssertExpectations(t)
				ar.AssertExpectations(t)
			}
			s := NewMgmtServiceService(db, ar, state, vmdb, vmClient)

			return ctx, s, teardown, vmdb
		}

		const (
			pgExporterID      = "/agent_id/00000000-0000-4000-8000-000000000003"
			pgStatStatementID = "/agent_id/00000000-0000-4000-8000-000000000004"
			PMMAgentID        = "/agent_id/00000000-0000-4000-8000-000000000007"
		)

		t.Run("Basic", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			s.vmClient.(*mockVictoriaMetricsClient).On("Query", ctx, mock.Anything, mock.Anything).Return(model.Vector{}, nil, nil).Times(3)
			s.r.(*mockAgentsRegistry).On("IsConnected", models.PMMServerAgentID).Return(true).Once() // PMM Server Agent
			s.r.(*mockAgentsRegistry).On("IsConnected", pgExporterID).Return(false).Once()           // PMM Server PostgreSQL exporter
			s.r.(*mockAgentsRegistry).On("IsConnected", pgStatStatementID).Return(false).Once()      // PMM Server PG Stat Statements agent
			response, err := s.ListServices(ctx, &servicev1beta1.ListServiceRequest{})

			require.NoError(t, err)
			assert.Len(t, response.Services, 1) // PMM Server PostgreSQL service
			assert.Len(t, response.Services[0].Agents, 3)
		})

		t.Run("RDS", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			node, err := models.CreateNode(s.db.Querier, models.RemoteRDSNodeType, &models.CreateNodeParams{
				NodeName: "test",
				Address:  "test-address",
				Region:   pointer.ToString("test-region"),
			})
			require.NoError(t, err)

			service, err := models.AddNewService(s.db.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
				ServiceName: "test-mysql",
				NodeID:      node.NodeID,
				Address:     pointer.ToString("127.0.0.1"),
				Port:        pointer.ToUint16(3306),
			})
			require.NoError(t, err)

			pmmAgent, err := models.CreatePMMAgent(s.db.Querier, models.PMMServerNodeID, nil)
			require.NoError(t, err)

			mysqldExporter, err := models.CreateAgent(s.db.Querier, models.MySQLdExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
				Password:   "password",
				Username:   "username",
			})
			require.NoError(t, err)

			rdsExporter, err := models.CreateAgent(s.db.Querier, models.RDSExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
			})
			require.NoError(t, err)

			s.vmClient.(*mockVictoriaMetricsClient).On("Query", ctx, mock.Anything, mock.Anything).Return(model.Vector{}, nil, nil).Times(7)
			s.r.(*mockAgentsRegistry).On("IsConnected", models.PMMServerAgentID).Return(true).Once() // PMM Server Agent
			s.r.(*mockAgentsRegistry).On("IsConnected", pmmAgent.AgentID).Return(true).Once()        // PMM Agent
			s.r.(*mockAgentsRegistry).On("IsConnected", pgExporterID).Return(false).Once()           // PMM Server PostgreSQL exporter
			s.r.(*mockAgentsRegistry).On("IsConnected", pgStatStatementID).Return(false).Once()      // PMM Server PG Stat Statements agent
			s.r.(*mockAgentsRegistry).On("IsConnected", PMMAgentID).Return(false)                    // PMM Agent 2
			s.r.(*mockAgentsRegistry).On("IsConnected", mysqldExporter.AgentID).Return(false).Once() // MySQLd exporter
			s.r.(*mockAgentsRegistry).On("IsConnected", rdsExporter.AgentID).Return(false).Once()    // RDS exporter

			response, err := s.ListServices(ctx, &servicev1beta1.ListServiceRequest{})

			require.NoError(t, err)
			assert.Len(t, response.Services, 2) // PMM Server PostgreSQL service, MySQL service
			assert.Len(t, response.Services[0].Agents, 4)
			assert.Len(t, response.Services[1].Agents, 2)
		})

		t.Run("Azure", func(t *testing.T) {
			ctx, s, teardown, _ := setup(t)
			defer teardown(t)

			node, err := models.CreateNode(s.db.Querier, models.RemoteAzureDatabaseNodeType, &models.CreateNodeParams{
				NodeName: "test",
				Address:  "test-address",
				Region:   pointer.ToString("test-region"),
			})
			require.NoError(t, err)

			service, err := models.AddNewService(s.db.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
				ServiceName: "test-mysql",
				NodeID:      node.NodeID,
				Address:     pointer.ToString("127.0.0.1"),
				Port:        pointer.ToUint16(3306),
			})
			require.NoError(t, err)

			pmmAgent, err := models.CreatePMMAgent(s.db.Querier, models.PMMServerNodeID, nil)
			require.NoError(t, err)

			mysqldExporter, err := models.CreateAgent(s.db.Querier, models.MySQLdExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
				Password:   "password",
				Username:   "username",
			})
			require.NoError(t, err)

			azureExporter, err := models.CreateAgent(s.db.Querier, models.AzureDatabaseExporterType, &models.CreateAgentParams{
				PMMAgentID: pmmAgent.AgentID,
				ServiceID:  service.ServiceID,
			})
			require.NoError(t, err)

			s.vmClient.(*mockVictoriaMetricsClient).On("Query", ctx, mock.Anything, mock.Anything).Return(model.Vector{}, nil, nil).Times(7)
			s.r.(*mockAgentsRegistry).On("IsConnected", models.PMMServerAgentID).Return(true).Once() // PMM Server Agent
			s.r.(*mockAgentsRegistry).On("IsConnected", pmmAgent.AgentID).Return(true).Once()        // PMM Agent
			s.r.(*mockAgentsRegistry).On("IsConnected", pgExporterID).Return(false).Once()           // PMM Server PostgreSQL exporter
			s.r.(*mockAgentsRegistry).On("IsConnected", pgStatStatementID).Return(false).Once()      // PMM Server PG Stat Statements agent
			s.r.(*mockAgentsRegistry).On("IsConnected", PMMAgentID).Return(false)                    // PMM Agent 2
			s.r.(*mockAgentsRegistry).On("IsConnected", mysqldExporter.AgentID).Return(false).Once() // MySQLd exporter
			s.r.(*mockAgentsRegistry).On("IsConnected", azureExporter.AgentID).Return(false).Once()  // Azure exporter

			response, err := s.ListServices(ctx, &servicev1beta1.ListServiceRequest{})

			require.NoError(t, err)
			assert.Len(t, response.Services, 2) // PMM Server PostgreSQL service, MySQL service
			assert.Len(t, response.Services[0].Agents, 4)
			assert.Len(t, response.Services[1].Agents, 2)
		})
	})
}
