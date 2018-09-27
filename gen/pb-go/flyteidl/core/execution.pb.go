// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/execution.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Indicates various phases of Workflow Execution
type WorkflowExecutionPhase int32

const (
	WorkflowExecutionPhase_WORKFLOW_PHASE_UNDEFINED  WorkflowExecutionPhase = 0
	WorkflowExecutionPhase_WORKFLOW_PHASE_RUNNING    WorkflowExecutionPhase = 1
	WorkflowExecutionPhase_WORKFLOW_PHASE_SUCCEEDING WorkflowExecutionPhase = 2
	WorkflowExecutionPhase_WORKFLOW_PHASE_SUCCEEDED  WorkflowExecutionPhase = 3
	WorkflowExecutionPhase_WORKFLOW_PHASE_FAILING    WorkflowExecutionPhase = 4
	WorkflowExecutionPhase_WORKFLOW_PHASE_FAILED     WorkflowExecutionPhase = 5
	WorkflowExecutionPhase_WORKFLOW_PHASE_TIMED_OUT  WorkflowExecutionPhase = 6
	WorkflowExecutionPhase_WORKFLOW_PHASE_ABORTED    WorkflowExecutionPhase = 7
	WorkflowExecutionPhase_WORKFLOW_PHASE_QUEUED     WorkflowExecutionPhase = 8
)

var WorkflowExecutionPhase_name = map[int32]string{
	0: "WORKFLOW_PHASE_UNDEFINED",
	1: "WORKFLOW_PHASE_RUNNING",
	2: "WORKFLOW_PHASE_SUCCEEDING",
	3: "WORKFLOW_PHASE_SUCCEEDED",
	4: "WORKFLOW_PHASE_FAILING",
	5: "WORKFLOW_PHASE_FAILED",
	6: "WORKFLOW_PHASE_TIMED_OUT",
	7: "WORKFLOW_PHASE_ABORTED",
	8: "WORKFLOW_PHASE_QUEUED",
}
var WorkflowExecutionPhase_value = map[string]int32{
	"WORKFLOW_PHASE_UNDEFINED":  0,
	"WORKFLOW_PHASE_RUNNING":    1,
	"WORKFLOW_PHASE_SUCCEEDING": 2,
	"WORKFLOW_PHASE_SUCCEEDED":  3,
	"WORKFLOW_PHASE_FAILING":    4,
	"WORKFLOW_PHASE_FAILED":     5,
	"WORKFLOW_PHASE_TIMED_OUT":  6,
	"WORKFLOW_PHASE_ABORTED":    7,
	"WORKFLOW_PHASE_QUEUED":     8,
}

func (x WorkflowExecutionPhase) String() string {
	return proto.EnumName(WorkflowExecutionPhase_name, int32(x))
}
func (WorkflowExecutionPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_7c221429e816b782, []int{0}
}

// Indicates various phases of Node Execution
type NodeExecutionPhase int32

const (
	NodeExecutionPhase_NODE_PHASE_UNDEFINED NodeExecutionPhase = 0
	NodeExecutionPhase_NODE_PHASE_RUNNING   NodeExecutionPhase = 1
	NodeExecutionPhase_NODE_PHASE_SUCCEEDED NodeExecutionPhase = 2
	NodeExecutionPhase_NODE_PHASE_FAILING   NodeExecutionPhase = 3
	NodeExecutionPhase_NODE_PHASE_FAILED    NodeExecutionPhase = 4
	NodeExecutionPhase_NODE_PHASE_TIMED_OUT NodeExecutionPhase = 5
	NodeExecutionPhase_NODE_PHASE_SKIPPED   NodeExecutionPhase = 6
	NodeExecutionPhase_NODE_PHASE_ABORTED   NodeExecutionPhase = 7
	NodeExecutionPhase_NODE_PHASE_QUEUED    NodeExecutionPhase = 8
)

var NodeExecutionPhase_name = map[int32]string{
	0: "NODE_PHASE_UNDEFINED",
	1: "NODE_PHASE_RUNNING",
	2: "NODE_PHASE_SUCCEEDED",
	3: "NODE_PHASE_FAILING",
	4: "NODE_PHASE_FAILED",
	5: "NODE_PHASE_TIMED_OUT",
	6: "NODE_PHASE_SKIPPED",
	7: "NODE_PHASE_ABORTED",
	8: "NODE_PHASE_QUEUED",
}
var NodeExecutionPhase_value = map[string]int32{
	"NODE_PHASE_UNDEFINED": 0,
	"NODE_PHASE_RUNNING":   1,
	"NODE_PHASE_SUCCEEDED": 2,
	"NODE_PHASE_FAILING":   3,
	"NODE_PHASE_FAILED":    4,
	"NODE_PHASE_TIMED_OUT": 5,
	"NODE_PHASE_SKIPPED":   6,
	"NODE_PHASE_ABORTED":   7,
	"NODE_PHASE_QUEUED":    8,
}

func (x NodeExecutionPhase) String() string {
	return proto.EnumName(NodeExecutionPhase_name, int32(x))
}
func (NodeExecutionPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_7c221429e816b782, []int{1}
}

// Represents the error message from the execution.
type ExecutionError struct {
	// Error code indicates a grouping of a type of error.
	// More Info: <Link>
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Detailed description of the error - including stack trace.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionError) Reset()         { *m = ExecutionError{} }
func (m *ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionError) ProtoMessage()    {}
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_7c221429e816b782, []int{0}
}
func (m *ExecutionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionError.Unmarshal(m, b)
}
func (m *ExecutionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionError.Marshal(b, m, deterministic)
}
func (dst *ExecutionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionError.Merge(dst, src)
}
func (m *ExecutionError) XXX_Size() int {
	return xxx_messageInfo_ExecutionError.Size(m)
}
func (m *ExecutionError) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionError.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionError proto.InternalMessageInfo

func (m *ExecutionError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ExecutionError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecutionError)(nil), "flyteidl.core.ExecutionError")
	proto.RegisterEnum("flyteidl.core.WorkflowExecutionPhase", WorkflowExecutionPhase_name, WorkflowExecutionPhase_value)
	proto.RegisterEnum("flyteidl.core.NodeExecutionPhase", NodeExecutionPhase_name, NodeExecutionPhase_value)
}

func init() {
	proto.RegisterFile("flyteidl/core/execution.proto", fileDescriptor_execution_7c221429e816b782)
}

var fileDescriptor_execution_7c221429e816b782 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0x87, 0xdd, 0xf8, 0xa7, 0x4d, 0x34, 0xb5, 0x11, 0x32, 0x8c, 0x24, 0xc6, 0x2b, 0x43, 0x22,
	0x33, 0x7a, 0x6f, 0x02, 0xf4, 0xa0, 0x0b, 0xd8, 0xcd, 0xc1, 0x42, 0xe2, 0x0d, 0x81, 0x51, 0x06,
	0x71, 0x50, 0x32, 0x46, 0x94, 0x67, 0xf1, 0x25, 0x7d, 0x04, 0x03, 0x71, 0x4c, 0x9a, 0xde, 0xb5,
	0xe7, 0xeb, 0xf9, 0x9a, 0xdf, 0xc9, 0x41, 0x95, 0x49, 0xb8, 0x89, 0xf9, 0x6c, 0x1c, 0x9a, 0xbe,
	0x88, 0xb8, 0xc9, 0xbf, 0xb8, 0xbf, 0x8e, 0x67, 0x62, 0x51, 0x5b, 0x46, 0x22, 0x16, 0xe4, 0x34,
	0xc1, 0xb5, 0x2d, 0xbe, 0x79, 0x42, 0x67, 0x90, 0xbc, 0x80, 0x28, 0x12, 0x11, 0x21, 0x28, 0xeb,
	0x8b, 0x31, 0x37, 0xb4, 0x6b, 0xed, 0xf6, 0xc4, 0xdd, 0x9d, 0x89, 0x81, 0x0a, 0x73, 0xbe, 0x5a,
	0x0d, 0x03, 0x6e, 0xe8, 0xbb, 0x72, 0x72, 0xad, 0x7e, 0xeb, 0xa8, 0xd4, 0x17, 0xd1, 0xc7, 0x24,
	0x14, 0x9f, 0x7b, 0x91, 0x33, 0x1d, 0xae, 0x38, 0xb9, 0x42, 0x46, 0xdf, 0x76, 0xdb, 0xad, 0x8e,
	0xdd, 0x1f, 0x38, 0x2f, 0xf5, 0x2e, 0x0c, 0x3c, 0x46, 0xa1, 0x65, 0x31, 0xa0, 0xf8, 0x88, 0x5c,
	0xa2, 0x92, 0x44, 0x5d, 0x8f, 0x31, 0x8b, 0x3d, 0x63, 0x8d, 0x54, 0x50, 0x59, 0x62, 0x5d, 0xaf,
	0xd9, 0x04, 0xa0, 0x5b, 0xac, 0x2b, 0xc4, 0x7f, 0x18, 0x28, 0xce, 0x28, 0xc4, 0xad, 0xba, 0xd5,
	0xd9, 0x76, 0x66, 0x49, 0x19, 0x15, 0x15, 0x0c, 0x28, 0xce, 0x29, 0xa4, 0x3d, 0xeb, 0x15, 0xe8,
	0xc0, 0xf6, 0x7a, 0x38, 0xaf, 0x90, 0xd6, 0x1b, 0xb6, 0xdb, 0x03, 0x8a, 0x0b, 0x0a, 0xe9, 0x9b,
	0x07, 0x1e, 0x50, 0x7c, 0x5c, 0xfd, 0xd1, 0x10, 0x61, 0x62, 0xcc, 0xa5, 0xc9, 0x18, 0xe8, 0x82,
	0xd9, 0x14, 0x14, 0x53, 0x29, 0x21, 0xf2, 0x8f, 0xa4, 0x13, 0x39, 0xec, 0x48, 0xe3, 0xea, 0x52,
	0x47, 0x12, 0x35, 0x43, 0x8a, 0xe8, 0x5c, 0xaa, 0x03, 0xc5, 0x59, 0x49, 0x94, 0x46, 0xcc, 0x49,
	0xa2, 0x6e, 0xdb, 0x72, 0x1c, 0xa0, 0x38, 0x2f, 0xd5, 0xd3, 0xd8, 0x87, 0x1f, 0x24, 0x91, 0x1b,
	0x0f, 0xef, 0xf7, 0xc1, 0x2c, 0x9e, 0xae, 0x47, 0x35, 0x5f, 0xcc, 0xcd, 0x70, 0x33, 0x89, 0xcd,
	0xfd, 0x42, 0x06, 0x7c, 0x61, 0x2e, 0x47, 0x77, 0x81, 0x30, 0x0f, 0x76, 0x74, 0x94, 0xdf, 0xad,
	0xe6, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4d, 0xb7, 0xfc, 0xbb, 0x02, 0x00, 0x00,
}