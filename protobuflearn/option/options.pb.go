// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: options.proto

package option

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExtObj struct {
	FooString            string   `protobuf:"bytes,1,opt,name=foo_string,json=fooString,proto3" json:"foo_string,omitempty"`
	BarInt               int64    `protobuf:"varint,2,opt,name=bar_int,json=barInt,proto3" json:"bar_int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtObj) Reset()         { *m = ExtObj{} }
func (m *ExtObj) String() string { return proto.CompactTextString(m) }
func (*ExtObj) ProtoMessage()    {}
func (*ExtObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}
func (m *ExtObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtObj.Unmarshal(m, b)
}
func (m *ExtObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtObj.Marshal(b, m, deterministic)
}
func (m *ExtObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtObj.Merge(m, src)
}
func (m *ExtObj) XXX_Size() int {
	return xxx_messageInfo_ExtObj.Size(m)
}
func (m *ExtObj) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtObj.DiscardUnknown(m)
}

var xxx_messageInfo_ExtObj proto.InternalMessageInfo

func (m *ExtObj) GetFooString() string {
	if m != nil {
		return m.FooString
	}
	return ""
}

func (m *ExtObj) GetBarInt() int64 {
	if m != nil {
		return m.BarInt
	}
	return 0
}

type MessageOption struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageOption) Reset()         { *m = MessageOption{} }
func (m *MessageOption) String() string { return proto.CompactTextString(m) }
func (*MessageOption) ProtoMessage()    {}
func (*MessageOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}
func (m *MessageOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageOption.Unmarshal(m, b)
}
func (m *MessageOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageOption.Marshal(b, m, deterministic)
}
func (m *MessageOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOption.Merge(m, src)
}
func (m *MessageOption) XXX_Size() int {
	return xxx_messageInfo_MessageOption.Size(m)
}
func (m *MessageOption) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOption.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOption proto.InternalMessageInfo

func (m *MessageOption) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

// 字段级别的option定义方式不使用option关键字，格式为：用[]包裹的用逗号分隔的k=v形式的数组。
type FieldOption struct {
	// 自定义的option
	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	// protobuf内置的option
	Bar                  string   `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldOption) Reset()         { *m = FieldOption{} }
func (m *FieldOption) String() string { return proto.CompactTextString(m) }
func (*FieldOption) ProtoMessage()    {}
func (*FieldOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}
func (m *FieldOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOption.Unmarshal(m, b)
}
func (m *FieldOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOption.Marshal(b, m, deterministic)
}
func (m *FieldOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOption.Merge(m, src)
}
func (m *FieldOption) XXX_Size() int {
	return xxx_messageInfo_FieldOption.Size(m)
}
func (m *FieldOption) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOption.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOption proto.InternalMessageInfo

func (m *FieldOption) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

// Deprecated: Do not use.
func (m *FieldOption) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

var E_FileOptString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1011,
	Name:          "example.option.file_opt_string",
	Tag:           "bytes,1011,opt,name=file_opt_string",
	Filename:      "options.proto",
}

var E_FileOptObj = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FileOptions)(nil),
	ExtensionType: (*ExtObj)(nil),
	Field:         1012,
	Name:          "example.option.file_opt_obj",
	Tag:           "bytes,1012,opt,name=file_opt_obj",
	Filename:      "options.proto",
}

var E_MsgOptString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1011,
	Name:          "example.option.msg_opt_string",
	Tag:           "bytes,1011,opt,name=msg_opt_string",
	Filename:      "options.proto",
}

var E_MsgOptObj = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*ExtObj)(nil),
	Field:         1012,
	Name:          "example.option.msg_opt_obj",
	Tag:           "bytes,1012,opt,name=msg_opt_obj",
	Filename:      "options.proto",
}

var E_FieldOptString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1021,
	Name:          "example.option.field_opt_string",
	Tag:           "bytes,1021,opt,name=field_opt_string",
	Filename:      "options.proto",
}

var E_FieldOptObj = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*ExtObj)(nil),
	Field:         1022,
	Name:          "example.option.field_opt_obj",
	Tag:           "bytes,1022,opt,name=field_opt_obj",
	Filename:      "options.proto",
}

func init() {
	proto.RegisterType((*ExtObj)(nil), "example.option.extObj")
	proto.RegisterType((*MessageOption)(nil), "example.option.MessageOption")
	proto.RegisterType((*FieldOption)(nil), "example.option.FieldOption")
	proto.RegisterExtension(E_FileOptString)
	proto.RegisterExtension(E_FileOptObj)
	proto.RegisterExtension(E_MsgOptString)
	proto.RegisterExtension(E_MsgOptObj)
	proto.RegisterExtension(E_FieldOptString)
	proto.RegisterExtension(E_FieldOptObj)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0xe5, 0x46, 0x24, 0xf2, 0xe4, 0x0f, 0x96, 0x85, 0xc0, 0x54, 0xad, 0x08, 0x11, 0x87,
	0x9e, 0x36, 0x2a, 0xdc, 0x72, 0x31, 0x8a, 0x44, 0x28, 0x07, 0x64, 0xc9, 0x1c, 0x40, 0xbd, 0x44,
	0xde, 0x66, 0xd7, 0xb2, 0xb5, 0xf1, 0x58, 0xbb, 0x8b, 0x28, 0xaf, 0xd2, 0x63, 0x1f, 0x8c, 0x07,
	0xe0, 0x08, 0x5c, 0x41, 0x68, 0x77, 0x1d, 0x37, 0x01, 0x84, 0x6f, 0xf6, 0x6f, 0x77, 0x3e, 0x7f,
	0x33, 0x63, 0x18, 0x63, 0xad, 0x0b, 0xac, 0x14, 0xa9, 0x25, 0x6a, 0x0c, 0x27, 0xec, 0x3a, 0xdb,
	0xd6, 0x82, 0x11, 0x17, 0x1f, 0x3f, 0xce, 0x11, 0x73, 0xc1, 0xe6, 0xf6, 0x94, 0x7e, 0xe4, 0xf3,
	0xac, 0xfa, 0xec, 0xae, 0x1e, 0x4f, 0xff, 0x3c, 0xda, 0x30, 0x75, 0x25, 0x8b, 0x5a, 0xa3, 0x74,
	0x37, 0x66, 0x2f, 0xa1, 0xcf, 0xae, 0x75, 0x42, 0xcb, 0xf0, 0x14, 0x80, 0x23, 0xae, 0x95, 0x96,
	0x45, 0x95, 0x47, 0xde, 0xd4, 0x3b, 0xf3, 0x53, 0x9f, 0x23, 0xbe, 0xb3, 0x41, 0xf8, 0x08, 0x06,
	0x34, 0x93, 0xeb, 0xa2, 0xd2, 0xd1, 0xd1, 0xd4, 0x3b, 0xeb, 0xa5, 0x7d, 0x9a, 0xc9, 0x37, 0x95,
	0x9e, 0xc5, 0x30, 0x7e, 0xcb, 0x94, 0xca, 0x72, 0x96, 0x58, 0x9f, 0x30, 0x80, 0x1e, 0x47, 0x6c,
	0x08, 0xe6, 0x71, 0x71, 0x72, 0x13, 0x8f, 0x2e, 0x98, 0x10, 0x38, 0xfd, 0x84, 0x52, 0x6c, 0x9e,
	0xde, 0xc6, 0x00, 0xf7, 0x38, 0xe2, 0xf9, 0x79, 0xf0, 0xc5, 0x9b, 0xbd, 0x86, 0xe1, 0xaa, 0x60,
	0x62, 0xd3, 0x94, 0xcf, 0xf6, 0xca, 0x97, 0xc1, 0xd7, 0xb8, 0x97, 0xd1, 0xab, 0x6f, 0xf1, 0xc0,
	0xa6, 0x81, 0x67, 0x81, 0xe1, 0x03, 0xe8, 0xd1, 0x4c, 0x5a, 0x11, 0x7f, 0x79, 0x14, 0x79, 0xa9,
	0x79, 0x5d, 0xbc, 0x82, 0xfb, 0xbc, 0x10, 0x6c, 0x8d, 0xb5, 0x6e, 0xda, 0x08, 0x4f, 0x88, 0x9b,
	0x00, 0xd9, 0x4d, 0x80, 0xac, 0x0a, 0xd1, 0x88, 0xaa, 0xe8, 0xfb, 0xc0, 0x3a, 0x8e, 0xb9, 0xcb,
	0x5c, 0xa7, 0x8b, 0xf7, 0x30, 0x6a, 0x31, 0x48, 0xcb, 0x0e, 0xc6, 0x0f, 0xc3, 0x18, 0x3e, 0x7f,
	0x48, 0x0e, 0xb7, 0x42, 0xdc, 0x54, 0x53, 0x68, 0xd8, 0x09, 0x2d, 0x17, 0x2b, 0x98, 0x6c, 0x55,
	0xbe, 0xaf, 0xf7, 0xe4, 0x2f, 0xf4, 0xc1, 0x28, 0x5b, 0xc3, 0xd1, 0x56, 0xe5, 0x77, 0x82, 0x1f,
	0x60, 0xb8, 0xe3, 0x18, 0xbf, 0x4e, 0x48, 0x87, 0xa2, 0xef, 0xe0, 0xc6, 0xf0, 0x02, 0x02, 0x6e,
	0x56, 0xb1, 0xef, 0x78, 0xfa, 0x8f, 0xf6, 0xdb, 0x6d, 0xa9, 0xe8, 0xa7, 0x33, 0x9c, 0xf0, 0x26,
	0x6c, 0x1c, 0x2f, 0x61, 0x7c, 0x47, 0x32, 0x96, 0x1d, 0x98, 0x5f, 0xff, 0x77, 0x1c, 0xee, 0xf0,
	0x09, 0x2d, 0x97, 0xcf, 0x6e, 0xe2, 0x76, 0x43, 0xa6, 0xf8, 0xb6, 0xfd, 0x3b, 0x2e, 0x7d, 0x42,
	0xe6, 0x2e, 0xa5, 0x7d, 0xfb, 0xa1, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xee, 0x25, 0xc1,
	0x8c, 0x3e, 0x03, 0x00, 0x00,
}