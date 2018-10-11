// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Kademlia.proto

package protobuf

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KademliaMessageType_Type int32

const (
	KademliaMessageType_CALL     KademliaMessageType_Type = 0
	KademliaMessageType_CALLBACK KademliaMessageType_Type = 1
)

var KademliaMessageType_Type_name = map[int32]string{
	0: "CALL",
	1: "CALLBACK",
}

var KademliaMessageType_Type_value = map[string]int32{
	"CALL":     0,
	"CALLBACK": 1,
}

func (x KademliaMessageType_Type) String() string {
	return proto.EnumName(KademliaMessageType_Type_name, int32(x))
}

func (KademliaMessageType_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{0, 0}
}

type KademliaMessageCall_Type int32

const (
	KademliaMessageCall_PING  KademliaMessageCall_Type = 0
	KademliaMessageCall_FINDC KademliaMessageCall_Type = 1
)

var KademliaMessageCall_Type_name = map[int32]string{
	0: "PING",
	1: "FINDC",
}

var KademliaMessageCall_Type_value = map[string]int32{
	"PING":  0,
	"FINDC": 1,
}

func (x KademliaMessageCall_Type) String() string {
	return proto.EnumName(KademliaMessageCall_Type_name, int32(x))
}

func (KademliaMessageCall_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{1, 0}
}

type KademliaMessageCallBack_Type int32

const (
	KademliaMessageCallBack_PING  KademliaMessageCallBack_Type = 0
	KademliaMessageCallBack_FINDC KademliaMessageCallBack_Type = 1
)

var KademliaMessageCallBack_Type_name = map[int32]string{
	0: "PING",
	1: "FINDC",
}

var KademliaMessageCallBack_Type_value = map[string]int32{
	"PING":  0,
	"FINDC": 1,
}

func (x KademliaMessageCallBack_Type) String() string {
	return proto.EnumName(KademliaMessageCallBack_Type_name, int32(x))
}

func (KademliaMessageCallBack_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{2, 0}
}

type KademliaMessageType struct {
	Type                 KademliaMessageType_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protobuf.KademliaMessageType_Type" json:"type,omitempty"`
	SenderC              *Contact                 `protobuf:"bytes,2,opt,name=senderC,proto3" json:"senderC,omitempty"`
	Call                 *KademliaMessageCall     `protobuf:"bytes,3,opt,name=Call,proto3" json:"Call,omitempty"`
	Callback             *KademliaMessageCallBack `protobuf:"bytes,4,opt,name=Callback,proto3" json:"Callback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *KademliaMessageType) Reset()         { *m = KademliaMessageType{} }
func (m *KademliaMessageType) String() string { return proto.CompactTextString(m) }
func (*KademliaMessageType) ProtoMessage()    {}
func (*KademliaMessageType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{0}
}

func (m *KademliaMessageType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KademliaMessageType.Unmarshal(m, b)
}
func (m *KademliaMessageType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KademliaMessageType.Marshal(b, m, deterministic)
}
func (m *KademliaMessageType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KademliaMessageType.Merge(m, src)
}
func (m *KademliaMessageType) XXX_Size() int {
	return xxx_messageInfo_KademliaMessageType.Size(m)
}
func (m *KademliaMessageType) XXX_DiscardUnknown() {
	xxx_messageInfo_KademliaMessageType.DiscardUnknown(m)
}

var xxx_messageInfo_KademliaMessageType proto.InternalMessageInfo

func (m *KademliaMessageType) GetType() KademliaMessageType_Type {
	if m != nil {
		return m.Type
	}
	return KademliaMessageType_CALL
}

func (m *KademliaMessageType) GetSenderC() *Contact {
	if m != nil {
		return m.SenderC
	}
	return nil
}

func (m *KademliaMessageType) GetCall() *KademliaMessageCall {
	if m != nil {
		return m.Call
	}
	return nil
}

func (m *KademliaMessageType) GetCallback() *KademliaMessageCallBack {
	if m != nil {
		return m.Callback
	}
	return nil
}

type KademliaMessageCall struct {
	Id                   int32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 KademliaMessageCall_Type `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.KademliaMessageCall_Type" json:"type,omitempty"`
	MessageString        string                   `protobuf:"bytes,3,opt,name=messageString,proto3" json:"messageString,omitempty"`
	Info                 string                   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *KademliaMessageCall) Reset()         { *m = KademliaMessageCall{} }
func (m *KademliaMessageCall) String() string { return proto.CompactTextString(m) }
func (*KademliaMessageCall) ProtoMessage()    {}
func (*KademliaMessageCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{1}
}

func (m *KademliaMessageCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KademliaMessageCall.Unmarshal(m, b)
}
func (m *KademliaMessageCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KademliaMessageCall.Marshal(b, m, deterministic)
}
func (m *KademliaMessageCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KademliaMessageCall.Merge(m, src)
}
func (m *KademliaMessageCall) XXX_Size() int {
	return xxx_messageInfo_KademliaMessageCall.Size(m)
}
func (m *KademliaMessageCall) XXX_DiscardUnknown() {
	xxx_messageInfo_KademliaMessageCall.DiscardUnknown(m)
}

var xxx_messageInfo_KademliaMessageCall proto.InternalMessageInfo

func (m *KademliaMessageCall) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KademliaMessageCall) GetType() KademliaMessageCall_Type {
	if m != nil {
		return m.Type
	}
	return KademliaMessageCall_PING
}

func (m *KademliaMessageCall) GetMessageString() string {
	if m != nil {
		return m.MessageString
	}
	return ""
}

func (m *KademliaMessageCall) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type KademliaMessageCallBack struct {
	Id                   int32                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 KademliaMessageCallBack_Type `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.KademliaMessageCallBack_Type" json:"type,omitempty"`
	Contacts             *Contact                     `protobuf:"bytes,3,opt,name=contacts,proto3" json:"contacts,omitempty"`
	Info                 string                       `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	CoontactList         []*Contact                   `protobuf:"bytes,5,rep,name=coontactList,proto3" json:"coontactList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *KademliaMessageCallBack) Reset()         { *m = KademliaMessageCallBack{} }
func (m *KademliaMessageCallBack) String() string { return proto.CompactTextString(m) }
func (*KademliaMessageCallBack) ProtoMessage()    {}
func (*KademliaMessageCallBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{2}
}

func (m *KademliaMessageCallBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KademliaMessageCallBack.Unmarshal(m, b)
}
func (m *KademliaMessageCallBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KademliaMessageCallBack.Marshal(b, m, deterministic)
}
func (m *KademliaMessageCallBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KademliaMessageCallBack.Merge(m, src)
}
func (m *KademliaMessageCallBack) XXX_Size() int {
	return xxx_messageInfo_KademliaMessageCallBack.Size(m)
}
func (m *KademliaMessageCallBack) XXX_DiscardUnknown() {
	xxx_messageInfo_KademliaMessageCallBack.DiscardUnknown(m)
}

var xxx_messageInfo_KademliaMessageCallBack proto.InternalMessageInfo

func (m *KademliaMessageCallBack) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KademliaMessageCallBack) GetType() KademliaMessageCallBack_Type {
	if m != nil {
		return m.Type
	}
	return KademliaMessageCallBack_PING
}

func (m *KademliaMessageCallBack) GetContacts() *Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *KademliaMessageCallBack) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *KademliaMessageCallBack) GetCoontactList() []*Contact {
	if m != nil {
		return m.CoontactList
	}
	return nil
}

type Contact struct {
	ContactID            string   `protobuf:"bytes,1,opt,name=contactID,proto3" json:"contactID,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Xor                  string   `protobuf:"bytes,3,opt,name=xor,proto3" json:"xor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_4162160cb5558cc0, []int{3}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetContactID() string {
	if m != nil {
		return m.ContactID
	}
	return ""
}

func (m *Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contact) GetXor() string {
	if m != nil {
		return m.Xor
	}
	return ""
}

func init() {
	proto.RegisterEnum("protobuf.KademliaMessageType_Type", KademliaMessageType_Type_name, KademliaMessageType_Type_value)
	proto.RegisterEnum("protobuf.KademliaMessageCall_Type", KademliaMessageCall_Type_name, KademliaMessageCall_Type_value)
	proto.RegisterEnum("protobuf.KademliaMessageCallBack_Type", KademliaMessageCallBack_Type_name, KademliaMessageCallBack_Type_value)
	proto.RegisterType((*KademliaMessageType)(nil), "protobuf.KademliaMessageType")
	proto.RegisterType((*KademliaMessageCall)(nil), "protobuf.KademliaMessageCall")
	proto.RegisterType((*KademliaMessageCallBack)(nil), "protobuf.KademliaMessageCallBack")
	proto.RegisterType((*Contact)(nil), "protobuf.contact")
}

func init() { proto.RegisterFile("Kademlia.proto", fileDescriptor_4162160cb5558cc0) }

var fileDescriptor_4162160cb5558cc0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x69, 0x29, 0xd2, 0x5e, 0x91, 0xd4, 0x71, 0x61, 0x13, 0x7f, 0x82, 0x8d, 0x31, 0x24,
	0xc6, 0x26, 0x62, 0x74, 0x61, 0xe2, 0x02, 0x4a, 0x34, 0x04, 0x24, 0x66, 0xf0, 0x05, 0x86, 0x76,
	0x20, 0x0d, 0xd0, 0x92, 0x4e, 0x4d, 0xe4, 0x1d, 0x7c, 0x1c, 0x1f, 0xcf, 0x85, 0xe9, 0xa5, 0xe5,
	0x27, 0x14, 0xd8, 0x74, 0x6e, 0x67, 0xbe, 0x33, 0x73, 0xcf, 0xb9, 0x50, 0x6e, 0x33, 0x97, 0x4f,
	0xc6, 0x1e, 0xb3, 0xa6, 0x61, 0x10, 0x05, 0x44, 0xc5, 0xa5, 0xff, 0x35, 0x30, 0x7f, 0x64, 0x38,
	0x49, 0x0f, 0xdf, 0xb9, 0x10, 0x6c, 0xc8, 0x3f, 0x67, 0x53, 0x4e, 0x9e, 0x40, 0x89, 0x66, 0x53,
	0x6e, 0x48, 0x15, 0xa9, 0x5a, 0xae, 0x99, 0x56, 0x2a, 0xb0, 0x32, 0x60, 0x2b, 0xfe, 0x50, 0xe4,
	0xc9, 0x2d, 0x14, 0x05, 0xf7, 0x5d, 0x1e, 0xda, 0x86, 0x5c, 0x91, 0xaa, 0x87, 0xb5, 0xe3, 0xa5,
	0xd4, 0x09, 0xfc, 0x88, 0x39, 0x11, 0x4d, 0x09, 0x72, 0x0f, 0x8a, 0xcd, 0xc6, 0x63, 0x23, 0x8f,
	0xe4, 0xc5, 0xd6, 0x47, 0x62, 0x88, 0x22, 0x4a, 0x5e, 0x40, 0x8d, 0xd7, 0x3e, 0x73, 0x46, 0x86,
	0x82, 0xb2, 0xab, 0x9d, 0xb2, 0x06, 0x73, 0x46, 0x74, 0x21, 0x31, 0x2f, 0x41, 0x41, 0x7b, 0x2a,
	0x28, 0x76, 0xbd, 0xd3, 0xd1, 0x73, 0xa4, 0x04, 0x6a, 0x5c, 0x35, 0xea, 0x76, 0x5b, 0x97, 0xcc,
	0x5f, 0x69, 0x23, 0x0e, 0x7c, 0xb6, 0x0c, 0xb2, 0xe7, 0x62, 0x18, 0x05, 0x2a, 0x7b, 0xee, 0x22,
	0x1e, 0x79, 0x4f, 0x3c, 0xb1, 0x78, 0x35, 0x9e, 0x6b, 0x38, 0x9a, 0xcc, 0x4f, 0x7a, 0x51, 0xe8,
	0xf9, 0x43, 0xb4, 0xae, 0xd1, 0xf5, 0x4d, 0x42, 0x40, 0xf1, 0xfc, 0x41, 0x80, 0x06, 0x35, 0x8a,
	0xb5, 0x79, 0xb6, 0xec, 0xfc, 0xa3, 0xd5, 0x7d, 0xd3, 0x73, 0x44, 0x83, 0xc2, 0x6b, 0xab, 0xdb,
	0xb4, 0x75, 0xc9, 0xfc, 0x93, 0xe0, 0x74, 0x8b, 0xf9, 0x8d, 0xd6, 0x9f, 0xd7, 0x5a, 0xbf, 0xd9,
	0x9b, 0xde, 0x6a, 0xfb, 0x77, 0xa0, 0x26, 0x43, 0x14, 0xc9, 0xd0, 0x32, 0xc6, 0xbb, 0x40, 0xb2,
	0x7c, 0x90, 0x47, 0x28, 0x39, 0xc1, 0x1c, 0xe8, 0x78, 0x22, 0x32, 0x0a, 0x95, 0x7c, 0xf6, 0x35,
	0x6b, 0xd8, 0x6e, 0xfb, 0x3d, 0x28, 0x26, 0x2a, 0x72, 0x0e, 0x5a, 0x52, 0xb6, 0x9a, 0x68, 0x5a,
	0xa3, 0xcb, 0x0d, 0x62, 0x40, 0x91, 0xb9, 0x6e, 0xc8, 0x85, 0x40, 0xfb, 0x1a, 0x4d, 0x7f, 0x89,
	0x0e, 0xf9, 0xef, 0x20, 0x4c, 0xc6, 0x11, 0x97, 0xfd, 0x03, 0xec, 0xe8, 0xe1, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0x9a, 0xe9, 0x47, 0x3c, 0x03, 0x00, 0x00,
}
