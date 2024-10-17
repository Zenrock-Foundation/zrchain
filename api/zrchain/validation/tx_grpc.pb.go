// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package validation

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateValidator defines a method for creating a new validator.
	CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error)
	// EditValidator defines a method for editing an existing validator.
	EditValidator(ctx context.Context, in *MsgEditValidator, opts ...grpc.CallOption) (*MsgEditValidatorResponse, error)
	// Delegate defines a method for performing a delegation of coins
	// from a delegator to a validator.
	Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error)
	// BeginRedelegate defines a method for performing a redelegation
	// of coins from a delegator and source validator to a destination validator.
	BeginRedelegate(ctx context.Context, in *MsgBeginRedelegate, opts ...grpc.CallOption) (*MsgBeginRedelegateResponse, error)
	// Undelegate defines a method for performing an undelegation from a
	// delegate and a validator.
	Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error)
	// Since: cosmos-sdk 0.46
	CancelUnbondingDelegation(ctx context.Context, in *MsgCancelUnbondingDelegation, opts ...grpc.CallOption) (*MsgCancelUnbondingDelegationResponse, error)
	// UpdateParams defines an operation for updating the x/staking module
	// parameters.
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error) {
	out := new(MsgCreateValidatorResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/CreateValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditValidator(ctx context.Context, in *MsgEditValidator, opts ...grpc.CallOption) (*MsgEditValidatorResponse, error) {
	out := new(MsgEditValidatorResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/EditValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error) {
	out := new(MsgDelegateResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/Delegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BeginRedelegate(ctx context.Context, in *MsgBeginRedelegate, opts ...grpc.CallOption) (*MsgBeginRedelegateResponse, error) {
	out := new(MsgBeginRedelegateResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/BeginRedelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error) {
	out := new(MsgUndelegateResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/Undelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUnbondingDelegation(ctx context.Context, in *MsgCancelUnbondingDelegation, opts ...grpc.CallOption) (*MsgCancelUnbondingDelegationResponse, error) {
	out := new(MsgCancelUnbondingDelegationResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/CancelUnbondingDelegation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/zrchain.validation.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateValidator defines a method for creating a new validator.
	CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error)
	// EditValidator defines a method for editing an existing validator.
	EditValidator(context.Context, *MsgEditValidator) (*MsgEditValidatorResponse, error)
	// Delegate defines a method for performing a delegation of coins
	// from a delegator to a validator.
	Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error)
	// BeginRedelegate defines a method for performing a redelegation
	// of coins from a delegator and source validator to a destination validator.
	BeginRedelegate(context.Context, *MsgBeginRedelegate) (*MsgBeginRedelegateResponse, error)
	// Undelegate defines a method for performing an undelegation from a
	// delegate and a validator.
	Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error)
	// Since: cosmos-sdk 0.46
	CancelUnbondingDelegation(context.Context, *MsgCancelUnbondingDelegation) (*MsgCancelUnbondingDelegationResponse, error)
	// UpdateParams defines an operation for updating the x/staking module
	// parameters.
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateValidator not implemented")
}
func (UnimplementedMsgServer) EditValidator(context.Context, *MsgEditValidator) (*MsgEditValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditValidator not implemented")
}
func (UnimplementedMsgServer) Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delegate not implemented")
}
func (UnimplementedMsgServer) BeginRedelegate(context.Context, *MsgBeginRedelegate) (*MsgBeginRedelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginRedelegate not implemented")
}
func (UnimplementedMsgServer) Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Undelegate not implemented")
}
func (UnimplementedMsgServer) CancelUnbondingDelegation(context.Context, *MsgCancelUnbondingDelegation) (*MsgCancelUnbondingDelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUnbondingDelegation not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
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

func _Msg_CreateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.validation.Msg/CreateValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateValidator(ctx, req.(*MsgCreateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.validation.Msg/EditValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditValidator(ctx, req.(*MsgEditValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Delegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.validation.Msg/Delegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delegate(ctx, req.(*MsgDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BeginRedelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBeginRedelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BeginRedelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.validation.Msg/BeginRedelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BeginRedelegate(ctx, req.(*MsgBeginRedelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Undelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Undelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.validation.Msg/Undelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Undelegate(ctx, req.(*MsgUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUnbondingDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUnbondingDelegation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUnbondingDelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.validation.Msg/CancelUnbondingDelegation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUnbondingDelegation(ctx, req.(*MsgCancelUnbondingDelegation))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/zrchain.validation.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zrchain.validation.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateValidator",
			Handler:    _Msg_CreateValidator_Handler,
		},
		{
			MethodName: "EditValidator",
			Handler:    _Msg_EditValidator_Handler,
		},
		{
			MethodName: "Delegate",
			Handler:    _Msg_Delegate_Handler,
		},
		{
			MethodName: "BeginRedelegate",
			Handler:    _Msg_BeginRedelegate_Handler,
		},
		{
			MethodName: "Undelegate",
			Handler:    _Msg_Undelegate_Handler,
		},
		{
			MethodName: "CancelUnbondingDelegation",
			Handler:    _Msg_CancelUnbondingDelegation_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zrchain/validation/tx.proto",
}
