// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secured_service.proto

package secured_servicepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SigninRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninRequest) Reset()         { *m = SigninRequest{} }
func (m *SigninRequest) String() string { return proto.CompactTextString(m) }
func (*SigninRequest) ProtoMessage()    {}
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{0}
}

func (m *SigninRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninRequest.Unmarshal(m, b)
}
func (m *SigninRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninRequest.Marshal(b, m, deterministic)
}
func (m *SigninRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninRequest.Merge(m, src)
}
func (m *SigninRequest) XXX_Size() int {
	return xxx_messageInfo_SigninRequest.Size(m)
}
func (m *SigninRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SigninRequest proto.InternalMessageInfo

type SigninResponse struct {
	// JWT token
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	// API Key
	ApiKey string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// OAuth2 token
	OauthToken           string   `protobuf:"bytes,3,opt,name=oauth_token,json=oauthToken,proto3" json:"oauth_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninResponse) Reset()         { *m = SigninResponse{} }
func (m *SigninResponse) String() string { return proto.CompactTextString(m) }
func (*SigninResponse) ProtoMessage()    {}
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{1}
}

func (m *SigninResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninResponse.Unmarshal(m, b)
}
func (m *SigninResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninResponse.Marshal(b, m, deterministic)
}
func (m *SigninResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninResponse.Merge(m, src)
}
func (m *SigninResponse) XXX_Size() int {
	return xxx_messageInfo_SigninResponse.Size(m)
}
func (m *SigninResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SigninResponse proto.InternalMessageInfo

func (m *SigninResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *SigninResponse) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *SigninResponse) GetOauthToken() string {
	if m != nil {
		return m.OauthToken
	}
	return ""
}

type SecureRequest struct {
	// Whether to force auth failure even with a valid JWT
	Fail                 bool     `protobuf:"varint,1,opt,name=fail,proto3" json:"fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecureRequest) Reset()         { *m = SecureRequest{} }
func (m *SecureRequest) String() string { return proto.CompactTextString(m) }
func (*SecureRequest) ProtoMessage()    {}
func (*SecureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{2}
}

func (m *SecureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecureRequest.Unmarshal(m, b)
}
func (m *SecureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecureRequest.Marshal(b, m, deterministic)
}
func (m *SecureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecureRequest.Merge(m, src)
}
func (m *SecureRequest) XXX_Size() int {
	return xxx_messageInfo_SecureRequest.Size(m)
}
func (m *SecureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecureRequest proto.InternalMessageInfo

func (m *SecureRequest) GetFail() bool {
	if m != nil {
		return m.Fail
	}
	return false
}

type SecureResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecureResponse) Reset()         { *m = SecureResponse{} }
func (m *SecureResponse) String() string { return proto.CompactTextString(m) }
func (*SecureResponse) ProtoMessage()    {}
func (*SecureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{3}
}

func (m *SecureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecureResponse.Unmarshal(m, b)
}
func (m *SecureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecureResponse.Marshal(b, m, deterministic)
}
func (m *SecureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecureResponse.Merge(m, src)
}
func (m *SecureResponse) XXX_Size() int {
	return xxx_messageInfo_SecureResponse.Size(m)
}
func (m *SecureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecureResponse proto.InternalMessageInfo

func (m *SecureResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type DoublySecureRequest struct {
	// API key
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoublySecureRequest) Reset()         { *m = DoublySecureRequest{} }
func (m *DoublySecureRequest) String() string { return proto.CompactTextString(m) }
func (*DoublySecureRequest) ProtoMessage()    {}
func (*DoublySecureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{4}
}

func (m *DoublySecureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoublySecureRequest.Unmarshal(m, b)
}
func (m *DoublySecureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoublySecureRequest.Marshal(b, m, deterministic)
}
func (m *DoublySecureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoublySecureRequest.Merge(m, src)
}
func (m *DoublySecureRequest) XXX_Size() int {
	return xxx_messageInfo_DoublySecureRequest.Size(m)
}
func (m *DoublySecureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoublySecureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoublySecureRequest proto.InternalMessageInfo

func (m *DoublySecureRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DoublySecureResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoublySecureResponse) Reset()         { *m = DoublySecureResponse{} }
func (m *DoublySecureResponse) String() string { return proto.CompactTextString(m) }
func (*DoublySecureResponse) ProtoMessage()    {}
func (*DoublySecureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{5}
}

func (m *DoublySecureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoublySecureResponse.Unmarshal(m, b)
}
func (m *DoublySecureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoublySecureResponse.Marshal(b, m, deterministic)
}
func (m *DoublySecureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoublySecureResponse.Merge(m, src)
}
func (m *DoublySecureResponse) XXX_Size() int {
	return xxx_messageInfo_DoublySecureResponse.Size(m)
}
func (m *DoublySecureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoublySecureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoublySecureResponse proto.InternalMessageInfo

func (m *DoublySecureResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type AlsoDoublySecureRequest struct {
	// Username used to perform signin
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Password used to perform signin
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// API key
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlsoDoublySecureRequest) Reset()         { *m = AlsoDoublySecureRequest{} }
func (m *AlsoDoublySecureRequest) String() string { return proto.CompactTextString(m) }
func (*AlsoDoublySecureRequest) ProtoMessage()    {}
func (*AlsoDoublySecureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{6}
}

func (m *AlsoDoublySecureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlsoDoublySecureRequest.Unmarshal(m, b)
}
func (m *AlsoDoublySecureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlsoDoublySecureRequest.Marshal(b, m, deterministic)
}
func (m *AlsoDoublySecureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlsoDoublySecureRequest.Merge(m, src)
}
func (m *AlsoDoublySecureRequest) XXX_Size() int {
	return xxx_messageInfo_AlsoDoublySecureRequest.Size(m)
}
func (m *AlsoDoublySecureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlsoDoublySecureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlsoDoublySecureRequest proto.InternalMessageInfo

func (m *AlsoDoublySecureRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AlsoDoublySecureRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AlsoDoublySecureRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type AlsoDoublySecureResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlsoDoublySecureResponse) Reset()         { *m = AlsoDoublySecureResponse{} }
func (m *AlsoDoublySecureResponse) String() string { return proto.CompactTextString(m) }
func (*AlsoDoublySecureResponse) ProtoMessage()    {}
func (*AlsoDoublySecureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76cb74962e1a109, []int{7}
}

func (m *AlsoDoublySecureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlsoDoublySecureResponse.Unmarshal(m, b)
}
func (m *AlsoDoublySecureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlsoDoublySecureResponse.Marshal(b, m, deterministic)
}
func (m *AlsoDoublySecureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlsoDoublySecureResponse.Merge(m, src)
}
func (m *AlsoDoublySecureResponse) XXX_Size() int {
	return xxx_messageInfo_AlsoDoublySecureResponse.Size(m)
}
func (m *AlsoDoublySecureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlsoDoublySecureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlsoDoublySecureResponse proto.InternalMessageInfo

func (m *AlsoDoublySecureResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func init() {
	proto.RegisterType((*SigninRequest)(nil), "multi_auth.SigninRequest")
	proto.RegisterType((*SigninResponse)(nil), "multi_auth.SigninResponse")
	proto.RegisterType((*SecureRequest)(nil), "multi_auth.SecureRequest")
	proto.RegisterType((*SecureResponse)(nil), "multi_auth.SecureResponse")
	proto.RegisterType((*DoublySecureRequest)(nil), "multi_auth.DoublySecureRequest")
	proto.RegisterType((*DoublySecureResponse)(nil), "multi_auth.DoublySecureResponse")
	proto.RegisterType((*AlsoDoublySecureRequest)(nil), "multi_auth.AlsoDoublySecureRequest")
	proto.RegisterType((*AlsoDoublySecureResponse)(nil), "multi_auth.AlsoDoublySecureResponse")
}

func init() { proto.RegisterFile("secured_service.proto", fileDescriptor_b76cb74962e1a109) }

var fileDescriptor_b76cb74962e1a109 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0xc6, 0x03, 0x53, 0xc4, 0xa3, 0x20, 0x16, 0x0c, 0x73, 0x37, 0x90, 0x61, 0xd4, 0x0b, 0x43,
	0x8c, 0x3e, 0x80, 0xd1, 0x78, 0xe7, 0x95, 0xe0, 0x95, 0xd1, 0x2c, 0x83, 0x1d, 0xb4, 0x32, 0xd6,
	0xb9, 0x6e, 0x92, 0xbd, 0x9f, 0x0f, 0x66, 0xda, 0x75, 0xb0, 0x4d, 0xd8, 0x5d, 0xcf, 0x9f, 0xfe,
	0xce, 0xd7, 0xf3, 0xa5, 0x70, 0xc2, 0x71, 0x1a, 0x05, 0xe8, 0x58, 0x1c, 0x83, 0x1f, 0x3a, 0xc5,
	0xa1, 0x1f, 0xb0, 0x90, 0x11, 0x58, 0x44, 0x6e, 0x48, 0x2d, 0x3b, 0x0a, 0x3f, 0xcd, 0x23, 0x68,
	0x8c, 0xe9, 0x87, 0x47, 0xbd, 0x11, 0x7e, 0x47, 0xc8, 0x43, 0xf3, 0x0d, 0x9a, 0x69, 0x82, 0xfb,
	0xcc, 0xe3, 0x48, 0x5a, 0xa0, 0x7d, 0x2d, 0x43, 0xbd, 0xd2, 0xaf, 0x5c, 0xee, 0x8f, 0xc4, 0x91,
	0x74, 0x61, 0xcf, 0xf6, 0xa9, 0x35, 0xc7, 0x58, 0xaf, 0xca, 0x6c, 0xcd, 0xf6, 0xe9, 0x13, 0xc6,
	0xa4, 0x07, 0x07, 0x4c, 0x60, 0xad, 0x90, 0xcd, 0xd1, 0xd3, 0x35, 0x59, 0x04, 0x99, 0x7a, 0x11,
	0x19, 0x73, 0x00, 0x8d, 0xb1, 0xd4, 0xa4, 0xc6, 0x11, 0x02, 0x3b, 0x33, 0x9b, 0xba, 0x92, 0x5e,
	0x1f, 0xc9, 0xb3, 0x79, 0x0e, 0xcd, 0xb4, 0x49, 0x49, 0xe8, 0xc0, 0xee, 0x8c, 0xa2, 0xeb, 0x28,
	0x11, 0x49, 0x60, 0x5e, 0x40, 0xfb, 0x91, 0x45, 0x13, 0x37, 0xce, 0x23, 0x5b, 0xa0, 0x09, 0x65,
	0x4a, 0xef, 0x1c, 0x63, 0xf3, 0x0a, 0x3a, 0xf9, 0xc6, 0x52, 0xec, 0x14, 0xba, 0xf7, 0x2e, 0x67,
	0x9b, 0xd0, 0x06, 0xd4, 0x23, 0x8e, 0x81, 0x67, 0x2f, 0x50, 0xdd, 0x59, 0xc5, 0xa2, 0xe6, 0xdb,
	0x9c, 0x2f, 0x59, 0xe0, 0xa8, 0xad, 0xac, 0xe2, 0x54, 0x92, 0xb6, 0x96, 0x74, 0x0d, 0xfa, 0xff,
	0x21, 0x65, 0xb2, 0x6e, 0x7e, 0xab, 0xe9, 0x5a, 0x9c, 0x71, 0x62, 0x27, 0xb9, 0x83, 0x5a, 0xe2,
	0x15, 0x39, 0x1d, 0xae, 0x3d, 0x1d, 0xe6, 0x0c, 0x35, 0x8c, 0x4d, 0x25, 0x35, 0x49, 0x00, 0x24,
	0xb2, 0x00, 0xc8, 0x3e, 0xba, 0x00, 0xc8, 0x4b, 0x7d, 0x86, 0xc3, 0xec, 0x13, 0x48, 0x2f, 0xdb,
	0xbb, 0x61, 0x83, 0x46, 0x7f, 0x7b, 0x83, 0x42, 0xbe, 0x43, 0xab, 0xb8, 0x19, 0x32, 0xc8, 0xde,
	0xda, 0x62, 0x8e, 0x71, 0x56, 0xde, 0x94, 0xe0, 0x1f, 0xda, 0xaf, 0xc7, 0x85, 0x5f, 0xe1, 0x4f,
	0x26, 0x35, 0xf9, 0x31, 0x6e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xab, 0xc7, 0xe7, 0xc1, 0x31,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecuredServiceClient is the client API for SecuredService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecuredServiceClient interface {
	// Creates a valid JWT
	Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninResponse, error)
	// This action is secured with the jwt scheme
	Secure(ctx context.Context, in *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// query string.
	DoublySecure(ctx context.Context, in *DoublySecureRequest, opts ...grpc.CallOption) (*DoublySecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// header.
	AlsoDoublySecure(ctx context.Context, in *AlsoDoublySecureRequest, opts ...grpc.CallOption) (*AlsoDoublySecureResponse, error)
}

type securedServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecuredServiceClient(cc *grpc.ClientConn) SecuredServiceClient {
	return &securedServiceClient{cc}
}

func (c *securedServiceClient) Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninResponse, error) {
	out := new(SigninResponse)
	err := c.cc.Invoke(ctx, "/multi_auth.SecuredService/Signin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securedServiceClient) Secure(ctx context.Context, in *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error) {
	out := new(SecureResponse)
	err := c.cc.Invoke(ctx, "/multi_auth.SecuredService/Secure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securedServiceClient) DoublySecure(ctx context.Context, in *DoublySecureRequest, opts ...grpc.CallOption) (*DoublySecureResponse, error) {
	out := new(DoublySecureResponse)
	err := c.cc.Invoke(ctx, "/multi_auth.SecuredService/DoublySecure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securedServiceClient) AlsoDoublySecure(ctx context.Context, in *AlsoDoublySecureRequest, opts ...grpc.CallOption) (*AlsoDoublySecureResponse, error) {
	out := new(AlsoDoublySecureResponse)
	err := c.cc.Invoke(ctx, "/multi_auth.SecuredService/AlsoDoublySecure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecuredServiceServer is the server API for SecuredService service.
type SecuredServiceServer interface {
	// Creates a valid JWT
	Signin(context.Context, *SigninRequest) (*SigninResponse, error)
	// This action is secured with the jwt scheme
	Secure(context.Context, *SecureRequest) (*SecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// query string.
	DoublySecure(context.Context, *DoublySecureRequest) (*DoublySecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// header.
	AlsoDoublySecure(context.Context, *AlsoDoublySecureRequest) (*AlsoDoublySecureResponse, error)
}

func RegisterSecuredServiceServer(s *grpc.Server, srv SecuredServiceServer) {
	s.RegisterService(&_SecuredService_serviceDesc, srv)
}

func _SecuredService_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multi_auth.SecuredService/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).Signin(ctx, req.(*SigninRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuredService_Secure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).Secure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multi_auth.SecuredService/Secure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).Secure(ctx, req.(*SecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuredService_DoublySecure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoublySecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).DoublySecure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multi_auth.SecuredService/DoublySecure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).DoublySecure(ctx, req.(*DoublySecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuredService_AlsoDoublySecure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlsoDoublySecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).AlsoDoublySecure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multi_auth.SecuredService/AlsoDoublySecure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).AlsoDoublySecure(ctx, req.(*AlsoDoublySecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecuredService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multi_auth.SecuredService",
	HandlerType: (*SecuredServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signin",
			Handler:    _SecuredService_Signin_Handler,
		},
		{
			MethodName: "Secure",
			Handler:    _SecuredService_Secure_Handler,
		},
		{
			MethodName: "DoublySecure",
			Handler:    _SecuredService_DoublySecure_Handler,
		},
		{
			MethodName: "AlsoDoublySecure",
			Handler:    _SecuredService_AlsoDoublySecure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secured_service.proto",
}
