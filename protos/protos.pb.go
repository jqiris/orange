// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: protos.proto

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

type MsgId int32

const (
	MsgId_MsgChan         MsgId = 0 //通道消息
	MsgId_MsgLogin        MsgId = 1 //登录
	MsgId_MsgLogout       MsgId = 2 //登出
	MsgId_MsgOffline      MsgId = 3 //断线
	MsgId_MsgMjRoomCreate MsgId = 4 //创建麻将房间
)

// Enum value maps for MsgId.
var (
	MsgId_name = map[int32]string{
		0: "MsgChan",
		1: "MsgLogin",
		2: "MsgLogout",
		3: "MsgOffline",
		4: "MsgMjRoomCreate",
	}
	MsgId_value = map[string]int32{
		"MsgChan":         0,
		"MsgLogin":        1,
		"MsgLogout":       2,
		"MsgOffline":      3,
		"MsgMjRoomCreate": 4,
	}
)

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}

func (x MsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_proto_enumTypes[0].Descriptor()
}

func (MsgId) Type() protoreflect.EnumType {
	return &file_protos_proto_enumTypes[0]
}

func (x MsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgId.Descriptor instead.
func (MsgId) EnumDescriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{0}
}

//request
type MsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId MsgId `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3,enum=protos.MsgId" json:"msg_id,omitempty"`
	// Types that are assignable to Any:
	//	*MsgRequest_LoginRequest
	//	*MsgRequest_LogoutRequest
	//	*MsgRequest_OfflineRequest
	//	*MsgRequest_MjRoomCreateRequest
	Any isMsgRequest_Any `protobuf_oneof:"any"`
}

func (x *MsgRequest) Reset() {
	*x = MsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRequest) ProtoMessage() {}

func (x *MsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRequest.ProtoReflect.Descriptor instead.
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{0}
}

func (x *MsgRequest) GetMsgId() MsgId {
	if x != nil {
		return x.MsgId
	}
	return MsgId_MsgChan
}

func (m *MsgRequest) GetAny() isMsgRequest_Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (x *MsgRequest) GetLoginRequest() *LoginRequest {
	if x, ok := x.GetAny().(*MsgRequest_LoginRequest); ok {
		return x.LoginRequest
	}
	return nil
}

func (x *MsgRequest) GetLogoutRequest() *LogoutRequest {
	if x, ok := x.GetAny().(*MsgRequest_LogoutRequest); ok {
		return x.LogoutRequest
	}
	return nil
}

func (x *MsgRequest) GetOfflineRequest() *OfflineRequest {
	if x, ok := x.GetAny().(*MsgRequest_OfflineRequest); ok {
		return x.OfflineRequest
	}
	return nil
}

func (x *MsgRequest) GetMjRoomCreateRequest() *MjRoomCreateRequest {
	if x, ok := x.GetAny().(*MsgRequest_MjRoomCreateRequest); ok {
		return x.MjRoomCreateRequest
	}
	return nil
}

type isMsgRequest_Any interface {
	isMsgRequest_Any()
}

type MsgRequest_LoginRequest struct {
	LoginRequest *LoginRequest `protobuf:"bytes,2,opt,name=login_request,json=loginRequest,proto3,oneof"` //登陆
}

type MsgRequest_LogoutRequest struct {
	LogoutRequest *LogoutRequest `protobuf:"bytes,3,opt,name=logout_request,json=logoutRequest,proto3,oneof"` //登出
}

type MsgRequest_OfflineRequest struct {
	OfflineRequest *OfflineRequest `protobuf:"bytes,4,opt,name=offline_request,json=offlineRequest,proto3,oneof"` //断线
}

type MsgRequest_MjRoomCreateRequest struct {
	MjRoomCreateRequest *MjRoomCreateRequest `protobuf:"bytes,5,opt,name=mj_room_create_request,json=mjRoomCreateRequest,proto3,oneof"` //创建麻将房间
}

func (*MsgRequest_LoginRequest) isMsgRequest_Any() {}

func (*MsgRequest_LogoutRequest) isMsgRequest_Any() {}

func (*MsgRequest_OfflineRequest) isMsgRequest_Any() {}

func (*MsgRequest_MjRoomCreateRequest) isMsgRequest_Any() {}

type MjRoomCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` //用户id
	Gid int32 `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"` //游戏id
}

func (x *MjRoomCreateRequest) Reset() {
	*x = MjRoomCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MjRoomCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MjRoomCreateRequest) ProtoMessage() {}

func (x *MjRoomCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MjRoomCreateRequest.ProtoReflect.Descriptor instead.
func (*MjRoomCreateRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{1}
}

func (x *MjRoomCreateRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *MjRoomCreateRequest) GetGid() int32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`          //用户ID
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"` //用户昵称
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`       //token key
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LoginRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *LoginRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` //用户id
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type OfflineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` //用户id
}

func (x *OfflineRequest) Reset() {
	*x = OfflineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineRequest) ProtoMessage() {}

func (x *OfflineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineRequest.ProtoReflect.Descriptor instead.
func (*OfflineRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{4}
}

func (x *OfflineRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

//response
type MsgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// Types that are assignable to Any:
	//	*MsgResponse_LoginResponse
	//	*MsgResponse_LogoutResponse
	//	*MsgResponse_MjRoomCreateResponse
	Any isMsgResponse_Any `protobuf_oneof:"any"`
}

func (x *MsgResponse) Reset() {
	*x = MsgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgResponse) ProtoMessage() {}

func (x *MsgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgResponse.ProtoReflect.Descriptor instead.
func (*MsgResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{5}
}

func (x *MsgResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (m *MsgResponse) GetAny() isMsgResponse_Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (x *MsgResponse) GetLoginResponse() *LoginResponse {
	if x, ok := x.GetAny().(*MsgResponse_LoginResponse); ok {
		return x.LoginResponse
	}
	return nil
}

func (x *MsgResponse) GetLogoutResponse() *LogoutResponse {
	if x, ok := x.GetAny().(*MsgResponse_LogoutResponse); ok {
		return x.LogoutResponse
	}
	return nil
}

func (x *MsgResponse) GetMjRoomCreateResponse() *MjRoomCreateResponse {
	if x, ok := x.GetAny().(*MsgResponse_MjRoomCreateResponse); ok {
		return x.MjRoomCreateResponse
	}
	return nil
}

type isMsgResponse_Any interface {
	isMsgResponse_Any()
}

type MsgResponse_LoginResponse struct {
	LoginResponse *LoginResponse `protobuf:"bytes,3,opt,name=login_response,json=loginResponse,proto3,oneof"`
}

type MsgResponse_LogoutResponse struct {
	LogoutResponse *LogoutResponse `protobuf:"bytes,4,opt,name=logout_response,json=logoutResponse,proto3,oneof"`
}

type MsgResponse_MjRoomCreateResponse struct {
	MjRoomCreateResponse *MjRoomCreateResponse `protobuf:"bytes,5,opt,name=mj_room_create_response,json=mjRoomCreateResponse,proto3,oneof"` //创建麻将房间回复
}

func (*MsgResponse_LoginResponse) isMsgResponse_Any() {}

func (*MsgResponse_LogoutResponse) isMsgResponse_Any() {}

func (*MsgResponse_MjRoomCreateResponse) isMsgResponse_Any() {}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                              // 用户id
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                     //昵称
	Avatar    string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                         //头像
	Gender    int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`                        //性别
	RegTime   int64  `protobuf:"varint,5,opt,name=reg_time,json=regTime,proto3" json:"reg_time,omitempty"`       //注册时间
	LoginTime int64  `protobuf:"varint,6,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"` //登录时间
	IsNew     bool   `protobuf:"varint,7,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`             //是否是新手
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{6}
}

func (x *LoginResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LoginResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *LoginResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *LoginResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *LoginResponse) GetRegTime() int64 {
	if x != nil {
		return x.RegTime
	}
	return 0
}

func (x *LoginResponse) GetLoginTime() int64 {
	if x != nil {
		return x.LoginTime
	}
	return 0
}

func (x *LoginResponse) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{7}
}

type MjRoomCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` //创建房间号
	Sid  string `protobuf:"bytes,2,opt,name=sid,proto3" json:"sid,omitempty"`    //创建服务器id
	Tid  int64  `protobuf:"varint,3,opt,name=tid,proto3" json:"tid,omitempty"`   //创建桌号
}

func (x *MjRoomCreateResponse) Reset() {
	*x = MjRoomCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MjRoomCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MjRoomCreateResponse) ProtoMessage() {}

func (x *MjRoomCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MjRoomCreateResponse.ProtoReflect.Descriptor instead.
func (*MjRoomCreateResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{8}
}

func (x *MjRoomCreateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MjRoomCreateResponse) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *MjRoomCreateResponse) GetTid() int64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

type ChessConfigMultiple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []int32 `protobuf:"varint,1,rep,packed,name=list,proto3" json:"list,omitempty"` //多选内容
}

func (x *ChessConfigMultiple) Reset() {
	*x = ChessConfigMultiple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessConfigMultiple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessConfigMultiple) ProtoMessage() {}

func (x *ChessConfigMultiple) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessConfigMultiple.ProtoReflect.Descriptor instead.
func (*ChessConfigMultiple) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{9}
}

func (x *ChessConfigMultiple) GetList() []int32 {
	if x != nil {
		return x.List
	}
	return nil
}

type ChessConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`                                                                                                //麻将类型
	Choices map[string]int32 `protobuf:"bytes,2,rep,name=choices,proto3" json:"choices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //选择
}

func (x *ChessConfig) Reset() {
	*x = ChessConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessConfig) ProtoMessage() {}

func (x *ChessConfig) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessConfig.ProtoReflect.Descriptor instead.
func (*ChessConfig) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{10}
}

func (x *ChessConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChessConfig) GetChoices() map[string]int32 {
	if x != nil {
		return x.Choices
	}
	return nil
}

var File_protos_proto protoreflect.FileDescriptor

var file_protos_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xcd, 0x02, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0d, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x6f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x16, 0x6d,
	0x6a, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x6a, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x13, 0x6d, 0x6a, 0x52, 0x6f,
	0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x05, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x22, 0x39, 0x0a, 0x13, 0x4d, 0x6a, 0x52, 0x6f, 0x6f, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x67, 0x69,
	0x64, 0x22, 0x52, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x21, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x94, 0x02, 0x0a,
	0x0b, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x3e, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x17, 0x6d, 0x6a, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4d, 0x6a, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x14, 0x6d, 0x6a, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x05, 0x0a, 0x03,
	0x61, 0x6e, 0x79, 0x22, 0xbe, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x4e, 0x65, 0x77, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x0a, 0x14, 0x4d, 0x6a, 0x52, 0x6f, 0x6f, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x74, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x56, 0x0a,
	0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x61,
	0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x03,
	0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x4d, 0x6a, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x10, 0x04, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_proto_rawDescOnce sync.Once
	file_protos_proto_rawDescData = file_protos_proto_rawDesc
)

func file_protos_proto_rawDescGZIP() []byte {
	file_protos_proto_rawDescOnce.Do(func() {
		file_protos_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_proto_rawDescData)
	})
	return file_protos_proto_rawDescData
}

var file_protos_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protos_proto_goTypes = []interface{}{
	(MsgId)(0),                   // 0: protos.MsgId
	(*MsgRequest)(nil),           // 1: protos.MsgRequest
	(*MjRoomCreateRequest)(nil),  // 2: protos.MjRoomCreateRequest
	(*LoginRequest)(nil),         // 3: protos.LoginRequest
	(*LogoutRequest)(nil),        // 4: protos.LogoutRequest
	(*OfflineRequest)(nil),       // 5: protos.OfflineRequest
	(*MsgResponse)(nil),          // 6: protos.MsgResponse
	(*LoginResponse)(nil),        // 7: protos.LoginResponse
	(*LogoutResponse)(nil),       // 8: protos.LogoutResponse
	(*MjRoomCreateResponse)(nil), // 9: protos.MjRoomCreateResponse
	(*ChessConfigMultiple)(nil),  // 10: protos.ChessConfigMultiple
	(*ChessConfig)(nil),          // 11: protos.ChessConfig
	nil,                          // 12: protos.ChessConfig.ChoicesEntry
}
var file_protos_proto_depIdxs = []int32{
	0,  // 0: protos.MsgRequest.msg_id:type_name -> protos.MsgId
	3,  // 1: protos.MsgRequest.login_request:type_name -> protos.LoginRequest
	4,  // 2: protos.MsgRequest.logout_request:type_name -> protos.LogoutRequest
	5,  // 3: protos.MsgRequest.offline_request:type_name -> protos.OfflineRequest
	2,  // 4: protos.MsgRequest.mj_room_create_request:type_name -> protos.MjRoomCreateRequest
	7,  // 5: protos.MsgResponse.login_response:type_name -> protos.LoginResponse
	8,  // 6: protos.MsgResponse.logout_response:type_name -> protos.LogoutResponse
	9,  // 7: protos.MsgResponse.mj_room_create_response:type_name -> protos.MjRoomCreateResponse
	12, // 8: protos.ChessConfig.choices:type_name -> protos.ChessConfig.ChoicesEntry
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_protos_proto_init() }
func file_protos_proto_init() {
	if File_protos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRequest); i {
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
		file_protos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MjRoomCreateRequest); i {
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
		file_protos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_protos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_protos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineRequest); i {
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
		file_protos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgResponse); i {
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
		file_protos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_protos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
		file_protos_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MjRoomCreateResponse); i {
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
		file_protos_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessConfigMultiple); i {
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
		file_protos_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessConfig); i {
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
	file_protos_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MsgRequest_LoginRequest)(nil),
		(*MsgRequest_LogoutRequest)(nil),
		(*MsgRequest_OfflineRequest)(nil),
		(*MsgRequest_MjRoomCreateRequest)(nil),
	}
	file_protos_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*MsgResponse_LoginResponse)(nil),
		(*MsgResponse_LogoutResponse)(nil),
		(*MsgResponse_MjRoomCreateResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_proto_goTypes,
		DependencyIndexes: file_protos_proto_depIdxs,
		EnumInfos:         file_protos_proto_enumTypes,
		MessageInfos:      file_protos_proto_msgTypes,
	}.Build()
	File_protos_proto = out.File
	file_protos_proto_rawDesc = nil
	file_protos_proto_goTypes = nil
	file_protos_proto_depIdxs = nil
}
