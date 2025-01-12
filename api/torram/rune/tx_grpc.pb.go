// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: torram/rune/tx.proto

package rune

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

const (
	Msg_UpdateParams_FullMethodName     = "/torram.rune.Msg/UpdateParams"
	Msg_StakeRune_FullMethodName        = "/torram.rune.Msg/StakeRune"
	Msg_UnstakeRune_FullMethodName      = "/torram.rune.Msg/UnstakeRune"
	Msg_UpdateRuneStatus_FullMethodName = "/torram.rune.Msg/UpdateRuneStatus"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	StakeRune(ctx context.Context, in *MsgStakeRune, opts ...grpc.CallOption) (*MsgStakeRuneResponse, error)
	UnstakeRune(ctx context.Context, in *MsgUnstakeRune, opts ...grpc.CallOption) (*MsgUnstakeRuneResponse, error)
	UpdateRuneStatus(ctx context.Context, in *MsgUpdateRuneStatus, opts ...grpc.CallOption) (*MsgUpdateRuneStatusResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StakeRune(ctx context.Context, in *MsgStakeRune, opts ...grpc.CallOption) (*MsgStakeRuneResponse, error) {
	out := new(MsgStakeRuneResponse)
	err := c.cc.Invoke(ctx, Msg_StakeRune_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnstakeRune(ctx context.Context, in *MsgUnstakeRune, opts ...grpc.CallOption) (*MsgUnstakeRuneResponse, error) {
	out := new(MsgUnstakeRuneResponse)
	err := c.cc.Invoke(ctx, Msg_UnstakeRune_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRuneStatus(ctx context.Context, in *MsgUpdateRuneStatus, opts ...grpc.CallOption) (*MsgUpdateRuneStatusResponse, error) {
	out := new(MsgUpdateRuneStatusResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRuneStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	StakeRune(context.Context, *MsgStakeRune) (*MsgStakeRuneResponse, error)
	UnstakeRune(context.Context, *MsgUnstakeRune) (*MsgUnstakeRuneResponse, error)
	UpdateRuneStatus(context.Context, *MsgUpdateRuneStatus) (*MsgUpdateRuneStatusResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) StakeRune(context.Context, *MsgStakeRune) (*MsgStakeRuneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakeRune not implemented")
}
func (UnimplementedMsgServer) UnstakeRune(context.Context, *MsgUnstakeRune) (*MsgUnstakeRuneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnstakeRune not implemented")
}
func (UnimplementedMsgServer) UpdateRuneStatus(context.Context, *MsgUpdateRuneStatus) (*MsgUpdateRuneStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRuneStatus not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StakeRune_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStakeRune)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StakeRune(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StakeRune_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StakeRune(ctx, req.(*MsgStakeRune))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnstakeRune_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnstakeRune)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnstakeRune(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UnstakeRune_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnstakeRune(ctx, req.(*MsgUnstakeRune))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRuneStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRuneStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRuneStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRuneStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRuneStatus(ctx, req.(*MsgUpdateRuneStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "torram.rune.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "StakeRune",
			Handler:    _Msg_StakeRune_Handler,
		},
		{
			MethodName: "UnstakeRune",
			Handler:    _Msg_UnstakeRune_Handler,
		},
		{
			MethodName: "UpdateRuneStatus",
			Handler:    _Msg_UpdateRuneStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "torram/rune/tx.proto",
}
