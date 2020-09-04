// Code generated by protoc-gen-go. DO NOT EDIT.
// source: container.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type ContainerType int32

const (
	ContainerType_UNKNOWN ContainerType = 0
	ContainerType_SAN     ContainerType = 1
	ContainerType_CSV     ContainerType = 2
	ContainerType_SMB     ContainerType = 3
	ContainerType_DAS     ContainerType = 4
)

var ContainerType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SAN",
	2: "CSV",
	3: "SMB",
	4: "DAS",
}

var ContainerType_value = map[string]int32{
	"UNKNOWN": 0,
	"SAN":     1,
	"CSV":     2,
	"SMB":     3,
	"DAS":     4,
}

func (x ContainerType) String() string {
	return proto.EnumName(ContainerType_name, int32(x))
}

func (ContainerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{0}
}

type ContainerRequest struct {
	Containers           []*Container     `protobuf:"bytes,1,rep,name=Containers,proto3" json:"Containers,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ContainerRequest) Reset()         { *m = ContainerRequest{} }
func (m *ContainerRequest) String() string { return proto.CompactTextString(m) }
func (*ContainerRequest) ProtoMessage()    {}
func (*ContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{0}
}

func (m *ContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRequest.Unmarshal(m, b)
}
func (m *ContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRequest.Marshal(b, m, deterministic)
}
func (m *ContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRequest.Merge(m, src)
}
func (m *ContainerRequest) XXX_Size() int {
	return xxx_messageInfo_ContainerRequest.Size(m)
}
func (m *ContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRequest proto.InternalMessageInfo

func (m *ContainerRequest) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *ContainerRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type ContainerResponse struct {
	Containers           []*Container        `protobuf:"bytes,1,rep,name=Containers,proto3" json:"Containers,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ContainerResponse) Reset()         { *m = ContainerResponse{} }
func (m *ContainerResponse) String() string { return proto.CompactTextString(m) }
func (*ContainerResponse) ProtoMessage()    {}
func (*ContainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{1}
}

func (m *ContainerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerResponse.Unmarshal(m, b)
}
func (m *ContainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerResponse.Marshal(b, m, deterministic)
}
func (m *ContainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerResponse.Merge(m, src)
}
func (m *ContainerResponse) XXX_Size() int {
	return xxx_messageInfo_ContainerResponse.Size(m)
}
func (m *ContainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerResponse proto.InternalMessageInfo

func (m *ContainerResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *ContainerResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ContainerResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Container struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status               *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,6,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Container) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Container) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.storage.ContainerType", ContainerType_name, ContainerType_value)
	proto.RegisterType((*ContainerRequest)(nil), "moc.nodeagent.storage.ContainerRequest")
	proto.RegisterType((*ContainerResponse)(nil), "moc.nodeagent.storage.ContainerResponse")
	proto.RegisterType((*Container)(nil), "moc.nodeagent.storage.Container")
}

func init() { proto.RegisterFile("container.proto", fileDescriptor_7afe31759757e49a) }

var fileDescriptor_7afe31759757e49a = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x37, 0x6d, 0x37, 0x4b, 0xa7, 0x6e, 0xcd, 0x0e, 0x2a, 0x31, 0x82, 0x94, 0x7a, 0x61,
	0xf1, 0x62, 0x22, 0xd1, 0x07, 0xd8, 0xb6, 0x16, 0x14, 0x31, 0x0b, 0xa9, 0xae, 0x20, 0x78, 0x91,
	0xa6, 0xa7, 0x69, 0xd8, 0x66, 0xce, 0x38, 0x33, 0x51, 0xfa, 0x0a, 0x82, 0xcf, 0xe0, 0x73, 0xf8,
	0x76, 0x92, 0x49, 0x9a, 0x8d, 0x7f, 0x60, 0x6f, 0xf6, 0xee, 0xcc, 0x39, 0xbf, 0xf3, 0xcd, 0x37,
	0x67, 0x0e, 0xb9, 0x9b, 0x20, 0xd7, 0x71, 0xc6, 0x41, 0x32, 0x21, 0x51, 0x23, 0xbd, 0x9f, 0x63,
	0xc2, 0x38, 0xae, 0x21, 0x4e, 0x81, 0x6b, 0xa6, 0x34, 0xca, 0x38, 0x05, 0xef, 0x71, 0x8a, 0x98,
	0xee, 0xc0, 0x37, 0xd0, 0xaa, 0xd8, 0xf8, 0xdf, 0x64, 0x2c, 0x04, 0x48, 0x55, 0xb5, 0x79, 0x77,
	0x12, 0xcc, 0x73, 0xe4, 0xf5, 0xe9, 0xd1, 0xdf, 0x34, 0xe4, 0x42, 0xef, 0xeb, 0x22, 0xe5, 0xa8,
	0xb3, 0x4d, 0x96, 0xc4, 0x3a, 0x3b, 0x34, 0x8c, 0xbf, 0x5b, 0xc4, 0x99, 0x1f, 0x9c, 0x44, 0xf0,
	0xa5, 0x00, 0xa5, 0xe9, 0x39, 0x21, 0x4d, 0x4e, 0xb9, 0xd6, 0xa8, 0x3b, 0x19, 0x04, 0x23, 0xf6,
	0x5f, 0x7f, 0xec, 0xba, 0xb9, 0xd5, 0x43, 0x5f, 0x92, 0xd3, 0x0b, 0x01, 0xd2, 0xdc, 0xf4, 0x7e,
	0x2f, 0xc0, 0xed, 0x8c, 0xac, 0xc9, 0x30, 0x18, 0x1a, 0x91, 0xa6, 0x12, 0xfd, 0x09, 0x8d, 0x7f,
	0x5a, 0xe4, 0xac, 0x65, 0x46, 0x09, 0xe4, 0x0a, 0x6e, 0xc1, 0x4d, 0x40, 0xec, 0x08, 0x54, 0xb1,
	0xd3, 0xc6, 0xc6, 0x20, 0xf0, 0x58, 0x35, 0x26, 0x76, 0x18, 0x13, 0x9b, 0x21, 0xee, 0x2e, 0xe3,
	0x5d, 0x01, 0x51, 0x4d, 0xd2, 0x7b, 0xe4, 0x78, 0x21, 0x25, 0x4a, 0xb7, 0x3b, 0xb2, 0x26, 0xfd,
	0xa8, 0x3a, 0x8c, 0x7f, 0x58, 0xa4, 0xdf, 0x08, 0x53, 0x4a, 0x7a, 0x3c, 0xce, 0xc1, 0xb5, 0x0c,
	0x62, 0x62, 0x3a, 0x24, 0x9d, 0x6c, 0x6d, 0xee, 0xe9, 0x47, 0x9d, 0x6c, 0x5d, 0x32, 0x22, 0xd6,
	0x5b, 0xb7, 0x57, 0x31, 0x65, 0x4c, 0x9f, 0x10, 0x5b, 0xe9, 0x58, 0x17, 0xca, 0x3d, 0x36, 0x7e,
	0x06, 0xe6, 0x35, 0x4b, 0x93, 0x8a, 0xea, 0x52, 0x09, 0x01, 0xd7, 0x99, 0xde, 0xbb, 0x76, 0x0b,
	0x5a, 0x98, 0x54, 0x54, 0x97, 0x9e, 0x9d, 0x93, 0xd3, 0xc6, 0x4e, 0x39, 0x42, 0x3a, 0x20, 0x27,
	0x1f, 0xc2, 0xb7, 0xe1, 0xc5, 0xc7, 0xd0, 0x39, 0xa2, 0x27, 0xa4, 0xbb, 0x9c, 0x86, 0x8e, 0x55,
	0x06, 0xf3, 0xe5, 0xa5, 0xd3, 0x31, 0x99, 0x77, 0x33, 0xa7, 0x5b, 0x06, 0xaf, 0xa6, 0x4b, 0xa7,
	0x17, 0xfc, 0xb2, 0xc8, 0xb0, 0x91, 0x98, 0x96, 0xc3, 0xa4, 0x9f, 0x89, 0xfd, 0x86, 0x7f, 0xc5,
	0x2b, 0xa0, 0x4f, 0x6f, 0x1c, 0x73, 0xb5, 0x31, 0xde, 0xe4, 0x66, 0xb0, 0xfa, 0xcd, 0xf1, 0x11,
	0x7d, 0x4d, 0xce, 0xe6, 0x5b, 0x48, 0xae, 0xc2, 0xd6, 0x36, 0xd2, 0x07, 0xff, 0x7c, 0xc9, 0xa2,
	0xdc, 0x5c, 0xef, 0xa1, 0x11, 0x6e, 0xa3, 0xd7, 0x4a, 0xb3, 0xe7, 0x9f, 0x58, 0x9a, 0xe9, 0x6d,
	0xb1, 0x62, 0x09, 0xe6, 0x7e, 0x9e, 0x25, 0x12, 0x15, 0x6e, 0xb4, 0x9f, 0x63, 0xe2, 0x4b, 0x91,
	0xf8, 0x8d, 0x1f, 0xbf, 0xf6, 0xb3, 0xb2, 0x8d, 0xfc, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6f, 0xab, 0x40, 0xab, 0x7e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContainerAgentClient is the client API for ContainerAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContainerAgentClient interface {
	Invoke(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*ContainerResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type containerAgentClient struct {
	cc *grpc.ClientConn
}

func NewContainerAgentClient(cc *grpc.ClientConn) ContainerAgentClient {
	return &containerAgentClient{cc}
}

func (c *containerAgentClient) Invoke(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*ContainerResponse, error) {
	out := new(ContainerResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.ContainerAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.ContainerAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerAgentServer is the server API for ContainerAgent service.
type ContainerAgentServer interface {
	Invoke(context.Context, *ContainerRequest) (*ContainerResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedContainerAgentServer can be embedded to have forward compatible implementations.
type UnimplementedContainerAgentServer struct {
}

func (*UnimplementedContainerAgentServer) Invoke(ctx context.Context, req *ContainerRequest) (*ContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedContainerAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterContainerAgentServer(s *grpc.Server, srv ContainerAgentServer) {
	s.RegisterService(&_ContainerAgent_serviceDesc, srv)
}

func _ContainerAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.ContainerAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerAgentServer).Invoke(ctx, req.(*ContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.ContainerAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContainerAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.storage.ContainerAgent",
	HandlerType: (*ContainerAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _ContainerAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _ContainerAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "container.proto",
}
