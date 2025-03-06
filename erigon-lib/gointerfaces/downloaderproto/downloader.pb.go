// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: downloader/downloader.proto

package downloaderproto

import (
	typesproto "github.com/erigontech/erigon-lib/gointerfaces/typesproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DownloadItem:
// - if Erigon created new snapshot and want seed it
// - if Erigon wnat download files - it fills only "torrent_hash" field
type AddItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Path          string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	TorrentHash   *typesproto.H160       `protobuf:"bytes,2,opt,name=torrent_hash,json=torrentHash,proto3" json:"torrent_hash,omitempty"` // will be resolved as magnet link
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItem) Reset() {
	*x = AddItem{}
	mi := &file_downloader_downloader_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItem) ProtoMessage() {}

func (x *AddItem) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItem.ProtoReflect.Descriptor instead.
func (*AddItem) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{0}
}

func (x *AddItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AddItem) GetTorrentHash() *typesproto.H160 {
	if x != nil {
		return x.TorrentHash
	}
	return nil
}

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*AddItem             `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // single hash will be resolved as magnet link
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{1}
}

func (x *AddRequest) GetItems() []*AddItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// DeleteRequest: stop seeding, delete file, delete .torrent
type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Paths         []string               `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type VerifyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{3}
}

type ProhibitNewDownloadsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProhibitNewDownloadsRequest) Reset() {
	*x = ProhibitNewDownloadsRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProhibitNewDownloadsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProhibitNewDownloadsRequest) ProtoMessage() {}

func (x *ProhibitNewDownloadsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProhibitNewDownloadsRequest.ProtoReflect.Descriptor instead.
func (*ProhibitNewDownloadsRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{4}
}

func (x *ProhibitNewDownloadsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// SetLogPrefixRequest: set downloader log prefix
type SetLogPrefixRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prefix        string                 `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetLogPrefixRequest) Reset() {
	*x = SetLogPrefixRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetLogPrefixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogPrefixRequest) ProtoMessage() {}

func (x *SetLogPrefixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogPrefixRequest.ProtoReflect.Descriptor instead.
func (*SetLogPrefixRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{5}
}

func (x *SetLogPrefixRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type CompletedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompletedRequest) Reset() {
	*x = CompletedRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedRequest) ProtoMessage() {}

func (x *CompletedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedRequest.ProtoReflect.Descriptor instead.
func (*CompletedRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{6}
}

// CompletedReply: return true if download is completed
type CompletedReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Completed     bool                   `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompletedReply) Reset() {
	*x = CompletedReply{}
	mi := &file_downloader_downloader_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedReply) ProtoMessage() {}

func (x *CompletedReply) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedReply.ProtoReflect.Descriptor instead.
func (*CompletedReply) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{7}
}

func (x *CompletedReply) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TorrentCompletedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TorrentCompletedRequest) Reset() {
	*x = TorrentCompletedRequest{}
	mi := &file_downloader_downloader_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TorrentCompletedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentCompletedRequest) ProtoMessage() {}

func (x *TorrentCompletedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentCompletedRequest.ProtoReflect.Descriptor instead.
func (*TorrentCompletedRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{8}
}

// Message: downloaded file data
type TorrentCompletedReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash          *typesproto.H160       `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TorrentCompletedReply) Reset() {
	*x = TorrentCompletedReply{}
	mi := &file_downloader_downloader_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TorrentCompletedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentCompletedReply) ProtoMessage() {}

func (x *TorrentCompletedReply) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentCompletedReply.ProtoReflect.Descriptor instead.
func (*TorrentCompletedReply) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{9}
}

func (x *TorrentCompletedReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TorrentCompletedReply) GetHash() *typesproto.H160 {
	if x != nil {
		return x.Hash
	}
	return nil
}

var File_downloader_downloader_proto protoreflect.FileDescriptor

var file_downloader_downloader_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x0c, 0x74, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x31, 0x36, 0x30, 0x52, 0x0b, 0x74, 0x6f, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x37, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x25, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x1b, 0x50, 0x72, 0x6f,
	0x68, 0x69, 0x62, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x13,
	0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x12, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2e, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0x19, 0x0a, 0x17, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x15, 0x54, 0x6f,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x31,
	0x36, 0x30, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x32, 0x90, 0x04, 0x0a, 0x0a, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x68, 0x69,
	0x62, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12,
	0x27, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x68, 0x69, 0x62, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1f, 0x2e, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x10, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x23, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x54,
	0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x2e,
	0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x3b, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_downloader_downloader_proto_rawDescOnce sync.Once
	file_downloader_downloader_proto_rawDescData []byte
)

func file_downloader_downloader_proto_rawDescGZIP() []byte {
	file_downloader_downloader_proto_rawDescOnce.Do(func() {
		file_downloader_downloader_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_downloader_downloader_proto_rawDesc), len(file_downloader_downloader_proto_rawDesc)))
	})
	return file_downloader_downloader_proto_rawDescData
}

var file_downloader_downloader_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_downloader_downloader_proto_goTypes = []any{
	(*AddItem)(nil),                     // 0: downloader.AddItem
	(*AddRequest)(nil),                  // 1: downloader.AddRequest
	(*DeleteRequest)(nil),               // 2: downloader.DeleteRequest
	(*VerifyRequest)(nil),               // 3: downloader.VerifyRequest
	(*ProhibitNewDownloadsRequest)(nil), // 4: downloader.ProhibitNewDownloadsRequest
	(*SetLogPrefixRequest)(nil),         // 5: downloader.SetLogPrefixRequest
	(*CompletedRequest)(nil),            // 6: downloader.CompletedRequest
	(*CompletedReply)(nil),              // 7: downloader.CompletedReply
	(*TorrentCompletedRequest)(nil),     // 8: downloader.TorrentCompletedRequest
	(*TorrentCompletedReply)(nil),       // 9: downloader.TorrentCompletedReply
	(*typesproto.H160)(nil),             // 10: types.H160
	(*emptypb.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_downloader_downloader_proto_depIdxs = []int32{
	10, // 0: downloader.AddItem.torrent_hash:type_name -> types.H160
	0,  // 1: downloader.AddRequest.items:type_name -> downloader.AddItem
	10, // 2: downloader.TorrentCompletedReply.hash:type_name -> types.H160
	4,  // 3: downloader.Downloader.ProhibitNewDownloads:input_type -> downloader.ProhibitNewDownloadsRequest
	1,  // 4: downloader.Downloader.Add:input_type -> downloader.AddRequest
	2,  // 5: downloader.Downloader.Delete:input_type -> downloader.DeleteRequest
	3,  // 6: downloader.Downloader.Verify:input_type -> downloader.VerifyRequest
	5,  // 7: downloader.Downloader.SetLogPrefix:input_type -> downloader.SetLogPrefixRequest
	6,  // 8: downloader.Downloader.Completed:input_type -> downloader.CompletedRequest
	8,  // 9: downloader.Downloader.TorrentCompleted:input_type -> downloader.TorrentCompletedRequest
	11, // 10: downloader.Downloader.ProhibitNewDownloads:output_type -> google.protobuf.Empty
	11, // 11: downloader.Downloader.Add:output_type -> google.protobuf.Empty
	11, // 12: downloader.Downloader.Delete:output_type -> google.protobuf.Empty
	11, // 13: downloader.Downloader.Verify:output_type -> google.protobuf.Empty
	11, // 14: downloader.Downloader.SetLogPrefix:output_type -> google.protobuf.Empty
	7,  // 15: downloader.Downloader.Completed:output_type -> downloader.CompletedReply
	9,  // 16: downloader.Downloader.TorrentCompleted:output_type -> downloader.TorrentCompletedReply
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_downloader_downloader_proto_init() }
func file_downloader_downloader_proto_init() {
	if File_downloader_downloader_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_downloader_downloader_proto_rawDesc), len(file_downloader_downloader_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_downloader_downloader_proto_goTypes,
		DependencyIndexes: file_downloader_downloader_proto_depIdxs,
		MessageInfos:      file_downloader_downloader_proto_msgTypes,
	}.Build()
	File_downloader_downloader_proto = out.File
	file_downloader_downloader_proto_goTypes = nil
	file_downloader_downloader_proto_depIdxs = nil
}
