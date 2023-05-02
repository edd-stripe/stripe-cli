// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: samples_list.proto

package rpc

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

type SamplesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SamplesListRequest) Reset() {
	*x = SamplesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplesListRequest) ProtoMessage() {}

func (x *SamplesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_samples_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplesListRequest.ProtoReflect.Descriptor instead.
func (*SamplesListRequest) Descriptor() ([]byte, []int) {
	return file_samples_list_proto_rawDescGZIP(), []int{0}
}

type SamplesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of available Stripe samples
	Samples []*SamplesListResponse_SampleData `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *SamplesListResponse) Reset() {
	*x = SamplesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplesListResponse) ProtoMessage() {}

func (x *SamplesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samples_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplesListResponse.ProtoReflect.Descriptor instead.
func (*SamplesListResponse) Descriptor() ([]byte, []int) {
	return file_samples_list_proto_rawDescGZIP(), []int{1}
}

func (x *SamplesListResponse) GetSamples() []*SamplesListResponse_SampleData {
	if x != nil {
		return x.Samples
	}
	return nil
}

type SamplesListResponse_SampleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the sample, e.g. accept-a-card-payment
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URL of the repo, e.g. https://github.com/stripe-samples/accept-a-card-payment
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Description of the sample, e.g. Learn how to accept a basic card payment
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SamplesListResponse_SampleData) Reset() {
	*x = SamplesListResponse_SampleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplesListResponse_SampleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplesListResponse_SampleData) ProtoMessage() {}

func (x *SamplesListResponse_SampleData) ProtoReflect() protoreflect.Message {
	mi := &file_samples_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplesListResponse_SampleData.ProtoReflect.Descriptor instead.
func (*SamplesListResponse_SampleData) Descriptor() ([]byte, []int) {
	return file_samples_list_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SamplesListResponse_SampleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SamplesListResponse_SampleData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SamplesListResponse_SampleData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_samples_list_proto protoreflect.FileDescriptor

var file_samples_list_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xaa, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x1a, 0x54, 0x0a, 0x0a, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samples_list_proto_rawDescOnce sync.Once
	file_samples_list_proto_rawDescData = file_samples_list_proto_rawDesc
)

func file_samples_list_proto_rawDescGZIP() []byte {
	file_samples_list_proto_rawDescOnce.Do(func() {
		file_samples_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_samples_list_proto_rawDescData)
	})
	return file_samples_list_proto_rawDescData
}

var file_samples_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_samples_list_proto_goTypes = []interface{}{
	(*SamplesListRequest)(nil),             // 0: rpc.SamplesListRequest
	(*SamplesListResponse)(nil),            // 1: rpc.SamplesListResponse
	(*SamplesListResponse_SampleData)(nil), // 2: rpc.SamplesListResponse.SampleData
}
var file_samples_list_proto_depIdxs = []int32{
	2, // 0: rpc.SamplesListResponse.samples:type_name -> rpc.SamplesListResponse.SampleData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_samples_list_proto_init() }
func file_samples_list_proto_init() {
	if File_samples_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samples_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplesListRequest); i {
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
		file_samples_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplesListResponse); i {
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
		file_samples_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplesListResponse_SampleData); i {
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
			RawDescriptor: file_samples_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_samples_list_proto_goTypes,
		DependencyIndexes: file_samples_list_proto_depIdxs,
		MessageInfos:      file_samples_list_proto_msgTypes,
	}.Build()
	File_samples_list_proto = out.File
	file_samples_list_proto_rawDesc = nil
	file_samples_list_proto_goTypes = nil
	file_samples_list_proto_depIdxs = nil
}
