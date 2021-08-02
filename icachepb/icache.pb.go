// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: icachepb/icache.proto

package icachepb

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icachepb_icache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_icachepb_icache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_icachepb_icache_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     []byte  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	MinuteQps float64 `protobuf:"fixed64,2,opt,name=minute_qps,json=minuteQps,proto3" json:"minute_qps,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icachepb_icache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_icachepb_icache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_icachepb_icache_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetResponse) GetMinuteQps() float64 {
	if x != nil {
		return x.MinuteQps
	}
	return 0
}

var File_icachepb_icache_proto protoreflect.FileDescriptor

var file_icachepb_icache_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2f, 0x69, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70,
	0x62, 0x22, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x5f, 0x71, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x51, 0x70, 0x73, 0x32, 0x3e, 0x0a, 0x06, 0x49,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x69,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_icachepb_icache_proto_rawDescOnce sync.Once
	file_icachepb_icache_proto_rawDescData = file_icachepb_icache_proto_rawDesc
)

func file_icachepb_icache_proto_rawDescGZIP() []byte {
	file_icachepb_icache_proto_rawDescOnce.Do(func() {
		file_icachepb_icache_proto_rawDescData = protoimpl.X.CompressGZIP(file_icachepb_icache_proto_rawDescData)
	})
	return file_icachepb_icache_proto_rawDescData
}

var file_icachepb_icache_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_icachepb_icache_proto_goTypes = []interface{}{
	(*GetRequest)(nil),  // 0: icachepb.GetRequest
	(*GetResponse)(nil), // 1: icachepb.GetResponse
}
var file_icachepb_icache_proto_depIdxs = []int32{
	0, // 0: icachepb.ICache.Get:input_type -> icachepb.GetRequest
	1, // 1: icachepb.ICache.Get:output_type -> icachepb.GetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_icachepb_icache_proto_init() }
func file_icachepb_icache_proto_init() {
	if File_icachepb_icache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_icachepb_icache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_icachepb_icache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
			RawDescriptor: file_icachepb_icache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_icachepb_icache_proto_goTypes,
		DependencyIndexes: file_icachepb_icache_proto_depIdxs,
		MessageInfos:      file_icachepb_icache_proto_msgTypes,
	}.Build()
	File_icachepb_icache_proto = out.File
	file_icachepb_icache_proto_rawDesc = nil
	file_icachepb_icache_proto_goTypes = nil
	file_icachepb_icache_proto_depIdxs = nil
}
