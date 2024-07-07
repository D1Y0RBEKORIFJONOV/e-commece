// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.5
// source: protos/soldiers/soldiers.proto

package soldiers1

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

type DeleteSoldierReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoldersId    string `protobuf:"bytes,1,opt,name=solders_id,json=soldersId,proto3" json:"solders_id,omitempty"`
	IsHardDelete bool   `protobuf:"varint,2,opt,name=is_hard_delete,json=isHardDelete,proto3" json:"is_hard_delete,omitempty"`
}

func (x *DeleteSoldierReq) Reset() {
	*x = DeleteSoldierReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSoldierReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSoldierReq) ProtoMessage() {}

func (x *DeleteSoldierReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSoldierReq.ProtoReflect.Descriptor instead.
func (*DeleteSoldierReq) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteSoldierReq) GetSoldersId() string {
	if x != nil {
		return x.SoldersId
	}
	return ""
}

func (x *DeleteSoldierReq) GetIsHardDelete() bool {
	if x != nil {
		return x.IsHardDelete
	}
	return false
}

type GetAllSoldierReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filed string `protobuf:"bytes,1,opt,name=filed,proto3" json:"filed,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Page  int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllSoldierReq) Reset() {
	*x = GetAllSoldierReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSoldierReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSoldierReq) ProtoMessage() {}

func (x *GetAllSoldierReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSoldierReq.ProtoReflect.Descriptor instead.
func (*GetAllSoldierReq) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllSoldierReq) GetFiled() string {
	if x != nil {
		return x.Filed
	}
	return ""
}

func (x *GetAllSoldierReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetAllSoldierReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllSoldierReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type UpdateSoldierReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoldersId string `protobuf:"bytes,1,opt,name=solders_id,json=soldersId,proto3" json:"solders_id,omitempty"`
	FnName    string `protobuf:"bytes,2,opt,name=fnName,proto3" json:"fnName,omitempty"`
	LnName    string `protobuf:"bytes,3,opt,name=lnName,proto3" json:"lnName,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	BirthDay  string `protobuf:"bytes,5,opt,name=birth_day,json=birthDay,proto3" json:"birth_day,omitempty"`
}

func (x *UpdateSoldierReq) Reset() {
	*x = UpdateSoldierReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSoldierReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSoldierReq) ProtoMessage() {}

func (x *UpdateSoldierReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSoldierReq.ProtoReflect.Descriptor instead.
func (*UpdateSoldierReq) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSoldierReq) GetSoldersId() string {
	if x != nil {
		return x.SoldersId
	}
	return ""
}

func (x *UpdateSoldierReq) GetFnName() string {
	if x != nil {
		return x.FnName
	}
	return ""
}

func (x *UpdateSoldierReq) GetLnName() string {
	if x != nil {
		return x.LnName
	}
	return ""
}

func (x *UpdateSoldierReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateSoldierReq) GetBirthDay() string {
	if x != nil {
		return x.BirthDay
	}
	return ""
}

type GetSoldierRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Soldiers []*Soldiers `protobuf:"bytes,1,rep,name=soldiers,proto3" json:"soldiers,omitempty"`
	Count    int64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetSoldierRequestResponse) Reset() {
	*x = GetSoldierRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSoldierRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSoldierRequestResponse) ProtoMessage() {}

func (x *GetSoldierRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSoldierRequestResponse.ProtoReflect.Descriptor instead.
func (*GetSoldierRequestResponse) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{3}
}

func (x *GetSoldierRequestResponse) GetSoldiers() []*Soldiers {
	if x != nil {
		return x.Soldiers
	}
	return nil
}

func (x *GetSoldierRequestResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetSoldierReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filed string `protobuf:"bytes,1,opt,name=filed,proto3" json:"filed,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetSoldierReq) Reset() {
	*x = GetSoldierReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSoldierReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSoldierReq) ProtoMessage() {}

func (x *GetSoldierReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSoldierReq.ProtoReflect.Descriptor instead.
func (*GetSoldierReq) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{4}
}

func (x *GetSoldierReq) GetFiled() string {
	if x != nil {
		return x.Filed
	}
	return ""
}

func (x *GetSoldierReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type LogerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogerResponse) Reset() {
	*x = LogerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogerResponse) ProtoMessage() {}

func (x *LogerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogerResponse.ProtoReflect.Descriptor instead.
func (*LogerResponse) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{5}
}

func (x *LogerResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RegisterAndLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	SecredCode string `protobuf:"bytes,2,opt,name=secred_code,json=secredCode,proto3" json:"secred_code,omitempty"`
}

func (x *RegisterAndLoginReq) Reset() {
	*x = RegisterAndLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAndLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAndLoginReq) ProtoMessage() {}

func (x *RegisterAndLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAndLoginReq.ProtoReflect.Descriptor instead.
func (*RegisterAndLoginReq) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterAndLoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterAndLoginReq) GetSecredCode() string {
	if x != nil {
		return x.SecredCode
	}
	return ""
}

type CreateSoldiersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FnName   string `protobuf:"bytes,1,opt,name=fnName,proto3" json:"fnName,omitempty"`
	LnName   string `protobuf:"bytes,2,opt,name=lnName,proto3" json:"lnName,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	BirhtDay string `protobuf:"bytes,5,opt,name=birht_day,json=birhtDay,proto3" json:"birht_day,omitempty"`
}

func (x *CreateSoldiersReq) Reset() {
	*x = CreateSoldiersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSoldiersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSoldiersReq) ProtoMessage() {}

func (x *CreateSoldiersReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSoldiersReq.ProtoReflect.Descriptor instead.
func (*CreateSoldiersReq) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{8}
}

func (x *CreateSoldiersReq) GetFnName() string {
	if x != nil {
		return x.FnName
	}
	return ""
}

func (x *CreateSoldiersReq) GetLnName() string {
	if x != nil {
		return x.LnName
	}
	return ""
}

func (x *CreateSoldiersReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateSoldiersReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateSoldiersReq) GetBirhtDay() string {
	if x != nil {
		return x.BirhtDay
	}
	return ""
}

type Soldiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FnName    string `protobuf:"bytes,2,opt,name=fnName,proto3" json:"fnName,omitempty"`
	LnName    string `protobuf:"bytes,3,opt,name=lnName,proto3" json:"lnName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	BirhtDay  string `protobuf:"bytes,6,opt,name=birht_day,json=birhtDay,proto3" json:"birht_day,omitempty"`
	DeletedAt string `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	JoinedAt  string `protobuf:"bytes,8,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Soldiers) Reset() {
	*x = Soldiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_soldiers_soldiers_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Soldiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Soldiers) ProtoMessage() {}

func (x *Soldiers) ProtoReflect() protoreflect.Message {
	mi := &file_protos_soldiers_soldiers_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Soldiers.ProtoReflect.Descriptor instead.
func (*Soldiers) Descriptor() ([]byte, []int) {
	return file_protos_soldiers_soldiers_proto_rawDescGZIP(), []int{9}
}

func (x *Soldiers) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Soldiers) GetFnName() string {
	if x != nil {
		return x.FnName
	}
	return ""
}

func (x *Soldiers) GetLnName() string {
	if x != nil {
		return x.LnName
	}
	return ""
}

func (x *Soldiers) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Soldiers) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Soldiers) GetBirhtDay() string {
	if x != nil {
		return x.BirhtDay
	}
	return ""
}

func (x *Soldiers) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Soldiers) GetJoinedAt() string {
	if x != nil {
		return x.JoinedAt
	}
	return ""
}

func (x *Soldiers) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Soldiers) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_protos_soldiers_soldiers_proto protoreflect.FileDescriptor

var file_protos_soldiers_soldiers_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72,
	0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x48,
	0x61, 0x72, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x68, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79,
	0x22, 0x58, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x64,
	0x69, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4c, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e,
	0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x92, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x68,
	0x74, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72,
	0x68, 0x74, 0x44, 0x61, 0x79, 0x22, 0x93, 0x02, 0x0a, 0x08, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65,
	0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x68, 0x74, 0x5f, 0x64, 0x61,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x68, 0x74, 0x44, 0x61,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xe6, 0x02, 0x0a, 0x0f,
	0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x12,
	0x2d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x4c, 0x6f, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x53,
	0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x53,
	0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x64, 0x69, 0x79, 0x6f, 0x72, 0x62, 0x65, 0x6b,
	0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x6f, 0x6c,
	0x64, 0x69, 0x65, 0x72, 0x73, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_soldiers_soldiers_proto_rawDescOnce sync.Once
	file_protos_soldiers_soldiers_proto_rawDescData = file_protos_soldiers_soldiers_proto_rawDesc
)

func file_protos_soldiers_soldiers_proto_rawDescGZIP() []byte {
	file_protos_soldiers_soldiers_proto_rawDescOnce.Do(func() {
		file_protos_soldiers_soldiers_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_soldiers_soldiers_proto_rawDescData)
	})
	return file_protos_soldiers_soldiers_proto_rawDescData
}

var file_protos_soldiers_soldiers_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_soldiers_soldiers_proto_goTypes = []interface{}{
	(*DeleteSoldierReq)(nil),          // 0: DeleteSoldierReq
	(*GetAllSoldierReq)(nil),          // 1: GetAllSoldierReq
	(*UpdateSoldierReq)(nil),          // 2: UpdateSoldierReq
	(*GetSoldierRequestResponse)(nil), // 3: GetSoldierRequestResponse
	(*GetSoldierReq)(nil),             // 4: GetSoldierReq
	(*LogerResponse)(nil),             // 5: LogerResponse
	(*Status)(nil),                    // 6: Status
	(*RegisterAndLoginReq)(nil),       // 7: RegisterAndLoginReq
	(*CreateSoldiersReq)(nil),         // 8: CreateSoldiersReq
	(*Soldiers)(nil),                  // 9: Soldiers
}
var file_protos_soldiers_soldiers_proto_depIdxs = []int32{
	9, // 0: GetSoldierRequestResponse.soldiers:type_name -> Soldiers
	8, // 1: SoldiersService.CreateSoldiers:input_type -> CreateSoldiersReq
	7, // 2: SoldiersService.RegisterUser:input_type -> RegisterAndLoginReq
	7, // 3: SoldiersService.Login:input_type -> RegisterAndLoginReq
	4, // 4: SoldiersService.GetSoldier:input_type -> GetSoldierReq
	1, // 5: SoldiersService.GetAllSoldiers:input_type -> GetAllSoldierReq
	2, // 6: SoldiersService.UpdateSoldier:input_type -> UpdateSoldierReq
	0, // 7: SoldiersService.DeleteSoldier:input_type -> DeleteSoldierReq
	6, // 8: SoldiersService.CreateSoldiers:output_type -> Status
	9, // 9: SoldiersService.RegisterUser:output_type -> Soldiers
	5, // 10: SoldiersService.Login:output_type -> LogerResponse
	9, // 11: SoldiersService.GetSoldier:output_type -> Soldiers
	3, // 12: SoldiersService.GetAllSoldiers:output_type -> GetSoldierRequestResponse
	9, // 13: SoldiersService.UpdateSoldier:output_type -> Soldiers
	6, // 14: SoldiersService.DeleteSoldier:output_type -> Status
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_soldiers_soldiers_proto_init() }
func file_protos_soldiers_soldiers_proto_init() {
	if File_protos_soldiers_soldiers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_soldiers_soldiers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSoldierReq); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSoldierReq); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSoldierReq); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSoldierRequestResponse); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSoldierReq); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogerResponse); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAndLoginReq); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSoldiersReq); i {
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
		file_protos_soldiers_soldiers_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Soldiers); i {
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
			RawDescriptor: file_protos_soldiers_soldiers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_soldiers_soldiers_proto_goTypes,
		DependencyIndexes: file_protos_soldiers_soldiers_proto_depIdxs,
		MessageInfos:      file_protos_soldiers_soldiers_proto_msgTypes,
	}.Build()
	File_protos_soldiers_soldiers_proto = out.File
	file_protos_soldiers_soldiers_proto_rawDesc = nil
	file_protos_soldiers_soldiers_proto_goTypes = nil
	file_protos_soldiers_soldiers_proto_depIdxs = nil
}
