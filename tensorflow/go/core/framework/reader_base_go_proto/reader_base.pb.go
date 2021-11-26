// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tensorflow/core/framework/reader_base.proto

package reader_base_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// For serializing and restoring the state of ReaderBase, see
// reader_base.h for details.
type ReaderBaseState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkStarted        int64  `protobuf:"varint,1,opt,name=work_started,json=workStarted,proto3" json:"work_started,omitempty"`
	WorkFinished       int64  `protobuf:"varint,2,opt,name=work_finished,json=workFinished,proto3" json:"work_finished,omitempty"`
	NumRecordsProduced int64  `protobuf:"varint,3,opt,name=num_records_produced,json=numRecordsProduced,proto3" json:"num_records_produced,omitempty"`
	CurrentWork        []byte `protobuf:"bytes,4,opt,name=current_work,json=currentWork,proto3" json:"current_work,omitempty"`
}

func (x *ReaderBaseState) Reset() {
	*x = ReaderBaseState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_reader_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReaderBaseState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReaderBaseState) ProtoMessage() {}

func (x *ReaderBaseState) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_reader_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReaderBaseState.ProtoReflect.Descriptor instead.
func (*ReaderBaseState) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_reader_base_proto_rawDescGZIP(), []int{0}
}

func (x *ReaderBaseState) GetWorkStarted() int64 {
	if x != nil {
		return x.WorkStarted
	}
	return 0
}

func (x *ReaderBaseState) GetWorkFinished() int64 {
	if x != nil {
		return x.WorkFinished
	}
	return 0
}

func (x *ReaderBaseState) GetNumRecordsProduced() int64 {
	if x != nil {
		return x.NumRecordsProduced
	}
	return 0
}

func (x *ReaderBaseState) GetCurrentWork() []byte {
	if x != nil {
		return x.CurrentWork
	}
	return nil
}

var File_tensorflow_core_framework_reader_base_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_reader_base_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x42, 0x85, 0x01, 0x0a, 0x18, 0x6f,
	0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x10, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42,
	0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_reader_base_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_reader_base_proto_rawDescData = file_tensorflow_core_framework_reader_base_proto_rawDesc
)

func file_tensorflow_core_framework_reader_base_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_reader_base_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_reader_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_reader_base_proto_rawDescData)
	})
	return file_tensorflow_core_framework_reader_base_proto_rawDescData
}

var file_tensorflow_core_framework_reader_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_framework_reader_base_proto_goTypes = []interface{}{
	(*ReaderBaseState)(nil), // 0: tensorflow.ReaderBaseState
}
var file_tensorflow_core_framework_reader_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_reader_base_proto_init() }
func file_tensorflow_core_framework_reader_base_proto_init() {
	if File_tensorflow_core_framework_reader_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_reader_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReaderBaseState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_reader_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_reader_base_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_reader_base_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_reader_base_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_reader_base_proto = out.File
	file_tensorflow_core_framework_reader_base_proto_rawDesc = nil
	file_tensorflow_core_framework_reader_base_proto_goTypes = nil
	file_tensorflow_core_framework_reader_base_proto_depIdxs = nil
}
