// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_service.proto

package pb

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "sufiyan.pcbook.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "sufiyan.pcbook.LoginResponse")
}

func init() { proto.RegisterFile("auth_service.proto", fileDescriptor_0f39bb026ca10b68) }

var fileDescriptor_0f39bb026ca10b68 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x4e, 0x86, 0x30,
	0x14, 0x85, 0xfd, 0x89, 0x1a, 0x2d, 0xe8, 0xd0, 0x38, 0x10, 0xa2, 0x89, 0x32, 0x39, 0x75, 0xc0,
	0x27, 0xd0, 0x81, 0xc9, 0x09, 0x75, 0x71, 0x21, 0xa5, 0x5e, 0x81, 0x90, 0xf6, 0xd6, 0xde, 0x56,
	0xc3, 0xdb, 0x1b, 0x28, 0x1a, 0x1d, 0xfe, 0xf1, 0x9c, 0xd3, 0x7c, 0xfd, 0x2e, 0xe3, 0x32, 0xf8,
	0xa1, 0x25, 0x70, 0x9f, 0xa3, 0x02, 0x61, 0x1d, 0x7a, 0xe4, 0xe7, 0x14, 0xde, 0xc7, 0x59, 0x1a,
	0x61, 0x55, 0x87, 0x38, 0x15, 0x17, 0x1a, 0x34, 0xba, 0xb9, 0xd5, 0x40, 0x24, 0xfb, 0xed, 0x55,
	0x59, 0xb3, 0xec, 0x11, 0xfb, 0xd1, 0x34, 0xf0, 0x11, 0x80, 0x3c, 0x2f, 0xd8, 0x49, 0x20, 0x70,
	0x46, 0x6a, 0xc8, 0x77, 0xd7, 0xbb, 0xdb, 0xd3, 0xe6, 0x37, 0x2f, 0x9b, 0x95, 0x44, 0x5f, 0xe8,
	0xde, 0xf2, 0x24, 0x6e, 0x3f, 0xb9, 0xac, 0xd8, 0xd9, 0xc6, 0x21, 0x8b, 0x86, 0x80, 0xdf, 0xb0,
	0x4c, 0x2a, 0x05, 0x44, 0xad, 0xc7, 0x09, 0xcc, 0x06, 0x4b, 0x63, 0xf7, 0xbc, 0x54, 0xd5, 0x0b,
	0x4b, 0xef, 0x83, 0x1f, 0x9e, 0xa2, 0x36, 0xaf, 0xd9, 0xd1, 0x8a, 0xe0, 0x97, 0xe2, 0xbf, 0xba,
	0xf8, 0x6b, 0x58, 0x5c, 0xed, 0x59, 0xe3, 0xbf, 0xe5, 0xc1, 0xc3, 0xe1, 0x6b, 0x62, 0xbb, 0xee,
	0x78, 0xbd, 0xef, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x3a, 0x3a, 0xda, 0x1b, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/sufiyan.pcbook.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sufiyan.pcbook.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sufiyan.pcbook.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
