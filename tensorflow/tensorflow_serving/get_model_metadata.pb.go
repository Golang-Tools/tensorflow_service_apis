// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: tensorflow_serving/apis/get_model_metadata.proto

package tensorflow_serving

import (
	protobuf "github.com/Golang-Tools/tensorflow_service_apis/tensorflow/tensorflow/core/protobuf"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Message returned for "signature_def" field.
type SignatureDefMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignatureDef map[string]*protobuf.SignatureDef `protobuf:"bytes,1,rep,name=signature_def,json=signatureDef,proto3" json:"signature_def,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SignatureDefMap) Reset() {
	*x = SignatureDefMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureDefMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureDefMap) ProtoMessage() {}

func (x *SignatureDefMap) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureDefMap.ProtoReflect.Descriptor instead.
func (*SignatureDefMap) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *SignatureDefMap) GetSignatureDef() map[string]*protobuf.SignatureDef {
	if x != nil {
		return x.SignatureDef
	}
	return nil
}

type GetModelMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification indicating which model we are querying for metadata.
	// If version is not specified, will use the latest (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Metadata fields to get. Currently supported: "signature_def".
	MetadataField []string `protobuf:"bytes,2,rep,name=metadata_field,json=metadataField,proto3" json:"metadata_field,omitempty"`
}

func (x *GetModelMetadataRequest) Reset() {
	*x = GetModelMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelMetadataRequest) ProtoMessage() {}

func (x *GetModelMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetModelMetadataRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *GetModelMetadataRequest) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *GetModelMetadataRequest) GetMetadataField() []string {
	if x != nil {
		return x.MetadataField
	}
	return nil
}

type GetModelMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification indicating which model this metadata belongs to.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Map of metadata field name to metadata field. The options for metadata
	// field name are listed in GetModelMetadataRequest. Currently supported:
	// "signature_def".
	Metadata map[string]*any.Any `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetModelMetadataResponse) Reset() {
	*x = GetModelMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelMetadataResponse) ProtoMessage() {}

func (x *GetModelMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetModelMetadataResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *GetModelMetadataResponse) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *GetModelMetadataResponse) GetMetadata() map[string]*any.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_tensorflow_serving_apis_get_model_metadata_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_get_model_metadata_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44,
	0x65, 0x66, 0x4d, 0x61, 0x70, 0x12, 0x5a, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x4d, 0x61,
	0x70, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65,
	0x66, 0x1a, 0x59, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65,
	0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65,
	0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7e, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x83, 0x02, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x56, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x22, 0x5a, 0x1d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_get_model_metadata_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_get_model_metadata_proto_rawDescData = file_tensorflow_serving_apis_get_model_metadata_proto_rawDesc
)

func file_tensorflow_serving_apis_get_model_metadata_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_get_model_metadata_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_get_model_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_get_model_metadata_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_get_model_metadata_proto_rawDescData
}

var file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_serving_apis_get_model_metadata_proto_goTypes = []interface{}{
	(*SignatureDefMap)(nil),          // 0: tensorflow.serving.SignatureDefMap
	(*GetModelMetadataRequest)(nil),  // 1: tensorflow.serving.GetModelMetadataRequest
	(*GetModelMetadataResponse)(nil), // 2: tensorflow.serving.GetModelMetadataResponse
	nil,                              // 3: tensorflow.serving.SignatureDefMap.SignatureDefEntry
	nil,                              // 4: tensorflow.serving.GetModelMetadataResponse.MetadataEntry
	(*ModelSpec)(nil),                // 5: tensorflow.serving.ModelSpec
	(*protobuf.SignatureDef)(nil),    // 6: tensorflow.SignatureDef
	(*any.Any)(nil),                  // 7: google.protobuf.Any
}
var file_tensorflow_serving_apis_get_model_metadata_proto_depIdxs = []int32{
	3, // 0: tensorflow.serving.SignatureDefMap.signature_def:type_name -> tensorflow.serving.SignatureDefMap.SignatureDefEntry
	5, // 1: tensorflow.serving.GetModelMetadataRequest.model_spec:type_name -> tensorflow.serving.ModelSpec
	5, // 2: tensorflow.serving.GetModelMetadataResponse.model_spec:type_name -> tensorflow.serving.ModelSpec
	4, // 3: tensorflow.serving.GetModelMetadataResponse.metadata:type_name -> tensorflow.serving.GetModelMetadataResponse.MetadataEntry
	6, // 4: tensorflow.serving.SignatureDefMap.SignatureDefEntry.value:type_name -> tensorflow.SignatureDef
	7, // 5: tensorflow.serving.GetModelMetadataResponse.MetadataEntry.value:type_name -> google.protobuf.Any
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_get_model_metadata_proto_init() }
func file_tensorflow_serving_apis_get_model_metadata_proto_init() {
	if File_tensorflow_serving_apis_get_model_metadata_proto != nil {
		return
	}
	file_tensorflow_serving_apis_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureDefMap); i {
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
		file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelMetadataRequest); i {
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
		file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelMetadataResponse); i {
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
			RawDescriptor: file_tensorflow_serving_apis_get_model_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_get_model_metadata_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_get_model_metadata_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_get_model_metadata_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_get_model_metadata_proto = out.File
	file_tensorflow_serving_apis_get_model_metadata_proto_rawDesc = nil
	file_tensorflow_serving_apis_get_model_metadata_proto_goTypes = nil
	file_tensorflow_serving_apis_get_model_metadata_proto_depIdxs = nil
}
