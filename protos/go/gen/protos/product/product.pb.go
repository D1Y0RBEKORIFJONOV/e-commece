// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.5
// source: protos/product/product.proto

package product1

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

type AddProductToUserBasketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddProductToUserBasketReq) Reset() {
	*x = AddProductToUserBasketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductToUserBasketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductToUserBasketReq) ProtoMessage() {}

func (x *AddProductToUserBasketReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductToUserBasketReq.ProtoReflect.Descriptor instead.
func (*AddProductToUserBasketReq) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *AddProductToUserBasketReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddProductToUserBasketReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *Empty) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *DeleteProductReq) Reset() {
	*x = DeleteProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductReq) ProtoMessage() {}

func (x *DeleteProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductReq.ProtoReflect.Descriptor instead.
func (*DeleteProductReq) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteProductReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName        string `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductDescription string `protobuf:"bytes,2,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	Price              string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Category           string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProductRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *CreateProductRequest) GetProductDescription() string {
	if x != nil {
		return x.ProductDescription
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId          string `protobuf:"bytes,6,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName        string `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductDescription string `protobuf:"bytes,2,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	Price              string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Category           string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	CreateAt           string `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Product) GetProductDescription() string {
	if x != nil {
		return x.ProductDescription
	}
	return ""
}

func (x *Product) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type GetProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *GetProductReq) Reset() {
	*x = GetProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReq) ProtoMessage() {}

func (x *GetProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReq.ProtoReflect.Descriptor instead.
func (*GetProductReq) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type GetAllProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Page  int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllProductReq) Reset() {
	*x = GetAllProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductReq) ProtoMessage() {}

func (x *GetAllProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductReq.ProtoReflect.Descriptor instead.
func (*GetAllProductReq) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllProductReq) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *GetAllProductReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetAllProductReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllProductReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllProductRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetAllProductRes) Reset() {
	*x = GetAllProductRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductRes) ProtoMessage() {}

func (x *GetAllProductRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductRes.ProtoReflect.Descriptor instead.
func (*GetAllProductRes) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllProductRes) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_protos_product_product_proto protoreflect.FileDescriptor

var file_protos_product_product_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53,
	0x0a, 0x19, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xcb, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x2e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x38,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xb9, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x26, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x08,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x30, 0x01, 0x12, 0x2a, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x12, 0x1a, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x1e, 0x5a, 0x1c, 0x64, 0x69, 0x79, 0x6f, 0x72, 0x62, 0x65, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_product_product_proto_rawDescOnce sync.Once
	file_protos_product_product_proto_rawDescData = file_protos_product_product_proto_rawDesc
)

func file_protos_product_product_proto_rawDescGZIP() []byte {
	file_protos_product_product_proto_rawDescOnce.Do(func() {
		file_protos_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_product_product_proto_rawDescData)
	})
	return file_protos_product_product_proto_rawDescData
}

var file_protos_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_product_product_proto_goTypes = []any{
	(*AddProductToUserBasketReq)(nil), // 0: AddProductToUserBasketReq
	(*Empty)(nil),                     // 1: Empty
	(*DeleteProductReq)(nil),          // 2: DeleteProductReq
	(*CreateProductRequest)(nil),      // 3: CreateProductRequest
	(*Product)(nil),                   // 4: Product
	(*GetProductReq)(nil),             // 5: GetProductReq
	(*GetAllProductReq)(nil),          // 6: GetAllProductReq
	(*GetAllProductRes)(nil),          // 7: GetAllProductRes
}
var file_protos_product_product_proto_depIdxs = []int32{
	4, // 0: GetAllProductRes.products:type_name -> Product
	3, // 1: ProductService.CreateProduct:input_type -> CreateProductRequest
	5, // 2: ProductService.GetProduct:input_type -> GetProductReq
	6, // 3: ProductService.GetAllProduct:input_type -> GetAllProductReq
	6, // 4: ProductService.ListProduct:input_type -> GetAllProductReq
	2, // 5: ProductService.DeleteProduct:input_type -> DeleteProductReq
	0, // 6: ProductService.AddProductToUserBasket:input_type -> AddProductToUserBasketReq
	4, // 7: ProductService.CreateProduct:output_type -> Product
	4, // 8: ProductService.GetProduct:output_type -> Product
	7, // 9: ProductService.GetAllProduct:output_type -> GetAllProductRes
	4, // 10: ProductService.ListProduct:output_type -> Product
	1, // 11: ProductService.DeleteProduct:output_type -> Empty
	1, // 12: ProductService.AddProductToUserBasket:output_type -> Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_product_product_proto_init() }
func file_protos_product_product_proto_init() {
	if File_protos_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_product_product_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddProductToUserBasketReq); i {
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
		file_protos_product_product_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_protos_product_product_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProductReq); i {
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
		file_protos_product_product_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProductRequest); i {
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
		file_protos_product_product_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Product); i {
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
		file_protos_product_product_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetProductReq); i {
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
		file_protos_product_product_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllProductReq); i {
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
		file_protos_product_product_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllProductRes); i {
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
			RawDescriptor: file_protos_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_product_product_proto_goTypes,
		DependencyIndexes: file_protos_product_product_proto_depIdxs,
		MessageInfos:      file_protos_product_product_proto_msgTypes,
	}.Build()
	File_protos_product_product_proto = out.File
	file_protos_product_product_proto_rawDesc = nil
	file_protos_product_product_proto_goTypes = nil
	file_protos_product_product_proto_depIdxs = nil
}
