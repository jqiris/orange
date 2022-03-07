// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: internel.proto

package protos

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

type InnerMsgId int32

const (
	InnerMsgId_InnerMsgChan       InnerMsgId = 0 //通道消息
	InnerMsgId_InnerMsgCreateRoom InnerMsgId = 1 //创建房间
	InnerMsgId_InnerMsgEnterRoom  InnerMsgId = 2 //进入房间
	InnerMsgId_InnerMsgRegServer  InnerMsgId = 3 //注册服务器
)

// Enum value maps for InnerMsgId.
var (
	InnerMsgId_name = map[int32]string{
		0: "InnerMsgChan",
		1: "InnerMsgCreateRoom",
		2: "InnerMsgEnterRoom",
		3: "InnerMsgRegServer",
	}
	InnerMsgId_value = map[string]int32{
		"InnerMsgChan":       0,
		"InnerMsgCreateRoom": 1,
		"InnerMsgEnterRoom":  2,
		"InnerMsgRegServer":  3,
	}
)

func (x InnerMsgId) Enum() *InnerMsgId {
	p := new(InnerMsgId)
	*p = x
	return p
}

func (x InnerMsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InnerMsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_internel_proto_enumTypes[0].Descriptor()
}

func (InnerMsgId) Type() protoreflect.EnumType {
	return &file_internel_proto_enumTypes[0]
}

func (x InnerMsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InnerMsgId.Descriptor instead.
func (InnerMsgId) EnumDescriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{0}
}

type InnerMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errcode int32  `protobuf:"varint,1,opt,name=errcode,proto3" json:"errcode,omitempty"` //错误码
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`    //错误消息
	// Types that are assignable to Any:
	//	*InnerMsgResp_CreateRoom
	//	*InnerMsgResp_EnterRoom
	//	*InnerMsgResp_RegServer
	Any isInnerMsgResp_Any `protobuf_oneof:"any"`
}

func (x *InnerMsgResp) Reset() {
	*x = InnerMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerMsgResp) ProtoMessage() {}

func (x *InnerMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerMsgResp.ProtoReflect.Descriptor instead.
func (*InnerMsgResp) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{0}
}

func (x *InnerMsgResp) GetErrcode() int32 {
	if x != nil {
		return x.Errcode
	}
	return 0
}

func (x *InnerMsgResp) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (m *InnerMsgResp) GetAny() isInnerMsgResp_Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (x *InnerMsgResp) GetCreateRoom() *InnerCreateRoomResp {
	if x, ok := x.GetAny().(*InnerMsgResp_CreateRoom); ok {
		return x.CreateRoom
	}
	return nil
}

func (x *InnerMsgResp) GetEnterRoom() *InnerEnterRoomResp {
	if x, ok := x.GetAny().(*InnerMsgResp_EnterRoom); ok {
		return x.EnterRoom
	}
	return nil
}

func (x *InnerMsgResp) GetRegServer() *InnerRegServerResp {
	if x, ok := x.GetAny().(*InnerMsgResp_RegServer); ok {
		return x.RegServer
	}
	return nil
}

type isInnerMsgResp_Any interface {
	isInnerMsgResp_Any()
}

type InnerMsgResp_CreateRoom struct {
	CreateRoom *InnerCreateRoomResp `protobuf:"bytes,3,opt,name=create_room,json=createRoom,proto3,oneof"` //创建房间
}

type InnerMsgResp_EnterRoom struct {
	EnterRoom *InnerEnterRoomResp `protobuf:"bytes,4,opt,name=enter_room,json=enterRoom,proto3,oneof"` //进入房间
}

type InnerMsgResp_RegServer struct {
	RegServer *InnerRegServerResp `protobuf:"bytes,5,opt,name=reg_server,json=regServer,proto3,oneof"` //注册服务器
}

func (*InnerMsgResp_CreateRoom) isInnerMsgResp_Any() {}

func (*InnerMsgResp_EnterRoom) isInnerMsgResp_Any() {}

func (*InnerMsgResp_RegServer) isInnerMsgResp_Any() {}

type InnerCreateRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"` //用户id
	Gems   int64  `protobuf:"varint,2,opt,name=gems,proto3" json:"gems,omitempty"`     //钻石数量
	Conf   string `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`      //配置内容
	Sign   string `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`      //签名
}

func (x *InnerCreateRoomReq) Reset() {
	*x = InnerCreateRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerCreateRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerCreateRoomReq) ProtoMessage() {}

func (x *InnerCreateRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerCreateRoomReq.ProtoReflect.Descriptor instead.
func (*InnerCreateRoomReq) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{1}
}

func (x *InnerCreateRoomReq) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *InnerCreateRoomReq) GetGems() int64 {
	if x != nil {
		return x.Gems
	}
	return 0
}

func (x *InnerCreateRoomReq) GetConf() string {
	if x != nil {
		return x.Conf
	}
	return ""
}

func (x *InnerCreateRoomReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type InnerCreateRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roomid   int32  `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty"`                    //房间id
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"` //服务器ID
}

func (x *InnerCreateRoomResp) Reset() {
	*x = InnerCreateRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerCreateRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerCreateRoomResp) ProtoMessage() {}

func (x *InnerCreateRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerCreateRoomResp.ProtoReflect.Descriptor instead.
func (*InnerCreateRoomResp) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{2}
}

func (x *InnerCreateRoomResp) GetRoomid() int32 {
	if x != nil {
		return x.Roomid
	}
	return 0
}

func (x *InnerCreateRoomResp) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type InnerEnterRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"` //用户id
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`      //用户名称
	Roomid int32  `protobuf:"varint,3,opt,name=roomid,proto3" json:"roomid,omitempty"` //房间id
	Sign   string `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`      //签名
}

func (x *InnerEnterRoomReq) Reset() {
	*x = InnerEnterRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerEnterRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerEnterRoomReq) ProtoMessage() {}

func (x *InnerEnterRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerEnterRoomReq.ProtoReflect.Descriptor instead.
func (*InnerEnterRoomReq) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{3}
}

func (x *InnerEnterRoomReq) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *InnerEnterRoomReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InnerEnterRoomReq) GetRoomid() int32 {
	if x != nil {
		return x.Roomid
	}
	return 0
}

func (x *InnerEnterRoomReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type InnerEnterRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //token
}

func (x *InnerEnterRoomResp) Reset() {
	*x = InnerEnterRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerEnterRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerEnterRoomResp) ProtoMessage() {}

func (x *InnerEnterRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerEnterRoomResp.ProtoReflect.Descriptor instead.
func (*InnerEnterRoomResp) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{4}
}

func (x *InnerEnterRoomResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type InnerRegServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                  //服务器id
	Clientip   string `protobuf:"bytes,2,opt,name=clientip,proto3" json:"clientip,omitempty"`      //客户端ip
	Clientport int32  `protobuf:"varint,3,opt,name=clientport,proto3" json:"clientport,omitempty"` //客户端port
	HttpPort   int32  `protobuf:"varint,4,opt,name=httpPort,proto3" json:"httpPort,omitempty"`     //httpport
	Load       int32  `protobuf:"varint,5,opt,name=load,proto3" json:"load,omitempty"`             //负载量
}

func (x *InnerRegServerReq) Reset() {
	*x = InnerRegServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerRegServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerRegServerReq) ProtoMessage() {}

func (x *InnerRegServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerRegServerReq.ProtoReflect.Descriptor instead.
func (*InnerRegServerReq) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{5}
}

func (x *InnerRegServerReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InnerRegServerReq) GetClientip() string {
	if x != nil {
		return x.Clientip
	}
	return ""
}

func (x *InnerRegServerReq) GetClientport() int32 {
	if x != nil {
		return x.Clientport
	}
	return 0
}

func (x *InnerRegServerReq) GetHttpPort() int32 {
	if x != nil {
		return x.HttpPort
	}
	return 0
}

func (x *InnerRegServerReq) GetLoad() int32 {
	if x != nil {
		return x.Load
	}
	return 0
}

type InnerRegServerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"` //ip
}

func (x *InnerRegServerResp) Reset() {
	*x = InnerRegServerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerRegServerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerRegServerResp) ProtoMessage() {}

func (x *InnerRegServerResp) ProtoReflect() protoreflect.Message {
	mi := &file_internel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerRegServerResp.ProtoReflect.Descriptor instead.
func (*InnerRegServerResp) Descriptor() ([]byte, []int) {
	return file_internel_proto_rawDescGZIP(), []int{6}
}

func (x *InnerRegServerResp) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_internel_proto protoreflect.FileDescriptor

var file_internel_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x81, 0x02, 0x0a, 0x0c, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x3e, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x09, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x22, 0x68, 0x0a, 0x12,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x65, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x4a, 0x0a, 0x13, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x6b, 0x0a, 0x11, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22,
	0x2a, 0x0a, 0x12, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x11,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x74, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x68, 0x74, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x24, 0x0a,
	0x12, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x2a, 0x64, 0x0a, 0x0a, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x61,
	0x6e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x03, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internel_proto_rawDescOnce sync.Once
	file_internel_proto_rawDescData = file_internel_proto_rawDesc
)

func file_internel_proto_rawDescGZIP() []byte {
	file_internel_proto_rawDescOnce.Do(func() {
		file_internel_proto_rawDescData = protoimpl.X.CompressGZIP(file_internel_proto_rawDescData)
	})
	return file_internel_proto_rawDescData
}

var file_internel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internel_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internel_proto_goTypes = []interface{}{
	(InnerMsgId)(0),             // 0: protos.InnerMsgId
	(*InnerMsgResp)(nil),        // 1: protos.InnerMsgResp
	(*InnerCreateRoomReq)(nil),  // 2: protos.InnerCreateRoomReq
	(*InnerCreateRoomResp)(nil), // 3: protos.InnerCreateRoomResp
	(*InnerEnterRoomReq)(nil),   // 4: protos.InnerEnterRoomReq
	(*InnerEnterRoomResp)(nil),  // 5: protos.InnerEnterRoomResp
	(*InnerRegServerReq)(nil),   // 6: protos.InnerRegServerReq
	(*InnerRegServerResp)(nil),  // 7: protos.InnerRegServerResp
}
var file_internel_proto_depIdxs = []int32{
	3, // 0: protos.InnerMsgResp.create_room:type_name -> protos.InnerCreateRoomResp
	5, // 1: protos.InnerMsgResp.enter_room:type_name -> protos.InnerEnterRoomResp
	7, // 2: protos.InnerMsgResp.reg_server:type_name -> protos.InnerRegServerResp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internel_proto_init() }
func file_internel_proto_init() {
	if File_internel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerMsgResp); i {
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
		file_internel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerCreateRoomReq); i {
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
		file_internel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerCreateRoomResp); i {
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
		file_internel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerEnterRoomReq); i {
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
		file_internel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerEnterRoomResp); i {
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
		file_internel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerRegServerReq); i {
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
		file_internel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerRegServerResp); i {
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
	file_internel_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InnerMsgResp_CreateRoom)(nil),
		(*InnerMsgResp_EnterRoom)(nil),
		(*InnerMsgResp_RegServer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internel_proto_goTypes,
		DependencyIndexes: file_internel_proto_depIdxs,
		EnumInfos:         file_internel_proto_enumTypes,
		MessageInfos:      file_internel_proto_msgTypes,
	}.Build()
	File_internel_proto = out.File
	file_internel_proto_rawDesc = nil
	file_internel_proto_goTypes = nil
	file_internel_proto_depIdxs = nil
}
