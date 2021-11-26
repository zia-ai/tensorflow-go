// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tensorflow/core/protobuf/graph_debug_info.proto

package for_core_protos_go_proto

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

type GraphDebugInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This stores all the source code file names and can be indexed by the
	// `file_index`.
	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// This maps a node name to a stack trace in the source code.
	// The map key is a mangling of the containing function and op name with
	// syntax:
	//   op.name '@' func_name
	// For ops in the top-level graph, the func_name is the empty string.
	// Note that op names are restricted to a small number of characters which
	// exclude '@', making it impossible to collide keys of this form. Function
	// names accept a much wider set of characters.
	// It would be preferable to avoid mangling and use a tuple key of (op.name,
	// func_name), but this is not supported with protocol buffers.
	Traces map[string]*GraphDebugInfo_StackTrace `protobuf:"bytes,2,rep,name=traces,proto3" json:"traces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GraphDebugInfo) Reset() {
	*x = GraphDebugInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphDebugInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphDebugInfo) ProtoMessage() {}

func (x *GraphDebugInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphDebugInfo.ProtoReflect.Descriptor instead.
func (*GraphDebugInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescGZIP(), []int{0}
}

func (x *GraphDebugInfo) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *GraphDebugInfo) GetTraces() map[string]*GraphDebugInfo_StackTrace {
	if x != nil {
		return x.Traces
	}
	return nil
}

// This represents a file/line location in the source code.
type GraphDebugInfo_FileLineCol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// File name index, which can be used to retrieve the file name string from
	// `files`. The value should be between 0 and (len(files)-1)
	FileIndex int32 `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
	// Line number in the file.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Col number in the file line.
	Col int32 `protobuf:"varint,3,opt,name=col,proto3" json:"col,omitempty"`
	// Name of function contains the file line.
	Func string `protobuf:"bytes,4,opt,name=func,proto3" json:"func,omitempty"`
	// Source code contained in this file line.
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GraphDebugInfo_FileLineCol) Reset() {
	*x = GraphDebugInfo_FileLineCol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphDebugInfo_FileLineCol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphDebugInfo_FileLineCol) ProtoMessage() {}

func (x *GraphDebugInfo_FileLineCol) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphDebugInfo_FileLineCol.ProtoReflect.Descriptor instead.
func (*GraphDebugInfo_FileLineCol) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GraphDebugInfo_FileLineCol) GetFileIndex() int32 {
	if x != nil {
		return x.FileIndex
	}
	return 0
}

func (x *GraphDebugInfo_FileLineCol) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *GraphDebugInfo_FileLineCol) GetCol() int32 {
	if x != nil {
		return x.Col
	}
	return 0
}

func (x *GraphDebugInfo_FileLineCol) GetFunc() string {
	if x != nil {
		return x.Func
	}
	return ""
}

func (x *GraphDebugInfo_FileLineCol) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// This represents a stack trace which is a ordered list of `FileLineCol`.
type GraphDebugInfo_StackTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each line in the stack trace.
	FileLineCols []*GraphDebugInfo_FileLineCol `protobuf:"bytes,1,rep,name=file_line_cols,json=fileLineCols,proto3" json:"file_line_cols,omitempty"`
}

func (x *GraphDebugInfo_StackTrace) Reset() {
	*x = GraphDebugInfo_StackTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphDebugInfo_StackTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphDebugInfo_StackTrace) ProtoMessage() {}

func (x *GraphDebugInfo_StackTrace) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphDebugInfo_StackTrace.ProtoReflect.Descriptor instead.
func (*GraphDebugInfo_StackTrace) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GraphDebugInfo_StackTrace) GetFileLineCols() []*GraphDebugInfo_FileLineCol {
	if x != nil {
		return x.FileLineCols
	}
	return nil
}

var File_tensorflow_core_protobuf_graph_debug_info_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_graph_debug_info_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x5f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xa0, 0x03,
	0x0a, 0x0e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x7a, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69,
	0x6e, 0x65, 0x43, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75,
	0x6e, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x1a, 0x5a, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x12, 0x4c, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6c,
	0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6c, 0x73, 0x1a, 0x60,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x8c, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x14, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescData = file_tensorflow_core_protobuf_graph_debug_info_proto_rawDesc
)

func file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_graph_debug_info_proto_rawDescData
}

var file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_protobuf_graph_debug_info_proto_goTypes = []interface{}{
	(*GraphDebugInfo)(nil),             // 0: tensorflow.GraphDebugInfo
	(*GraphDebugInfo_FileLineCol)(nil), // 1: tensorflow.GraphDebugInfo.FileLineCol
	(*GraphDebugInfo_StackTrace)(nil),  // 2: tensorflow.GraphDebugInfo.StackTrace
	nil,                                // 3: tensorflow.GraphDebugInfo.TracesEntry
}
var file_tensorflow_core_protobuf_graph_debug_info_proto_depIdxs = []int32{
	3, // 0: tensorflow.GraphDebugInfo.traces:type_name -> tensorflow.GraphDebugInfo.TracesEntry
	1, // 1: tensorflow.GraphDebugInfo.StackTrace.file_line_cols:type_name -> tensorflow.GraphDebugInfo.FileLineCol
	2, // 2: tensorflow.GraphDebugInfo.TracesEntry.value:type_name -> tensorflow.GraphDebugInfo.StackTrace
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_graph_debug_info_proto_init() }
func file_tensorflow_core_protobuf_graph_debug_info_proto_init() {
	if File_tensorflow_core_protobuf_graph_debug_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphDebugInfo); i {
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
		file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphDebugInfo_FileLineCol); i {
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
		file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphDebugInfo_StackTrace); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_graph_debug_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_graph_debug_info_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_graph_debug_info_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_graph_debug_info_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_graph_debug_info_proto = out.File
	file_tensorflow_core_protobuf_graph_debug_info_proto_rawDesc = nil
	file_tensorflow_core_protobuf_graph_debug_info_proto_goTypes = nil
	file_tensorflow_core_protobuf_graph_debug_info_proto_depIdxs = nil
}
