// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tensorflow/core/protobuf/trackable_object_graph.proto

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

type TrackableObjectGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*TrackableObjectGraph_TrackableObject `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *TrackableObjectGraph) Reset() {
	*x = TrackableObjectGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackableObjectGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackableObjectGraph) ProtoMessage() {}

func (x *TrackableObjectGraph) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackableObjectGraph.ProtoReflect.Descriptor instead.
func (*TrackableObjectGraph) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescGZIP(), []int{0}
}

func (x *TrackableObjectGraph) GetNodes() []*TrackableObjectGraph_TrackableObject {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type TrackableObjectGraph_TrackableObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Objects which this object depends on.
	Children []*TrackableObjectGraph_TrackableObject_ObjectReference `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
	// Serialized data specific to this object.
	Attributes []*TrackableObjectGraph_TrackableObject_SerializedTensor `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// Slot variables owned by this object.
	SlotVariables []*TrackableObjectGraph_TrackableObject_SlotVariableReference `protobuf:"bytes,3,rep,name=slot_variables,json=slotVariables,proto3" json:"slot_variables,omitempty"`
}

func (x *TrackableObjectGraph_TrackableObject) Reset() {
	*x = TrackableObjectGraph_TrackableObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackableObjectGraph_TrackableObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackableObjectGraph_TrackableObject) ProtoMessage() {}

func (x *TrackableObjectGraph_TrackableObject) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackableObjectGraph_TrackableObject.ProtoReflect.Descriptor instead.
func (*TrackableObjectGraph_TrackableObject) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TrackableObjectGraph_TrackableObject) GetChildren() []*TrackableObjectGraph_TrackableObject_ObjectReference {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *TrackableObjectGraph_TrackableObject) GetAttributes() []*TrackableObjectGraph_TrackableObject_SerializedTensor {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *TrackableObjectGraph_TrackableObject) GetSlotVariables() []*TrackableObjectGraph_TrackableObject_SlotVariableReference {
	if x != nil {
		return x.SlotVariables
	}
	return nil
}

type TrackableObjectGraph_TrackableObject_ObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An index into `TrackableObjectGraph.nodes`, indicating the object
	// being referenced.
	NodeId int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// A user-provided name for the edge.
	LocalName string `protobuf:"bytes,2,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
}

func (x *TrackableObjectGraph_TrackableObject_ObjectReference) Reset() {
	*x = TrackableObjectGraph_TrackableObject_ObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackableObjectGraph_TrackableObject_ObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackableObjectGraph_TrackableObject_ObjectReference) ProtoMessage() {}

func (x *TrackableObjectGraph_TrackableObject_ObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackableObjectGraph_TrackableObject_ObjectReference.ProtoReflect.Descriptor instead.
func (*TrackableObjectGraph_TrackableObject_ObjectReference) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *TrackableObjectGraph_TrackableObject_ObjectReference) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *TrackableObjectGraph_TrackableObject_ObjectReference) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

type TrackableObjectGraph_TrackableObject_SerializedTensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A name for the Tensor. Simple variables have only one
	// `SerializedTensor` named "VARIABLE_VALUE" by convention. This value may
	// be restored on object creation as an optimization.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The full name of the variable/tensor, if applicable. Used to allow
	// name-based loading of checkpoints which were saved using an
	// object-based API. Should match the checkpoint key which would have been
	// assigned by tf.train.Saver.
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// The generated name of the Tensor in the checkpoint.
	CheckpointKey string `protobuf:"bytes,3,opt,name=checkpoint_key,json=checkpointKey,proto3" json:"checkpoint_key,omitempty"`
	// Whether checkpoints should be considered as matching even without this
	// value restored. Used for non-critical values which don't affect the
	// TensorFlow graph, such as layer configurations.
	OptionalRestore bool `protobuf:"varint,4,opt,name=optional_restore,json=optionalRestore,proto3" json:"optional_restore,omitempty"`
}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) Reset() {
	*x = TrackableObjectGraph_TrackableObject_SerializedTensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackableObjectGraph_TrackableObject_SerializedTensor) ProtoMessage() {}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackableObjectGraph_TrackableObject_SerializedTensor.ProtoReflect.Descriptor instead.
func (*TrackableObjectGraph_TrackableObject_SerializedTensor) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) GetCheckpointKey() string {
	if x != nil {
		return x.CheckpointKey
	}
	return ""
}

func (x *TrackableObjectGraph_TrackableObject_SerializedTensor) GetOptionalRestore() bool {
	if x != nil {
		return x.OptionalRestore
	}
	return false
}

type TrackableObjectGraph_TrackableObject_SlotVariableReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An index into `TrackableObjectGraph.nodes`, indicating the
	// variable object this slot was created for.
	OriginalVariableNodeId int32 `protobuf:"varint,1,opt,name=original_variable_node_id,json=originalVariableNodeId,proto3" json:"original_variable_node_id,omitempty"`
	// The name of the slot (e.g. "m"/"v").
	SlotName string `protobuf:"bytes,2,opt,name=slot_name,json=slotName,proto3" json:"slot_name,omitempty"`
	// An index into `TrackableObjectGraph.nodes`, indicating the
	// `Object` with the value of the slot variable.
	SlotVariableNodeId int32 `protobuf:"varint,3,opt,name=slot_variable_node_id,json=slotVariableNodeId,proto3" json:"slot_variable_node_id,omitempty"`
}

func (x *TrackableObjectGraph_TrackableObject_SlotVariableReference) Reset() {
	*x = TrackableObjectGraph_TrackableObject_SlotVariableReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackableObjectGraph_TrackableObject_SlotVariableReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackableObjectGraph_TrackableObject_SlotVariableReference) ProtoMessage() {}

func (x *TrackableObjectGraph_TrackableObject_SlotVariableReference) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackableObjectGraph_TrackableObject_SlotVariableReference.ProtoReflect.Descriptor instead.
func (*TrackableObjectGraph_TrackableObject_SlotVariableReference) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *TrackableObjectGraph_TrackableObject_SlotVariableReference) GetOriginalVariableNodeId() int32 {
	if x != nil {
		return x.OriginalVariableNodeId
	}
	return 0
}

func (x *TrackableObjectGraph_TrackableObject_SlotVariableReference) GetSlotName() string {
	if x != nil {
		return x.SlotName
	}
	return ""
}

func (x *TrackableObjectGraph_TrackableObject_SlotVariableReference) GetSlotVariableNodeId() int32 {
	if x != nil {
		return x.SlotVariableNodeId
	}
	return 0
}

var File_tensorflow_core_protobuf_trackable_object_graph_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0xaa, 0x06, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x46, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x1a, 0xc9, 0x05, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x5c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x61, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x6d, 0x0a, 0x0e, 0x73, 0x6c, 0x6f,
	0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x46, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0d, 0x73, 0x6c, 0x6f, 0x74, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x49, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x95, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0xa2, 0x01, 0x0a, 0x15,
	0x53, 0x6c, 0x6f, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6c, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x15, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x73, 0x6c,
	0x6f, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x42, 0x5a, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescData = file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDesc
)

func file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDescData
}

var file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_core_protobuf_trackable_object_graph_proto_goTypes = []interface{}{
	(*TrackableObjectGraph)(nil),                                       // 0: tensorflow.TrackableObjectGraph
	(*TrackableObjectGraph_TrackableObject)(nil),                       // 1: tensorflow.TrackableObjectGraph.TrackableObject
	(*TrackableObjectGraph_TrackableObject_ObjectReference)(nil),       // 2: tensorflow.TrackableObjectGraph.TrackableObject.ObjectReference
	(*TrackableObjectGraph_TrackableObject_SerializedTensor)(nil),      // 3: tensorflow.TrackableObjectGraph.TrackableObject.SerializedTensor
	(*TrackableObjectGraph_TrackableObject_SlotVariableReference)(nil), // 4: tensorflow.TrackableObjectGraph.TrackableObject.SlotVariableReference
}
var file_tensorflow_core_protobuf_trackable_object_graph_proto_depIdxs = []int32{
	1, // 0: tensorflow.TrackableObjectGraph.nodes:type_name -> tensorflow.TrackableObjectGraph.TrackableObject
	2, // 1: tensorflow.TrackableObjectGraph.TrackableObject.children:type_name -> tensorflow.TrackableObjectGraph.TrackableObject.ObjectReference
	3, // 2: tensorflow.TrackableObjectGraph.TrackableObject.attributes:type_name -> tensorflow.TrackableObjectGraph.TrackableObject.SerializedTensor
	4, // 3: tensorflow.TrackableObjectGraph.TrackableObject.slot_variables:type_name -> tensorflow.TrackableObjectGraph.TrackableObject.SlotVariableReference
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_trackable_object_graph_proto_init() }
func file_tensorflow_core_protobuf_trackable_object_graph_proto_init() {
	if File_tensorflow_core_protobuf_trackable_object_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackableObjectGraph); i {
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
		file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackableObjectGraph_TrackableObject); i {
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
		file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackableObjectGraph_TrackableObject_ObjectReference); i {
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
		file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackableObjectGraph_TrackableObject_SerializedTensor); i {
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
		file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackableObjectGraph_TrackableObject_SlotVariableReference); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_trackable_object_graph_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_trackable_object_graph_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_trackable_object_graph_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_trackable_object_graph_proto = out.File
	file_tensorflow_core_protobuf_trackable_object_graph_proto_rawDesc = nil
	file_tensorflow_core_protobuf_trackable_object_graph_proto_goTypes = nil
	file_tensorflow_core_protobuf_trackable_object_graph_proto_depIdxs = nil
}
