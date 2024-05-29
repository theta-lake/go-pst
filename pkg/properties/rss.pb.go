// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright 2023 Marten Mooij
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate msgp -tests=false

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: rss.proto

package properties

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

type RSS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains the contents of the title field from the XML of the Atom feed or RSS channel.
	PostRssChannel *string `protobuf:"bytes,1,opt,name=post_rss_channel,json=postRssChannel,proto3,oneof" json:"post_rss_channel,omitempty" msg:"27136431,omitempty"`  
	// Contains the URL of the RSS or Atom feed from which the XML file came.
	PostRssChannelLink *string `protobuf:"bytes,2,opt,name=post_rss_channel_link,json=postRssChannelLink,proto3,oneof" json:"post_rss_channel_link,omitempty" msg:"27136031,omitempty"`  
	// Contains a unique identifier for the RSS object.
	PostRssItemGuid *string `protobuf:"bytes,3,opt,name=post_rss_item_guid,json=postRssItemGuid,proto3,oneof" json:"post_rss_item_guid,omitempty" msg:"27136331,omitempty"`  
	// Contains a hash of the feed XML computed by using an implementation-dependent algorithm.
	PostRssItemHash *int32 `protobuf:"varint,4,opt,name=post_rss_item_hash,json=postRssItemHash,proto3,oneof" json:"post_rss_item_hash,omitempty" msg:"2713623,omitempty"`  
	// Contains the URL of the link from an RSS or Atom item.
	PostRssItemLink *string `protobuf:"bytes,5,opt,name=post_rss_item_link,json=postRssItemLink,proto3,oneof" json:"post_rss_item_link,omitempty" msg:"27136131,omitempty"`  
	// Contains the item element and all of its sub-elements from an RSS feed, or the entry element and all of its sub-elements from an Atom feed.
	PostRssItemXml *string `protobuf:"bytes,6,opt,name=post_rss_item_xml,json=postRssItemXml,proto3,oneof" json:"post_rss_item_xml,omitempty" msg:"27136531,omitempty"`  
	// Contains the user's preferred name for the RSS or Atom subscription.
	PostRssSubscription *string `protobuf:"bytes,7,opt,name=post_rss_subscription,json=postRssSubscription,proto3,oneof" json:"post_rss_subscription,omitempty" msg:"27136631,omitempty"`  
}

func (x *RSS) Reset() {
	*x = RSS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RSS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RSS) ProtoMessage() {}

func (x *RSS) ProtoReflect() protoreflect.Message {
	mi := &file_rss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RSS.ProtoReflect.Descriptor instead.
func (*RSS) Descriptor() ([]byte, []int) {
	return file_rss_proto_rawDescGZIP(), []int{0}
}

func (x *RSS) GetPostRssChannel() string {
	if x != nil && x.PostRssChannel != nil {
		return *x.PostRssChannel
	}
	return ""
}

func (x *RSS) GetPostRssChannelLink() string {
	if x != nil && x.PostRssChannelLink != nil {
		return *x.PostRssChannelLink
	}
	return ""
}

func (x *RSS) GetPostRssItemGuid() string {
	if x != nil && x.PostRssItemGuid != nil {
		return *x.PostRssItemGuid
	}
	return ""
}

func (x *RSS) GetPostRssItemHash() int32 {
	if x != nil && x.PostRssItemHash != nil {
		return *x.PostRssItemHash
	}
	return 0
}

func (x *RSS) GetPostRssItemLink() string {
	if x != nil && x.PostRssItemLink != nil {
		return *x.PostRssItemLink
	}
	return ""
}

func (x *RSS) GetPostRssItemXml() string {
	if x != nil && x.PostRssItemXml != nil {
		return *x.PostRssItemXml
	}
	return ""
}

func (x *RSS) GetPostRssSubscription() string {
	if x != nil && x.PostRssSubscription != nil {
		return *x.PostRssSubscription
	}
	return ""
}

var File_rss_proto protoreflect.FileDescriptor

var file_rss_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x04, 0x0a, 0x03,
	0x52, 0x53, 0x53, 0x12, 0x2d, 0x0a, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0e, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x73, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x36, 0x0a, 0x15, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x73, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x12, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x73,
	0x73, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x12,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0f, 0x70, 0x6f, 0x73, 0x74,
	0x52, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x30,
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0f, 0x70, 0x6f,
	0x73, 0x74, 0x52, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x88, 0x01, 0x01,
	0x12, 0x2e, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x78, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0e, 0x70,
	0x6f, 0x73, 0x74, 0x52, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x58, 0x6d, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x37, 0x0a, 0x15, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x06, 0x52, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x73, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x18,
	0x0a, 0x16, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x42,
	0x15, 0x0a, 0x13, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x14, 0x0a,
	0x12, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x78, 0x6d, 0x6c, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x73, 0x73,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x2d, 0x6c, 0x61, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rss_proto_rawDescOnce sync.Once
	file_rss_proto_rawDescData = file_rss_proto_rawDesc
)

func file_rss_proto_rawDescGZIP() []byte {
	file_rss_proto_rawDescOnce.Do(func() {
		file_rss_proto_rawDescData = protoimpl.X.CompressGZIP(file_rss_proto_rawDescData)
	})
	return file_rss_proto_rawDescData
}

var file_rss_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rss_proto_goTypes = []interface{}{
	(*RSS)(nil), // 0: RSS
}
var file_rss_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rss_proto_init() }
func file_rss_proto_init() {
	if File_rss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RSS); i {
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
	file_rss_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rss_proto_goTypes,
		DependencyIndexes: file_rss_proto_depIdxs,
		MessageInfos:      file_rss_proto_msgTypes,
	}.Build()
	File_rss_proto = out.File
	file_rss_proto_rawDesc = nil
	file_rss_proto_goTypes = nil
	file_rss_proto_depIdxs = nil
}
