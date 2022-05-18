// Input used in serving APIs.  Based on the tensorflow.Example family of
// feature representations.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: tensorflow_serving/apis/input.proto

package apis

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	example "github.com/Golang-Tools/tensorflow_service_apis/v2/tensorflow/core/example"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specifies one or more fully independent input Examples.
// See examples at:
//     https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/example/example.proto
type ExampleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Examples []*example.Example `protobuf:"bytes,1,rep,name=examples,proto3" json:"examples,omitempty"`
}

func (x *ExampleList) Reset() {
	*x = ExampleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleList) ProtoMessage() {}

func (x *ExampleList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleList.ProtoReflect.Descriptor instead.
func (*ExampleList) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_input_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleList) GetExamples() []*example.Example {
	if x != nil {
		return x.Examples
	}
	return nil
}

// Specifies one or more independent input Examples, with a common context
// Example.
//
// The common use case for context is to cleanly and optimally specify some
// features that are common across multiple examples.
//
// See example below with a search query as the context and multiple restaurants
// to perform some inference on.
//
// context: {
//   features: {
//     feature: {
//       key  : "query"
//       value: {
//         bytes_list: {
//           value: [ "pizza" ]
//         }
//       }
//     }
//   }
// }
// examples: {
//   features: {
//     feature: {
//       key  : "cuisine"
//       value: {
//         bytes_list: {
//           value: [ "Pizzeria" ]
//         }
//       }
//     }
//   }
// }
// examples: {
//   features: {
//     feature: {
//       key  : "cuisine"
//       value: {
//         bytes_list: {
//           value: [ "Taqueria" ]
//         }
//       }
//     }
//   }
// }
//
// Implementations of ExampleListWithContext merge the context Example into each
// of the Examples. Note that feature keys must not be duplicated between the
// Examples and context Example, or the behavior is undefined.
//
// See also:
//     tensorflow/core/example/example.proto
//     https://developers.google.com/protocol-buffers/docs/proto3#maps
type ExampleListWithContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Examples []*example.Example `protobuf:"bytes,1,rep,name=examples,proto3" json:"examples,omitempty"`
	Context  *example.Example   `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ExampleListWithContext) Reset() {
	*x = ExampleListWithContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleListWithContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleListWithContext) ProtoMessage() {}

func (x *ExampleListWithContext) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleListWithContext.ProtoReflect.Descriptor instead.
func (*ExampleListWithContext) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_input_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleListWithContext) GetExamples() []*example.Example {
	if x != nil {
		return x.Examples
	}
	return nil
}

func (x *ExampleListWithContext) GetContext() *example.Example {
	if x != nil {
		return x.Context
	}
	return nil
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Input_ExampleList
	//	*Input_ExampleListWithContext
	Kind isInput_Kind `protobuf_oneof:"kind"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_input_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_input_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_input_proto_rawDescGZIP(), []int{2}
}

func (m *Input) GetKind() isInput_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Input) GetExampleList() *ExampleList {
	if x, ok := x.GetKind().(*Input_ExampleList); ok {
		return x.ExampleList
	}
	return nil
}

func (x *Input) GetExampleListWithContext() *ExampleListWithContext {
	if x, ok := x.GetKind().(*Input_ExampleListWithContext); ok {
		return x.ExampleListWithContext
	}
	return nil
}

type isInput_Kind interface {
	isInput_Kind()
}

type Input_ExampleList struct {
	ExampleList *ExampleList `protobuf:"bytes,1,opt,name=example_list,json=exampleList,proto3,oneof"`
}

type Input_ExampleListWithContext struct {
	ExampleListWithContext *ExampleListWithContext `protobuf:"bytes,2,opt,name=example_list_with_context,json=exampleListWithContext,proto3,oneof"`
}

func (*Input_ExampleList) isInput_Kind() {}

func (*Input_ExampleListWithContext) isInput_Kind() {}

var File_tensorflow_serving_apis_input_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_input_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3e, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x22, 0x78, 0x0a, 0x16, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x05, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x48, 0x0a, 0x0c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x02, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6b,
	0x0a, 0x19, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x02, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x16, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x42, 0x1c, 0x5a, 0x17, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0xf8, 0x01,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_input_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_input_proto_rawDescData = file_tensorflow_serving_apis_input_proto_rawDesc
)

func file_tensorflow_serving_apis_input_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_input_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_input_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_input_proto_rawDescData
}

var file_tensorflow_serving_apis_input_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_serving_apis_input_proto_goTypes = []interface{}{
	(*ExampleList)(nil),            // 0: tensorflow.serving.ExampleList
	(*ExampleListWithContext)(nil), // 1: tensorflow.serving.ExampleListWithContext
	(*Input)(nil),                  // 2: tensorflow.serving.Input
	(*example.Example)(nil),        // 3: tensorflow.Example
}
var file_tensorflow_serving_apis_input_proto_depIdxs = []int32{
	3, // 0: tensorflow.serving.ExampleList.examples:type_name -> tensorflow.Example
	3, // 1: tensorflow.serving.ExampleListWithContext.examples:type_name -> tensorflow.Example
	3, // 2: tensorflow.serving.ExampleListWithContext.context:type_name -> tensorflow.Example
	0, // 3: tensorflow.serving.Input.example_list:type_name -> tensorflow.serving.ExampleList
	1, // 4: tensorflow.serving.Input.example_list_with_context:type_name -> tensorflow.serving.ExampleListWithContext
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_input_proto_init() }
func file_tensorflow_serving_apis_input_proto_init() {
	if File_tensorflow_serving_apis_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleList); i {
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
		file_tensorflow_serving_apis_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleListWithContext); i {
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
		file_tensorflow_serving_apis_input_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
	file_tensorflow_serving_apis_input_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Input_ExampleList)(nil),
		(*Input_ExampleListWithContext)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_input_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_input_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_input_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_input_proto = out.File
	file_tensorflow_serving_apis_input_proto_rawDesc = nil
	file_tensorflow_serving_apis_input_proto_goTypes = nil
	file_tensorflow_serving_apis_input_proto_depIdxs = nil
}
