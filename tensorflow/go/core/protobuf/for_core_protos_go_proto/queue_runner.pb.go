// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/queue_runner.proto

package for_core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Protocol buffer representing a QueueRunner.
type QueueRunnerDef struct {
	// Queue name.
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	// A list of enqueue operations.
	EnqueueOpName []string `protobuf:"bytes,2,rep,name=enqueue_op_name,json=enqueueOpName,proto3" json:"enqueue_op_name,omitempty"`
	// The operation to run to close the queue.
	CloseOpName string `protobuf:"bytes,3,opt,name=close_op_name,json=closeOpName,proto3" json:"close_op_name,omitempty"`
	// The operation to run to cancel the queue.
	CancelOpName string `protobuf:"bytes,4,opt,name=cancel_op_name,json=cancelOpName,proto3" json:"cancel_op_name,omitempty"`
	// A list of exception types considered to signal a safely closed queue
	// if raised during enqueue operations.
	QueueClosedExceptionTypes []Code   `protobuf:"varint,5,rep,packed,name=queue_closed_exception_types,json=queueClosedExceptionTypes,proto3,enum=tensorflow.error.Code" json:"queue_closed_exception_types,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *QueueRunnerDef) Reset()         { *m = QueueRunnerDef{} }
func (m *QueueRunnerDef) String() string { return proto.CompactTextString(m) }
func (*QueueRunnerDef) ProtoMessage()    {}
func (*QueueRunnerDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af35200d68d14ae, []int{0}
}

func (m *QueueRunnerDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRunnerDef.Unmarshal(m, b)
}
func (m *QueueRunnerDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRunnerDef.Marshal(b, m, deterministic)
}
func (m *QueueRunnerDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRunnerDef.Merge(m, src)
}
func (m *QueueRunnerDef) XXX_Size() int {
	return xxx_messageInfo_QueueRunnerDef.Size(m)
}
func (m *QueueRunnerDef) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRunnerDef.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRunnerDef proto.InternalMessageInfo

func (m *QueueRunnerDef) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRunnerDef) GetEnqueueOpName() []string {
	if m != nil {
		return m.EnqueueOpName
	}
	return nil
}

func (m *QueueRunnerDef) GetCloseOpName() string {
	if m != nil {
		return m.CloseOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetCancelOpName() string {
	if m != nil {
		return m.CancelOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetQueueClosedExceptionTypes() []Code {
	if m != nil {
		return m.QueueClosedExceptionTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueRunnerDef)(nil), "tensorflow.QueueRunnerDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/queue_runner.proto", fileDescriptor_7af35200d68d14ae)
}

var fileDescriptor_7af35200d68d14ae = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0xab, 0x42, 0x57, 0x5b, 0xb1, 0x07, 0xa9, 0xa2, 0x50, 0x8a, 0x48, 0x51, 0x48,
	0x40, 0xdf, 0xa0, 0xd5, 0xab, 0xd6, 0xa0, 0x08, 0x5e, 0x96, 0x74, 0x33, 0x89, 0xc5, 0x64, 0x27,
	0xce, 0x26, 0x54, 0x1f, 0xc1, 0x37, 0x16, 0x4f, 0x92, 0xd9, 0x68, 0x82, 0xe8, 0x6d, 0xf8, 0xf9,
	0xbe, 0x19, 0xe6, 0x97, 0x67, 0x05, 0x18, 0x8b, 0x14, 0xa7, 0xb8, 0xf2, 0x35, 0x12, 0xf8, 0x39,
	0x61, 0x81, 0x8b, 0x32, 0xf6, 0x5f, 0x4a, 0x28, 0x41, 0x51, 0x69, 0x0c, 0x90, 0xc7, 0xe9, 0x40,
	0x36, 0xf0, 0xc1, 0xe9, 0xbf, 0x22, 0x10, 0x21, 0x29, 0x8d, 0x11, 0x58, 0xe7, 0x8d, 0x3f, 0x85,
	0xec, 0xdf, 0x56, 0xeb, 0x02, 0xde, 0x76, 0x09, 0xf1, 0xe0, 0x48, 0x4a, 0x77, 0xc0, 0x84, 0x19,
	0x0c, 0xc5, 0x48, 0x4c, 0xba, 0x41, 0x97, 0x93, 0xeb, 0x30, 0x83, 0xc1, 0x89, 0xdc, 0x01, 0xe3,
	0x00, 0xcc, 0x1d, 0xb3, 0x36, 0xea, 0x4c, 0xba, 0x41, 0xaf, 0x8e, 0x6f, 0x72, 0xe6, 0xc6, 0xb2,
	0xa7, 0x53, 0xb4, 0x0d, 0xd5, 0xe1, 0x4d, 0x5b, 0x1c, 0xd6, 0xcc, 0xb1, 0xec, 0xeb, 0xd0, 0x68,
	0x48, 0x7f, 0xa0, 0x75, 0x86, 0xb6, 0x5d, 0x5a, 0x53, 0x0f, 0xf2, 0xd0, 0xdd, 0x63, 0x35, 0x52,
	0xf0, 0xaa, 0x21, 0x2f, 0x96, 0x68, 0x54, 0xf1, 0x96, 0x83, 0x1d, 0x6e, 0x8c, 0x3a, 0x93, 0xfe,
	0xf9, 0x9e, 0xd7, 0xbc, 0xed, 0xf1, 0xa3, 0xde, 0x0c, 0x23, 0x08, 0xf6, 0xd9, 0x9d, 0xb1, 0x7a,
	0xf5, 0x6d, 0xde, 0x55, 0xe2, 0xf4, 0x5d, 0xc8, 0x21, 0x52, 0xd2, 0x16, 0x63, 0x0a, 0x33, 0x58,
	0x21, 0x3d, 0x4f, 0x77, 0x5b, 0xb5, 0xcc, 0xab, 0xae, 0xec, 0x5c, 0x3c, 0xde, 0x27, 0xcb, 0xe2,
	0xa9, 0x5c, 0x78, 0x1a, 0x33, 0xbf, 0xd5, 0xf2, 0xdf, 0x63, 0x82, 0xbf, 0xea, 0x8f, 0xb9, 0x7c,
	0x02, 0xc5, 0x89, 0x55, 0x09, 0xba, 0xe9, 0x43, 0x88, 0xc5, 0x26, 0x4f, 0x17, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x6d, 0xb6, 0x77, 0xf6, 0x01, 0x00, 0x00,
}