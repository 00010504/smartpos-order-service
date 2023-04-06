// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: validation.proto

package validation_service

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

type GetImportValidationLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Row     int32  `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
}

func (x *GetImportValidationLogResponse) Reset() {
	*x = GetImportValidationLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImportValidationLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImportValidationLogResponse) ProtoMessage() {}

func (x *GetImportValidationLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImportValidationLogResponse.ProtoReflect.Descriptor instead.
func (*GetImportValidationLogResponse) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{0}
}

func (x *GetImportValidationLogResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetImportValidationLogResponse) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

type GetImportValidationLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*GetImportValidationLogResponse `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetImportValidationLogsResponse) Reset() {
	*x = GetImportValidationLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImportValidationLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImportValidationLogsResponse) ProtoMessage() {}

func (x *GetImportValidationLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImportValidationLogsResponse.ProtoReflect.Descriptor instead.
func (*GetImportValidationLogsResponse) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{1}
}

func (x *GetImportValidationLogsResponse) GetLogs() []*GetImportValidationLogResponse {
	if x != nil {
		return x.Logs
	}
	return nil
}

type CreateImportErrorReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportId string `protobuf:"bytes,1,opt,name=import_id,json=importId,proto3" json:"import_id,omitempty"`
}

func (x *CreateImportErrorReq) Reset() {
	*x = CreateImportErrorReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImportErrorReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImportErrorReq) ProtoMessage() {}

func (x *CreateImportErrorReq) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImportErrorReq.ProtoReflect.Descriptor instead.
func (*CreateImportErrorReq) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{2}
}

func (x *CreateImportErrorReq) GetImportId() string {
	if x != nil {
		return x.ImportId
	}
	return ""
}

type CreateImportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Store            string            `protobuf:"bytes,2,opt,name=store,proto3" json:"store,omitempty"`
	Url              string            `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	GenerateBarcode  bool              `protobuf:"varint,4,opt,name=generate_barcode,json=generateBarcode,proto3" json:"generate_barcode,omitempty"`
	GenerateSku      bool              `protobuf:"varint,5,opt,name=generate_sku,json=generateSku,proto3" json:"generate_sku,omitempty"`
	Type             string            `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	ImportProperties []*ItemProperties `protobuf:"bytes,7,rep,name=import_properties,json=importProperties,proto3" json:"import_properties,omitempty"`
	Quantity         int32             `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Total            int32             `protobuf:"varint,9,opt,name=total,proto3" json:"total,omitempty"`
	Status           string            `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Date             string            `protobuf:"bytes,11,opt,name=date,proto3" json:"date,omitempty"`
	CreatedBy        string            `protobuf:"bytes,12,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CompletedBy      string            `protobuf:"bytes,13,opt,name=completed_by,json=completedBy,proto3" json:"completed_by,omitempty"`
}

func (x *CreateImportReq) Reset() {
	*x = CreateImportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImportReq) ProtoMessage() {}

func (x *CreateImportReq) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImportReq.ProtoReflect.Descriptor instead.
func (*CreateImportReq) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{3}
}

func (x *CreateImportReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateImportReq) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *CreateImportReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateImportReq) GetGenerateBarcode() bool {
	if x != nil {
		return x.GenerateBarcode
	}
	return false
}

func (x *CreateImportReq) GetGenerateSku() bool {
	if x != nil {
		return x.GenerateSku
	}
	return false
}

func (x *CreateImportReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateImportReq) GetImportProperties() []*ItemProperties {
	if x != nil {
		return x.ImportProperties
	}
	return nil
}

func (x *CreateImportReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateImportReq) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CreateImportReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateImportReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *CreateImportReq) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *CreateImportReq) GetCompletedBy() string {
	if x != nil {
		return x.CompletedBy
	}
	return ""
}

type CreateImportRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateImportRes) Reset() {
	*x = CreateImportRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImportRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImportRes) ProtoMessage() {}

func (x *CreateImportRes) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImportRes.ProtoReflect.Descriptor instead.
func (*CreateImportRes) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{4}
}

func (x *CreateImportRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConfirmImportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportId string `protobuf:"bytes,1,opt,name=import_id,json=importId,proto3" json:"import_id,omitempty"`
}

func (x *ConfirmImportReq) Reset() {
	*x = ConfirmImportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmImportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmImportReq) ProtoMessage() {}

func (x *ConfirmImportReq) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmImportReq.ProtoReflect.Descriptor instead.
func (*ConfirmImportReq) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmImportReq) GetImportId() string {
	if x != nil {
		return x.ImportId
	}
	return ""
}

type ItemProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FieldName         string `protobuf:"bytes,2,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	IsAttribute       bool   `protobuf:"varint,3,opt,name=is_attribute,json=isAttribute,proto3" json:"is_attribute,omitempty"`
	IsCharacteristics bool   `protobuf:"varint,4,opt,name=is_characteristics,json=isCharacteristics,proto3" json:"is_characteristics,omitempty"`
	IsNew             bool   `protobuf:"varint,5,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsUploadable      bool   `protobuf:"varint,6,opt,name=is_uploadable,json=isUploadable,proto3" json:"is_uploadable,omitempty"`
	SequenceNumber    int32  `protobuf:"varint,7,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
}

func (x *ItemProperties) Reset() {
	*x = ItemProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemProperties) ProtoMessage() {}

func (x *ItemProperties) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemProperties.ProtoReflect.Descriptor instead.
func (*ItemProperties) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{6}
}

func (x *ItemProperties) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemProperties) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *ItemProperties) GetIsAttribute() bool {
	if x != nil {
		return x.IsAttribute
	}
	return false
}

func (x *ItemProperties) GetIsCharacteristics() bool {
	if x != nil {
		return x.IsCharacteristics
	}
	return false
}

func (x *ItemProperties) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *ItemProperties) GetIsUploadable() bool {
	if x != nil {
		return x.IsUploadable
	}
	return false
}

func (x *ItemProperties) GetSequenceNumber() int32 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

var File_validation_proto protoreflect.FileDescriptor

var file_validation_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x77,
	0x22, 0x56, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x8d, 0x03,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x29, 0x0a,
	0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x3c, 0x0a, 0x11, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x2b, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0xf6, 0x01, 0x0a, 0x0e,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x2d, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_validation_proto_rawDescOnce sync.Once
	file_validation_proto_rawDescData = file_validation_proto_rawDesc
)

func file_validation_proto_rawDescGZIP() []byte {
	file_validation_proto_rawDescOnce.Do(func() {
		file_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_validation_proto_rawDescData)
	})
	return file_validation_proto_rawDescData
}

var file_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_validation_proto_goTypes = []interface{}{
	(*GetImportValidationLogResponse)(nil),  // 0: GetImportValidationLogResponse
	(*GetImportValidationLogsResponse)(nil), // 1: GetImportValidationLogsResponse
	(*CreateImportErrorReq)(nil),            // 2: CreateImportErrorReq
	(*CreateImportReq)(nil),                 // 3: CreateImportReq
	(*CreateImportRes)(nil),                 // 4: CreateImportRes
	(*ConfirmImportReq)(nil),                // 5: ConfirmImportReq
	(*ItemProperties)(nil),                  // 6: ItemProperties
}
var file_validation_proto_depIdxs = []int32{
	0, // 0: GetImportValidationLogsResponse.logs:type_name -> GetImportValidationLogResponse
	6, // 1: CreateImportReq.import_properties:type_name -> ItemProperties
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_validation_proto_init() }
func file_validation_proto_init() {
	if File_validation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImportValidationLogResponse); i {
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
		file_validation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImportValidationLogsResponse); i {
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
		file_validation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImportErrorReq); i {
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
		file_validation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImportReq); i {
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
		file_validation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImportRes); i {
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
		file_validation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmImportReq); i {
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
		file_validation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemProperties); i {
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
			RawDescriptor: file_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_validation_proto_goTypes,
		DependencyIndexes: file_validation_proto_depIdxs,
		MessageInfos:      file_validation_proto_msgTypes,
	}.Build()
	File_validation_proto = out.File
	file_validation_proto_rawDesc = nil
	file_validation_proto_goTypes = nil
	file_validation_proto_depIdxs = nil
}
