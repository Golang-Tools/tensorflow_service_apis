// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: tensorflow_serving/apis/regression.proto

package tensorflow_serving

import (
	proto "github.com/golang/protobuf/proto"
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

// Regression result for a single item (tensorflow.Example).
type Regression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Regression) Reset() {
	*x = Regression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regression) ProtoMessage() {}

func (x *Regression) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regression.ProtoReflect.Descriptor instead.
func (*Regression) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_regression_proto_rawDescGZIP(), []int{0}
}

func (x *Regression) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Contains one result per input example, in the same order as the input in
// RegressionRequest.
type RegressionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regressions []*Regression `protobuf:"bytes,1,rep,name=regressions,proto3" json:"regressions,omitempty"`
}

func (x *RegressionResult) Reset() {
	*x = RegressionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegressionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegressionResult) ProtoMessage() {}

func (x *RegressionResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegressionResult.ProtoReflect.Descriptor instead.
func (*RegressionResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_regression_proto_rawDescGZIP(), []int{1}
}

func (x *RegressionResult) GetRegressions() []*Regression {
	if x != nil {
		return x.Regressions
	}
	return nil
}

type RegressionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input data.
	Input *Input `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *RegressionRequest) Reset() {
	*x = RegressionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegressionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegressionRequest) ProtoMessage() {}

func (x *RegressionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegressionRequest.ProtoReflect.Descriptor instead.
func (*RegressionRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_regression_proto_rawDescGZIP(), []int{2}
}

func (x *RegressionRequest) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *RegressionRequest) GetInput() *Input {
	if x != nil {
		return x.Input
	}
	return nil
}

type RegressionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective Model Specification used for regression.
	ModelSpec *ModelSpec        `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	Result    *RegressionResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RegressionResponse) Reset() {
	*x = RegressionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegressionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegressionResponse) ProtoMessage() {}

func (x *RegressionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_regression_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegressionResponse.ProtoReflect.Descriptor instead.
func (*RegressionResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_regression_proto_rawDescGZIP(), []int{3}
}

func (x *RegressionResponse) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *RegressionResponse) GetResult() *RegressionResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_tensorflow_serving_apis_regression_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_regression_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x23,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x40, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x22, 0x5a, 0x1d, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_regression_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_regression_proto_rawDescData = file_tensorflow_serving_apis_regression_proto_rawDesc
)

func file_tensorflow_serving_apis_regression_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_regression_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_regression_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_regression_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_regression_proto_rawDescData
}

var file_tensorflow_serving_apis_regression_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_serving_apis_regression_proto_goTypes = []interface{}{
	(*Regression)(nil),         // 0: tensorflow.serving.Regression
	(*RegressionResult)(nil),   // 1: tensorflow.serving.RegressionResult
	(*RegressionRequest)(nil),  // 2: tensorflow.serving.RegressionRequest
	(*RegressionResponse)(nil), // 3: tensorflow.serving.RegressionResponse
	(*ModelSpec)(nil),          // 4: tensorflow.serving.ModelSpec
	(*Input)(nil),              // 5: tensorflow.serving.Input
}
var file_tensorflow_serving_apis_regression_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.RegressionResult.regressions:type_name -> tensorflow.serving.Regression
	4, // 1: tensorflow.serving.RegressionRequest.model_spec:type_name -> tensorflow.serving.ModelSpec
	5, // 2: tensorflow.serving.RegressionRequest.input:type_name -> tensorflow.serving.Input
	4, // 3: tensorflow.serving.RegressionResponse.model_spec:type_name -> tensorflow.serving.ModelSpec
	1, // 4: tensorflow.serving.RegressionResponse.result:type_name -> tensorflow.serving.RegressionResult
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_regression_proto_init() }
func file_tensorflow_serving_apis_regression_proto_init() {
	if File_tensorflow_serving_apis_regression_proto != nil {
		return
	}
	file_tensorflow_serving_apis_input_proto_init()
	file_tensorflow_serving_apis_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_regression_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regression); i {
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
		file_tensorflow_serving_apis_regression_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegressionResult); i {
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
		file_tensorflow_serving_apis_regression_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegressionRequest); i {
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
		file_tensorflow_serving_apis_regression_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegressionResponse); i {
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
			RawDescriptor: file_tensorflow_serving_apis_regression_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_regression_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_regression_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_regression_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_regression_proto = out.File
	file_tensorflow_serving_apis_regression_proto_rawDesc = nil
	file_tensorflow_serving_apis_regression_proto_goTypes = nil
	file_tensorflow_serving_apis_regression_proto_depIdxs = nil
}
