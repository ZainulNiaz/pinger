// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	// ListChecks fetches a list of checks registered.
	ListChecks(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*CheckList, error)
	// PushCheck creates a new check. If the check already exists it simply
	// updates the check.
	PushCheck(ctx context.Context, in *Check, opts ...grpc.CallOption) (*BoolResponse, error)
	// RemoveCheck removes the check.
	RemoveCheck(ctx context.Context, in *CheckID, opts ...grpc.CallOption) (*BoolResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) ListChecks(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*CheckList, error) {
	out := new(CheckList)
	err := c.cc.Invoke(ctx, "/proto.Agent/ListChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PushCheck(ctx context.Context, in *Check, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/proto.Agent/PushCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RemoveCheck(ctx context.Context, in *CheckID, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/proto.Agent/RemoveCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	// ListChecks fetches a list of checks registered.
	ListChecks(context.Context, *Nil) (*CheckList, error)
	// PushCheck creates a new check. If the check already exists it simply
	// updates the check.
	PushCheck(context.Context, *Check) (*BoolResponse, error)
	// RemoveCheck removes the check.
	RemoveCheck(context.Context, *CheckID) (*BoolResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) ListChecks(context.Context, *Nil) (*CheckList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChecks not implemented")
}
func (UnimplementedAgentServer) PushCheck(context.Context, *Check) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushCheck not implemented")
}
func (UnimplementedAgentServer) RemoveCheck(context.Context, *CheckID) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCheck not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_ListChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ListChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/ListChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ListChecks(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PushCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Check)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).PushCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/PushCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PushCheck(ctx, req.(*Check))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RemoveCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RemoveCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/RemoveCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RemoveCheck(ctx, req.(*CheckID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListChecks",
			Handler:    _Agent_ListChecks_Handler,
		},
		{
			MethodName: "PushCheck",
			Handler:    _Agent_PushCheck_Handler,
		},
		{
			MethodName: "RemoveCheck",
			Handler:    _Agent_RemoveCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}