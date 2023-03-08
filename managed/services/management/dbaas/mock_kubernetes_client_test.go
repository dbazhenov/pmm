// Code generated by mockery v1.0.0. DO NOT EDIT.

package dbaas

import (
	context "context"

	v1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	v1 "github.com/percona/dbaas-operator/api/v1"
	mock "github.com/stretchr/testify/mock"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"

	dbaasv1beta1 "github.com/percona/pmm/api/managementpb/dbaas"
	kubernetes "github.com/percona/pmm/managed/services/dbaas/kubernetes"
)

// mockKubernetesClient is an autogenerated mock type for the kubernetesClient type
type mockKubernetesClient struct {
	mock.Mock
}

// CreateDatabaseCluster provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) CreateDatabaseCluster(_a0 *v1.DatabaseCluster) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.DatabaseCluster) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePMMSecret provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) CreatePMMSecret(_a0 string, _a1 map[string][]byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string][]byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRestore provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) CreateRestore(_a0 *v1.DatabaseClusterRestore) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.DatabaseClusterRestore) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDatabaseCluster provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) DeleteDatabaseCluster(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllClusterResources provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockKubernetesClient) GetAllClusterResources(_a0 context.Context, _a1 kubernetes.ClusterType, _a2 *corev1.PersistentVolumeList) (uint64, uint64, uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, kubernetes.ClusterType, *corev1.PersistentVolumeList) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, kubernetes.ClusterType, *corev1.PersistentVolumeList) uint64); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 uint64
	if rf, ok := ret.Get(2).(func(context.Context, kubernetes.ClusterType, *corev1.PersistentVolumeList) uint64); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, kubernetes.ClusterType, *corev1.PersistentVolumeList) error); ok {
		r3 = rf(_a0, _a1, _a2)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetClusterType provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) GetClusterType(_a0 context.Context) (kubernetes.ClusterType, error) {
	ret := _m.Called(_a0)

	var r0 kubernetes.ClusterType
	if rf, ok := ret.Get(0).(func(context.Context) kubernetes.ClusterType); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(kubernetes.ClusterType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsumedCPUAndMemory provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) GetConsumedCPUAndMemory(_a0 context.Context, _a1 string) (uint64, uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, string) uint64); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetConsumedDiskBytes provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockKubernetesClient) GetConsumedDiskBytes(_a0 context.Context, _a1 kubernetes.ClusterType, _a2 *corev1.PersistentVolumeList) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, kubernetes.ClusterType, *corev1.PersistentVolumeList) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kubernetes.ClusterType, *corev1.PersistentVolumeList) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabaseCluster provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) GetDatabaseCluster(_a0 context.Context, _a1 string) (*v1.DatabaseCluster, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.DatabaseCluster
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1.DatabaseCluster); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.DatabaseCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultStorageClassName provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) GetDefaultStorageClassName(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPSMDBOperatorVersion provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) GetPSMDBOperatorVersion(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPXCOperatorVersion provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) GetPXCOperatorVersion(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPersistentVolumes provides a mock function with given fields: ctx
func (_m *mockKubernetesClient) GetPersistentVolumes(ctx context.Context) (*corev1.PersistentVolumeList, error) {
	ret := _m.Called(ctx)

	var r0 *corev1.PersistentVolumeList
	if rf, ok := ret.Get(0).(func(context.Context) *corev1.PersistentVolumeList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecret provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) GetSecret(_a0 context.Context, _a1 string) (*corev1.Secret, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *corev1.Secret
	if rf, ok := ret.Get(0).(func(context.Context, string) *corev1.Secret); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageClasses provides a mock function with given fields: ctx
func (_m *mockKubernetesClient) GetStorageClasses(ctx context.Context) (*storagev1.StorageClassList, error) {
	ret := _m.Called(ctx)

	var r0 *storagev1.StorageClassList
	if rf, ok := ret.Get(0).(func(context.Context) *storagev1.StorageClassList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.StorageClassList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallOLMOperator provides a mock function with given fields: ctx
func (_m *mockKubernetesClient) InstallOLMOperator(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstallOperator provides a mock function with given fields: ctx, req
func (_m *mockKubernetesClient) InstallOperator(ctx context.Context, req kubernetes.InstallOperatorRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, kubernetes.InstallOperatorRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListDatabaseClusters provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) ListDatabaseClusters(_a0 context.Context) (*v1.DatabaseClusterList, error) {
	ret := _m.Called(_a0)

	var r0 *v1.DatabaseClusterList
	if rf, ok := ret.Get(0).(func(context.Context) *v1.DatabaseClusterList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.DatabaseClusterList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecrets provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) ListSecrets(_a0 context.Context) (*corev1.SecretList, error) {
	ret := _m.Called(_a0)

	var r0 *corev1.SecretList
	if rf, ok := ret.Get(0).(func(context.Context) *corev1.SecretList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.SecretList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSubscriptions provides a mock function with given fields: ctx, namespace
func (_m *mockKubernetesClient) ListSubscriptions(ctx context.Context, namespace string) (*v1alpha1.SubscriptionList, error) {
	ret := _m.Called(ctx, namespace)

	var r0 *v1alpha1.SubscriptionList
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1alpha1.SubscriptionList); ok {
		r0 = rf(ctx, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.SubscriptionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTemplates provides a mock function with given fields: ctx, engine, namespace
func (_m *mockKubernetesClient) ListTemplates(ctx context.Context, engine string, namespace string) ([]*dbaasv1beta1.Template, error) {
	ret := _m.Called(ctx, engine, namespace)

	var r0 []*dbaasv1beta1.Template
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*dbaasv1beta1.Template); ok {
		r0 = rf(ctx, engine, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbaasv1beta1.Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, engine, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchDatabaseCluster provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) PatchDatabaseCluster(_a0 *v1.DatabaseCluster) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.DatabaseCluster) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RestartDatabaseCluster provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) RestartDatabaseCluster(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetKubeconfig provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) SetKubeconfig(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpgradeOperator provides a mock function with given fields: ctx, namespace, name
func (_m *mockKubernetesClient) UpgradeOperator(ctx context.Context, namespace string, name string) error {
	ret := _m.Called(ctx, namespace, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, namespace, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
