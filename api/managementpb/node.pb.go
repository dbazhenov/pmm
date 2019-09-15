// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/node.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterNodeRequest struct {
	// Node type to be registered.
	NodeType inventorypb.NodeType `protobuf:"varint,1,opt,name=node_type,json=nodeType,proto3,enum=inventory.NodeType" json:"node_type,omitempty"`
	// Unique across all Nodes user-defined name.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Node address (DNS name or IP).
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Linux machine-id.
	MachineId string `protobuf:"bytes,4,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	// Linux distribution name and version.
	Distro string `protobuf:"bytes,5,opt,name=distro,proto3" json:"distro,omitempty"`
	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerId string `protobuf:"bytes,6,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// Container name.
	ContainerName string `protobuf:"bytes,7,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	// Node model.
	NodeModel string `protobuf:"bytes,8,opt,name=node_model,json=nodeModel,proto3" json:"node_model,omitempty"`
	// Node region.
	Region string `protobuf:"bytes,9,opt,name=region,proto3" json:"region,omitempty"`
	// Node availability zone.
	Az string `protobuf:"bytes,10,opt,name=az,proto3" json:"az,omitempty"`
	// Custom user-assigned labels.
	CustomLabels map[string]string `protobuf:"bytes,11,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If true, and Node with that name already exist, it will be removed with all dependent Services and Agents.
	Reregister           bool     `protobuf:"varint,12,opt,name=reregister,proto3" json:"reregister,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16aa3e5e785d25d0, []int{0}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetNodeType() inventorypb.NodeType {
	if m != nil {
		return m.NodeType
	}
	return inventorypb.NodeType_NODE_TYPE_INVALID
}

func (m *RegisterNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RegisterNodeRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegisterNodeRequest) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

func (m *RegisterNodeRequest) GetDistro() string {
	if m != nil {
		return m.Distro
	}
	return ""
}

func (m *RegisterNodeRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *RegisterNodeRequest) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *RegisterNodeRequest) GetNodeModel() string {
	if m != nil {
		return m.NodeModel
	}
	return ""
}

func (m *RegisterNodeRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *RegisterNodeRequest) GetAz() string {
	if m != nil {
		return m.Az
	}
	return ""
}

func (m *RegisterNodeRequest) GetCustomLabels() map[string]string {
	if m != nil {
		return m.CustomLabels
	}
	return nil
}

func (m *RegisterNodeRequest) GetReregister() bool {
	if m != nil {
		return m.Reregister
	}
	return false
}

type RegisterNodeResponse struct {
	GenericNode          *inventorypb.GenericNode   `protobuf:"bytes,1,opt,name=generic_node,json=genericNode,proto3" json:"generic_node,omitempty"`
	ContainerNode        *inventorypb.ContainerNode `protobuf:"bytes,2,opt,name=container_node,json=containerNode,proto3" json:"container_node,omitempty"`
	PmmAgent             *inventorypb.PMMAgent      `protobuf:"bytes,3,opt,name=pmm_agent,json=pmmAgent,proto3" json:"pmm_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RegisterNodeResponse) Reset()         { *m = RegisterNodeResponse{} }
func (m *RegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeResponse) ProtoMessage()    {}
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16aa3e5e785d25d0, []int{1}
}

func (m *RegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeResponse.Unmarshal(m, b)
}
func (m *RegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeResponse.Merge(m, src)
}
func (m *RegisterNodeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeResponse.Size(m)
}
func (m *RegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeResponse proto.InternalMessageInfo

func (m *RegisterNodeResponse) GetGenericNode() *inventorypb.GenericNode {
	if m != nil {
		return m.GenericNode
	}
	return nil
}

func (m *RegisterNodeResponse) GetContainerNode() *inventorypb.ContainerNode {
	if m != nil {
		return m.ContainerNode
	}
	return nil
}

func (m *RegisterNodeResponse) GetPmmAgent() *inventorypb.PMMAgent {
	if m != nil {
		return m.PmmAgent
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterNodeRequest)(nil), "management.RegisterNodeRequest")
	proto.RegisterMapType((map[string]string)(nil), "management.RegisterNodeRequest.CustomLabelsEntry")
	proto.RegisterType((*RegisterNodeResponse)(nil), "management.RegisterNodeResponse")
}

func init() { proto.RegisterFile("managementpb/node.proto", fileDescriptor_16aa3e5e785d25d0) }

var fileDescriptor_16aa3e5e785d25d0 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdf, 0x4a, 0x1b, 0x4f,
	0x14, 0x66, 0xa3, 0xc6, 0xe4, 0x24, 0xca, 0xef, 0x37, 0x8a, 0x2e, 0x41, 0x6b, 0x8c, 0x14, 0xd3,
	0x3f, 0xd9, 0xd5, 0x14, 0x4a, 0xeb, 0x8d, 0x44, 0x91, 0x22, 0xa8, 0x94, 0xa5, 0x94, 0xd2, 0x9b,
	0x30, 0xd9, 0x3d, 0xae, 0x83, 0xbb, 0x33, 0xdb, 0xd9, 0x49, 0x24, 0x5e, 0x7a, 0xd3, 0xfb, 0xf6,
	0x95, 0xfa, 0x06, 0x7d, 0x80, 0x42, 0xe9, 0x83, 0x94, 0x99, 0xdd, 0x24, 0x2b, 0xb6, 0xf4, 0x2a,
	0x73, 0xce, 0xf7, 0x7d, 0x67, 0xbe, 0x9c, 0x73, 0x66, 0x61, 0x3d, 0xa6, 0x9c, 0x86, 0x18, 0x23,
	0x57, 0xc9, 0xc0, 0xe5, 0x22, 0x40, 0x27, 0x91, 0x42, 0x09, 0x02, 0x33, 0xa0, 0xf1, 0x32, 0x64,
	0xea, 0x6a, 0x38, 0x70, 0x7c, 0x11, 0xbb, 0xf1, 0x0d, 0x53, 0xd7, 0xe2, 0xc6, 0x0d, 0x45, 0xc7,
	0x10, 0x3b, 0x23, 0x1a, 0xb1, 0x80, 0x2a, 0x21, 0x53, 0x77, 0x7a, 0xcc, 0x6a, 0x34, 0x36, 0x42,
	0x21, 0xc2, 0x08, 0x5d, 0x9a, 0x30, 0x97, 0x72, 0x2e, 0x14, 0x55, 0x4c, 0xf0, 0x34, 0x47, 0x6d,
	0xc6, 0x47, 0xc8, 0x95, 0x90, 0xe3, 0x64, 0xe0, 0xd2, 0x10, 0xb9, 0x9a, 0x20, 0xeb, 0x45, 0x44,
	0x7b, 0x9a, 0x00, 0xcf, 0xcd, 0x8f, 0xdf, 0x09, 0x91, 0x77, 0xd2, 0x1b, 0x1a, 0x86, 0x28, 0x5d,
	0x91, 0x98, 0xa2, 0x0f, 0x2f, 0x68, 0x7d, 0x9e, 0x87, 0x15, 0x0f, 0x43, 0x96, 0x2a, 0x94, 0x17,
	0x22, 0x40, 0x0f, 0x3f, 0x0d, 0x31, 0x55, 0x64, 0x0f, 0xaa, 0xba, 0x68, 0x5f, 0x8d, 0x13, 0xb4,
	0xad, 0xa6, 0xd5, 0x5e, 0xee, 0xae, 0x38, 0xd3, 0x2b, 0x1d, 0x4d, 0x7d, 0x37, 0x4e, 0xd0, 0xab,
	0xf0, 0xfc, 0x44, 0x76, 0x72, 0x05, 0xa7, 0x31, 0xda, 0xa5, 0xa6, 0xd5, 0xae, 0x1e, 0x95, 0x7f,
	0xfe, 0xd8, 0x2a, 0x7d, 0xb0, 0x32, 0xd2, 0x05, 0x8d, 0x91, 0xd8, 0xb0, 0x48, 0x83, 0x40, 0x62,
	0x9a, 0xda, 0x73, 0x9a, 0xe2, 0x4d, 0x42, 0xb2, 0x09, 0x10, 0x53, 0xff, 0x8a, 0x71, 0xec, 0xb3,
	0xc0, 0x9e, 0x37, 0x60, 0x35, 0xcf, 0x9c, 0x06, 0x64, 0x0d, 0xca, 0x01, 0x4b, 0x95, 0x14, 0xf6,
	0x82, 0x81, 0xf2, 0x88, 0x6c, 0x43, 0xdd, 0x17, 0x5c, 0x51, 0xc6, 0x51, 0x6a, 0x61, 0xd9, 0xa0,
	0xb5, 0x69, 0xee, 0x34, 0x20, 0x8f, 0x61, 0x79, 0x46, 0x31, 0xee, 0x16, 0x0d, 0x69, 0x69, 0x9a,
	0x35, 0xd6, 0x36, 0x01, 0x8c, 0xff, 0x58, 0x04, 0x18, 0xd9, 0x95, 0xcc, 0x80, 0xce, 0x9c, 0xeb,
	0x84, 0x36, 0x20, 0x31, 0x64, 0x82, 0xdb, 0xd5, 0xcc, 0x40, 0x16, 0x91, 0x65, 0x28, 0xd1, 0x5b,
	0x1b, 0x4c, 0xae, 0x44, 0x6f, 0xc9, 0x7b, 0x58, 0xf2, 0x87, 0xa9, 0x12, 0x71, 0x3f, 0xa2, 0x03,
	0x8c, 0x52, 0xbb, 0xd6, 0x9c, 0x6b, 0xd7, 0xba, 0xfb, 0xce, 0x6c, 0x57, 0x9c, 0x3f, 0x34, 0xdc,
	0x39, 0x36, 0xa2, 0x33, 0xa3, 0x39, 0xe1, 0x4a, 0x8e, 0xbd, 0xba, 0x5f, 0x48, 0x91, 0x47, 0x00,
	0x12, 0x65, 0x2e, 0xb4, 0xeb, 0x4d, 0xab, 0x5d, 0xf1, 0x0a, 0x99, 0xc6, 0x21, 0xfc, 0xff, 0xa0,
	0x04, 0xf9, 0x0f, 0xe6, 0xae, 0x71, 0x6c, 0xe6, 0x57, 0xf5, 0xf4, 0x91, 0xac, 0xc2, 0xc2, 0x88,
	0x46, 0xc3, 0x7c, 0x42, 0x5e, 0x16, 0x1c, 0x94, 0x5e, 0x59, 0xad, 0x6f, 0x16, 0xac, 0xde, 0x37,
	0x96, 0x26, 0x82, 0xa7, 0x48, 0x5e, 0x43, 0x3d, 0x44, 0x8e, 0x92, 0xf9, 0x7d, 0xdd, 0x0e, 0x53,
	0xad, 0xd6, 0x5d, 0x2b, 0x6c, 0xc3, 0x9b, 0x0c, 0x36, 0xaa, 0x5a, 0x38, 0x0b, 0xc8, 0xe1, 0xbd,
	0xd6, 0x6b, 0x71, 0xc9, 0x88, 0xed, 0x82, 0xf8, 0x78, 0x3a, 0x05, 0x2d, 0x2f, 0x0c, 0x45, 0x17,
	0xd8, 0x83, 0x6a, 0x12, 0xc7, 0x7d, 0xb3, 0xf9, 0x66, 0x63, 0x6a, 0xf7, 0xd6, 0xf0, 0xed, 0xf9,
	0x79, 0x4f, 0x43, 0x5e, 0x25, 0x89, 0x63, 0x73, 0xea, 0xde, 0x59, 0x30, 0x6f, 0xa4, 0xb7, 0x50,
	0x2f, 0xfe, 0x1d, 0xb2, 0xf5, 0x8f, 0x09, 0x34, 0x9a, 0x7f, 0x27, 0x64, 0x9d, 0x68, 0xed, 0xde,
	0x7d, 0xff, 0xf5, 0xb5, 0xb4, 0xdd, 0xda, 0x70, 0x47, 0xfb, 0xee, 0x8c, 0xec, 0x6a, 0x92, 0x3b,
	0x51, 0x1c, 0x58, 0x4f, 0x8f, 0xfa, 0x5f, 0x7a, 0x17, 0xde, 0x19, 0x2c, 0x06, 0x78, 0x49, 0x87,
	0x91, 0x22, 0x3d, 0x20, 0x3d, 0xde, 0x44, 0x29, 0x85, 0x6c, 0xca, 0xbc, 0x9a, 0x43, 0x9e, 0xc1,
	0x93, 0xc6, 0xee, 0x8e, 0x1b, 0xe0, 0x25, 0xe3, 0x2c, 0x7b, 0xa0, 0xc5, 0x2f, 0xcd, 0x89, 0xa6,
	0x4f, 0xee, 0xfe, 0x58, 0x2f, 0x42, 0x83, 0xb2, 0x79, 0xbd, 0x2f, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0xea, 0xd7, 0xe3, 0x9b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	// RegisterNode registers a new Node and pmm-agent.
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/management.Node/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	// RegisterNode registers a new Node and pmm-agent.
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) RegisterNode(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Node/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Node_RegisterNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/node.proto",
}
