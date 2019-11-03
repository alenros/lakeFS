// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type APICredentials_Type int32

const (
	APICredentials_CREDENTIAL_TYPE_USER        APICredentials_Type = 0
	APICredentials_CREDENTIAL_TYPE_APPLICATION APICredentials_Type = 1
)

var APICredentials_Type_name = map[int32]string{
	0: "CREDENTIAL_TYPE_USER",
	1: "CREDENTIAL_TYPE_APPLICATION",
}

var APICredentials_Type_value = map[string]int32{
	"CREDENTIAL_TYPE_USER":        0,
	"CREDENTIAL_TYPE_APPLICATION": 1,
}

func (x APICredentials_Type) String() string {
	return proto.EnumName(APICredentials_Type_name, int32(x))
}

func (APICredentials_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0, 0}
}

type APICredentials struct {
	AccessKeyId          string              `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	AccessSecretKey      string              `protobuf:"bytes,2,opt,name=access_secret_key,json=accessSecretKey,proto3" json:"access_secret_key,omitempty"`
	CredentialType       APICredentials_Type `protobuf:"varint,3,opt,name=credential_type,json=credentialType,proto3,enum=APICredentials_Type" json:"credential_type,omitempty"`
	ClientId             string              `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	EntityId             string              `protobuf:"bytes,5,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	IssuedDate           int64               `protobuf:"varint,6,opt,name=issued_date,json=issuedDate,proto3" json:"issued_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *APICredentials) Reset()         { *m = APICredentials{} }
func (m *APICredentials) String() string { return proto.CompactTextString(m) }
func (*APICredentials) ProtoMessage()    {}
func (*APICredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *APICredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APICredentials.Unmarshal(m, b)
}
func (m *APICredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APICredentials.Marshal(b, m, deterministic)
}
func (m *APICredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APICredentials.Merge(m, src)
}
func (m *APICredentials) XXX_Size() int {
	return xxx_messageInfo_APICredentials.Size(m)
}
func (m *APICredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_APICredentials.DiscardUnknown(m)
}

var xxx_messageInfo_APICredentials proto.InternalMessageInfo

func (m *APICredentials) GetAccessKeyId() string {
	if m != nil {
		return m.AccessKeyId
	}
	return ""
}

func (m *APICredentials) GetAccessSecretKey() string {
	if m != nil {
		return m.AccessSecretKey
	}
	return ""
}

func (m *APICredentials) GetCredentialType() APICredentials_Type {
	if m != nil {
		return m.CredentialType
	}
	return APICredentials_CREDENTIAL_TYPE_USER
}

func (m *APICredentials) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *APICredentials) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *APICredentials) GetIssuedDate() int64 {
	if m != nil {
		return m.IssuedDate
	}
	return 0
}

type Client struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type User struct {
	ClientId             string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string            `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FullName             string            `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	PasswdHash           string            `protobuf:"bytes,5,opt,name=passwd_hash,json=passwdHash,proto3" json:"passwd_hash,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Groups               []string          `protobuf:"bytes,7,rep,name=groups,proto3" json:"groups,omitempty"`
	Roles                []string          `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetPasswdHash() string {
	if m != nil {
		return m.PasswdHash
	}
	return ""
}

func (m *User) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *User) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *User) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Application struct {
	ClientId             string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AccessKeyId          string            `protobuf:"bytes,4,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	AccessKeySecret      string            `protobuf:"bytes,5,opt,name=access_key_secret,json=accessKeySecret,proto3" json:"access_key_secret,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Groups               []string          `protobuf:"bytes,7,rep,name=groups,proto3" json:"groups,omitempty"`
	Roles                []string          `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetAccessKeyId() string {
	if m != nil {
		return m.AccessKeyId
	}
	return ""
}

func (m *Application) GetAccessKeySecret() string {
	if m != nil {
		return m.AccessKeySecret
	}
	return ""
}

func (m *Application) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Application) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Application) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Group struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Roles                []string `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Group) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	ClientId             string    `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Id                   string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Policies             []*Policy `protobuf:"bytes,4,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Policy struct {
	Permission           string   `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	Arn                  string   `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{6}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func (m *Policy) GetArn() string {
	if m != nil {
		return m.Arn
	}
	return ""
}

func init() {
	proto.RegisterEnum("APICredentials_Type", APICredentials_Type_name, APICredentials_Type_value)
	proto.RegisterType((*APICredentials)(nil), "APICredentials")
	proto.RegisterType((*Client)(nil), "Client")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterMapType((map[string]string)(nil), "User.MetadataEntry")
	proto.RegisterType((*Application)(nil), "Application")
	proto.RegisterMapType((map[string]string)(nil), "Application.MetadataEntry")
	proto.RegisterType((*Group)(nil), "Group")
	proto.RegisterType((*Role)(nil), "Role")
	proto.RegisterType((*Policy)(nil), "Policy")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x5f, 0x6b, 0xd4, 0x4e,
	0x14, 0xfd, 0xe5, 0x4f, 0xd3, 0xdd, 0x1b, 0xba, 0xed, 0x6f, 0x2c, 0x12, 0x5a, 0xb0, 0x4b, 0x7c,
	0x59, 0x44, 0x22, 0x54, 0x10, 0xa9, 0xf8, 0x10, 0xb6, 0x8b, 0x86, 0xd6, 0xba, 0xa4, 0xdb, 0x07,
	0x5f, 0x0c, 0x63, 0x72, 0x75, 0x07, 0x27, 0x7f, 0xc8, 0xcc, 0x2a, 0xf9, 0x3c, 0x7e, 0x0e, 0x1f,
	0xfc, 0x66, 0x32, 0x99, 0x34, 0xdd, 0x6d, 0x41, 0x90, 0xbe, 0xf8, 0x36, 0xf7, 0xdc, 0x93, 0x7b,
	0xce, 0x3d, 0x5c, 0x02, 0x6e, 0x5e, 0x66, 0xc8, 0x83, 0xaa, 0x2e, 0x65, 0xe9, 0xff, 0x34, 0x61,
	0x14, 0xce, 0xa3, 0x69, 0x8d, 0x19, 0x16, 0x92, 0x51, 0x2e, 0x88, 0x0f, 0x3b, 0x34, 0x4d, 0x51,
	0x88, 0xe4, 0x2b, 0x36, 0x09, 0xcb, 0x3c, 0x63, 0x6c, 0x4c, 0x86, 0xb1, 0xab, 0xc1, 0x33, 0x6c,
	0xa2, 0x8c, 0x3c, 0x81, 0xff, 0x3b, 0x8e, 0xc0, 0xb4, 0x46, 0xa9, 0xa8, 0x9e, 0xd9, 0xf2, 0x76,
	0x75, 0xe3, 0xb2, 0xc5, 0xcf, 0xb0, 0x21, 0xaf, 0x61, 0x37, 0xed, 0xc7, 0x27, 0xb2, 0xa9, 0xd0,
	0xb3, 0xc6, 0xc6, 0x64, 0x74, 0xbc, 0x1f, 0x6c, 0x2a, 0x07, 0x8b, 0xa6, 0xc2, 0x78, 0x74, 0x43,
	0x56, 0x35, 0x39, 0x84, 0x61, 0xca, 0x19, 0x16, 0x52, 0x59, 0xb1, 0x5b, 0x89, 0x81, 0x06, 0xa2,
	0x4c, 0x35, 0x15, 0x55, 0xb6, 0x3e, 0xb7, 0x74, 0x53, 0x03, 0x51, 0x46, 0x8e, 0xc0, 0x65, 0x42,
	0xac, 0x30, 0x4b, 0x32, 0x2a, 0xd1, 0x73, 0xc6, 0xc6, 0xc4, 0x8a, 0x41, 0x43, 0xa7, 0x54, 0xa2,
	0x1f, 0x82, 0xdd, 0x4a, 0x78, 0xb0, 0x3f, 0x8d, 0x67, 0xa7, 0xb3, 0x8b, 0x45, 0x14, 0x9e, 0x27,
	0x8b, 0x0f, 0xf3, 0x59, 0x72, 0x75, 0x39, 0x8b, 0xf7, 0xfe, 0x23, 0x47, 0x70, 0x78, 0xbb, 0x13,
	0xce, 0xe7, 0xe7, 0xd1, 0x34, 0x5c, 0x44, 0xef, 0x2f, 0xf6, 0x0c, 0xff, 0x29, 0x38, 0xd3, 0xd6,
	0x0c, 0x19, 0x81, 0xd9, 0x67, 0x65, 0xb2, 0x8c, 0x10, 0xb0, 0x0b, 0x9a, 0x63, 0x97, 0x4a, 0xfb,
	0xf6, 0x7f, 0x98, 0x60, 0x5f, 0x09, 0xac, 0x37, 0x97, 0x32, 0x6e, 0x2d, 0xa5, 0x27, 0x99, 0xfd,
	0xa4, 0x7d, 0xd8, 0xc2, 0x9c, 0x32, 0xde, 0xc6, 0x36, 0x8c, 0x75, 0xa1, 0x46, 0x7c, 0x5e, 0x71,
	0x9e, 0xb4, 0x22, 0x5d, 0x2e, 0x0a, 0xb8, 0xa0, 0x39, 0xaa, 0xd5, 0x2b, 0x2a, 0xc4, 0xf7, 0x2c,
	0x59, 0x52, 0xb1, 0xec, 0x92, 0x01, 0x0d, 0xbd, 0xa5, 0x62, 0x49, 0x9e, 0xc1, 0x20, 0x47, 0x49,
	0x33, 0x2a, 0xa9, 0xe7, 0x8c, 0xad, 0x89, 0x7b, 0xfc, 0x20, 0x50, 0xce, 0x82, 0x77, 0x1d, 0x3a,
	0x2b, 0x64, 0xdd, 0xc4, 0x3d, 0x89, 0x3c, 0x04, 0xe7, 0x4b, 0x5d, 0xae, 0x2a, 0xe1, 0x6d, 0x8f,
	0xad, 0xc9, 0x30, 0xee, 0x2a, 0x65, 0xae, 0x2e, 0x39, 0x0a, 0x6f, 0xd0, 0xc2, 0xba, 0x38, 0x78,
	0x05, 0x3b, 0x1b, 0x83, 0xc8, 0x1e, 0x58, 0xea, 0x44, 0xf4, 0xaa, 0xea, 0xa9, 0x3e, 0xfc, 0x46,
	0xf9, 0xea, 0x3a, 0x20, 0x5d, 0x9c, 0x98, 0x2f, 0x0d, 0xff, 0x97, 0x09, 0x6e, 0x58, 0x55, 0x9c,
	0xa5, 0x54, 0xb2, 0xb2, 0xf8, 0xbb, 0xb0, 0xae, 0x63, 0xb7, 0x6e, 0x62, 0xbf, 0x7b, 0xd1, 0xf6,
	0x9f, 0x2e, 0x5a, 0x71, 0xf4, 0x55, 0x77, 0xb9, 0xed, 0xf6, 0x3c, 0x7d, 0xd4, 0xe4, 0xc5, 0x9d,
	0xf0, 0x0e, 0x82, 0x35, 0xc3, 0xff, 0x42, 0x86, 0x1f, 0x61, 0xeb, 0x8d, 0x1a, 0x7e, 0xff, 0xf0,
	0x7a, 0x73, 0xf6, 0x9a, 0x39, 0x9f, 0x83, 0x1d, 0x97, 0x1c, 0xef, 0x3f, 0xfe, 0x31, 0x0c, 0xaa,
	0x92, 0xb3, 0x94, 0x75, 0x0a, 0xee, 0xf1, 0x76, 0x30, 0x57, 0x40, 0x13, 0xf7, 0x0d, 0xff, 0x04,
	0x1c, 0x8d, 0x91, 0x47, 0x00, 0x15, 0xd6, 0x39, 0x13, 0x82, 0x95, 0x45, 0x27, 0xb8, 0x86, 0xa8,
	0x8c, 0x68, 0x5d, 0x74, 0x9a, 0xea, 0xf9, 0xc9, 0x69, 0x7f, 0x74, 0xcf, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0x5e, 0x1e, 0xf7, 0xf7, 0x04, 0x00, 0x00,
}