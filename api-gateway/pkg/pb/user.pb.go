// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/proto/user.proto

package pb

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

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32   `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	City      string  `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	FirstName string  `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	Phone     string  `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Height    float32 `protobuf:"fixed32,5,opt,name=Height,proto3" json:"Height,omitempty"`
	Married   bool    `protobuf:"varint,6,opt,name=Married,proto3" json:"Married,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserData) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserData) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserData) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UserData) GetMarried() bool {
	if x != nil {
		return x.Married
	}
	return false
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *UserData `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Status   int32     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Response string    `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Error    string    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserResponse) GetResult() *UserData {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetUserResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetUserResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *GetUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetUserDataListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIdList []int32 `protobuf:"varint,1,rep,packed,name=UserIdList,proto3" json:"UserIdList,omitempty"`
}

func (x *GetUserDataListRequest) Reset() {
	*x = GetUserDataListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataListRequest) ProtoMessage() {}

func (x *GetUserDataListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataListRequest.ProtoReflect.Descriptor instead.
func (*GetUserDataListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserDataListRequest) GetUserIdList() []int32 {
	if x != nil {
		return x.UserIdList
	}
	return nil
}

type GetUserDataListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   []*UserData `protobuf:"bytes,1,rep,name=Result,proto3" json:"Result,omitempty"`
	NotFound []int32     `protobuf:"varint,2,rep,packed,name=NotFound,proto3" json:"NotFound,omitempty"`
	Status   int32       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Response string      `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	Error    string      `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetUserDataListResponse) Reset() {
	*x = GetUserDataListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataListResponse) ProtoMessage() {}

func (x *GetUserDataListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataListResponse.ProtoReflect.Descriptor instead.
func (*GetUserDataListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserDataListResponse) GetResult() []*UserData {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetUserDataListResponse) GetNotFound() []int32 {
	if x != nil {
		return x.NotFound
	}
	return nil
}

func (x *GetUserDataListResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetUserDataListResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *GetUserDataListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_proto_user_proto protoreflect.FileDescriptor

var file_pkg_proto_user_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x72,
	0x72, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x4d, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0xa5, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_user_proto_rawDescOnce sync.Once
	file_pkg_proto_user_proto_rawDescData = file_pkg_proto_user_proto_rawDesc
)

func file_pkg_proto_user_proto_rawDescGZIP() []byte {
	file_pkg_proto_user_proto_rawDescOnce.Do(func() {
		file_pkg_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_user_proto_rawDescData)
	})
	return file_pkg_proto_user_proto_rawDescData
}

var file_pkg_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_user_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),          // 0: pb.GetUserRequest
	(*UserData)(nil),                // 1: pb.UserData
	(*GetUserResponse)(nil),         // 2: pb.GetUserResponse
	(*GetUserDataListRequest)(nil),  // 3: pb.GetUserDataListRequest
	(*GetUserDataListResponse)(nil), // 4: pb.GetUserDataListResponse
}
var file_pkg_proto_user_proto_depIdxs = []int32{
	1, // 0: pb.GetUserResponse.Result:type_name -> pb.UserData
	1, // 1: pb.GetUserDataListResponse.Result:type_name -> pb.UserData
	0, // 2: pb.UserService.GetUserData:input_type -> pb.GetUserRequest
	3, // 3: pb.UserService.GetUserDataList:input_type -> pb.GetUserDataListRequest
	2, // 4: pb.UserService.GetUserData:output_type -> pb.GetUserResponse
	4, // 5: pb.UserService.GetUserDataList:output_type -> pb.GetUserDataListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_user_proto_init() }
func file_pkg_proto_user_proto_init() {
	if File_pkg_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_pkg_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_pkg_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_pkg_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataListRequest); i {
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
		file_pkg_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataListResponse); i {
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
			RawDescriptor: file_pkg_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_user_proto_goTypes,
		DependencyIndexes: file_pkg_proto_user_proto_depIdxs,
		MessageInfos:      file_pkg_proto_user_proto_msgTypes,
	}.Build()
	File_pkg_proto_user_proto = out.File
	file_pkg_proto_user_proto_rawDesc = nil
	file_pkg_proto_user_proto_goTypes = nil
	file_pkg_proto_user_proto_depIdxs = nil
}
