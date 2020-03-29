// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo_list_activity_manager.proto

package todolist

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Tags int32

const (
	Tags_PRIORITY    Tags = 0
	Tags_NORMAL      Tags = 1
	Tags_MONEY       Tags = 2
	Tags_SPORTS      Tags = 3
	Tags_COKKING     Tags = 4
	Tags_SOCAILIZING Tags = 5
	Tags_GENERAL     Tags = 6
)

var Tags_name = map[int32]string{
	0: "PRIORITY",
	1: "NORMAL",
	2: "MONEY",
	3: "SPORTS",
	4: "COKKING",
	5: "SOCAILIZING",
	6: "GENERAL",
}

var Tags_value = map[string]int32{
	"PRIORITY":    0,
	"NORMAL":      1,
	"MONEY":       2,
	"SPORTS":      3,
	"COKKING":     4,
	"SOCAILIZING": 5,
	"GENERAL":     6,
}

func (x Tags) String() string {
	return proto.EnumName(Tags_name, int32(x))
}

func (Tags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{0}
}

type CreateTodoListRequest struct {
	UserID               string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DueDate              string        `protobuf:"bytes,4,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	CustomDueDate        string        `protobuf:"bytes,5,opt,name=customDueDate,proto3" json:"customDueDate,omitempty"`
	RepeatMode           bool          `protobuf:"varint,6,opt,name=repeatMode,proto3" json:"repeatMode,omitempty"`
	Tag                  Tags          `protobuf:"varint,7,opt,name=tag,proto3,enum=todolist.Tags" json:"tag,omitempty"`
	Notification         *Notification `protobuf:"bytes,8,opt,name=notification,proto3" json:"notification,omitempty"`
	SubTask              []*SubTask    `protobuf:"bytes,9,rep,name=subTask,proto3" json:"subTask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateTodoListRequest) Reset()         { *m = CreateTodoListRequest{} }
func (m *CreateTodoListRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoListRequest) ProtoMessage()    {}
func (*CreateTodoListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{0}
}

func (m *CreateTodoListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListRequest.Unmarshal(m, b)
}
func (m *CreateTodoListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListRequest.Merge(m, src)
}
func (m *CreateTodoListRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListRequest.Size(m)
}
func (m *CreateTodoListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListRequest proto.InternalMessageInfo

func (m *CreateTodoListRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CreateTodoListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTodoListRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTodoListRequest) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *CreateTodoListRequest) GetCustomDueDate() string {
	if m != nil {
		return m.CustomDueDate
	}
	return ""
}

func (m *CreateTodoListRequest) GetRepeatMode() bool {
	if m != nil {
		return m.RepeatMode
	}
	return false
}

func (m *CreateTodoListRequest) GetTag() Tags {
	if m != nil {
		return m.Tag
	}
	return Tags_PRIORITY
}

func (m *CreateTodoListRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *CreateTodoListRequest) GetSubTask() []*SubTask {
	if m != nil {
		return m.SubTask
	}
	return nil
}

type CreateTodoListResponse struct {
	Data                 []*CreateTodoListResponse_CreateTodoListUserInformationResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Errors               []*CreateTodoListResponse_ErrorResponse                         `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *CreateTodoListResponse) Reset()         { *m = CreateTodoListResponse{} }
func (m *CreateTodoListResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoListResponse) ProtoMessage()    {}
func (*CreateTodoListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{1}
}

func (m *CreateTodoListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListResponse.Unmarshal(m, b)
}
func (m *CreateTodoListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListResponse.Merge(m, src)
}
func (m *CreateTodoListResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListResponse.Size(m)
}
func (m *CreateTodoListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListResponse proto.InternalMessageInfo

func (m *CreateTodoListResponse) GetData() []*CreateTodoListResponse_CreateTodoListUserInformationResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateTodoListResponse) GetErrors() []*CreateTodoListResponse_ErrorResponse {
	if m != nil {
		return m.Errors
	}
	return nil
}

type CreateTodoListResponse_CreateTodoListUserInformationResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) Reset() {
	*m = CreateTodoListResponse_CreateTodoListUserInformationResponse{}
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) String() string {
	return proto.CompactTextString(m)
}
func (*CreateTodoListResponse_CreateTodoListUserInformationResponse) ProtoMessage() {}
func (*CreateTodoListResponse_CreateTodoListUserInformationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{1, 0}
}

func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Unmarshal(m, b)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Merge(m, src)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Size(m)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse proto.InternalMessageInfo

func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CreateTodoListResponse_ErrorResponse struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	StatusCode           string   `protobuf:"bytes,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoListResponse_ErrorResponse) Reset()         { *m = CreateTodoListResponse_ErrorResponse{} }
func (m *CreateTodoListResponse_ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoListResponse_ErrorResponse) ProtoMessage()    {}
func (*CreateTodoListResponse_ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{1, 1}
}

func (m *CreateTodoListResponse_ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Unmarshal(m, b)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Merge(m, src)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Size(m)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListResponse_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListResponse_ErrorResponse proto.InternalMessageInfo

func (m *CreateTodoListResponse_ErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *CreateTodoListResponse_ErrorResponse) GetStatusCode() string {
	if m != nil {
		return m.StatusCode
	}
	return ""
}

type SubTask struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Offset               int32    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask) Reset()         { *m = SubTask{} }
func (m *SubTask) String() string { return proto.CompactTextString(m) }
func (*SubTask) ProtoMessage()    {}
func (*SubTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{2}
}

func (m *SubTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask.Unmarshal(m, b)
}
func (m *SubTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask.Marshal(b, m, deterministic)
}
func (m *SubTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask.Merge(m, src)
}
func (m *SubTask) XXX_Size() int {
	return xxx_messageInfo_SubTask.Size(m)
}
func (m *SubTask) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask proto.InternalMessageInfo

func (m *SubTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubTask) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SubTask) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SubTask) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Notification struct {
	Email                *Email       `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Message              *Textmessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{3}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetEmail() *Email {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Notification) GetMessage() *Textmessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type Email struct {
	Notification         bool     `protobuf:"varint,1,opt,name=notification,proto3" json:"notification,omitempty"`
	EmailID              string   `protobuf:"bytes,2,opt,name=emailID,proto3" json:"emailID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Email) Reset()         { *m = Email{} }
func (m *Email) String() string { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()    {}
func (*Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{4}
}

func (m *Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Email.Unmarshal(m, b)
}
func (m *Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Email.Marshal(b, m, deterministic)
}
func (m *Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Email.Merge(m, src)
}
func (m *Email) XXX_Size() int {
	return xxx_messageInfo_Email.Size(m)
}
func (m *Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Email proto.InternalMessageInfo

func (m *Email) GetNotification() bool {
	if m != nil {
		return m.Notification
	}
	return false
}

func (m *Email) GetEmailID() string {
	if m != nil {
		return m.EmailID
	}
	return ""
}

type Textmessage struct {
	Notification         bool     `protobuf:"varint,1,opt,name=notification,proto3" json:"notification,omitempty"`
	CountryCode          string   `protobuf:"bytes,2,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	MobileNumber         string   `protobuf:"bytes,3,opt,name=mobileNumber,proto3" json:"mobileNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Textmessage) Reset()         { *m = Textmessage{} }
func (m *Textmessage) String() string { return proto.CompactTextString(m) }
func (*Textmessage) ProtoMessage()    {}
func (*Textmessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{5}
}

func (m *Textmessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Textmessage.Unmarshal(m, b)
}
func (m *Textmessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Textmessage.Marshal(b, m, deterministic)
}
func (m *Textmessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Textmessage.Merge(m, src)
}
func (m *Textmessage) XXX_Size() int {
	return xxx_messageInfo_Textmessage.Size(m)
}
func (m *Textmessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Textmessage.DiscardUnknown(m)
}

var xxx_messageInfo_Textmessage proto.InternalMessageInfo

func (m *Textmessage) GetNotification() bool {
	if m != nil {
		return m.Notification
	}
	return false
}

func (m *Textmessage) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *Textmessage) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

func init() {
	proto.RegisterEnum("todolist.Tags", Tags_name, Tags_value)
	proto.RegisterType((*CreateTodoListRequest)(nil), "todolist.CreateTodoListRequest")
	proto.RegisterType((*CreateTodoListResponse)(nil), "todolist.CreateTodoListResponse")
	proto.RegisterType((*CreateTodoListResponse_CreateTodoListUserInformationResponse)(nil), "todolist.CreateTodoListResponse.CreateTodoListUserInformationResponse")
	proto.RegisterType((*CreateTodoListResponse_ErrorResponse)(nil), "todolist.CreateTodoListResponse.ErrorResponse")
	proto.RegisterType((*SubTask)(nil), "todolist.subTask")
	proto.RegisterType((*Notification)(nil), "todolist.notification")
	proto.RegisterType((*Email)(nil), "todolist.email")
	proto.RegisterType((*Textmessage)(nil), "todolist.textmessage")
}

func init() { proto.RegisterFile("todo_list_activity_manager.proto", fileDescriptor_516e47f255818132) }

var fileDescriptor_516e47f255818132 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x6f, 0xda, 0x3e,
	0x14, 0x6d, 0x02, 0x04, 0x7a, 0xd3, 0x3f, 0xfc, 0xfc, 0x53, 0x51, 0xd4, 0x87, 0x2d, 0x8a, 0x56,
	0x09, 0x6d, 0x12, 0x93, 0xd8, 0xc3, 0xa4, 0xbd, 0x31, 0xa0, 0x55, 0x54, 0xfe, 0x74, 0x86, 0x3d,
	0xb4, 0x2f, 0x95, 0x01, 0xc3, 0xa2, 0x35, 0x31, 0xb3, 0x9d, 0x6d, 0xfd, 0x8a, 0x7b, 0xda, 0x47,
	0x9a, 0x6c, 0x0c, 0x49, 0xba, 0x55, 0xed, 0x5b, 0xee, 0x39, 0x27, 0xc7, 0xf6, 0xf5, 0xf1, 0x05,
	0x5f, 0xb2, 0x05, 0xbb, 0xbd, 0x8b, 0x84, 0xbc, 0x25, 0x73, 0x19, 0x7d, 0x8f, 0xe4, 0xfd, 0x6d,
	0x4c, 0x12, 0xb2, 0xa2, 0xbc, 0xb5, 0xe6, 0x4c, 0x32, 0x54, 0x53, 0x0a, 0x25, 0x08, 0x7e, 0xdb,
	0x70, 0xd2, 0xe5, 0x94, 0x48, 0x3a, 0x65, 0x0b, 0x36, 0x88, 0x84, 0xc4, 0xf4, 0x5b, 0x4a, 0x85,
	0x44, 0x0d, 0x70, 0x52, 0x41, 0x79, 0xd8, 0xf3, 0x2c, 0xdf, 0x6a, 0xee, 0x63, 0x53, 0x21, 0x04,
	0xe5, 0x84, 0xc4, 0xd4, 0xb3, 0x35, 0xaa, 0xbf, 0x91, 0x0f, 0xee, 0x82, 0x8a, 0x39, 0x8f, 0xd6,
	0x32, 0x62, 0x89, 0x57, 0xd2, 0x54, 0x1e, 0x42, 0x1e, 0x54, 0x17, 0x29, 0xed, 0x11, 0x49, 0xbd,
	0xb2, 0x66, 0xb7, 0x25, 0x7a, 0x05, 0x87, 0xf3, 0x54, 0x48, 0x16, 0xf7, 0x0c, 0x5f, 0xd1, 0x7c,
	0x11, 0x44, 0x2f, 0x00, 0x38, 0x5d, 0x53, 0x22, 0x87, 0x6c, 0x41, 0x3d, 0xc7, 0xb7, 0x9a, 0x35,
	0x9c, 0x43, 0x90, 0x0f, 0x25, 0x49, 0x56, 0x5e, 0xd5, 0xb7, 0x9a, 0x47, 0xed, 0xa3, 0xd6, 0xf6,
	0x7c, 0xad, 0x29, 0x59, 0x09, 0xac, 0x28, 0xf4, 0x01, 0x0e, 0x12, 0x26, 0xa3, 0x65, 0x34, 0x27,
	0x7a, 0x93, 0x35, 0xdf, 0x6a, 0xba, 0xed, 0x46, 0x26, 0xcd, 0xb3, 0xb8, 0xa0, 0x45, 0x6f, 0xa0,
	0x2a, 0xd2, 0xd9, 0x94, 0x88, 0xaf, 0xde, 0xbe, 0x5f, 0x6a, 0xba, 0xed, 0xff, 0xb2, 0xdf, 0x0c,
	0x81, 0xb7, 0x8a, 0xe0, 0x97, 0x0d, 0x8d, 0x87, 0x2d, 0x15, 0x6b, 0x96, 0x08, 0x8a, 0x6e, 0xa0,
	0xbc, 0x20, 0x92, 0x78, 0x96, 0x36, 0x39, 0xcf, 0x4c, 0xfe, 0xad, 0x7f, 0x00, 0x7f, 0x56, 0x17,
	0x90, 0x2c, 0x19, 0x8f, 0x37, 0x7b, 0x34, 0x2a, 0xac, 0x3d, 0xd1, 0x39, 0x38, 0x94, 0x73, 0xc6,
	0x85, 0x67, 0x6b, 0xf7, 0xd6, 0x93, 0xee, 0x7d, 0x25, 0xdf, 0xb9, 0x98, 0xbf, 0x4f, 0xdf, 0xc3,
	0xd9, 0xb3, 0x96, 0x45, 0x47, 0x60, 0xef, 0xc2, 0x61, 0x87, 0xbd, 0xd3, 0x09, 0x1c, 0x16, 0x1c,
	0x51, 0x00, 0x07, 0xda, 0x73, 0x48, 0x85, 0x20, 0x2b, 0x6a, 0xa4, 0x05, 0x4c, 0xdd, 0xab, 0x90,
	0x44, 0xa6, 0xa2, 0xab, 0xee, 0x75, 0x93, 0xa9, 0x1c, 0x12, 0xb0, 0x5d, 0xe7, 0x77, 0xc1, 0xb3,
	0x1e, 0x0f, 0x9e, 0xfd, 0x77, 0xf0, 0x1a, 0xe0, 0x6c, 0xec, 0x74, 0x2a, 0x6b, 0xd8, 0x54, 0x0a,
	0x67, 0xcb, 0xa5, 0xa0, 0x52, 0xe7, 0xb1, 0x82, 0x4d, 0x15, 0x2c, 0x8b, 0x31, 0x41, 0x67, 0x50,
	0xa1, 0x31, 0x89, 0xee, 0xf4, 0xb2, 0x6e, 0xfb, 0x38, 0xeb, 0xaa, 0x86, 0xf1, 0x86, 0x45, 0x6f,
	0xa1, 0x1a, 0x9b, 0x63, 0xda, 0x5a, 0x78, 0x92, 0x09, 0x25, 0xfd, 0x29, 0x0d, 0x89, 0xb7, 0xaa,
	0xa0, 0x6f, 0x7c, 0x55, 0x97, 0x0a, 0xb9, 0xb4, 0xf4, 0x36, 0x8b, 0x9b, 0xf0, 0xa0, 0xaa, 0xc5,
	0x61, 0xcf, 0x1c, 0x71, 0x5b, 0x06, 0x3f, 0xc0, 0xcd, 0xd9, 0x3f, 0xcb, 0xcc, 0x07, 0x77, 0xce,
	0xd2, 0x44, 0xf2, 0xfb, 0x5c, 0xcf, 0xf3, 0x90, 0x72, 0x89, 0xd9, 0x2c, 0xba, 0xa3, 0xa3, 0x34,
	0x9e, 0x51, 0x6e, 0xde, 0x73, 0x01, 0x7b, 0x3d, 0x83, 0xb2, 0x7a, 0x5b, 0xe8, 0x00, 0x6a, 0x57,
	0x38, 0x1c, 0xe3, 0x70, 0x7a, 0x5d, 0xdf, 0x43, 0x00, 0xce, 0x68, 0x8c, 0x87, 0x9d, 0x41, 0xdd,
	0x42, 0xfb, 0x50, 0x19, 0x8e, 0x47, 0xfd, 0xeb, 0xba, 0xad, 0xe0, 0xc9, 0xd5, 0x18, 0x4f, 0x27,
	0xf5, 0x12, 0x72, 0xa1, 0xda, 0x1d, 0x5f, 0x5e, 0x86, 0xa3, 0x8b, 0x7a, 0x19, 0x1d, 0x83, 0x3b,
	0x19, 0x77, 0x3b, 0xe1, 0x20, 0xbc, 0x51, 0x40, 0x45, 0xb1, 0x17, 0xfd, 0x51, 0x1f, 0x77, 0x06,
	0x75, 0xa7, 0xfd, 0x05, 0xfe, 0x57, 0x21, 0xec, 0x98, 0x21, 0x36, 0xdc, 0xcc, 0x30, 0xf4, 0x09,
	0x20, 0x4b, 0x28, 0x7a, 0xf9, 0x78, 0xce, 0xf5, 0x20, 0x3b, 0xf5, 0x9f, 0x7a, 0x08, 0xc1, 0xde,
	0x47, 0xb8, 0xd9, 0x8d, 0xc4, 0x99, 0xa3, 0x67, 0xe4, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0x54, 0x3e, 0x68, 0x47, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoActivityManagerClient is the client API for TodoActivityManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoActivityManagerClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoListRequest, opts ...grpc.CallOption) (*CreateTodoListResponse, error)
}

type todoActivityManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoActivityManagerClient(cc grpc.ClientConnInterface) TodoActivityManagerClient {
	return &todoActivityManagerClient{cc}
}

func (c *todoActivityManagerClient) CreateTodo(ctx context.Context, in *CreateTodoListRequest, opts ...grpc.CallOption) (*CreateTodoListResponse, error) {
	out := new(CreateTodoListResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoActivityManager/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoActivityManagerServer is the server API for TodoActivityManager service.
type TodoActivityManagerServer interface {
	CreateTodo(context.Context, *CreateTodoListRequest) (*CreateTodoListResponse, error)
}

// UnimplementedTodoActivityManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTodoActivityManagerServer struct {
}

func (*UnimplementedTodoActivityManagerServer) CreateTodo(ctx context.Context, req *CreateTodoListRequest) (*CreateTodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}

func RegisterTodoActivityManagerServer(s *grpc.Server, srv TodoActivityManagerServer) {
	s.RegisterService(&_TodoActivityManager_serviceDesc, srv)
}

func _TodoActivityManager_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoActivityManagerServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoActivityManager/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoActivityManagerServer).CreateTodo(ctx, req.(*CreateTodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoActivityManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.TodoActivityManager",
	HandlerType: (*TodoActivityManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoActivityManager_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo_list_activity_manager.proto",
}