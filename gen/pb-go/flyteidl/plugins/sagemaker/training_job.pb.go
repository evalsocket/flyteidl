// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/training_job.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InputMode_Value int32

const (
	InputMode_FILE InputMode_Value = 0
	InputMode_PIPE InputMode_Value = 1
)

var InputMode_Value_name = map[int32]string{
	0: "FILE",
	1: "PIPE",
}

var InputMode_Value_value = map[string]int32{
	"FILE": 0,
	"PIPE": 1,
}

func (x InputMode_Value) String() string {
	return proto.EnumName(InputMode_Value_name, int32(x))
}

func (InputMode_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0, 0}
}

type AlgorithmName_Value int32

const (
	AlgorithmName_CUSTOM  AlgorithmName_Value = 0
	AlgorithmName_XGBOOST AlgorithmName_Value = 1
)

var AlgorithmName_Value_name = map[int32]string{
	0: "CUSTOM",
	1: "XGBOOST",
}

var AlgorithmName_Value_value = map[string]int32{
	"CUSTOM":  0,
	"XGBOOST": 1,
}

func (x AlgorithmName_Value) String() string {
	return proto.EnumName(AlgorithmName_Value_name, int32(x))
}

func (AlgorithmName_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1, 0}
}

type InputMode struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputMode) Reset()         { *m = InputMode{} }
func (m *InputMode) String() string { return proto.CompactTextString(m) }
func (*InputMode) ProtoMessage()    {}
func (*InputMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0}
}

func (m *InputMode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputMode.Unmarshal(m, b)
}
func (m *InputMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputMode.Marshal(b, m, deterministic)
}
func (m *InputMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputMode.Merge(m, src)
}
func (m *InputMode) XXX_Size() int {
	return xxx_messageInfo_InputMode.Size(m)
}
func (m *InputMode) XXX_DiscardUnknown() {
	xxx_messageInfo_InputMode.DiscardUnknown(m)
}

var xxx_messageInfo_InputMode proto.InternalMessageInfo

type AlgorithmName struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlgorithmName) Reset()         { *m = AlgorithmName{} }
func (m *AlgorithmName) String() string { return proto.CompactTextString(m) }
func (*AlgorithmName) ProtoMessage()    {}
func (*AlgorithmName) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1}
}

func (m *AlgorithmName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmName.Unmarshal(m, b)
}
func (m *AlgorithmName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmName.Marshal(b, m, deterministic)
}
func (m *AlgorithmName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmName.Merge(m, src)
}
func (m *AlgorithmName) XXX_Size() int {
	return xxx_messageInfo_AlgorithmName.Size(m)
}
func (m *AlgorithmName) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmName.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmName proto.InternalMessageInfo

// Specifies a metric that the training algorithm writes to stderr or stdout.
// This object is a pass-through.
// See this for details: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_MetricDefinition.html
type MetricDefinition struct {
	// User-defined name of the metric
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SageMaker hyperparameter tuning parses your algorithm’s stdout and stderr streams to find algorithm metrics
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricDefinition) Reset()         { *m = MetricDefinition{} }
func (m *MetricDefinition) String() string { return proto.CompactTextString(m) }
func (*MetricDefinition) ProtoMessage()    {}
func (*MetricDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{2}
}

func (m *MetricDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDefinition.Unmarshal(m, b)
}
func (m *MetricDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDefinition.Marshal(b, m, deterministic)
}
func (m *MetricDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDefinition.Merge(m, src)
}
func (m *MetricDefinition) XXX_Size() int {
	return xxx_messageInfo_MetricDefinition.Size(m)
}
func (m *MetricDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDefinition proto.InternalMessageInfo

func (m *MetricDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricDefinition) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// Specifies the training algorithm to be used in the training job
// This object is mostly a pass-through, with a couple of exceptions include: (1) in Flyte, users don't need to specify
// TrainingImage; either use the built-in algorithm mode by using Flytekit's Simple Training Job and specifying an algorithm
// name and an algorithm version or (2) when users want to supply custom algorithms they should set algorithm_name field to
// CUSTOM. In this case, the value of the algorithm_version field has no effect
// For pass-through use cases: refer to this AWS official document for more details
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
type AlgorithmSpecification struct {
	// The input mode can be either PIPE or FILE
	InputMode InputMode_Value `protobuf:"varint,1,opt,name=input_mode,json=inputMode,proto3,enum=flyteidl.plugins.sagemaker.InputMode_Value" json:"input_mode,omitempty"`
	// The algorithm name is used for deciding which pre-built image to point to
	// This is only needed for use cases where SageMaker's built-in algorithm mode is chosen
	AlgorithmName AlgorithmName_Value `protobuf:"varint,2,opt,name=algorithm_name,json=algorithmName,proto3,enum=flyteidl.plugins.sagemaker.AlgorithmName_Value" json:"algorithm_name,omitempty"`
	// The algorithm version field is used for deciding which pre-built image to point to
	// This is only needed for use cases where SageMaker's built-in algorithm mode is chosen
	AlgorithmVersion string `protobuf:"bytes,3,opt,name=algorithm_version,json=algorithmVersion,proto3" json:"algorithm_version,omitempty"`
	// A list of metric definitions for SageMaker to evaluate/track on the progress of the training job
	// See this: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
	// and this: https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics.html
	MetricDefinitions    []*MetricDefinition `protobuf:"bytes,4,rep,name=metric_definitions,json=metricDefinitions,proto3" json:"metric_definitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AlgorithmSpecification) Reset()         { *m = AlgorithmSpecification{} }
func (m *AlgorithmSpecification) String() string { return proto.CompactTextString(m) }
func (*AlgorithmSpecification) ProtoMessage()    {}
func (*AlgorithmSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{3}
}

func (m *AlgorithmSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmSpecification.Unmarshal(m, b)
}
func (m *AlgorithmSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmSpecification.Marshal(b, m, deterministic)
}
func (m *AlgorithmSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmSpecification.Merge(m, src)
}
func (m *AlgorithmSpecification) XXX_Size() int {
	return xxx_messageInfo_AlgorithmSpecification.Size(m)
}
func (m *AlgorithmSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmSpecification proto.InternalMessageInfo

func (m *AlgorithmSpecification) GetInputMode() InputMode_Value {
	if m != nil {
		return m.InputMode
	}
	return InputMode_FILE
}

func (m *AlgorithmSpecification) GetAlgorithmName() AlgorithmName_Value {
	if m != nil {
		return m.AlgorithmName
	}
	return AlgorithmName_CUSTOM
}

func (m *AlgorithmSpecification) GetAlgorithmVersion() string {
	if m != nil {
		return m.AlgorithmVersion
	}
	return ""
}

func (m *AlgorithmSpecification) GetMetricDefinitions() []*MetricDefinition {
	if m != nil {
		return m.MetricDefinitions
	}
	return nil
}

// TrainingJobResourceConfig is a pass-through, specifying the instance type to use for the training job, the
// number of instances to launch, and the size of the ML storage volume the user wants to provision
// Refer to SageMaker official doc for more details: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
type TrainingJobResourceConfig struct {
	// The number of ML compute instances to use. For distributed training, provide a value greater than 1.
	InstanceCount int64 `protobuf:"varint,1,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	// The ML compute instance type
	InstanceType string `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// The size of the ML storage volume that you want to provision.
	VolumeSizeInGb       int64    `protobuf:"varint,3,opt,name=volume_size_in_gb,json=volumeSizeInGb,proto3" json:"volume_size_in_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainingJobResourceConfig) Reset()         { *m = TrainingJobResourceConfig{} }
func (m *TrainingJobResourceConfig) String() string { return proto.CompactTextString(m) }
func (*TrainingJobResourceConfig) ProtoMessage()    {}
func (*TrainingJobResourceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{4}
}

func (m *TrainingJobResourceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJobResourceConfig.Unmarshal(m, b)
}
func (m *TrainingJobResourceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJobResourceConfig.Marshal(b, m, deterministic)
}
func (m *TrainingJobResourceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJobResourceConfig.Merge(m, src)
}
func (m *TrainingJobResourceConfig) XXX_Size() int {
	return xxx_messageInfo_TrainingJobResourceConfig.Size(m)
}
func (m *TrainingJobResourceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJobResourceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJobResourceConfig proto.InternalMessageInfo

func (m *TrainingJobResourceConfig) GetInstanceCount() int64 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

func (m *TrainingJobResourceConfig) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *TrainingJobResourceConfig) GetVolumeSizeInGb() int64 {
	if m != nil {
		return m.VolumeSizeInGb
	}
	return 0
}

// The spec of a training job. This is mostly a pass-through object
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
type TrainingJob struct {
	AlgorithmSpecification    *AlgorithmSpecification    `protobuf:"bytes,1,opt,name=algorithm_specification,json=algorithmSpecification,proto3" json:"algorithm_specification,omitempty"`
	TrainingJobResourceConfig *TrainingJobResourceConfig `protobuf:"bytes,2,opt,name=training_job_resource_config,json=trainingJobResourceConfig,proto3" json:"training_job_resource_config,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                   `json:"-"`
	XXX_unrecognized          []byte                     `json:"-"`
	XXX_sizecache             int32                      `json:"-"`
}

func (m *TrainingJob) Reset()         { *m = TrainingJob{} }
func (m *TrainingJob) String() string { return proto.CompactTextString(m) }
func (*TrainingJob) ProtoMessage()    {}
func (*TrainingJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{5}
}

func (m *TrainingJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJob.Unmarshal(m, b)
}
func (m *TrainingJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJob.Marshal(b, m, deterministic)
}
func (m *TrainingJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJob.Merge(m, src)
}
func (m *TrainingJob) XXX_Size() int {
	return xxx_messageInfo_TrainingJob.Size(m)
}
func (m *TrainingJob) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJob.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJob proto.InternalMessageInfo

func (m *TrainingJob) GetAlgorithmSpecification() *AlgorithmSpecification {
	if m != nil {
		return m.AlgorithmSpecification
	}
	return nil
}

func (m *TrainingJob) GetTrainingJobResourceConfig() *TrainingJobResourceConfig {
	if m != nil {
		return m.TrainingJobResourceConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.InputMode_Value", InputMode_Value_name, InputMode_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.AlgorithmName_Value", AlgorithmName_Value_name, AlgorithmName_Value_value)
	proto.RegisterType((*InputMode)(nil), "flyteidl.plugins.sagemaker.InputMode")
	proto.RegisterType((*AlgorithmName)(nil), "flyteidl.plugins.sagemaker.AlgorithmName")
	proto.RegisterType((*MetricDefinition)(nil), "flyteidl.plugins.sagemaker.MetricDefinition")
	proto.RegisterType((*AlgorithmSpecification)(nil), "flyteidl.plugins.sagemaker.AlgorithmSpecification")
	proto.RegisterType((*TrainingJobResourceConfig)(nil), "flyteidl.plugins.sagemaker.TrainingJobResourceConfig")
	proto.RegisterType((*TrainingJob)(nil), "flyteidl.plugins.sagemaker.TrainingJob")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/training_job.proto", fileDescriptor_6a68f64d8fd9fe30)
}

var fileDescriptor_6a68f64d8fd9fe30 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xad, 0x9b, 0xb4, 0xf7, 0x66, 0x42, 0x22, 0x67, 0x85, 0x4a, 0x5a, 0x10, 0x8a, 0x8c, 0x90,
	0x82, 0x4a, 0x6d, 0x91, 0xaa, 0x6f, 0xbc, 0xd0, 0x50, 0xaa, 0x54, 0x84, 0x54, 0x4e, 0x88, 0x10,
	0x3c, 0x58, 0xb6, 0x33, 0xde, 0x2e, 0xb5, 0x77, 0x2d, 0x7b, 0x1d, 0x91, 0x7e, 0x06, 0x1f, 0xc5,
	0x07, 0xf1, 0x05, 0x28, 0x9b, 0xd8, 0x4d, 0x23, 0x12, 0xf1, 0xb6, 0x3e, 0x33, 0x3e, 0x73, 0xe6,
	0xcc, 0x0c, 0x9c, 0x04, 0xe1, 0x4c, 0x22, 0x9b, 0x84, 0x56, 0x1c, 0x66, 0x94, 0xf1, 0xd4, 0x4a,
	0x5d, 0x8a, 0x91, 0x7b, 0x8b, 0x89, 0x25, 0x13, 0x97, 0x71, 0xc6, 0xa9, 0xf3, 0x5d, 0x78, 0x66,
	0x9c, 0x08, 0x29, 0xc8, 0x51, 0x9e, 0x6e, 0x2e, 0xd3, 0xcd, 0x22, 0xfd, 0xe8, 0x39, 0x15, 0x82,
	0x86, 0x68, 0xa9, 0x4c, 0x2f, 0x0b, 0xac, 0x49, 0x96, 0xb8, 0x92, 0x09, 0xbe, 0xf8, 0xd7, 0x68,
	0x43, 0xa5, 0xc7, 0xe3, 0x4c, 0xf6, 0xc5, 0x04, 0x8d, 0xa7, 0xb0, 0x37, 0x76, 0xc3, 0x0c, 0xc9,
	0xff, 0x50, 0xfe, 0xd0, 0xfb, 0x78, 0xa1, 0xef, 0xcc, 0x5f, 0xd7, 0xbd, 0xeb, 0x0b, 0x5d, 0x33,
	0xde, 0x40, 0xed, 0x5d, 0x48, 0x45, 0xc2, 0xe4, 0x4d, 0xf4, 0xc9, 0x8d, 0xd0, 0x68, 0xe5, 0xd9,
	0x00, 0xfb, 0xdd, 0xcf, 0xc3, 0xd1, 0xa0, 0xaf, 0xef, 0x90, 0x2a, 0xfc, 0xf7, 0xe5, 0xf2, 0x7c,
	0x30, 0x18, 0x8e, 0x74, 0xcd, 0x78, 0x0b, 0x7a, 0x1f, 0x65, 0xc2, 0xfc, 0xf7, 0x18, 0x30, 0xce,
	0xe6, 0x65, 0x09, 0x81, 0x32, 0x77, 0x23, 0x6c, 0x6a, 0x2d, 0xad, 0x5d, 0xb1, 0xd5, 0x9b, 0x3c,
	0x86, 0xbd, 0x04, 0x29, 0xfe, 0x68, 0xee, 0x2a, 0x70, 0xf1, 0x61, 0xfc, 0xda, 0x85, 0x83, 0xa2,
	0xe2, 0x30, 0x46, 0x9f, 0x05, 0xcc, 0x57, 0xda, 0xc9, 0x15, 0x00, 0x9b, 0xab, 0x76, 0x22, 0x31,
	0x59, 0x50, 0xd5, 0x3b, 0xc7, 0xe6, 0x66, 0x1b, 0xcc, 0xa2, 0x47, 0x53, 0x49, 0xb6, 0x2b, 0x2c,
	0x07, 0xc8, 0x18, 0xea, 0x6e, 0x5e, 0xc5, 0x51, 0xd2, 0x76, 0x15, 0x9f, 0xb5, 0x8d, 0xef, 0x81,
	0x13, 0x4b, 0xce, 0x9a, 0xbb, 0x0a, 0x92, 0x63, 0x68, 0xdc, 0xf3, 0x4e, 0x31, 0x49, 0x99, 0xe0,
	0xcd, 0x92, 0x6a, 0x50, 0x2f, 0x02, 0xe3, 0x05, 0x4e, 0xbe, 0x01, 0x89, 0x94, 0x53, 0xce, 0xa4,
	0xb0, 0x2a, 0x6d, 0x96, 0x5b, 0xa5, 0x76, 0xb5, 0xf3, 0x7a, 0x9b, 0x90, 0x75, 0x7f, 0xed, 0x46,
	0xb4, 0x86, 0xa4, 0xc6, 0x4f, 0x0d, 0x0e, 0x47, 0xcb, 0xb5, 0xb9, 0x12, 0x9e, 0x8d, 0xa9, 0xc8,
	0x12, 0x1f, 0xbb, 0x82, 0x07, 0x8c, 0x92, 0x97, 0x50, 0x67, 0x3c, 0x95, 0x2e, 0xf7, 0xd1, 0xf1,
	0x45, 0xc6, 0xa5, 0xf2, 0xb3, 0x64, 0xd7, 0x72, 0xb4, 0x3b, 0x07, 0xc9, 0x0b, 0x28, 0x00, 0x47,
	0xce, 0x62, 0x5c, 0xce, 0xea, 0x51, 0x0e, 0x8e, 0x66, 0x31, 0x92, 0x57, 0xd0, 0x98, 0x8a, 0x30,
	0x8b, 0xd0, 0x49, 0xd9, 0x1d, 0x3a, 0x8c, 0x3b, 0xd4, 0x53, 0x3d, 0x97, 0xec, 0xfa, 0x22, 0x30,
	0x64, 0x77, 0xd8, 0xe3, 0x97, 0x9e, 0xf1, 0x5b, 0x83, 0xea, 0x8a, 0x28, 0x72, 0x0b, 0x4f, 0xee,
	0xed, 0x4a, 0x57, 0xa7, 0xad, 0xf4, 0x54, 0x3b, 0x9d, 0x7f, 0x9a, 0xc7, 0x83, 0x3d, 0xb1, 0x0f,
	0xdc, 0xbf, 0xef, 0xcf, 0x14, 0x9e, 0xad, 0xde, 0x91, 0x93, 0x2c, 0x2d, 0x71, 0x7c, 0xe5, 0x89,
	0xea, 0xad, 0xda, 0x39, 0xdb, 0x56, 0x71, 0xa3, 0xa1, 0xf6, 0xa1, 0xdc, 0x14, 0x3a, 0x3f, 0xfb,
	0x7a, 0x4a, 0x99, 0xbc, 0xc9, 0x3c, 0xd3, 0x17, 0x91, 0x15, 0xce, 0x02, 0x69, 0x15, 0xa7, 0x4e,
	0x91, 0x5b, 0xb1, 0x77, 0x42, 0x85, 0xb5, 0x7e, 0xfd, 0xde, 0xbe, 0xba, 0xd5, 0xd3, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x32, 0x95, 0x03, 0x39, 0x18, 0x04, 0x00, 0x00,
}