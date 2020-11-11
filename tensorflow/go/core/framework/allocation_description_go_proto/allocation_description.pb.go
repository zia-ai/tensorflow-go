// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/allocation_description.proto

package allocation_description_go_proto

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

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr                  uint64   `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationDescription) Reset()         { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()    {}
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_1254702e9f0c7d2f, []int{0}
}

func (m *AllocationDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationDescription.Unmarshal(m, b)
}
func (m *AllocationDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationDescription.Marshal(b, m, deterministic)
}
func (m *AllocationDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationDescription.Merge(m, src)
}
func (m *AllocationDescription) XXX_Size() int {
	return xxx_messageInfo_AllocationDescription.Size(m)
}
func (m *AllocationDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationDescription.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationDescription proto.InternalMessageInfo

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/allocation_description.proto", fileDescriptor_1254702e9f0c7d2f)
}

var fileDescriptor_1254702e9f0c7d2f = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0x5b, 0x8b, 0x2e, 0x56, 0x65, 0x51, 0x58, 0xf0, 0x12, 0x14, 0x31, 0xa7, 0x44,
	0x10, 0xbc, 0x1b, 0xbc, 0x78, 0x91, 0x12, 0x6f, 0x82, 0x2c, 0x9b, 0x64, 0xf2, 0x81, 0x49, 0x26,
	0xce, 0x6e, 0x29, 0xfe, 0x16, 0xff, 0xa8, 0xde, 0x24, 0xa9, 0xdd, 0xf4, 0xd0, 0x83, 0xb7, 0x97,
	0x67, 0x9e, 0xfd, 0xe0, 0x1d, 0x7e, 0x6f, 0xa1, 0x35, 0x48, 0x79, 0x8d, 0xab, 0x30, 0x45, 0x82,
	0x30, 0x27, 0xdd, 0xc0, 0x0a, 0xe9, 0x3d, 0xd4, 0x75, 0x8d, 0xa9, 0xb6, 0x15, 0xb6, 0x2a, 0x03,
	0x93, 0x52, 0xd5, 0xf5, 0x39, 0xe8, 0x08, 0x2d, 0x0a, 0x3e, 0x9e, 0xbb, 0xfc, 0x61, 0xfc, 0xfc,
	0xc1, 0xc9, 0x8f, 0xa3, 0x2b, 0x6e, 0xf8, 0x09, 0xc1, 0xc7, 0x12, 0x8c, 0x85, 0x4c, 0x25, 0x9f,
	0x16, 0x8c, 0x64, 0x1e, 0xf3, 0x27, 0xf1, 0xb1, 0xc3, 0x51, 0x4f, 0x7b, 0xf1, 0xef, 0x39, 0x27,
	0xee, 0xad, 0x45, 0x87, 0xd7, 0xe2, 0x35, 0xdf, 0x10, 0x24, 0xd5, 0xea, 0x06, 0xe4, 0xc4, 0x63,
	0xfe, 0x61, 0x3c, 0x77, 0xf4, 0x59, 0x37, 0x20, 0xae, 0xf8, 0x7c, 0xeb, 0xfb, 0x55, 0x26, 0xa7,
	0xc3, 0x6d, 0x47, 0x23, 0x7c, 0xca, 0xc4, 0x2d, 0x3f, 0x2b, 0xb5, 0x51, 0xa6, 0x6a, 0x8b, 0x1a,
	0x14, 0x41, 0x0e, 0x04, 0x6d, 0x0a, 0x72, 0xdf, 0x63, 0xfe, 0x41, 0x2c, 0x4a, 0x6d, 0x5e, 0x86,
	0x51, 0xbc, 0x99, 0x88, 0x53, 0x3e, 0xe9, 0x2c, 0xc9, 0x99, 0xc7, 0xfc, 0x69, 0xdc, 0xc7, 0xe8,
	0x8b, 0x71, 0x89, 0x54, 0x04, 0x63, 0x1d, 0x81, 0x6b, 0x30, 0xba, 0xd8, 0xd9, 0xca, 0xa2, 0x2f,
	0xd0, 0x2c, 0xd8, 0xeb, 0x5b, 0x51, 0xd9, 0x72, 0x99, 0x04, 0x29, 0x36, 0xe1, 0xd6, 0x1a, 0x76,
	0xc7, 0x02, 0xff, 0xb7, 0x1f, 0x55, 0xa0, 0x1a, 0x56, 0xf4, 0xcd, 0x58, 0x32, 0x1b, 0xd2, 0xdd,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xc0, 0xbc, 0x20, 0xe6, 0x01, 0x00, 0x00,
}