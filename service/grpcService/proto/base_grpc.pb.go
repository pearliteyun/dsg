// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	GetIDByMobile(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*UserID, error)
	GetInfoByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Info, error)
	ListByIDs(ctx context.Context, opts ...grpc.CallOption) (UserService_ListByIDsClient, error)
	//  登录
	Login(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Token, error)
	//  刷新token
	ExpireToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  刷新token
	ExpireTokenByID(ctx context.Context, in *Token, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 修改手机号
	UpdateMobile(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetIDByMobile(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetIDByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetInfoByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetInfoByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListByIDs(ctx context.Context, opts ...grpc.CallOption) (UserService_ListByIDsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/proto.UserService/ListByIDs", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListByIDsClient{stream}
	return x, nil
}

type UserService_ListByIDsClient interface {
	Send(*UserID) error
	Recv() (*UserList, error)
	grpc.ClientStream
}

type userServiceListByIDsClient struct {
	grpc.ClientStream
}

func (x *userServiceListByIDsClient) Send(m *UserID) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceListByIDsClient) Recv() (*UserList, error) {
	m := new(UserList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ExpireToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserService/ExpireToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ExpireTokenByID(ctx context.Context, in *Token, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserService/ExpireTokenByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateMobile(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserService/UpdateMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetByID(context.Context, *UserID) (*User, error)
	GetIDByMobile(context.Context, *Mobile) (*UserID, error)
	GetInfoByToken(context.Context, *Token) (*Info, error)
	ListByIDs(UserService_ListByIDsServer) error
	//  登录
	Login(context.Context, *Login) (*Token, error)
	//  刷新token
	ExpireToken(context.Context, *Token) (*emptypb.Empty, error)
	//  刷新token
	ExpireTokenByID(context.Context, *Token) (*emptypb.Empty, error)
	// 修改手机号
	UpdateMobile(context.Context, *User) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetByID(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServiceServer) GetIDByMobile(context.Context, *Mobile) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIDByMobile not implemented")
}
func (UnimplementedUserServiceServer) GetInfoByToken(context.Context, *Token) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByToken not implemented")
}
func (UnimplementedUserServiceServer) ListByIDs(UserService_ListByIDsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListByIDs not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *Login) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) ExpireToken(context.Context, *Token) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpireToken not implemented")
}
func (UnimplementedUserServiceServer) ExpireTokenByID(context.Context, *Token) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpireTokenByID not implemented")
}
func (UnimplementedUserServiceServer) UpdateMobile(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMobile not implemented")
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

func _UserService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByID(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetIDByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mobile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetIDByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetIDByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetIDByMobile(ctx, req.(*Mobile))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetInfoByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInfoByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetInfoByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInfoByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListByIDs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).ListByIDs(&userServiceListByIDsServer{stream})
}

type UserService_ListByIDsServer interface {
	Send(*UserList) error
	Recv() (*UserID, error)
	grpc.ServerStream
}

type userServiceListByIDsServer struct {
	grpc.ServerStream
}

func (x *userServiceListByIDsServer) Send(m *UserList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceListByIDsServer) Recv() (*UserID, error) {
	m := new(UserID)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ExpireToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ExpireToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ExpireToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ExpireToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ExpireTokenByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ExpireTokenByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ExpireTokenByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ExpireTokenByID(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UpdateMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateMobile(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _UserService_GetByID_Handler,
		},
		{
			MethodName: "GetIDByMobile",
			Handler:    _UserService_GetIDByMobile_Handler,
		},
		{
			MethodName: "GetInfoByToken",
			Handler:    _UserService_GetInfoByToken_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "ExpireToken",
			Handler:    _UserService_ExpireToken_Handler,
		},
		{
			MethodName: "ExpireTokenByID",
			Handler:    _UserService_ExpireTokenByID_Handler,
		},
		{
			MethodName: "UpdateMobile",
			Handler:    _UserService_UpdateMobile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListByIDs",
			Handler:       _UserService_ListByIDs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service/grpcService/proto/base.proto",
}

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgServiceClient interface {
	//  发送验证码
	SendMsgCode(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*MsgCodeBack, error)
	//  校验验证码
	CheckMsgCode(ctx context.Context, in *MsgCode, opts ...grpc.CallOption) (*BoolF, error)
}

type msgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgServiceClient(cc grpc.ClientConnInterface) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) SendMsgCode(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*MsgCodeBack, error) {
	out := new(MsgCodeBack)
	err := c.cc.Invoke(ctx, "/proto.MsgService/SendMsgCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CheckMsgCode(ctx context.Context, in *MsgCode, opts ...grpc.CallOption) (*BoolF, error) {
	out := new(BoolF)
	err := c.cc.Invoke(ctx, "/proto.MsgService/CheckMsgCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
// All implementations must embed UnimplementedMsgServiceServer
// for forward compatibility
type MsgServiceServer interface {
	//  发送验证码
	SendMsgCode(context.Context, *Mobile) (*MsgCodeBack, error)
	//  校验验证码
	CheckMsgCode(context.Context, *MsgCode) (*BoolF, error)
	mustEmbedUnimplementedMsgServiceServer()
}

// UnimplementedMsgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (UnimplementedMsgServiceServer) SendMsgCode(context.Context, *Mobile) (*MsgCodeBack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgCode not implemented")
}
func (UnimplementedMsgServiceServer) CheckMsgCode(context.Context, *MsgCode) (*BoolF, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMsgCode not implemented")
}
func (UnimplementedMsgServiceServer) mustEmbedUnimplementedMsgServiceServer() {}

// UnsafeMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServiceServer will
// result in compilation errors.
type UnsafeMsgServiceServer interface {
	mustEmbedUnimplementedMsgServiceServer()
}

func RegisterMsgServiceServer(s grpc.ServiceRegistrar, srv MsgServiceServer) {
	s.RegisterService(&MsgService_ServiceDesc, srv)
}

func _MsgService_SendMsgCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mobile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SendMsgCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MsgService/SendMsgCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SendMsgCode(ctx, req.(*Mobile))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CheckMsgCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CheckMsgCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MsgService/CheckMsgCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CheckMsgCode(ctx, req.(*MsgCode))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgService_ServiceDesc is the grpc.ServiceDesc for MsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsgCode",
			Handler:    _MsgService_SendMsgCode_Handler,
		},
		{
			MethodName: "CheckMsgCode",
			Handler:    _MsgService_CheckMsgCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/grpcService/proto/base.proto",
}

// WechatServiceClient is the client API for WechatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WechatServiceClient interface {
	//  登录
	Login(ctx context.Context, in *WechatLogin, opts ...grpc.CallOption) (*MiniLoginBack, error)
	// 获取access_token
	GetAccessToken(ctx context.Context, in *AppID, opts ...grpc.CallOption) (*Token, error)
	//  从用户系统获取
	GetAccessTokenFromUser(ctx context.Context, in *AppID, opts ...grpc.CallOption) (*Token, error)
	//  微信绑定用户信息
	BindUser(ctx context.Context, in *WechatMobile, opts ...grpc.CallOption) (*UserID, error)
}

type wechatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWechatServiceClient(cc grpc.ClientConnInterface) WechatServiceClient {
	return &wechatServiceClient{cc}
}

func (c *wechatServiceClient) Login(ctx context.Context, in *WechatLogin, opts ...grpc.CallOption) (*MiniLoginBack, error) {
	out := new(MiniLoginBack)
	err := c.cc.Invoke(ctx, "/proto.WechatService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatServiceClient) GetAccessToken(ctx context.Context, in *AppID, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.WechatService/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatServiceClient) GetAccessTokenFromUser(ctx context.Context, in *AppID, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.WechatService/GetAccessTokenFromUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatServiceClient) BindUser(ctx context.Context, in *WechatMobile, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/proto.WechatService/BindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatServiceServer is the server API for WechatService service.
// All implementations must embed UnimplementedWechatServiceServer
// for forward compatibility
type WechatServiceServer interface {
	//  登录
	Login(context.Context, *WechatLogin) (*MiniLoginBack, error)
	// 获取access_token
	GetAccessToken(context.Context, *AppID) (*Token, error)
	//  从用户系统获取
	GetAccessTokenFromUser(context.Context, *AppID) (*Token, error)
	//  微信绑定用户信息
	BindUser(context.Context, *WechatMobile) (*UserID, error)
	mustEmbedUnimplementedWechatServiceServer()
}

// UnimplementedWechatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWechatServiceServer struct {
}

func (UnimplementedWechatServiceServer) Login(context.Context, *WechatLogin) (*MiniLoginBack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedWechatServiceServer) GetAccessToken(context.Context, *AppID) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedWechatServiceServer) GetAccessTokenFromUser(context.Context, *AppID) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessTokenFromUser not implemented")
}
func (UnimplementedWechatServiceServer) BindUser(context.Context, *WechatMobile) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindUser not implemented")
}
func (UnimplementedWechatServiceServer) mustEmbedUnimplementedWechatServiceServer() {}

// UnsafeWechatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WechatServiceServer will
// result in compilation errors.
type UnsafeWechatServiceServer interface {
	mustEmbedUnimplementedWechatServiceServer()
}

func RegisterWechatServiceServer(s grpc.ServiceRegistrar, srv WechatServiceServer) {
	s.RegisterService(&WechatService_ServiceDesc, srv)
}

func _WechatService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WechatLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WechatService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatServiceServer).Login(ctx, req.(*WechatLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatService_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatServiceServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WechatService/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatServiceServer).GetAccessToken(ctx, req.(*AppID))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatService_GetAccessTokenFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatServiceServer).GetAccessTokenFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WechatService/GetAccessTokenFromUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatServiceServer).GetAccessTokenFromUser(ctx, req.(*AppID))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatService_BindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WechatMobile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatServiceServer).BindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WechatService/BindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatServiceServer).BindUser(ctx, req.(*WechatMobile))
	}
	return interceptor(ctx, in, info, handler)
}

// WechatService_ServiceDesc is the grpc.ServiceDesc for WechatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WechatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WechatService",
	HandlerType: (*WechatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _WechatService_Login_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _WechatService_GetAccessToken_Handler,
		},
		{
			MethodName: "GetAccessTokenFromUser",
			Handler:    _WechatService_GetAccessTokenFromUser_Handler,
		},
		{
			MethodName: "BindUser",
			Handler:    _WechatService_BindUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/grpcService/proto/base.proto",
}

// UserCardClient is the client API for UserCard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCardClient interface {
	// 获取信息
	GetUserCard(ctx context.Context, in *UserCardReq, opts ...grpc.CallOption) (*UserCardRes, error)
	// 修改身份信息
	UpdateUserCard(ctx context.Context, in *UpdateUserCardReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取地址
	GetAddressList(ctx context.Context, in *UserCardReq, opts ...grpc.CallOption) (*UserCardRes, error)
	// 更新地址
	UpAddress(ctx context.Context, in *UpdateUserCardReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 插入地址
	InstallAddress(ctx context.Context, in *UpdateUserCardReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userCardClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCardClient(cc grpc.ClientConnInterface) UserCardClient {
	return &userCardClient{cc}
}

func (c *userCardClient) GetUserCard(ctx context.Context, in *UserCardReq, opts ...grpc.CallOption) (*UserCardRes, error) {
	out := new(UserCardRes)
	err := c.cc.Invoke(ctx, "/proto.UserCard/GetUserCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCardClient) UpdateUserCard(ctx context.Context, in *UpdateUserCardReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserCard/UpdateUserCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCardClient) GetAddressList(ctx context.Context, in *UserCardReq, opts ...grpc.CallOption) (*UserCardRes, error) {
	out := new(UserCardRes)
	err := c.cc.Invoke(ctx, "/proto.UserCard/GetAddressList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCardClient) UpAddress(ctx context.Context, in *UpdateUserCardReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserCard/UpAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCardClient) InstallAddress(ctx context.Context, in *UpdateUserCardReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserCard/InstallAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCardServer is the server API for UserCard service.
// All implementations must embed UnimplementedUserCardServer
// for forward compatibility
type UserCardServer interface {
	// 获取信息
	GetUserCard(context.Context, *UserCardReq) (*UserCardRes, error)
	// 修改身份信息
	UpdateUserCard(context.Context, *UpdateUserCardReq) (*emptypb.Empty, error)
	// 获取地址
	GetAddressList(context.Context, *UserCardReq) (*UserCardRes, error)
	// 更新地址
	UpAddress(context.Context, *UpdateUserCardReq) (*emptypb.Empty, error)
	// 插入地址
	InstallAddress(context.Context, *UpdateUserCardReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserCardServer()
}

// UnimplementedUserCardServer must be embedded to have forward compatible implementations.
type UnimplementedUserCardServer struct {
}

func (UnimplementedUserCardServer) GetUserCard(context.Context, *UserCardReq) (*UserCardRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCard not implemented")
}
func (UnimplementedUserCardServer) UpdateUserCard(context.Context, *UpdateUserCardReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserCard not implemented")
}
func (UnimplementedUserCardServer) GetAddressList(context.Context, *UserCardReq) (*UserCardRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressList not implemented")
}
func (UnimplementedUserCardServer) UpAddress(context.Context, *UpdateUserCardReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpAddress not implemented")
}
func (UnimplementedUserCardServer) InstallAddress(context.Context, *UpdateUserCardReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallAddress not implemented")
}
func (UnimplementedUserCardServer) mustEmbedUnimplementedUserCardServer() {}

// UnsafeUserCardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCardServer will
// result in compilation errors.
type UnsafeUserCardServer interface {
	mustEmbedUnimplementedUserCardServer()
}

func RegisterUserCardServer(s grpc.ServiceRegistrar, srv UserCardServer) {
	s.RegisterService(&UserCard_ServiceDesc, srv)
}

func _UserCard_GetUserCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCardServer).GetUserCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserCard/GetUserCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCardServer).GetUserCard(ctx, req.(*UserCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCard_UpdateUserCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCardServer).UpdateUserCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserCard/UpdateUserCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCardServer).UpdateUserCard(ctx, req.(*UpdateUserCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCard_GetAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCardServer).GetAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserCard/GetAddressList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCardServer).GetAddressList(ctx, req.(*UserCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCard_UpAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCardServer).UpAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserCard/UpAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCardServer).UpAddress(ctx, req.(*UpdateUserCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCard_InstallAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCardServer).InstallAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserCard/InstallAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCardServer).InstallAddress(ctx, req.(*UpdateUserCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCard_ServiceDesc is the grpc.ServiceDesc for UserCard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserCard",
	HandlerType: (*UserCardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserCard",
			Handler:    _UserCard_GetUserCard_Handler,
		},
		{
			MethodName: "UpdateUserCard",
			Handler:    _UserCard_UpdateUserCard_Handler,
		},
		{
			MethodName: "GetAddressList",
			Handler:    _UserCard_GetAddressList_Handler,
		},
		{
			MethodName: "UpAddress",
			Handler:    _UserCard_UpAddress_Handler,
		},
		{
			MethodName: "InstallAddress",
			Handler:    _UserCard_InstallAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/grpcService/proto/base.proto",
}

// UserHistoryClient is the client API for UserHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserHistoryClient interface {
	Add(ctx context.Context, in *UserHistoryAddReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Query(ctx context.Context, in *UserHistoryQueryReq, opts ...grpc.CallOption) (*UserHistoryQueryRes, error)
}

type userHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewUserHistoryClient(cc grpc.ClientConnInterface) UserHistoryClient {
	return &userHistoryClient{cc}
}

func (c *userHistoryClient) Add(ctx context.Context, in *UserHistoryAddReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.UserHistory/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHistoryClient) Query(ctx context.Context, in *UserHistoryQueryReq, opts ...grpc.CallOption) (*UserHistoryQueryRes, error) {
	out := new(UserHistoryQueryRes)
	err := c.cc.Invoke(ctx, "/proto.UserHistory/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHistoryServer is the server API for UserHistory service.
// All implementations must embed UnimplementedUserHistoryServer
// for forward compatibility
type UserHistoryServer interface {
	Add(context.Context, *UserHistoryAddReq) (*emptypb.Empty, error)
	Query(context.Context, *UserHistoryQueryReq) (*UserHistoryQueryRes, error)
	mustEmbedUnimplementedUserHistoryServer()
}

// UnimplementedUserHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedUserHistoryServer struct {
}

func (UnimplementedUserHistoryServer) Add(context.Context, *UserHistoryAddReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedUserHistoryServer) Query(context.Context, *UserHistoryQueryReq) (*UserHistoryQueryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedUserHistoryServer) mustEmbedUnimplementedUserHistoryServer() {}

// UnsafeUserHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserHistoryServer will
// result in compilation errors.
type UnsafeUserHistoryServer interface {
	mustEmbedUnimplementedUserHistoryServer()
}

func RegisterUserHistoryServer(s grpc.ServiceRegistrar, srv UserHistoryServer) {
	s.RegisterService(&UserHistory_ServiceDesc, srv)
}

func _UserHistory_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHistoryAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHistoryServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserHistory/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHistoryServer).Add(ctx, req.(*UserHistoryAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHistory_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHistoryQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHistoryServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserHistory/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHistoryServer).Query(ctx, req.(*UserHistoryQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserHistory_ServiceDesc is the grpc.ServiceDesc for UserHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserHistory",
	HandlerType: (*UserHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserHistory_Add_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _UserHistory_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/grpcService/proto/base.proto",
}