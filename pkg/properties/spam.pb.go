// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright (C) 2022  Marten Mooij
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

//go:generate msgp -tests=false

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: spam.proto

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

type Spam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether email recipients are to be added to the safe senders list.
	JunkAddRecipientsToSafeSendersList *int32 `protobuf:"varint,2,opt,name=junk_add_recipients_to_safe_senders_list,json=junkAddRecipientsToSafeSendersList,proto3,oneof" json:"junk_add_recipients_to_safe_senders_list,omitempty" msg:"248353"` // @gotags: msg:"248353"
	// Indicates whether email addresses of the contacts in the Contacts folder are treated in a special way with respect to the spam filter.
	JunkIncludeContacts *int32 `protobuf:"varint,3,opt,name=junk_include_contacts,json=junkIncludeContacts,proto3,oneof" json:"junk_include_contacts,omitempty" msg:"248323"` // @gotags: msg:"248323"
	// Indicates whether messages identified as spam can be permanently deleted.
	JunkPermanentlyDelete *int32 `protobuf:"varint,4,opt,name=junk_permanently_delete,json=junkPermanentlyDelete,proto3,oneof" json:"junk_permanently_delete,omitempty" msg:"248343"` // @gotags: msg:"248343"
	// Indicated whether the phishing stamp on a message is to be ignored.
	JunkPhishingEnableLinks *bool `protobuf:"varint,5,opt,name=junk_phishing_enable_links,json=junkPhishingEnableLinks,proto3,oneof" json:"junk_phishing_enable_links,omitempty" msg:"2483911"` // @gotags: msg:"2483911"
	// Indicates how aggressively incoming email is to be sent to the Junk Email folder.
	JunkThreshold *int32 `protobuf:"varint,6,opt,name=junk_threshold,json=junkThreshold,proto3,oneof" json:"junk_threshold,omitempty" msg:"248333"` // @gotags: msg:"248333"
}

func (x *Spam) Reset() {
	*x = Spam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spam) ProtoMessage() {}

func (x *Spam) ProtoReflect() protoreflect.Message {
	mi := &file_spam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spam.ProtoReflect.Descriptor instead.
func (*Spam) Descriptor() ([]byte, []int) {
	return file_spam_proto_rawDescGZIP(), []int{0}
}

func (x *Spam) GetJunkAddRecipientsToSafeSendersList() int32 {
	if x != nil && x.JunkAddRecipientsToSafeSendersList != nil {
		return *x.JunkAddRecipientsToSafeSendersList
	}
	return 0
}

func (x *Spam) GetJunkIncludeContacts() int32 {
	if x != nil && x.JunkIncludeContacts != nil {
		return *x.JunkIncludeContacts
	}
	return 0
}

func (x *Spam) GetJunkPermanentlyDelete() int32 {
	if x != nil && x.JunkPermanentlyDelete != nil {
		return *x.JunkPermanentlyDelete
	}
	return 0
}

func (x *Spam) GetJunkPhishingEnableLinks() bool {
	if x != nil && x.JunkPhishingEnableLinks != nil {
		return *x.JunkPhishingEnableLinks
	}
	return false
}

func (x *Spam) GetJunkThreshold() int32 {
	if x != nil && x.JunkThreshold != nil {
		return *x.JunkThreshold
	}
	return 0
}

var File_spam_proto protoreflect.FileDescriptor

var file_spam_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x70, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x03, 0x0a,
	0x04, 0x53, 0x70, 0x61, 0x6d, 0x12, 0x59, 0x0a, 0x28, 0x6a, 0x75, 0x6e, 0x6b, 0x5f, 0x61, 0x64,
	0x64, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x73, 0x61, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x22, 0x6a, 0x75, 0x6e, 0x6b, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x6f, 0x53, 0x61,
	0x66, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x37, 0x0a, 0x15, 0x6a, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x01, 0x52, 0x13, 0x6a, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x17, 0x6a, 0x75, 0x6e,
	0x6b, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x15, 0x6a, 0x75,
	0x6e, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x1a, 0x6a, 0x75, 0x6e, 0x6b, 0x5f, 0x70,
	0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x17, 0x6a, 0x75,
	0x6e, 0x6b, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6a, 0x75, 0x6e, 0x6b,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x04, 0x52, 0x0d, 0x6a, 0x75, 0x6e, 0x6b, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x2b, 0x0a, 0x29, 0x5f, 0x6a, 0x75, 0x6e, 0x6b, 0x5f, 0x61, 0x64,
	0x64, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x73, 0x61, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x6a, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x5f,
	0x6a, 0x75, 0x6e, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x6a, 0x75, 0x6e, 0x6b,
	0x5f, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6a, 0x75, 0x6e, 0x6b, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f, 0x69, 0x6a, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spam_proto_rawDescOnce sync.Once
	file_spam_proto_rawDescData = file_spam_proto_rawDesc
)

func file_spam_proto_rawDescGZIP() []byte {
	file_spam_proto_rawDescOnce.Do(func() {
		file_spam_proto_rawDescData = protoimpl.X.CompressGZIP(file_spam_proto_rawDescData)
	})
	return file_spam_proto_rawDescData
}

var file_spam_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spam_proto_goTypes = []interface{}{
	(*Spam)(nil), // 0: Spam
}
var file_spam_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spam_proto_init() }
func file_spam_proto_init() {
	if File_spam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spam); i {
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
	file_spam_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spam_proto_goTypes,
		DependencyIndexes: file_spam_proto_depIdxs,
		MessageInfos:      file_spam_proto_msgTypes,
	}.Build()
	File_spam_proto = out.File
	file_spam_proto_rawDesc = nil
	file_spam_proto_goTypes = nil
	file_spam_proto_depIdxs = nil
}
