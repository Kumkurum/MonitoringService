// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: message.proto

package monitoringGrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error_ErrorCode int32

const (
	Error_OK              Error_ErrorCode = 0
	Error_ERROR           Error_ErrorCode = 1
	Error_UNKNOWN_REQUEST Error_ErrorCode = 2
	Error_UNKNOWN_ID      Error_ErrorCode = 3
	Error_UNKNOWN_TOKEN   Error_ErrorCode = 4
)

// Enum value maps for Error_ErrorCode.
var (
	Error_ErrorCode_name = map[int32]string{
		0: "OK",
		1: "ERROR",
		2: "UNKNOWN_REQUEST",
		3: "UNKNOWN_ID",
		4: "UNKNOWN_TOKEN",
	}
	Error_ErrorCode_value = map[string]int32{
		"OK":              0,
		"ERROR":           1,
		"UNKNOWN_REQUEST": 2,
		"UNKNOWN_ID":      3,
		"UNKNOWN_TOKEN":   4,
	}
)

func (x Error_ErrorCode) Enum() *Error_ErrorCode {
	p := new(Error_ErrorCode)
	*p = x
	return p
}

func (x Error_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Error_ErrorCode) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Error_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_ErrorCode.Descriptor instead.
func (Error_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2, 0}
}

type ProcessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CpuPercent float64 `protobuf:"fixed64,2,opt,name=cpuPercent,proto3" json:"cpuPercent,omitempty"`
	MemPercent float32 `protobuf:"fixed32,3,opt,name=memPercent,proto3" json:"memPercent,omitempty"`
	Status     string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProcessInfo) Reset() {
	*x = ProcessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInfo) ProtoMessage() {}

func (x *ProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInfo.ProtoReflect.Descriptor instead.
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessInfo) GetCpuPercent() float64 {
	if x != nil {
		return x.CpuPercent
	}
	return 0
}

func (x *ProcessInfo) GetMemPercent() float32 {
	if x != nil {
		return x.MemPercent
	}
	return 0
}

func (x *ProcessInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AllInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ProcessName []string `protobuf:"bytes,2,rep,name=processName,proto3" json:"processName,omitempty"`
}

func (x *AllInfoRequest) Reset() {
	*x = AllInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInfoRequest) ProtoMessage() {}

func (x *AllInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInfoRequest.ProtoReflect.Descriptor instead.
func (*AllInfoRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *AllInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AllInfoRequest) GetProcessName() []string {
	if x != nil {
		return x.ProcessName
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Error_ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=Error_ErrorCode" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() Error_ErrorCode {
	if x != nil {
		return x.Code
	}
	return Error_OK
}

type AllInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*AllInfoResponse_Error
	//	*AllInfoResponse_Success
	Response isAllInfoResponse_Response `protobuf_oneof:"response"`
}

func (x *AllInfoResponse) Reset() {
	*x = AllInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInfoResponse) ProtoMessage() {}

func (x *AllInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInfoResponse.ProtoReflect.Descriptor instead.
func (*AllInfoResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (m *AllInfoResponse) GetResponse() isAllInfoResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *AllInfoResponse) GetError() *Error {
	if x, ok := x.GetResponse().(*AllInfoResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *AllInfoResponse) GetSuccess() *AllInfoResponseSuccess {
	if x, ok := x.GetResponse().(*AllInfoResponse_Success); ok {
		return x.Success
	}
	return nil
}

type isAllInfoResponse_Response interface {
	isAllInfoResponse_Response()
}

type AllInfoResponse_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type AllInfoResponse_Success struct {
	Success *AllInfoResponseSuccess `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

func (*AllInfoResponse_Error) isAllInfoResponse_Response() {}

func (*AllInfoResponse_Success) isAllInfoResponse_Response() {}

type AllInfoResponseSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuPercent  float64                `protobuf:"fixed64,1,opt,name=cpuPercent,proto3" json:"cpuPercent,omitempty"`
	MemPercent  float64                `protobuf:"fixed64,2,opt,name=memPercent,proto3" json:"memPercent,omitempty"`
	DiskPercent float64                `protobuf:"fixed64,3,opt,name=diskPercent,proto3" json:"diskPercent,omitempty"`
	NetInfoSent uint64                 `protobuf:"varint,4,opt,name=netInfoSent,proto3" json:"netInfoSent,omitempty"`
	NetInfoRecv uint64                 `protobuf:"varint,5,opt,name=netInfoRecv,proto3" json:"netInfoRecv,omitempty"`
	ListenPorts []uint32               `protobuf:"varint,6,rep,packed,name=listenPorts,proto3" json:"listenPorts,omitempty"`
	ProcessInfo []*ProcessInfo         `protobuf:"bytes,7,rep,name=processInfo,proto3" json:"processInfo,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *AllInfoResponseSuccess) Reset() {
	*x = AllInfoResponseSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInfoResponseSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInfoResponseSuccess) ProtoMessage() {}

func (x *AllInfoResponseSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInfoResponseSuccess.ProtoReflect.Descriptor instead.
func (*AllInfoResponseSuccess) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *AllInfoResponseSuccess) GetCpuPercent() float64 {
	if x != nil {
		return x.CpuPercent
	}
	return 0
}

func (x *AllInfoResponseSuccess) GetMemPercent() float64 {
	if x != nil {
		return x.MemPercent
	}
	return 0
}

func (x *AllInfoResponseSuccess) GetDiskPercent() float64 {
	if x != nil {
		return x.DiskPercent
	}
	return 0
}

func (x *AllInfoResponseSuccess) GetNetInfoSent() uint64 {
	if x != nil {
		return x.NetInfoSent
	}
	return 0
}

func (x *AllInfoResponseSuccess) GetNetInfoRecv() uint64 {
	if x != nil {
		return x.NetInfoRecv
	}
	return 0
}

func (x *AllInfoResponseSuccess) GetListenPorts() []uint32 {
	if x != nil {
		return x.ListenPorts
	}
	return nil
}

func (x *AllInfoResponseSuccess) GetProcessInfo() []*ProcessInfo {
	if x != nil {
		return x.ProcessInfo
	}
	return nil
}

func (x *AllInfoResponseSuccess) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x79, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a, 0x0e, 0x41,
	0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x56, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x04, 0x22, 0x72, 0x0a,
	0x0f, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x33, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xc0, 0x02, 0x0a, 0x16, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x69, 0x73, 0x6b, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x63, 0x76, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x63, 0x76, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(Error_ErrorCode)(0),           // 0: Error.ErrorCode
	(*ProcessInfo)(nil),            // 1: ProcessInfo
	(*AllInfoRequest)(nil),         // 2: AllInfoRequest
	(*Error)(nil),                  // 3: Error
	(*AllInfoResponse)(nil),        // 4: AllInfoResponse
	(*AllInfoResponseSuccess)(nil), // 5: AllInfoResponseSuccess
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_message_proto_depIdxs = []int32{
	0, // 0: Error.code:type_name -> Error.ErrorCode
	3, // 1: AllInfoResponse.error:type_name -> Error
	5, // 2: AllInfoResponse.success:type_name -> AllInfoResponseSuccess
	1, // 3: AllInfoResponseSuccess.processInfo:type_name -> ProcessInfo
	6, // 4: AllInfoResponseSuccess.time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInfo); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInfoRequest); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInfoResponse); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInfoResponseSuccess); i {
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
	file_message_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AllInfoResponse_Error)(nil),
		(*AllInfoResponse_Success)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
