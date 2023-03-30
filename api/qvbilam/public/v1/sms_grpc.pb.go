// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: sms.proto

package publicV1

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

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsClient interface {
	SendLogin(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckLogin(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendLogout(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckLogout(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendPassword(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckPassword(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Send(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Check(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type smsClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsClient(cc grpc.ClientConnInterface) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) SendLogin(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/SendLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) CheckLogin(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/CheckLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) SendLogout(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/SendLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) CheckLogout(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/CheckLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) SendPassword(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/SendPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) CheckPassword(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) Send(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) Check(ctx context.Context, in *CheckSmsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Sms/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
// All implementations must embed UnimplementedSmsServer
// for forward compatibility
type SmsServer interface {
	SendLogin(context.Context, *SendSmsRequest) (*emptypb.Empty, error)
	CheckLogin(context.Context, *CheckSmsRequest) (*emptypb.Empty, error)
	SendLogout(context.Context, *SendSmsRequest) (*emptypb.Empty, error)
	CheckLogout(context.Context, *CheckSmsRequest) (*emptypb.Empty, error)
	SendPassword(context.Context, *SendSmsRequest) (*emptypb.Empty, error)
	CheckPassword(context.Context, *CheckSmsRequest) (*emptypb.Empty, error)
	Send(context.Context, *SendSmsRequest) (*emptypb.Empty, error)
	Check(context.Context, *CheckSmsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSmsServer()
}

// UnimplementedSmsServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (UnimplementedSmsServer) SendLogin(context.Context, *SendSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogin not implemented")
}
func (UnimplementedSmsServer) CheckLogin(context.Context, *CheckSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogin not implemented")
}
func (UnimplementedSmsServer) SendLogout(context.Context, *SendSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogout not implemented")
}
func (UnimplementedSmsServer) CheckLogout(context.Context, *CheckSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogout not implemented")
}
func (UnimplementedSmsServer) SendPassword(context.Context, *SendSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPassword not implemented")
}
func (UnimplementedSmsServer) CheckPassword(context.Context, *CheckSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}
func (UnimplementedSmsServer) Send(context.Context, *SendSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSmsServer) Check(context.Context, *CheckSmsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedSmsServer) mustEmbedUnimplementedSmsServer() {}

// UnsafeSmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServer will
// result in compilation errors.
type UnsafeSmsServer interface {
	mustEmbedUnimplementedSmsServer()
}

func RegisterSmsServer(s grpc.ServiceRegistrar, srv SmsServer) {
	s.RegisterService(&Sms_ServiceDesc, srv)
}

func _Sms_SendLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/SendLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendLogin(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).CheckLogin(ctx, req.(*CheckSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_SendLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/SendLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendLogout(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_CheckLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).CheckLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/CheckLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).CheckLogout(ctx, req.(*CheckSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_SendPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/SendPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendPassword(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).CheckPassword(ctx, req.(*CheckSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).Send(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Sms/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).Check(ctx, req.(*CheckSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sms_ServiceDesc is the grpc.ServiceDesc for Sms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "publicPb.v1.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogin",
			Handler:    _Sms_SendLogin_Handler,
		},
		{
			MethodName: "CheckLogin",
			Handler:    _Sms_CheckLogin_Handler,
		},
		{
			MethodName: "SendLogout",
			Handler:    _Sms_SendLogout_Handler,
		},
		{
			MethodName: "CheckLogout",
			Handler:    _Sms_CheckLogout_Handler,
		},
		{
			MethodName: "SendPassword",
			Handler:    _Sms_SendPassword_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _Sms_CheckPassword_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Sms_Send_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Sms_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}
