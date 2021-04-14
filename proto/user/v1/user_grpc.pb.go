// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	ScheduleReminder(ctx context.Context, in *ScheduleReminderRequest, opts ...grpc.CallOption) (*ScheduleReminderResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	InsertUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ScheduleReminder(ctx context.Context, in *ScheduleReminderRequest, opts ...grpc.CallOption) (*ScheduleReminderResponse, error) {
	out := new(ScheduleReminderResponse)
	err := c.cc.Invoke(ctx, "/pack.user.v1.UserService/ScheduleReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/pack.user.v1.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) InsertUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error) {
	out := new(InsertUserResponse)
	err := c.cc.Invoke(ctx, "/pack.user.v1.UserService/InsertUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	ScheduleReminder(context.Context, *ScheduleReminderRequest) (*ScheduleReminderResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*UsersResponse, error)
	InsertUser(context.Context, *InsertUserRequest) (*InsertUserResponse, error)
	//mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) ScheduleReminder(context.Context, *ScheduleReminderRequest) (*ScheduleReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleReminder not implemented")
}
func (UnimplementedUserServiceServer) GetUsers(context.Context, *GetUsersRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServiceServer) InsertUser(context.Context, *InsertUserRequest) (*InsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_ScheduleReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ScheduleReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pack.user.v1.UserService/ScheduleReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ScheduleReminder(ctx, req.(*ScheduleReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pack.user.v1.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pack.user.v1.UserService/InsertUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).InsertUser(ctx, req.(*InsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pack.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleReminder",
			Handler:    _UserService_ScheduleReminder_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "InsertUser",
			Handler:    _UserService_InsertUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/v1/user.proto",
}
