// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: Scraping.proto

package Scraping

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

type GetScrapingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetLink string `protobuf:"bytes,1,opt,name=target_link,json=targetLink,proto3" json:"target_link,omitempty"`
}

func (x *GetScrapingRequest) Reset() {
	*x = GetScrapingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Scraping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScrapingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScrapingRequest) ProtoMessage() {}

func (x *GetScrapingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Scraping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScrapingRequest.ProtoReflect.Descriptor instead.
func (*GetScrapingRequest) Descriptor() ([]byte, []int) {
	return file_Scraping_proto_rawDescGZIP(), []int{0}
}

func (x *GetScrapingRequest) GetTargetLink() string {
	if x != nil {
		return x.TargetLink
	}
	return ""
}

type ScrapingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Tell         string `protobuf:"bytes,3,opt,name=tell,proto3" json:"tell,omitempty"`
	OpeningHours string `protobuf:"bytes,4,opt,name=opening_hours,json=openingHours,proto3" json:"opening_hours,omitempty"`
}

func (x *ScrapingResponse) Reset() {
	*x = ScrapingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Scraping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapingResponse) ProtoMessage() {}

func (x *ScrapingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Scraping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapingResponse.ProtoReflect.Descriptor instead.
func (*ScrapingResponse) Descriptor() ([]byte, []int) {
	return file_Scraping_proto_rawDescGZIP(), []int{1}
}

func (x *ScrapingResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScrapingResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ScrapingResponse) GetTell() string {
	if x != nil {
		return x.Tell
	}
	return ""
}

func (x *ScrapingResponse) GetOpeningHours() string {
	if x != nil {
		return x.OpeningHours
	}
	return ""
}

var File_Scraping_proto protoreflect.FileDescriptor

var file_Scraping_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x79, 0x0a, 0x10, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x6c, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75,
	0x72, 0x73, 0x32, 0x43, 0x0a, 0x08, 0x53, 0x63, 0x72, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x37,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Scraping_proto_rawDescOnce sync.Once
	file_Scraping_proto_rawDescData = file_Scraping_proto_rawDesc
)

func file_Scraping_proto_rawDescGZIP() []byte {
	file_Scraping_proto_rawDescOnce.Do(func() {
		file_Scraping_proto_rawDescData = protoimpl.X.CompressGZIP(file_Scraping_proto_rawDescData)
	})
	return file_Scraping_proto_rawDescData
}

var file_Scraping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Scraping_proto_goTypes = []interface{}{
	(*GetScrapingRequest)(nil), // 0: GetScrapingRequest
	(*ScrapingResponse)(nil),   // 1: ScrapingResponse
}
var file_Scraping_proto_depIdxs = []int32{
	0, // 0: Scraping.GetShopInfo:input_type -> GetScrapingRequest
	1, // 1: Scraping.GetShopInfo:output_type -> ScrapingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Scraping_proto_init() }
func file_Scraping_proto_init() {
	if File_Scraping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Scraping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScrapingRequest); i {
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
		file_Scraping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapingResponse); i {
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
			RawDescriptor: file_Scraping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Scraping_proto_goTypes,
		DependencyIndexes: file_Scraping_proto_depIdxs,
		MessageInfos:      file_Scraping_proto_msgTypes,
	}.Build()
	File_Scraping_proto = out.File
	file_Scraping_proto_rawDesc = nil
	file_Scraping_proto_goTypes = nil
	file_Scraping_proto_depIdxs = nil
}
