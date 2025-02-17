// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: parser.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ParseService_GetTable_FullMethodName = "/server.ParseService/GetTable"
)

// ParseServiceClient is the client API for ParseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Определение сервиса
type ParseServiceClient interface {
	GetTable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Table, error)
}

type parseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParseServiceClient(cc grpc.ClientConnInterface) ParseServiceClient {
	return &parseServiceClient{cc}
}

func (c *parseServiceClient) GetTable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Table, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Table)
	err := c.cc.Invoke(ctx, ParseService_GetTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParseServiceServer is the server API for ParseService service.
// All implementations must embed UnimplementedParseServiceServer
// for forward compatibility.
//
// Определение сервиса
type ParseServiceServer interface {
	GetTable(context.Context, *emptypb.Empty) (*Table, error)
	mustEmbedUnimplementedParseServiceServer()
}

// UnimplementedParseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParseServiceServer struct{}

func (UnimplementedParseServiceServer) GetTable(context.Context, *emptypb.Empty) (*Table, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTable not implemented")
}
func (UnimplementedParseServiceServer) mustEmbedUnimplementedParseServiceServer() {}
func (UnimplementedParseServiceServer) testEmbeddedByValue()                      {}

// UnsafeParseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParseServiceServer will
// result in compilation errors.
type UnsafeParseServiceServer interface {
	mustEmbedUnimplementedParseServiceServer()
}

func RegisterParseServiceServer(s grpc.ServiceRegistrar, srv ParseServiceServer) {
	// If the following call pancis, it indicates UnimplementedParseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParseService_ServiceDesc, srv)
}

func _ParseService_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParseServiceServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParseService_GetTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParseServiceServer).GetTable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ParseService_ServiceDesc is the grpc.ServiceDesc for ParseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.ParseService",
	HandlerType: (*ParseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTable",
			Handler:    _ParseService_GetTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser.proto",
}
