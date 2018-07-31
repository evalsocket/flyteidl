// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/admin.proto

package lyft_flyte_flyteadmin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import admin "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error)
	ListTaskIds(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error)
	ListTasks(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.TaskList, error)
	CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error)
	ListWorkflowIds(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error)
	ListWorkflows(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.WorkflowList, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error) {
	out := new(admin.TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListTaskIds(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error) {
	out := new(admin.IdentifierList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListTaskIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListTasks(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.TaskList, error) {
	out := new(admin.TaskList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error) {
	out := new(admin.WorkflowCreateResponse)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListWorkflowIds(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error) {
	out := new(admin.IdentifierList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListWorkflowIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListWorkflows(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.WorkflowList, error) {
	out := new(admin.WorkflowList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	CreateTask(context.Context, *admin.TaskCreateRequest) (*admin.TaskCreateResponse, error)
	ListTaskIds(context.Context, *admin.TaskListRequest) (*admin.IdentifierList, error)
	ListTasks(context.Context, *admin.TaskListRequest) (*admin.TaskList, error)
	CreateWorkflow(context.Context, *admin.WorkflowCreateRequest) (*admin.WorkflowCreateResponse, error)
	ListWorkflowIds(context.Context, *admin.WorkflowListRequest) (*admin.IdentifierList, error)
	ListWorkflows(context.Context, *admin.WorkflowListRequest) (*admin.WorkflowList, error)
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateTask(ctx, req.(*admin.TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListTaskIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListTaskIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListTaskIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListTaskIds(ctx, req.(*admin.TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListTasks(ctx, req.(*admin.TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, req.(*admin.WorkflowCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListWorkflowIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListWorkflowIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListWorkflowIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListWorkflowIds(ctx, req.(*admin.WorkflowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListWorkflows(ctx, req.(*admin.WorkflowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lyft.flyte.flyteadmin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _AdminService_CreateTask_Handler,
		},
		{
			MethodName: "ListTaskIds",
			Handler:    _AdminService_ListTaskIds_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _AdminService_ListTasks_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _AdminService_CreateWorkflow_Handler,
		},
		{
			MethodName: "ListWorkflowIds",
			Handler:    _AdminService_ListWorkflowIds_Handler,
		},
		{
			MethodName: "ListWorkflows",
			Handler:    _AdminService_ListWorkflows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/admin.proto",
}

func init() { proto.RegisterFile("flyteidl/service/admin.proto", fileDescriptor_admin_8b3022176527d448) }

var fileDescriptor_admin_8b3022176527d448 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0xc9, 0x5d, 0xdc, 0xcb, 0x1d, 0xad, 0xe2, 0x80, 0x52, 0x6b, 0x54, 0x8c, 0xe8, 0xc2,
	0x45, 0x06, 0x5b, 0x11, 0x29, 0x22, 0x88, 0xab, 0x82, 0x2b, 0x15, 0x84, 0x6e, 0x64, 0x6c, 0x26,
	0x65, 0x4c, 0x32, 0x93, 0x66, 0xc6, 0x96, 0x52, 0xb2, 0x71, 0xe9, 0x56, 0xf0, 0x31, 0xdc, 0xfb,
	0x10, 0xae, 0x7c, 0x05, 0x1f, 0x44, 0x32, 0x99, 0x11, 0x0d, 0xb6, 0x36, 0x6e, 0xda, 0x30, 0xff,
	0x7f, 0xce, 0xf9, 0xce, 0xfc, 0x09, 0xb0, 0xfd, 0x70, 0x28, 0x09, 0xf5, 0x42, 0x24, 0x48, 0xd2,
	0xa7, 0x1d, 0x82, 0xb0, 0x17, 0x51, 0xe6, 0xc6, 0x09, 0x97, 0x1c, 0x2e, 0x86, 0x43, 0x5f, 0xba,
	0xca, 0x92, 0xff, 0x2a, 0xb1, 0x66, 0x77, 0x39, 0xef, 0x86, 0x04, 0xe1, 0x98, 0x22, 0xcc, 0x18,
	0x97, 0x58, 0x52, 0xce, 0x44, 0x5e, 0x54, 0x5b, 0xfe, 0x68, 0xa9, 0xdc, 0x48, 0x62, 0x11, 0x68,
	0x69, 0xb5, 0x20, 0x0d, 0x78, 0x12, 0xf8, 0x21, 0x1f, 0x68, 0x79, 0xa5, 0x20, 0x77, 0x78, 0x14,
	0x71, 0xcd, 0x52, 0x7f, 0xfa, 0x07, 0x66, 0x8f, 0xb3, 0xe3, 0xf3, 0x1c, 0x14, 0x46, 0x00, 0x9c,
	0x24, 0x04, 0x4b, 0x72, 0x81, 0x45, 0x00, 0x37, 0x5c, 0x53, 0xec, 0xe6, 0x1b, 0x64, 0xa7, 0xb9,
	0x7e, 0x46, 0x7a, 0xb7, 0x44, 0xc8, 0x9a, 0x33, 0xc9, 0x22, 0x62, 0xce, 0x04, 0x71, 0xaa, 0x77,
	0xaf, 0x6f, 0x0f, 0x7f, 0xa0, 0x53, 0x51, 0xcb, 0xf5, 0x77, 0x15, 0xbd, 0x68, 0x5a, 0x3b, 0xf0,
	0xde, 0x02, 0x33, 0xa7, 0x54, 0xc8, 0xac, 0xa8, 0xe5, 0x09, 0xb8, 0xfe, 0x5d, 0xb7, 0xcc, 0x60,
	0xc6, 0xad, 0x15, 0x0d, 0x2d, 0x8f, 0x30, 0x49, 0x7d, 0x4a, 0x92, 0xcc, 0xe6, 0x34, 0xd5, 0xa8,
	0x3d, 0x58, 0x37, 0xa3, 0xe2, 0x84, 0xdf, 0x90, 0x8e, 0x44, 0x23, 0xfd, 0x90, 0x22, 0x8f, 0x47,
	0x98, 0x32, 0x34, 0xca, 0xff, 0x53, 0x05, 0x73, 0x45, 0x3d, 0x01, 0x9f, 0x2d, 0xf0, 0xdf, 0xc0,
	0x4c, 0x81, 0x52, 0x1d, 0x67, 0x70, 0x7a, 0x0a, 0x22, 0x80, 0xa8, 0x1c, 0x84, 0x68, 0x1f, 0xc0,
	0xfd, 0x92, 0x25, 0x68, 0xc4, 0x70, 0x44, 0x52, 0x98, 0x82, 0xb9, 0xfc, 0xd2, 0x2f, 0x75, 0xfa,
	0x70, 0xab, 0x88, 0x67, 0x94, 0xaf, 0xf9, 0x6d, 0xff, 0x64, 0xd3, 0x19, 0xda, 0x6a, 0xa7, 0x25,
	0x67, 0xc1, 0x00, 0x9a, 0xd7, 0x4c, 0xe5, 0xf8, 0x68, 0x81, 0xf9, 0x6c, 0x75, 0x53, 0x9c, 0x65,
	0xb9, 0x39, 0xae, 0x73, 0x99, 0x3c, 0x8f, 0xd4, 0xd8, 0x52, 0xf7, 0x62, 0xc0, 0x54, 0xa6, 0x2f,
	0x16, 0xa8, 0x7c, 0x06, 0x9b, 0x12, 0xcb, 0x9e, 0x64, 0x72, 0x86, 0x0a, 0x4a, 0xc0, 0x46, 0x79,
	0x28, 0xd1, 0x3e, 0x84, 0xcd, 0x5f, 0x94, 0xe9, 0x9c, 0xaf, 0xff, 0xaa, 0xef, 0xb6, 0xf1, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0xb3, 0x32, 0x0f, 0x63, 0x04, 0x00, 0x00,
}