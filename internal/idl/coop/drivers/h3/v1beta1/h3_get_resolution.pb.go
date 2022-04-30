// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: coop/drivers/h3/v1beta1/h3_get_resolution.proto

package v1beta1

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

type H3GetResolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *H3GetResolutionRequest) Reset() {
	*x = H3GetResolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H3GetResolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H3GetResolutionRequest) ProtoMessage() {}

func (x *H3GetResolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H3GetResolutionRequest.ProtoReflect.Descriptor instead.
func (*H3GetResolutionRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescGZIP(), []int{0}
}

func (x *H3GetResolutionRequest) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

type H3GetResolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resolution int32 `protobuf:"varint,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
}

func (x *H3GetResolutionResponse) Reset() {
	*x = H3GetResolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H3GetResolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H3GetResolutionResponse) ProtoMessage() {}

func (x *H3GetResolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H3GetResolutionResponse.ProtoReflect.Descriptor instead.
func (*H3GetResolutionResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescGZIP(), []int{1}
}

func (x *H3GetResolutionResponse) GetResolution() int32 {
	if x != nil {
		return x.Resolution
	}
	return 0
}

var File_coop_drivers_h3_v1beta1_h3_get_resolution_proto protoreflect.FileDescriptor

var file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x68,
	0x33, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x68, 0x33, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x33, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x2e, 0x0a, 0x16, 0x48, 0x33,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x39, 0x0a, 0x17, 0x48, 0x33,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69, 0x6e, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c,
	0x63, 0x68, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x65, 0x6f, 0x2d, 0x68, 0x33, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6f,
	0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x33, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescOnce sync.Once
	file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescData = file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDesc
)

func file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescGZIP() []byte {
	file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescOnce.Do(func() {
		file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescData = protoimpl.X.CompressGZIP(file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescData)
	})
	return file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDescData
}

var file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_goTypes = []interface{}{
	(*H3GetResolutionRequest)(nil),  // 0: coop.drivers.h3.v1beta1.H3GetResolutionRequest
	(*H3GetResolutionResponse)(nil), // 1: coop.drivers.h3.v1beta1.H3GetResolutionResponse
}
var file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_init() }
func file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_init() {
	if File_coop_drivers_h3_v1beta1_h3_get_resolution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H3GetResolutionRequest); i {
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
		file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H3GetResolutionResponse); i {
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
			RawDescriptor: file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_goTypes,
		DependencyIndexes: file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_depIdxs,
		MessageInfos:      file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_msgTypes,
	}.Build()
	File_coop_drivers_h3_v1beta1_h3_get_resolution_proto = out.File
	file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_rawDesc = nil
	file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_goTypes = nil
	file_coop_drivers_h3_v1beta1_h3_get_resolution_proto_depIdxs = nil
}