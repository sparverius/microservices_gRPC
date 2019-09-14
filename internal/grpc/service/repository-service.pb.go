// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microservices_gRPC/internal/proto-files/service/repository-service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	domain "microservices_gRPC/internal/grpc/domain"
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

type AddRepositoryResponse struct {
	AddedRepository      *domain.Repository `protobuf:"bytes,1,opt,name=addedRepository,proto3" json:"addedRepository,omitempty"`
	Error                *Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddRepositoryResponse) Reset()         { *m = AddRepositoryResponse{} }
func (m *AddRepositoryResponse) String() string { return proto.CompactTextString(m) }
func (*AddRepositoryResponse) ProtoMessage()    {}
func (*AddRepositoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f24b145bb638a69, []int{0}
}

func (m *AddRepositoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRepositoryResponse.Unmarshal(m, b)
}
func (m *AddRepositoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRepositoryResponse.Marshal(b, m, deterministic)
}
func (m *AddRepositoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRepositoryResponse.Merge(m, src)
}
func (m *AddRepositoryResponse) XXX_Size() int {
	return xxx_messageInfo_AddRepositoryResponse.Size(m)
}
func (m *AddRepositoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRepositoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRepositoryResponse proto.InternalMessageInfo

func (m *AddRepositoryResponse) GetAddedRepository() *domain.Repository {
	if m != nil {
		return m.AddedRepository
	}
	return nil
}

func (m *AddRepositoryResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f24b145bb638a69, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRepositoryResponse)(nil), "service.AddRepositoryResponse")
	proto.RegisterType((*Error)(nil), "service.Error")
}

func init() {
	proto.RegisterFile("microservices_gRPC/internal/proto-files/service/repository-service.proto", fileDescriptor_9f24b145bb638a69)
}

var fileDescriptor_9f24b145bb638a69 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xc8, 0xcd, 0x4c, 0x2e,
	0xca, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0x8e, 0x4f, 0x0f, 0x0a, 0x70, 0xd6, 0xcf,
	0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcb,
	0xcc, 0x49, 0x2d, 0xd6, 0x87, 0x2a, 0xd1, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0x2f,
	0xaa, 0xd4, 0x85, 0x0a, 0xe9, 0x81, 0x95, 0x09, 0xb1, 0x43, 0xb9, 0x52, 0xf6, 0xc4, 0x1a, 0x99,
	0x92, 0x9f, 0x9b, 0x98, 0x99, 0x87, 0x64, 0x22, 0xc4, 0x24, 0xa5, 0x6a, 0x2e, 0x51, 0xc7, 0x94,
	0x94, 0x20, 0xb8, 0x70, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x0d, 0x17, 0x7f,
	0x62, 0x4a, 0x4a, 0x2a, 0x92, 0x94, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x90, 0x1e, 0xc4,
	0x2c, 0x3d, 0x24, 0x4d, 0xe8, 0x4a, 0x85, 0x54, 0xb8, 0x58, 0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0x24,
	0x98, 0xc0, 0x7a, 0xf8, 0xf4, 0x60, 0xee, 0x77, 0x05, 0x89, 0x06, 0x41, 0x24, 0x95, 0x4c, 0xb9,
	0x58, 0xc1, 0x7c, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0xb0, 0x0d, 0x9c, 0x41, 0x60,
	0xb6, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0x2a, 0xd8, 0x10, 0xce, 0x20,
	0x18, 0xd7, 0xc8, 0x8f, 0x4b, 0x10, 0x61, 0x55, 0x30, 0xc4, 0x60, 0x21, 0x4b, 0x2e, 0xe6, 0xc4,
	0x94, 0x14, 0x21, 0x2c, 0xae, 0x93, 0x92, 0x83, 0xdb, 0x8e, 0xd5, 0xab, 0x4e, 0x5a, 0x51, 0x1a,
	0xf8, 0x82, 0x31, 0xbd, 0xa8, 0x20, 0x19, 0x16, 0x25, 0x49, 0x6c, 0xe0, 0x60, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x45, 0xac, 0x85, 0x80, 0xcc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	Add(ctx context.Context, in *domain.Repository, opts ...grpc.CallOption) (*AddRepositoryResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) Add(ctx context.Context, in *domain.Repository, opts ...grpc.CallOption) (*AddRepositoryResponse, error) {
	out := new(AddRepositoryResponse)
	err := c.cc.Invoke(ctx, "/service.RepositoryService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
type RepositoryServiceServer interface {
	Add(context.Context, *domain.Repository) (*AddRepositoryResponse, error)
}

// UnimplementedRepositoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepositoryServiceServer struct {
}

func (*UnimplementedRepositoryServiceServer) Add(ctx context.Context, req *domain.Repository) (*AddRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RepositoryService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Add(ctx, req.(*domain.Repository))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _RepositoryService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservices_gRPC/internal/proto-files/service/repository-service.proto",
}
