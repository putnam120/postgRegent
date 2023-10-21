// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/api/postgregent/postgregent.proto

package postgregent

import (
	context "context"
	postgregent "github.com/putnam120/postgRegent/proto/api/postgregent"
	rbac "github.com/putnam120/postgRegent/proto/api/rbac"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostgRegentService_CreatePermission_FullMethodName = "/postgregent.PostgRegentService/CreatePermission"
	PostgRegentService_CreateRole_FullMethodName       = "/postgregent.PostgRegentService/CreateRole"
	PostgRegentService_CreateUSer_FullMethodName       = "/postgregent.PostgRegentService/CreateUSer"
)

// PostgRegentServiceClient is the client API for PostgRegentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgRegentServiceClient interface {
	CreatePermission(ctx context.Context, in *rbac.Permission, opts ...grpc.CallOption) (*postgregent.Status, error)
	CreateRole(ctx context.Context, in *rbac.Role, opts ...grpc.CallOption) (*postgregent.Status, error)
	CreateUSer(ctx context.Context, in *rbac.User, opts ...grpc.CallOption) (*postgregent.Status, error)
}

type postgRegentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgRegentServiceClient(cc grpc.ClientConnInterface) PostgRegentServiceClient {
	return &postgRegentServiceClient{cc}
}

func (c *postgRegentServiceClient) CreatePermission(ctx context.Context, in *rbac.Permission, opts ...grpc.CallOption) (*postgregent.Status, error) {
	out := new(postgregent.Status)
	err := c.cc.Invoke(ctx, PostgRegentService_CreatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgRegentServiceClient) CreateRole(ctx context.Context, in *rbac.Role, opts ...grpc.CallOption) (*postgregent.Status, error) {
	out := new(postgregent.Status)
	err := c.cc.Invoke(ctx, PostgRegentService_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgRegentServiceClient) CreateUSer(ctx context.Context, in *rbac.User, opts ...grpc.CallOption) (*postgregent.Status, error) {
	out := new(postgregent.Status)
	err := c.cc.Invoke(ctx, PostgRegentService_CreateUSer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgRegentServiceServer is the server API for PostgRegentService service.
// All implementations must embed UnimplementedPostgRegentServiceServer
// for forward compatibility
type PostgRegentServiceServer interface {
	CreatePermission(context.Context, *rbac.Permission) (*postgregent.Status, error)
	CreateRole(context.Context, *rbac.Role) (*postgregent.Status, error)
	CreateUSer(context.Context, *rbac.User) (*postgregent.Status, error)
	mustEmbedUnimplementedPostgRegentServiceServer()
}

// UnimplementedPostgRegentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostgRegentServiceServer struct {
}

func (UnimplementedPostgRegentServiceServer) CreatePermission(context.Context, *rbac.Permission) (*postgregent.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedPostgRegentServiceServer) CreateRole(context.Context, *rbac.Role) (*postgregent.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedPostgRegentServiceServer) CreateUSer(context.Context, *rbac.User) (*postgregent.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUSer not implemented")
}
func (UnimplementedPostgRegentServiceServer) mustEmbedUnimplementedPostgRegentServiceServer() {}

// UnsafePostgRegentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgRegentServiceServer will
// result in compilation errors.
type UnsafePostgRegentServiceServer interface {
	mustEmbedUnimplementedPostgRegentServiceServer()
}

func RegisterPostgRegentServiceServer(s grpc.ServiceRegistrar, srv PostgRegentServiceServer) {
	s.RegisterService(&PostgRegentService_ServiceDesc, srv)
}

func _PostgRegentService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rbac.Permission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgRegentServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgRegentService_CreatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgRegentServiceServer).CreatePermission(ctx, req.(*rbac.Permission))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgRegentService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rbac.Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgRegentServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgRegentService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgRegentServiceServer).CreateRole(ctx, req.(*rbac.Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgRegentService_CreateUSer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rbac.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgRegentServiceServer).CreateUSer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgRegentService_CreateUSer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgRegentServiceServer).CreateUSer(ctx, req.(*rbac.User))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgRegentService_ServiceDesc is the grpc.ServiceDesc for PostgRegentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgRegentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postgregent.PostgRegentService",
	HandlerType: (*PostgRegentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePermission",
			Handler:    _PostgRegentService_CreatePermission_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _PostgRegentService_CreateRole_Handler,
		},
		{
			MethodName: "CreateUSer",
			Handler:    _PostgRegentService_CreateUSer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/postgregent/postgregent.proto",
}