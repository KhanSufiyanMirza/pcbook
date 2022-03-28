// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter_message.proto

package pb

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

type Filter struct {
	MaxPriceUsd          float64  `protobuf:"fixed64,1,opt,name=max_price_usd,json=maxPriceUsd,proto3" json:"max_price_usd,omitempty"`
	MinCpuCores          uint32   `protobuf:"varint,2,opt,name=min_cpu_cores,json=minCpuCores,proto3" json:"min_cpu_cores,omitempty"`
	MinCpuGhz            float64  `protobuf:"fixed64,3,opt,name=min_cpu_ghz,json=minCpuGhz,proto3" json:"min_cpu_ghz,omitempty"`
	MinRam               *Memory  `protobuf:"bytes,4,opt,name=min_ram,json=minRam,proto3" json:"min_ram,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_02dd12c5821a9fa1, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetMaxPriceUsd() float64 {
	if m != nil {
		return m.MaxPriceUsd
	}
	return 0
}

func (m *Filter) GetMinCpuCores() uint32 {
	if m != nil {
		return m.MinCpuCores
	}
	return 0
}

func (m *Filter) GetMinCpuGhz() float64 {
	if m != nil {
		return m.MinCpuGhz
	}
	return 0
}

func (m *Filter) GetMinRam() *Memory {
	if m != nil {
		return m.MinRam
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "sufiyan.pcbook.Filter")
}

func init() { proto.RegisterFile("filter_message.proto", fileDescriptor_02dd12c5821a9fa1) }

var fileDescriptor_02dd12c5821a9fa1 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x2b, 0x2e, 0x4d, 0xcb, 0xac, 0x4c, 0xcc, 0xd3, 0x2b, 0x48, 0x4e, 0xca, 0xcf, 0xcf,
	0x96, 0x12, 0xc9, 0x4d, 0xcd, 0xcd, 0x2f, 0xaa, 0x44, 0x55, 0xa5, 0xb4, 0x90, 0x91, 0x8b, 0xcd,
	0x0d, 0xac, 0x5d, 0x48, 0x89, 0x8b, 0x37, 0x37, 0xb1, 0x22, 0xbe, 0xa0, 0x28, 0x33, 0x39, 0x35,
	0xbe, 0xb4, 0x38, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x31, 0x88, 0x3b, 0x37, 0xb1, 0x22, 0x00,
	0x24, 0x16, 0x5a, 0x9c, 0x02, 0x56, 0x93, 0x99, 0x17, 0x9f, 0x5c, 0x50, 0x1a, 0x9f, 0x9c, 0x5f,
	0x94, 0x5a, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4, 0x9d, 0x9b, 0x99, 0xe7, 0x5c, 0x50,
	0xea, 0x0c, 0x12, 0x12, 0x92, 0xe3, 0xe2, 0x86, 0xa9, 0x49, 0xcf, 0xa8, 0x92, 0x60, 0x06, 0x9b,
	0xc2, 0x09, 0x51, 0xe1, 0x9e, 0x51, 0x25, 0xa4, 0xcf, 0xc5, 0x0e, 0x92, 0x2f, 0x4a, 0xcc, 0x95,
	0x60, 0x51, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd3, 0x43, 0x75, 0xaa, 0x9e, 0x2f, 0xd8, 0xa5, 0x41,
	0x6c, 0xb9, 0x99, 0x79, 0x41, 0x89, 0xb9, 0x4e, 0x2c, 0x51, 0x4c, 0x05, 0x49, 0x49, 0x6c, 0x60,
	0x07, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x55, 0xaf, 0x21, 0xed, 0xee, 0x00, 0x00, 0x00,
}