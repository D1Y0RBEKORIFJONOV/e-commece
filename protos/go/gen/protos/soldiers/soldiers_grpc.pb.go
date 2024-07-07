// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.5
// source: protos/soldiers/soldiers.proto

package soldiers1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SoldiersService_CreateSoldiers_FullMethodName = "/SoldiersService/CreateSoldiers"
	SoldiersService_RegisterUser_FullMethodName   = "/SoldiersService/RegisterUser"
	SoldiersService_Login_FullMethodName          = "/SoldiersService/Login"
	SoldiersService_GetSoldier_FullMethodName     = "/SoldiersService/GetSoldier"
	SoldiersService_GetAllSoldiers_FullMethodName = "/SoldiersService/GetAllSoldiers"
	SoldiersService_UpdateSoldier_FullMethodName  = "/SoldiersService/UpdateSoldier"
	SoldiersService_DeleteSoldier_FullMethodName  = "/SoldiersService/DeleteSoldier"
)

// SoldiersServiceClient is the client API for SoldiersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoldiersServiceClient interface {
	CreateSoldiers(ctx context.Context, in *CreateSoldiersReq, opts ...grpc.CallOption) (*Status, error)
	RegisterUser(ctx context.Context, in *RegisterAndLoginReq, opts ...grpc.CallOption) (*Soldiers, error)
	Login(ctx context.Context, in *RegisterAndLoginReq, opts ...grpc.CallOption) (*LogerResponse, error)
	GetSoldier(ctx context.Context, in *GetSoldierReq, opts ...grpc.CallOption) (*Soldiers, error)
	GetAllSoldiers(ctx context.Context, in *GetAllSoldierReq, opts ...grpc.CallOption) (*GetSoldierRequestResponse, error)
	UpdateSoldier(ctx context.Context, in *UpdateSoldierReq, opts ...grpc.CallOption) (*Soldiers, error)
	DeleteSoldier(ctx context.Context, in *DeleteSoldierReq, opts ...grpc.CallOption) (*Status, error)
}

type soldiersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSoldiersServiceClient(cc grpc.ClientConnInterface) SoldiersServiceClient {
	return &soldiersServiceClient{cc}
}

func (c *soldiersServiceClient) CreateSoldiers(ctx context.Context, in *CreateSoldiersReq, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, SoldiersService_CreateSoldiers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soldiersServiceClient) RegisterUser(ctx context.Context, in *RegisterAndLoginReq, opts ...grpc.CallOption) (*Soldiers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Soldiers)
	err := c.cc.Invoke(ctx, SoldiersService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soldiersServiceClient) Login(ctx context.Context, in *RegisterAndLoginReq, opts ...grpc.CallOption) (*LogerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogerResponse)
	err := c.cc.Invoke(ctx, SoldiersService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soldiersServiceClient) GetSoldier(ctx context.Context, in *GetSoldierReq, opts ...grpc.CallOption) (*Soldiers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Soldiers)
	err := c.cc.Invoke(ctx, SoldiersService_GetSoldier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soldiersServiceClient) GetAllSoldiers(ctx context.Context, in *GetAllSoldierReq, opts ...grpc.CallOption) (*GetSoldierRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSoldierRequestResponse)
	err := c.cc.Invoke(ctx, SoldiersService_GetAllSoldiers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soldiersServiceClient) UpdateSoldier(ctx context.Context, in *UpdateSoldierReq, opts ...grpc.CallOption) (*Soldiers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Soldiers)
	err := c.cc.Invoke(ctx, SoldiersService_UpdateSoldier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soldiersServiceClient) DeleteSoldier(ctx context.Context, in *DeleteSoldierReq, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, SoldiersService_DeleteSoldier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SoldiersServiceServer is the server API for SoldiersService service.
// All implementations must embed UnimplementedSoldiersServiceServer
// for forward compatibility
type SoldiersServiceServer interface {
	CreateSoldiers(context.Context, *CreateSoldiersReq) (*Status, error)
	RegisterUser(context.Context, *RegisterAndLoginReq) (*Soldiers, error)
	Login(context.Context, *RegisterAndLoginReq) (*LogerResponse, error)
	GetSoldier(context.Context, *GetSoldierReq) (*Soldiers, error)
	GetAllSoldiers(context.Context, *GetAllSoldierReq) (*GetSoldierRequestResponse, error)
	UpdateSoldier(context.Context, *UpdateSoldierReq) (*Soldiers, error)
	DeleteSoldier(context.Context, *DeleteSoldierReq) (*Status, error)
	mustEmbedUnimplementedSoldiersServiceServer()
}

// UnimplementedSoldiersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSoldiersServiceServer struct {
}

func (UnimplementedSoldiersServiceServer) CreateSoldiers(context.Context, *CreateSoldiersReq) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSoldiers not implemented")
}
func (UnimplementedSoldiersServiceServer) RegisterUser(context.Context, *RegisterAndLoginReq) (*Soldiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedSoldiersServiceServer) Login(context.Context, *RegisterAndLoginReq) (*LogerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSoldiersServiceServer) GetSoldier(context.Context, *GetSoldierReq) (*Soldiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoldier not implemented")
}
func (UnimplementedSoldiersServiceServer) GetAllSoldiers(context.Context, *GetAllSoldierReq) (*GetSoldierRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSoldiers not implemented")
}
func (UnimplementedSoldiersServiceServer) UpdateSoldier(context.Context, *UpdateSoldierReq) (*Soldiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSoldier not implemented")
}
func (UnimplementedSoldiersServiceServer) DeleteSoldier(context.Context, *DeleteSoldierReq) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSoldier not implemented")
}
func (UnimplementedSoldiersServiceServer) mustEmbedUnimplementedSoldiersServiceServer() {}

// UnsafeSoldiersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoldiersServiceServer will
// result in compilation errors.
type UnsafeSoldiersServiceServer interface {
	mustEmbedUnimplementedSoldiersServiceServer()
}

func RegisterSoldiersServiceServer(s grpc.ServiceRegistrar, srv SoldiersServiceServer) {
	s.RegisterService(&SoldiersService_ServiceDesc, srv)
}

func _SoldiersService_CreateSoldiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSoldiersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).CreateSoldiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_CreateSoldiers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).CreateSoldiers(ctx, req.(*CreateSoldiersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoldiersService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAndLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).RegisterUser(ctx, req.(*RegisterAndLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoldiersService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAndLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).Login(ctx, req.(*RegisterAndLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoldiersService_GetSoldier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSoldierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).GetSoldier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_GetSoldier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).GetSoldier(ctx, req.(*GetSoldierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoldiersService_GetAllSoldiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSoldierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).GetAllSoldiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_GetAllSoldiers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).GetAllSoldiers(ctx, req.(*GetAllSoldierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoldiersService_UpdateSoldier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSoldierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).UpdateSoldier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_UpdateSoldier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).UpdateSoldier(ctx, req.(*UpdateSoldierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoldiersService_DeleteSoldier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSoldierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoldiersServiceServer).DeleteSoldier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoldiersService_DeleteSoldier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoldiersServiceServer).DeleteSoldier(ctx, req.(*DeleteSoldierReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SoldiersService_ServiceDesc is the grpc.ServiceDesc for SoldiersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SoldiersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SoldiersService",
	HandlerType: (*SoldiersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSoldiers",
			Handler:    _SoldiersService_CreateSoldiers_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _SoldiersService_RegisterUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SoldiersService_Login_Handler,
		},
		{
			MethodName: "GetSoldier",
			Handler:    _SoldiersService_GetSoldier_Handler,
		},
		{
			MethodName: "GetAllSoldiers",
			Handler:    _SoldiersService_GetAllSoldiers_Handler,
		},
		{
			MethodName: "UpdateSoldier",
			Handler:    _SoldiersService_UpdateSoldier_Handler,
		},
		{
			MethodName: "DeleteSoldier",
			Handler:    _SoldiersService_DeleteSoldier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/soldiers/soldiers.proto",
}