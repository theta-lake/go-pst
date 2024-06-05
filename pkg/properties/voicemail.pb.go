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
// 	protoc        v5.27.0
// source: voicemail.proto

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

type Voicemail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains textual annotations to a voice message after it has been delivered to the user's mailbox.
	AudioNotes *string `protobuf:"bytes,1,opt,name=audio_notes,json=audioNotes,proto3,oneof" json:"audio_notes,omitempty"`
	// Contains a unique identifier associated with the phone call.
	XCallId *string `protobuf:"bytes,3,opt,name=x_call_id,json=xCallId,proto3,oneof" json:"x_call_id,omitempty"`
	// Specifies how many discrete pages are contained within an attachment representing a facsimile message.
	XFaxNumberOfPages *int32 `protobuf:"varint,4,opt,name=x_fax_number_of_pages,json=xFaxNumberOfPages,proto3,oneof" json:"x_fax_number_of_pages,omitempty"`
	// Indicates that the client only renders the message on a phone.
	XRequireProtectedPlayOnPhone *bool `protobuf:"varint,5,opt,name=x_require_protected_play_on_phone,json=xRequireProtectedPlayOnPhone,proto3,oneof" json:"x_require_protected_play_on_phone,omitempty"`
	// Contains the telephone number of the caller associated with a voice mail message.
	XSenderTelephoneNumber *string `protobuf:"bytes,6,opt,name=x_sender_telephone_number,json=xSenderTelephoneNumber,proto3,oneof" json:"x_sender_telephone_number,omitempty"`
	// Contains the list of names for the audio file attachments that are to be played as part of a message, in reverse order.
	XVoiceMessageAttachmentOrder *string `protobuf:"bytes,7,opt,name=x_voice_message_attachment_order,json=xVoiceMessageAttachmentOrder,proto3,oneof" json:"x_voice_message_attachment_order,omitempty"`
	// Specifies the length of the attached audio message, in seconds.
	XVoiceMessageDuration *int32 `protobuf:"varint,8,opt,name=x_voice_message_duration,json=xVoiceMessageDuration,proto3,oneof" json:"x_voice_message_duration,omitempty"`
	// Contains the name of the caller who left the attached voice message, as provided by the voice network's caller ID system.
	XVoiceMessageSenderName *string `protobuf:"bytes,9,opt,name=x_voice_message_sender_name,json=xVoiceMessageSenderName,proto3,oneof" json:"x_voice_message_sender_name,omitempty"`
	// Contains a unique identifier associated with the phone call.
	CallId *string `protobuf:"bytes,10,opt,name=call_id,json=callId,proto3,oneof" json:"call_id,omitempty" msg:"2663031,omitempty"`  
	// Contains the number of pages in a Fax object.
	FaxNumberOfPages *int32 `protobuf:"varint,11,opt,name=fax_number_of_pages,json=faxNumberOfPages,proto3,oneof" json:"fax_number_of_pages,omitempty" msg:"266283,omitempty"`  
	// Contains the telephone number of the caller associated with a voice mail message.
	SenderTelephoneNumber *string `protobuf:"bytes,12,opt,name=sender_telephone_number,json=senderTelephoneNumber,proto3,oneof" json:"sender_telephone_number,omitempty" msg:"2662631,omitempty"`  
	// Contains a list of file names for the audio file attachments that are to be played as part of a message.
	VoiceMessageAttachmentOrder *string `protobuf:"bytes,13,opt,name=voice_message_attachment_order,json=voiceMessageAttachmentOrder,proto3,oneof" json:"voice_message_attachment_order,omitempty" msg:"2662931,omitempty"`  
	// Specifies the length of the attached audio message, in seconds.
	VoiceMessageDuration *int32 `protobuf:"varint,14,opt,name=voice_message_duration,json=voiceMessageDuration,proto3,oneof" json:"voice_message_duration,omitempty" msg:"266253,omitempty"`  
	// Specifies the name of the caller who left the attached voice message, as provided by the voice network's caller ID system.
	VoiceMessageSenderName *string `protobuf:"bytes,15,opt,name=voice_message_sender_name,json=voiceMessageSenderName,proto3,oneof" json:"voice_message_sender_name,omitempty" msg:"2662731,omitempty"`  
}

func (x *Voicemail) Reset() {
	*x = Voicemail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voicemail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voicemail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voicemail) ProtoMessage() {}

func (x *Voicemail) ProtoReflect() protoreflect.Message {
	mi := &file_voicemail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voicemail.ProtoReflect.Descriptor instead.
func (*Voicemail) Descriptor() ([]byte, []int) {
	return file_voicemail_proto_rawDescGZIP(), []int{0}
}

func (x *Voicemail) GetAudioNotes() string {
	if x != nil && x.AudioNotes != nil {
		return *x.AudioNotes
	}
	return ""
}

func (x *Voicemail) GetXCallId() string {
	if x != nil && x.XCallId != nil {
		return *x.XCallId
	}
	return ""
}

func (x *Voicemail) GetXFaxNumberOfPages() int32 {
	if x != nil && x.XFaxNumberOfPages != nil {
		return *x.XFaxNumberOfPages
	}
	return 0
}

func (x *Voicemail) GetXRequireProtectedPlayOnPhone() bool {
	if x != nil && x.XRequireProtectedPlayOnPhone != nil {
		return *x.XRequireProtectedPlayOnPhone
	}
	return false
}

func (x *Voicemail) GetXSenderTelephoneNumber() string {
	if x != nil && x.XSenderTelephoneNumber != nil {
		return *x.XSenderTelephoneNumber
	}
	return ""
}

func (x *Voicemail) GetXVoiceMessageAttachmentOrder() string {
	if x != nil && x.XVoiceMessageAttachmentOrder != nil {
		return *x.XVoiceMessageAttachmentOrder
	}
	return ""
}

func (x *Voicemail) GetXVoiceMessageDuration() int32 {
	if x != nil && x.XVoiceMessageDuration != nil {
		return *x.XVoiceMessageDuration
	}
	return 0
}

func (x *Voicemail) GetXVoiceMessageSenderName() string {
	if x != nil && x.XVoiceMessageSenderName != nil {
		return *x.XVoiceMessageSenderName
	}
	return ""
}

func (x *Voicemail) GetCallId() string {
	if x != nil && x.CallId != nil {
		return *x.CallId
	}
	return ""
}

func (x *Voicemail) GetFaxNumberOfPages() int32 {
	if x != nil && x.FaxNumberOfPages != nil {
		return *x.FaxNumberOfPages
	}
	return 0
}

func (x *Voicemail) GetSenderTelephoneNumber() string {
	if x != nil && x.SenderTelephoneNumber != nil {
		return *x.SenderTelephoneNumber
	}
	return ""
}

func (x *Voicemail) GetVoiceMessageAttachmentOrder() string {
	if x != nil && x.VoiceMessageAttachmentOrder != nil {
		return *x.VoiceMessageAttachmentOrder
	}
	return ""
}

func (x *Voicemail) GetVoiceMessageDuration() int32 {
	if x != nil && x.VoiceMessageDuration != nil {
		return *x.VoiceMessageDuration
	}
	return 0
}

func (x *Voicemail) GetVoiceMessageSenderName() string {
	if x != nil && x.VoiceMessageSenderName != nil {
		return *x.VoiceMessageSenderName
	}
	return ""
}

var File_voicemail_proto protoreflect.FileDescriptor

var file_voicemail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb3, 0x09, 0x0a, 0x09, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x24, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4e, 0x6f, 0x74,
	0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x09, 0x78, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x78, 0x43, 0x61, 0x6c,
	0x6c, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x15, 0x78, 0x5f, 0x66, 0x61, 0x78, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x11, 0x78, 0x46, 0x61, 0x78, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a,
	0x21, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x1c, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x61,
	0x79, 0x4f, 0x6e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x19, 0x78,
	0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x16, 0x78, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x20, 0x78,
	0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x1c, 0x78, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x18, 0x78, 0x5f, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x15, 0x78, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x1b, 0x78, 0x5f, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x17, 0x78,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x06, 0x63, 0x61,
	0x6c, 0x6c, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x13, 0x66, 0x61, 0x78, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x09, 0x52, 0x10, 0x66, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x50, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x17, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x15,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x1e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x0b, 0x52, 0x1b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x39, 0x0a, 0x16, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x0c, 0x52, 0x14, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a,
	0x19, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x0d, 0x52, 0x16, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x78, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x42, 0x18, 0x0a, 0x16, 0x5f,
	0x78, 0x5f, 0x66, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x42, 0x24, 0x0a, 0x22, 0x5f, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x1c, 0x0a, 0x1a, 0x5f,
	0x78, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x23, 0x0a, 0x21, 0x5f, 0x78, 0x5f,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x1b,
	0x0a, 0x19, 0x5f, 0x78, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1e, 0x0a, 0x1c, 0x5f,
	0x78, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x66, 0x61, 0x78, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x42,
	0x1a, 0x0a, 0x18, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x21, 0x0a, 0x1f, 0x5f,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x19,
	0x0a, 0x17, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1c, 0x0a, 0x1a, 0x5f, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x2d, 0x6c, 0x61, 0x6b, 0x65,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voicemail_proto_rawDescOnce sync.Once
	file_voicemail_proto_rawDescData = file_voicemail_proto_rawDesc
)

func file_voicemail_proto_rawDescGZIP() []byte {
	file_voicemail_proto_rawDescOnce.Do(func() {
		file_voicemail_proto_rawDescData = protoimpl.X.CompressGZIP(file_voicemail_proto_rawDescData)
	})
	return file_voicemail_proto_rawDescData
}

var file_voicemail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_voicemail_proto_goTypes = []interface{}{
	(*Voicemail)(nil), // 0: Voicemail
}
var file_voicemail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_voicemail_proto_init() }
func file_voicemail_proto_init() {
	if File_voicemail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voicemail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voicemail); i {
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
	file_voicemail_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voicemail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_voicemail_proto_goTypes,
		DependencyIndexes: file_voicemail_proto_depIdxs,
		MessageInfos:      file_voicemail_proto_msgTypes,
	}.Build()
	File_voicemail_proto = out.File
	file_voicemail_proto_rawDesc = nil
	file_voicemail_proto_goTypes = nil
	file_voicemail_proto_depIdxs = nil
}
