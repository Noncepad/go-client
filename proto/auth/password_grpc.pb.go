// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: password.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	base "github.com/noncepad/client/proto/base"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// a new customer registers
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*base.Empty, error)
	// as part of registration, a user creates a 2FA token (Google Authenticator)
	Create2FALocal(ctx context.Context, in *SecondFactorLocalRequest, opts ...grpc.CallOption) (*SecondFactorLocalResponse, error)
	SendSmsCode(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*OneTimeCodeResponse, error)
	VerifySmsCode(ctx context.Context, in *VerifySMSRequest, opts ...grpc.CallOption) (*base.Empty, error)
	SendEmailCode(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*OneTimeCodeResponse, error)
	VerifyEmailCode(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*base.Empty, error)
	// customer logs in
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// customer submits cookie
	ParseJWT(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Create2FALocal(ctx context.Context, in *SecondFactorLocalRequest, opts ...grpc.CallOption) (*SecondFactorLocalResponse, error) {
	out := new(SecondFactorLocalResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Create2FALocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendSmsCode(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*OneTimeCodeResponse, error) {
	out := new(OneTimeCodeResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SendSmsCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifySmsCode(ctx context.Context, in *VerifySMSRequest, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/VerifySmsCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendEmailCode(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*OneTimeCodeResponse, error) {
	out := new(OneTimeCodeResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SendEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyEmailCode(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/VerifyEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ParseJWT(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/ParseJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// a new customer registers
	Register(context.Context, *RegisterRequest) (*base.Empty, error)
	// as part of registration, a user creates a 2FA token (Google Authenticator)
	Create2FALocal(context.Context, *SecondFactorLocalRequest) (*SecondFactorLocalResponse, error)
	SendSmsCode(context.Context, *SendSMSRequest) (*OneTimeCodeResponse, error)
	VerifySmsCode(context.Context, *VerifySMSRequest) (*base.Empty, error)
	SendEmailCode(context.Context, *SendEmailRequest) (*OneTimeCodeResponse, error)
	VerifyEmailCode(context.Context, *VerifyEmailRequest) (*base.Empty, error)
	// customer logs in
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// customer submits cookie
	ParseJWT(context.Context, *ParseRequest) (*ParseResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Register(context.Context, *RegisterRequest) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Create2FALocal(context.Context, *SecondFactorLocalRequest) (*SecondFactorLocalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create2FALocal not implemented")
}
func (UnimplementedAuthServer) SendSmsCode(context.Context, *SendSMSRequest) (*OneTimeCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsCode not implemented")
}
func (UnimplementedAuthServer) VerifySmsCode(context.Context, *VerifySMSRequest) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySmsCode not implemented")
}
func (UnimplementedAuthServer) SendEmailCode(context.Context, *SendEmailRequest) (*OneTimeCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedAuthServer) VerifyEmailCode(context.Context, *VerifyEmailRequest) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmailCode not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) ParseJWT(context.Context, *ParseRequest) (*ParseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseJWT not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Create2FALocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecondFactorLocalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Create2FALocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Create2FALocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Create2FALocal(ctx, req.(*SecondFactorLocalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendSmsCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendSmsCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SendSmsCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendSmsCode(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifySmsCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifySmsCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/VerifySmsCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifySmsCode(ctx, req.(*VerifySMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SendEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendEmailCode(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/VerifyEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyEmailCode(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ParseJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ParseJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ParseJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ParseJWT(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Create2FALocal",
			Handler:    _Auth_Create2FALocal_Handler,
		},
		{
			MethodName: "SendSmsCode",
			Handler:    _Auth_SendSmsCode_Handler,
		},
		{
			MethodName: "VerifySmsCode",
			Handler:    _Auth_VerifySmsCode_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _Auth_SendEmailCode_Handler,
		},
		{
			MethodName: "VerifyEmailCode",
			Handler:    _Auth_VerifyEmailCode_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "ParseJWT",
			Handler:    _Auth_ParseJWT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "password.proto",
}
