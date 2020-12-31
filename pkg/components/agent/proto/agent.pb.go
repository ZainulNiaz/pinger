// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.13.0
// source: agent.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2c,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x69, 0x6c, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09,
	0x50, 0x75, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_agent_proto_goTypes = []interface{}{
	(*Nil)(nil),          // 0: proto.Nil
	(*Check)(nil),        // 1: proto.Check
	(*CheckID)(nil),      // 2: proto.CheckID
	(*CheckList)(nil),    // 3: proto.CheckList
	(*BoolResponse)(nil), // 4: proto.BoolResponse
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: proto.Agent.ListChecks:input_type -> proto.Nil
	1, // 1: proto.Agent.PushCheck:input_type -> proto.Check
	2, // 2: proto.Agent.RemoveCheck:input_type -> proto.CheckID
	3, // 3: proto.Agent.ListChecks:output_type -> proto.CheckList
	4, // 4: proto.Agent.PushCheck:output_type -> proto.BoolResponse
	4, // 5: proto.Agent.RemoveCheck:output_type -> proto.BoolResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	file_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
type AgentServer interface {
	// ListChecks fetches a list of checks registered.
	ListChecks(context.Context, *Nil) (*CheckList, error)
	// PushCheck creates a new check. If the check already exists it simply
	// updates the check.
	PushCheck(context.Context, *Check) (*BoolResponse, error)
	// RemoveCheck removes the check.
	RemoveCheck(context.Context, *CheckID) (*BoolResponse, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) ListChecks(context.Context, *Nil) (*CheckList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChecks not implemented")
}
func (*UnimplementedAgentServer) PushCheck(context.Context, *Check) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushCheck not implemented")
}
func (*UnimplementedAgentServer) RemoveCheck(context.Context, *CheckID) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCheck not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
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