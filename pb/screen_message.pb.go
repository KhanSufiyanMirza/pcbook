// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

package sufiyan_pb

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

type Screen_Panel int32

const (
	Screen_UNKOWN Screen_Panel = 0
	Screen_IPS    Screen_Panel = 1
	Screen_OLED   Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKOWN": 0,
	"IPS":    1,
	"OLED":   2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=sufiyan.pcbook.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("sufiyan.pcbook.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "sufiyan.pcbook.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "sufiyan.pcbook.Screen.Resolution")
}

func init() { proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7) }

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x4d, 0x9b, 0xbc, 0x71, 0xb4, 0x25, 0x0c, 0x45, 0x82, 0x8a, 0xc4, 0x1e, 0x24,
	0xa7, 0x3d, 0xc4, 0x9b, 0x37, 0x45, 0x0f, 0x45, 0x69, 0xcb, 0x16, 0x11, 0xbc, 0x94, 0x64, 0x5d,
	0xbb, 0x8b, 0xe9, 0x6e, 0xc8, 0x6e, 0x10, 0xfd, 0x34, 0x7e, 0x54, 0x69, 0x22, 0xb5, 0x1e, 0x3c,
	0xce, 0xc3, 0xef, 0xf9, 0xc3, 0xc0, 0xc8, 0xf2, 0x5a, 0x08, 0xbd, 0x5c, 0x0b, 0x6b, 0xf3, 0x95,
	0xa0, 0x55, 0x6d, 0x9c, 0xc1, 0xa1, 0x6d, 0x5e, 0xd4, 0x7b, 0xae, 0x69, 0xc5, 0x0b, 0x63, 0x5e,
	0xc7, 0x9f, 0x1e, 0x04, 0x8b, 0x16, 0xc4, 0x63, 0xd8, 0xb3, 0xea, 0x43, 0x2c, 0x95, 0xe6, 0x32,
	0x26, 0x09, 0x49, 0x3d, 0x16, 0x6e, 0x84, 0x89, 0xe6, 0x12, 0xaf, 0x00, 0x6a, 0x61, 0x4d, 0xd9,
	0x38, 0x65, 0x74, 0xec, 0x25, 0x24, 0xdd, 0xcf, 0xce, 0xe8, 0xef, 0x30, 0xda, 0x05, 0x51, 0xb6,
	0x05, 0xd9, 0x8e, 0x09, 0x33, 0xf0, 0xab, 0x5c, 0x8b, 0x32, 0xee, 0x25, 0x24, 0x1d, 0x66, 0x27,
	0x7f, 0xb8, 0xe7, 0x1b, 0x86, 0x75, 0x28, 0x9e, 0x02, 0xac, 0x9b, 0xd2, 0x29, 0x67, 0x1a, 0x2e,
	0xe3, 0x7e, 0x42, 0xd2, 0x90, 0xed, 0x28, 0x47, 0x97, 0x00, 0x3f, 0x6d, 0x38, 0x02, 0xff, 0x4d,
	0x3d, 0xbb, 0x6e, 0xfd, 0x80, 0x75, 0x07, 0x1e, 0x42, 0x20, 0x85, 0x5a, 0x49, 0xd7, 0xce, 0x1e,
	0xb0, 0xef, 0x6b, 0x7c, 0x0e, 0x7e, 0xdb, 0x85, 0x00, 0xc1, 0xc3, 0xf4, 0x6e, 0xf6, 0x38, 0x8d,
	0xfe, 0xe1, 0x7f, 0xe8, 0x4d, 0xe6, 0x8b, 0x88, 0x60, 0x08, 0xfd, 0xd9, 0xfd, 0xed, 0x4d, 0xe4,
	0x5d, 0x1f, 0x3c, 0xc1, 0x76, 0x69, 0x51, 0x04, 0xed, 0x1f, 0x2f, 0xbe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x30, 0x48, 0x24, 0x8c, 0x5f, 0x01, 0x00, 0x00,
}
