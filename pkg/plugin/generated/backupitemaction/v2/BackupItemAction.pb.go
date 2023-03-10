// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: backupitemaction/v2/BackupItemAction.proto

package v2

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	generated "github.com/vmware-tanzu/velero/pkg/plugin/generated"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Item   []byte `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Backup []byte `protobuf:"bytes,3,opt,name=backup,proto3" json:"backup,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *ExecuteRequest) GetItem() []byte {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *ExecuteRequest) GetBackup() []byte {
	if x != nil {
		return x.Backup
	}
	return nil
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item            []byte                          `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	AdditionalItems []*generated.ResourceIdentifier `protobuf:"bytes,2,rep,name=additionalItems,proto3" json:"additionalItems,omitempty"`
	OperationID     string                          `protobuf:"bytes,3,opt,name=operationID,proto3" json:"operationID,omitempty"`
	ItemsToUpdate   []*generated.ResourceIdentifier `protobuf:"bytes,4,rep,name=itemsToUpdate,proto3" json:"itemsToUpdate,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteResponse) GetItem() []byte {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *ExecuteResponse) GetAdditionalItems() []*generated.ResourceIdentifier {
	if x != nil {
		return x.AdditionalItems
	}
	return nil
}

func (x *ExecuteResponse) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *ExecuteResponse) GetItemsToUpdate() []*generated.ResourceIdentifier {
	if x != nil {
		return x.ItemsToUpdate
	}
	return nil
}

type BackupItemActionAppliesToRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *BackupItemActionAppliesToRequest) Reset() {
	*x = BackupItemActionAppliesToRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupItemActionAppliesToRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupItemActionAppliesToRequest) ProtoMessage() {}

func (x *BackupItemActionAppliesToRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupItemActionAppliesToRequest.ProtoReflect.Descriptor instead.
func (*BackupItemActionAppliesToRequest) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{2}
}

func (x *BackupItemActionAppliesToRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

type BackupItemActionAppliesToResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceSelector *generated.ResourceSelector `protobuf:"bytes,1,opt,name=ResourceSelector,proto3" json:"ResourceSelector,omitempty"`
}

func (x *BackupItemActionAppliesToResponse) Reset() {
	*x = BackupItemActionAppliesToResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupItemActionAppliesToResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupItemActionAppliesToResponse) ProtoMessage() {}

func (x *BackupItemActionAppliesToResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupItemActionAppliesToResponse.ProtoReflect.Descriptor instead.
func (*BackupItemActionAppliesToResponse) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{3}
}

func (x *BackupItemActionAppliesToResponse) GetResourceSelector() *generated.ResourceSelector {
	if x != nil {
		return x.ResourceSelector
	}
	return nil
}

type BackupItemActionProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin      string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	OperationID string `protobuf:"bytes,2,opt,name=operationID,proto3" json:"operationID,omitempty"`
	Backup      []byte `protobuf:"bytes,3,opt,name=backup,proto3" json:"backup,omitempty"`
}

func (x *BackupItemActionProgressRequest) Reset() {
	*x = BackupItemActionProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupItemActionProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupItemActionProgressRequest) ProtoMessage() {}

func (x *BackupItemActionProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupItemActionProgressRequest.ProtoReflect.Descriptor instead.
func (*BackupItemActionProgressRequest) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{4}
}

func (x *BackupItemActionProgressRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *BackupItemActionProgressRequest) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *BackupItemActionProgressRequest) GetBackup() []byte {
	if x != nil {
		return x.Backup
	}
	return nil
}

type BackupItemActionProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress *generated.OperationProgress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *BackupItemActionProgressResponse) Reset() {
	*x = BackupItemActionProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupItemActionProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupItemActionProgressResponse) ProtoMessage() {}

func (x *BackupItemActionProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupItemActionProgressResponse.ProtoReflect.Descriptor instead.
func (*BackupItemActionProgressResponse) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{5}
}

func (x *BackupItemActionProgressResponse) GetProgress() *generated.OperationProgress {
	if x != nil {
		return x.Progress
	}
	return nil
}

type BackupItemActionCancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin      string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	OperationID string `protobuf:"bytes,2,opt,name=operationID,proto3" json:"operationID,omitempty"`
	Backup      []byte `protobuf:"bytes,3,opt,name=backup,proto3" json:"backup,omitempty"`
}

func (x *BackupItemActionCancelRequest) Reset() {
	*x = BackupItemActionCancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupItemActionCancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupItemActionCancelRequest) ProtoMessage() {}

func (x *BackupItemActionCancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backupitemaction_v2_BackupItemAction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupItemActionCancelRequest.ProtoReflect.Descriptor instead.
func (*BackupItemActionCancelRequest) Descriptor() ([]byte, []int) {
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP(), []int{6}
}

func (x *BackupItemActionCancelRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *BackupItemActionCancelRequest) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *BackupItemActionCancelRequest) GetBackup() []byte {
	if x != nil {
		return x.Backup
	}
	return nil
}

var File_backupitemaction_v2_BackupItemAction_proto protoreflect.FileDescriptor

var file_backupitemaction_v2_BackupItemAction_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x32,
	0x1a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x47, 0x0a, 0x0f, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x54, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x54, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x20, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x6c, 0x0a, 0x21, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49,
	0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x54, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x22, 0x73, 0x0a, 0x1f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65,
	0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x5c, 0x0a, 0x20, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x71, 0x0a, 0x1d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x32, 0xbc, 0x02, 0x0a, 0x10, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58,
	0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x12, 0x24, 0x2e, 0x76, 0x32,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65,
	0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x54, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x76, 0x32, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x21, 0x2e,
	0x76, 0x32, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x61,
	0x6e, 0x7a, 0x75, 0x2f, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backupitemaction_v2_BackupItemAction_proto_rawDescOnce sync.Once
	file_backupitemaction_v2_BackupItemAction_proto_rawDescData = file_backupitemaction_v2_BackupItemAction_proto_rawDesc
)

func file_backupitemaction_v2_BackupItemAction_proto_rawDescGZIP() []byte {
	file_backupitemaction_v2_BackupItemAction_proto_rawDescOnce.Do(func() {
		file_backupitemaction_v2_BackupItemAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_backupitemaction_v2_BackupItemAction_proto_rawDescData)
	})
	return file_backupitemaction_v2_BackupItemAction_proto_rawDescData
}

var file_backupitemaction_v2_BackupItemAction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_backupitemaction_v2_BackupItemAction_proto_goTypes = []interface{}{
	(*ExecuteRequest)(nil),                    // 0: v2.ExecuteRequest
	(*ExecuteResponse)(nil),                   // 1: v2.ExecuteResponse
	(*BackupItemActionAppliesToRequest)(nil),  // 2: v2.BackupItemActionAppliesToRequest
	(*BackupItemActionAppliesToResponse)(nil), // 3: v2.BackupItemActionAppliesToResponse
	(*BackupItemActionProgressRequest)(nil),   // 4: v2.BackupItemActionProgressRequest
	(*BackupItemActionProgressResponse)(nil),  // 5: v2.BackupItemActionProgressResponse
	(*BackupItemActionCancelRequest)(nil),     // 6: v2.BackupItemActionCancelRequest
	(*generated.ResourceIdentifier)(nil),      // 7: generated.ResourceIdentifier
	(*generated.ResourceSelector)(nil),        // 8: generated.ResourceSelector
	(*generated.OperationProgress)(nil),       // 9: generated.OperationProgress
	(*emptypb.Empty)(nil),                     // 10: google.protobuf.Empty
}
var file_backupitemaction_v2_BackupItemAction_proto_depIdxs = []int32{
	7,  // 0: v2.ExecuteResponse.additionalItems:type_name -> generated.ResourceIdentifier
	7,  // 1: v2.ExecuteResponse.itemsToUpdate:type_name -> generated.ResourceIdentifier
	8,  // 2: v2.BackupItemActionAppliesToResponse.ResourceSelector:type_name -> generated.ResourceSelector
	9,  // 3: v2.BackupItemActionProgressResponse.progress:type_name -> generated.OperationProgress
	2,  // 4: v2.BackupItemAction.AppliesTo:input_type -> v2.BackupItemActionAppliesToRequest
	0,  // 5: v2.BackupItemAction.Execute:input_type -> v2.ExecuteRequest
	4,  // 6: v2.BackupItemAction.Progress:input_type -> v2.BackupItemActionProgressRequest
	6,  // 7: v2.BackupItemAction.Cancel:input_type -> v2.BackupItemActionCancelRequest
	3,  // 8: v2.BackupItemAction.AppliesTo:output_type -> v2.BackupItemActionAppliesToResponse
	1,  // 9: v2.BackupItemAction.Execute:output_type -> v2.ExecuteResponse
	5,  // 10: v2.BackupItemAction.Progress:output_type -> v2.BackupItemActionProgressResponse
	10, // 11: v2.BackupItemAction.Cancel:output_type -> google.protobuf.Empty
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_backupitemaction_v2_BackupItemAction_proto_init() }
func file_backupitemaction_v2_BackupItemAction_proto_init() {
	if File_backupitemaction_v2_BackupItemAction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupItemActionAppliesToRequest); i {
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
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupItemActionAppliesToResponse); i {
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
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupItemActionProgressRequest); i {
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
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupItemActionProgressResponse); i {
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
		file_backupitemaction_v2_BackupItemAction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupItemActionCancelRequest); i {
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
			RawDescriptor: file_backupitemaction_v2_BackupItemAction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backupitemaction_v2_BackupItemAction_proto_goTypes,
		DependencyIndexes: file_backupitemaction_v2_BackupItemAction_proto_depIdxs,
		MessageInfos:      file_backupitemaction_v2_BackupItemAction_proto_msgTypes,
	}.Build()
	File_backupitemaction_v2_BackupItemAction_proto = out.File
	file_backupitemaction_v2_BackupItemAction_proto_rawDesc = nil
	file_backupitemaction_v2_BackupItemAction_proto_goTypes = nil
	file_backupitemaction_v2_BackupItemAction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BackupItemActionClient is the client API for BackupItemAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupItemActionClient interface {
	AppliesTo(ctx context.Context, in *BackupItemActionAppliesToRequest, opts ...grpc.CallOption) (*BackupItemActionAppliesToResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	Progress(ctx context.Context, in *BackupItemActionProgressRequest, opts ...grpc.CallOption) (*BackupItemActionProgressResponse, error)
	Cancel(ctx context.Context, in *BackupItemActionCancelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type backupItemActionClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupItemActionClient(cc grpc.ClientConnInterface) BackupItemActionClient {
	return &backupItemActionClient{cc}
}

func (c *backupItemActionClient) AppliesTo(ctx context.Context, in *BackupItemActionAppliesToRequest, opts ...grpc.CallOption) (*BackupItemActionAppliesToResponse, error) {
	out := new(BackupItemActionAppliesToResponse)
	err := c.cc.Invoke(ctx, "/v2.BackupItemAction/AppliesTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupItemActionClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/v2.BackupItemAction/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupItemActionClient) Progress(ctx context.Context, in *BackupItemActionProgressRequest, opts ...grpc.CallOption) (*BackupItemActionProgressResponse, error) {
	out := new(BackupItemActionProgressResponse)
	err := c.cc.Invoke(ctx, "/v2.BackupItemAction/Progress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupItemActionClient) Cancel(ctx context.Context, in *BackupItemActionCancelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/v2.BackupItemAction/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupItemActionServer is the server API for BackupItemAction service.
type BackupItemActionServer interface {
	AppliesTo(context.Context, *BackupItemActionAppliesToRequest) (*BackupItemActionAppliesToResponse, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	Progress(context.Context, *BackupItemActionProgressRequest) (*BackupItemActionProgressResponse, error)
	Cancel(context.Context, *BackupItemActionCancelRequest) (*emptypb.Empty, error)
}

// UnimplementedBackupItemActionServer can be embedded to have forward compatible implementations.
type UnimplementedBackupItemActionServer struct {
}

func (*UnimplementedBackupItemActionServer) AppliesTo(context.Context, *BackupItemActionAppliesToRequest) (*BackupItemActionAppliesToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliesTo not implemented")
}
func (*UnimplementedBackupItemActionServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (*UnimplementedBackupItemActionServer) Progress(context.Context, *BackupItemActionProgressRequest) (*BackupItemActionProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Progress not implemented")
}
func (*UnimplementedBackupItemActionServer) Cancel(context.Context, *BackupItemActionCancelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterBackupItemActionServer(s *grpc.Server, srv BackupItemActionServer) {
	s.RegisterService(&_BackupItemAction_serviceDesc, srv)
}

func _BackupItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupItemActionAppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.BackupItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).AppliesTo(ctx, req.(*BackupItemActionAppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.BackupItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupItemAction_Progress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupItemActionProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).Progress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.BackupItemAction/Progress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).Progress(ctx, req.(*BackupItemActionProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupItemAction_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupItemActionCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.BackupItemAction/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).Cancel(ctx, req.(*BackupItemActionCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2.BackupItemAction",
	HandlerType: (*BackupItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _BackupItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _BackupItemAction_Execute_Handler,
		},
		{
			MethodName: "Progress",
			Handler:    _BackupItemAction_Progress_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _BackupItemAction_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backupitemaction/v2/BackupItemAction.proto",
}
