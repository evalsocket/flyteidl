// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/common.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Indicates various states of an Execution - both task and workflow
type ExecutionPhase int32

const (
	ExecutionPhase_UNDEFINED ExecutionPhase = 0
	ExecutionPhase_RUNNING   ExecutionPhase = 1
	ExecutionPhase_SUCCEEDED ExecutionPhase = 2
	ExecutionPhase_FAILED    ExecutionPhase = 3
	ExecutionPhase_TIMED_OUT ExecutionPhase = 4
	ExecutionPhase_ABORTED   ExecutionPhase = 5
	ExecutionPhase_QUEUED    ExecutionPhase = 6
)

var ExecutionPhase_name = map[int32]string{
	0: "UNDEFINED",
	1: "RUNNING",
	2: "SUCCEEDED",
	3: "FAILED",
	4: "TIMED_OUT",
	5: "ABORTED",
	6: "QUEUED",
}
var ExecutionPhase_value = map[string]int32{
	"UNDEFINED": 0,
	"RUNNING":   1,
	"SUCCEEDED": 2,
	"FAILED":    3,
	"TIMED_OUT": 4,
	"ABORTED":   5,
	"QUEUED":    6,
}

func (x ExecutionPhase) String() string {
	return proto.EnumName(ExecutionPhase_name, int32(x))
}
func (ExecutionPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{0}
}

type Notification_Type int32

const (
	Notification_UNDEFINED  Notification_Type = 0
	Notification_EMAIL      Notification_Type = 1
	Notification_PAGER_DUTY Notification_Type = 2
	Notification_SLACK      Notification_Type = 3
)

var Notification_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "EMAIL",
	2: "PAGER_DUTY",
	3: "SLACK",
}
var Notification_Type_value = map[string]int32{
	"UNDEFINED":  0,
	"EMAIL":      1,
	"PAGER_DUTY": 2,
	"SLACK":      3,
}

func (x Notification_Type) String() string {
	return proto.EnumName(Notification_Type_name, int32(x))
}
func (Notification_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{4, 0}
}

// Encapsulation of fields that identifies a Flyte resource
type Identifier struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{0}
}
func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (dst *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(dst, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Identifier) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Identifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IdentifierList struct {
	Entities             []*Identifier `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IdentifierList) Reset()         { *m = IdentifierList{} }
func (m *IdentifierList) String() string { return proto.CompactTextString(m) }
func (*IdentifierList) ProtoMessage()    {}
func (*IdentifierList) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{1}
}
func (m *IdentifierList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentifierList.Unmarshal(m, b)
}
func (m *IdentifierList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentifierList.Marshal(b, m, deterministic)
}
func (dst *IdentifierList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifierList.Merge(dst, src)
}
func (m *IdentifierList) XXX_Size() int {
	return xxx_messageInfo_IdentifierList.Size(m)
}
func (m *IdentifierList) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifierList.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifierList proto.InternalMessageInfo

func (m *IdentifierList) GetEntities() []*Identifier {
	if m != nil {
		return m.Entities
	}
	return nil
}

// Declares an input parameter.
type Parameter struct {
	// +required Variable. Defines a name and a type to reference/compare through out the system.
	Var *core.Variable `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	// +optional
	//
	// Types that are valid to be assigned to Default:
	//	*Parameter_Value
	//	*Parameter_Required
	Default              isParameter_Default `protobuf_oneof:"default"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{2}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (dst *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(dst, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetVar() *core.Variable {
	if m != nil {
		return m.Var
	}
	return nil
}

type isParameter_Default interface {
	isParameter_Default()
}

type Parameter_Value struct {
	Value *core.Literal `protobuf:"bytes,2,opt,name=value,proto3,oneof"`
}

type Parameter_Required struct {
	Required bool `protobuf:"varint,3,opt,name=required,proto3,oneof"`
}

func (*Parameter_Value) isParameter_Default() {}

func (*Parameter_Required) isParameter_Default() {}

func (m *Parameter) GetDefault() isParameter_Default {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *Parameter) GetValue() *core.Literal {
	if x, ok := m.GetDefault().(*Parameter_Value); ok {
		return x.Value
	}
	return nil
}

func (m *Parameter) GetRequired() bool {
	if x, ok := m.GetDefault().(*Parameter_Required); ok {
		return x.Required
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Parameter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Parameter_OneofMarshaler, _Parameter_OneofUnmarshaler, _Parameter_OneofSizer, []interface{}{
		(*Parameter_Value)(nil),
		(*Parameter_Required)(nil),
	}
}

func _Parameter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Parameter)
	// default
	switch x := m.Default.(type) {
	case *Parameter_Value:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Value); err != nil {
			return err
		}
	case *Parameter_Required:
		t := uint64(0)
		if x.Required {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Parameter.Default has unexpected type %T", x)
	}
	return nil
}

func _Parameter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Parameter)
	switch tag {
	case 2: // default.value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.Literal)
		err := b.DecodeMessage(msg)
		m.Default = &Parameter_Value{msg}
		return true, err
	case 3: // default.required
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Default = &Parameter_Required{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Parameter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Parameter)
	// default
	switch x := m.Default.(type) {
	case *Parameter_Value:
		s := proto.Size(x.Value)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Parameter_Required:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A map of Parameters.
type ParameterMap struct {
	Parameters           map[string]*Parameter `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ParameterMap) Reset()         { *m = ParameterMap{} }
func (m *ParameterMap) String() string { return proto.CompactTextString(m) }
func (*ParameterMap) ProtoMessage()    {}
func (*ParameterMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{3}
}
func (m *ParameterMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterMap.Unmarshal(m, b)
}
func (m *ParameterMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterMap.Marshal(b, m, deterministic)
}
func (dst *ParameterMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterMap.Merge(dst, src)
}
func (m *ParameterMap) XXX_Size() int {
	return xxx_messageInfo_ParameterMap.Size(m)
}
func (m *ParameterMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterMap.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterMap proto.InternalMessageInfo

func (m *ParameterMap) GetParameters() map[string]*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Structure of notifications based on execution status.
type Notification struct {
	Type                 Notification_Type `protobuf:"varint,1,opt,name=type,proto3,enum=flyteidl.admin.Notification_Type" json:"type,omitempty"`
	Phases               []ExecutionPhase  `protobuf:"varint,2,rep,packed,name=phases,proto3,enum=flyteidl.admin.ExecutionPhase" json:"phases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{4}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (dst *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(dst, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetType() Notification_Type {
	if m != nil {
		return m.Type
	}
	return Notification_UNDEFINED
}

func (m *Notification) GetPhases() []ExecutionPhase {
	if m != nil {
		return m.Phases
	}
	return nil
}

type GetObjectRequest struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectRequest) Reset()         { *m = GetObjectRequest{} }
func (m *GetObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectRequest) ProtoMessage()    {}
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_cb502e21af75c35f, []int{5}
}
func (m *GetObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectRequest.Unmarshal(m, b)
}
func (m *GetObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectRequest.Marshal(b, m, deterministic)
}
func (dst *GetObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectRequest.Merge(dst, src)
}
func (m *GetObjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectRequest.Size(m)
}
func (m *GetObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectRequest proto.InternalMessageInfo

func (m *GetObjectRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func init() {
	proto.RegisterType((*Identifier)(nil), "flyteidl.admin.Identifier")
	proto.RegisterType((*IdentifierList)(nil), "flyteidl.admin.IdentifierList")
	proto.RegisterType((*Parameter)(nil), "flyteidl.admin.Parameter")
	proto.RegisterType((*ParameterMap)(nil), "flyteidl.admin.ParameterMap")
	proto.RegisterMapType((map[string]*Parameter)(nil), "flyteidl.admin.ParameterMap.ParametersEntry")
	proto.RegisterType((*Notification)(nil), "flyteidl.admin.Notification")
	proto.RegisterType((*GetObjectRequest)(nil), "flyteidl.admin.GetObjectRequest")
	proto.RegisterEnum("flyteidl.admin.ExecutionPhase", ExecutionPhase_name, ExecutionPhase_value)
	proto.RegisterEnum("flyteidl.admin.Notification_Type", Notification_Type_name, Notification_Type_value)
}

func init() { proto.RegisterFile("flyteidl/admin/common.proto", fileDescriptor_common_cb502e21af75c35f) }

var fileDescriptor_common_cb502e21af75c35f = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xa6, 0xeb, 0xd6, 0xd3, 0x51, 0x22, 0x5f, 0x8c, 0x52, 0x06, 0x1a, 0x11, 0x17,
	0x03, 0x41, 0x22, 0x3a, 0x31, 0x21, 0xb8, 0xea, 0x16, 0x6f, 0xab, 0xe8, 0xba, 0xe1, 0x35, 0x08,
	0xb8, 0x99, 0xdc, 0xf4, 0x74, 0x33, 0xe4, 0xdf, 0x1c, 0x67, 0xa2, 0x4f, 0xc1, 0xdb, 0x70, 0xc1,
	0xd3, 0x21, 0xa7, 0xff, 0xd6, 0x4a, 0xdc, 0xf9, 0xf8, 0xfb, 0x7d, 0xc7, 0xfe, 0x4e, 0x62, 0x78,
	0x32, 0x0a, 0xc7, 0x0a, 0xc5, 0x30, 0x74, 0xf9, 0x30, 0x12, 0xb1, 0x1b, 0x24, 0x51, 0x94, 0xc4,
	0x4e, 0x2a, 0x13, 0x95, 0x90, 0xfa, 0x4c, 0x74, 0x0a, 0xb1, 0xf9, 0x74, 0x0e, 0x07, 0x89, 0x44,
	0x57, 0xc4, 0x0a, 0xe5, 0x88, 0x07, 0x38, 0xc1, 0x9b, 0x3b, 0xcb, 0x72, 0x28, 0x14, 0x4a, 0x1e,
	0x66, 0x13, 0xd5, 0x66, 0x00, 0x9d, 0x21, 0xc6, 0x4a, 0x8c, 0x04, 0x4a, 0xd2, 0x80, 0x8d, 0x54,
	0x26, 0x3f, 0x30, 0x50, 0x0d, 0x63, 0xd7, 0xd8, 0xab, 0xb2, 0x59, 0x49, 0xb6, 0xa1, 0x32, 0x4c,
	0x22, 0x2e, 0xe2, 0x46, 0xa9, 0x10, 0xa6, 0x15, 0x21, 0x50, 0x8e, 0x79, 0x84, 0x0d, 0xb3, 0xd8,
	0x2d, 0xd6, 0xf6, 0x29, 0xd4, 0x17, 0x3d, 0xbb, 0x22, 0x53, 0xe4, 0x00, 0x36, 0x75, 0xad, 0x04,
	0x66, 0x0d, 0x63, 0xd7, 0xdc, 0xab, 0xb5, 0x9a, 0xce, 0x72, 0x0a, 0x67, 0xe1, 0x60, 0x73, 0xd6,
	0xfe, 0x6d, 0x40, 0xf5, 0x82, 0x4b, 0x1e, 0xa1, 0x42, 0x49, 0x5e, 0x82, 0x79, 0xc7, 0x65, 0x71,
	0xb3, 0x5a, 0xeb, 0xd1, 0xa2, 0x81, 0xce, 0xe5, 0x7c, 0xe1, 0x52, 0xf0, 0x41, 0x88, 0x4c, 0x33,
	0xc4, 0x81, 0xf5, 0x3b, 0x1e, 0xe6, 0x58, 0xdc, 0xb6, 0xd6, 0xda, 0x5e, 0x81, 0xbb, 0x93, 0x21,
	0x9c, 0xae, 0xb1, 0x09, 0x46, 0x76, 0x60, 0x53, 0xe2, 0x6d, 0x2e, 0x24, 0x0e, 0x8b, 0x28, 0x9b,
	0xa7, 0x6b, 0x6c, 0xbe, 0x73, 0x58, 0x85, 0x8d, 0x21, 0x8e, 0x78, 0x1e, 0x2a, 0xfb, 0x8f, 0x01,
	0x5b, 0xf3, 0x1b, 0x9d, 0xf1, 0x94, 0x74, 0x01, 0xd2, 0x59, 0x3d, 0x0b, 0xf7, 0x7a, 0x35, 0xdc,
	0x7d, 0xc7, 0xa2, 0xc8, 0x68, 0xac, 0xe4, 0x98, 0xdd, 0xf3, 0x37, 0xbf, 0xc2, 0xc3, 0x15, 0x99,
	0x58, 0x60, 0xfe, 0xc4, 0xf1, 0xf4, 0x7b, 0xe8, 0x25, 0x71, 0x97, 0xc3, 0x3d, 0xfe, 0xef, 0x69,
	0xd3, 0x74, 0x1f, 0x4a, 0xef, 0x0d, 0xfb, 0xaf, 0x01, 0x5b, 0xbd, 0x44, 0x4f, 0x38, 0xe0, 0x4a,
	0x24, 0x31, 0x79, 0x07, 0x65, 0x35, 0x4e, 0xb1, 0x68, 0x5c, 0x6f, 0x3d, 0x5f, 0x6d, 0x72, 0x9f,
	0x75, 0xfa, 0xe3, 0x14, 0x59, 0x81, 0x93, 0x03, 0xa8, 0xa4, 0x37, 0x3c, 0xc3, 0xac, 0x51, 0xda,
	0x35, 0xf7, 0xea, 0xad, 0x67, 0xab, 0x46, 0xfa, 0x0b, 0x83, 0x5c, 0xbb, 0x2e, 0x34, 0xc6, 0xa6,
	0xb4, 0xfd, 0x11, 0xca, 0xba, 0x0b, 0x79, 0x00, 0x55, 0xbf, 0xe7, 0xd1, 0xe3, 0x4e, 0x8f, 0x7a,
	0xd6, 0x1a, 0xa9, 0xc2, 0x3a, 0x3d, 0x6b, 0x77, 0xba, 0x96, 0x41, 0xea, 0x00, 0x17, 0xed, 0x13,
	0xca, 0xae, 0x3c, 0xbf, 0xff, 0xcd, 0x2a, 0x69, 0xe9, 0xb2, 0xdb, 0x3e, 0xfa, 0x64, 0x99, 0xf6,
	0x0b, 0xb0, 0x4e, 0x50, 0x9d, 0x0f, 0xf4, 0xaf, 0xc8, 0xf0, 0x36, 0xc7, 0x4c, 0xe9, 0xb9, 0xe4,
	0x32, 0x9e, 0xcd, 0x25, 0x97, 0xf1, 0xab, 0x04, 0xea, 0xcb, 0x87, 0xaf, 0x1e, 0x56, 0x83, 0x0d,
	0xe6, 0xf7, 0x7a, 0x9d, 0xde, 0x89, 0x65, 0x68, 0xed, 0xd2, 0x3f, 0x3a, 0xa2, 0xd4, 0xa3, 0x9e,
	0x55, 0x22, 0x00, 0x95, 0xe3, 0x76, 0xa7, 0x4b, 0x3d, 0xcb, 0xd4, 0x52, 0xbf, 0x73, 0x46, 0xbd,
	0xab, 0x73, 0xbf, 0x6f, 0x95, 0xb5, 0xad, 0x7d, 0x78, 0xce, 0xfa, 0xd4, 0xb3, 0xd6, 0x35, 0xf7,
	0xd9, 0xa7, 0x3e, 0xf5, 0xac, 0xca, 0xe1, 0xfe, 0xf7, 0xb7, 0xd7, 0x42, 0xdd, 0xe4, 0x03, 0x27,
	0x48, 0x22, 0x37, 0x1c, 0x8f, 0x94, 0x3b, 0x7f, 0x6c, 0xd7, 0x18, 0xbb, 0xe9, 0xe0, 0xcd, 0x75,
	0xe2, 0x2e, 0xbf, 0xe5, 0x41, 0xa5, 0x78, 0x78, 0xfb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0x0a, 0x50, 0xdc, 0xe4, 0x03, 0x00, 0x00,
}
