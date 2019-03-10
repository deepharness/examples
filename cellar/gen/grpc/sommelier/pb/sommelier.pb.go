// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sommelier.proto

package sommelierpb

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

type PickRequest struct {
	// Name of bottle to pick
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `protobuf:"bytes,2,rep,name=varietal,proto3" json:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery               string   `protobuf:"bytes,3,opt,name=winery,proto3" json:"winery,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PickRequest) Reset()         { *m = PickRequest{} }
func (m *PickRequest) String() string { return proto.CompactTextString(m) }
func (*PickRequest) ProtoMessage()    {}
func (*PickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ab3ffe67c1f95d, []int{0}
}

func (m *PickRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PickRequest.Unmarshal(m, b)
}
func (m *PickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PickRequest.Marshal(b, m, deterministic)
}
func (m *PickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PickRequest.Merge(m, src)
}
func (m *PickRequest) XXX_Size() int {
	return xxx_messageInfo_PickRequest.Size(m)
}
func (m *PickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PickRequest proto.InternalMessageInfo

func (m *PickRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PickRequest) GetVarietal() []string {
	if m != nil {
		return m.Varietal
	}
	return nil
}

func (m *PickRequest) GetWinery() string {
	if m != nil {
		return m.Winery
	}
	return ""
}

type StoredBottleCollection struct {
	Field                []*StoredBottle `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StoredBottleCollection) Reset()         { *m = StoredBottleCollection{} }
func (m *StoredBottleCollection) String() string { return proto.CompactTextString(m) }
func (*StoredBottleCollection) ProtoMessage()    {}
func (*StoredBottleCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ab3ffe67c1f95d, []int{1}
}

func (m *StoredBottleCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoredBottleCollection.Unmarshal(m, b)
}
func (m *StoredBottleCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoredBottleCollection.Marshal(b, m, deterministic)
}
func (m *StoredBottleCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredBottleCollection.Merge(m, src)
}
func (m *StoredBottleCollection) XXX_Size() int {
	return xxx_messageInfo_StoredBottleCollection.Size(m)
}
func (m *StoredBottleCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredBottleCollection.DiscardUnknown(m)
}

var xxx_messageInfo_StoredBottleCollection proto.InternalMessageInfo

func (m *StoredBottleCollection) GetField() []*StoredBottle {
	if m != nil {
		return m.Field
	}
	return nil
}

// A StoredBottle describes a bottle retrieved by the storage service.
type StoredBottle struct {
	// ID is the unique id of the bottle.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name of bottle
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Winery that produces wine
	Winery *Winery `protobuf:"bytes,2,opt,name=winery,proto3" json:"winery,omitempty"`
	// Vintage of bottle
	Vintage uint32 `protobuf:"varint,3,opt,name=vintage,proto3" json:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*Component `protobuf:"bytes,4,rep,name=composition,proto3" json:"composition,omitempty"`
	// Description of bottle
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating               uint32   `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoredBottle) Reset()         { *m = StoredBottle{} }
func (m *StoredBottle) String() string { return proto.CompactTextString(m) }
func (*StoredBottle) ProtoMessage()    {}
func (*StoredBottle) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ab3ffe67c1f95d, []int{2}
}

func (m *StoredBottle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoredBottle.Unmarshal(m, b)
}
func (m *StoredBottle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoredBottle.Marshal(b, m, deterministic)
}
func (m *StoredBottle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredBottle.Merge(m, src)
}
func (m *StoredBottle) XXX_Size() int {
	return xxx_messageInfo_StoredBottle.Size(m)
}
func (m *StoredBottle) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredBottle.DiscardUnknown(m)
}

var xxx_messageInfo_StoredBottle proto.InternalMessageInfo

func (m *StoredBottle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoredBottle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoredBottle) GetWinery() *Winery {
	if m != nil {
		return m.Winery
	}
	return nil
}

func (m *StoredBottle) GetVintage() uint32 {
	if m != nil {
		return m.Vintage
	}
	return 0
}

func (m *StoredBottle) GetComposition() []*Component {
	if m != nil {
		return m.Composition
	}
	return nil
}

func (m *StoredBottle) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StoredBottle) GetRating() uint32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type Winery struct {
	// Name of winery
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Region of winery
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Country of winery
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	// Winery website URL
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Winery) Reset()         { *m = Winery{} }
func (m *Winery) String() string { return proto.CompactTextString(m) }
func (*Winery) ProtoMessage()    {}
func (*Winery) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ab3ffe67c1f95d, []int{3}
}

func (m *Winery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Winery.Unmarshal(m, b)
}
func (m *Winery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Winery.Marshal(b, m, deterministic)
}
func (m *Winery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Winery.Merge(m, src)
}
func (m *Winery) XXX_Size() int {
	return xxx_messageInfo_Winery.Size(m)
}
func (m *Winery) XXX_DiscardUnknown() {
	xxx_messageInfo_Winery.DiscardUnknown(m)
}

var xxx_messageInfo_Winery proto.InternalMessageInfo

func (m *Winery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Winery) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Winery) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Winery) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Component struct {
	// Grape varietal
	Varietal string `protobuf:"bytes,1,opt,name=varietal,proto3" json:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage           uint32   `protobuf:"varint,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Component) Reset()         { *m = Component{} }
func (m *Component) String() string { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()    {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ab3ffe67c1f95d, []int{4}
}

func (m *Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Component.Unmarshal(m, b)
}
func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Component.Marshal(b, m, deterministic)
}
func (m *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(m, src)
}
func (m *Component) XXX_Size() int {
	return xxx_messageInfo_Component.Size(m)
}
func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

func (m *Component) GetVarietal() string {
	if m != nil {
		return m.Varietal
	}
	return ""
}

func (m *Component) GetPercentage() uint32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func init() {
	proto.RegisterType((*PickRequest)(nil), "cellar.PickRequest")
	proto.RegisterType((*StoredBottleCollection)(nil), "cellar.StoredBottleCollection")
	proto.RegisterType((*StoredBottle)(nil), "cellar.StoredBottle")
	proto.RegisterType((*Winery)(nil), "cellar.Winery")
	proto.RegisterType((*Component)(nil), "cellar.Component")
}

func init() { proto.RegisterFile("sommelier.proto", fileDescriptor_d0ab3ffe67c1f95d) }

var fileDescriptor_d0ab3ffe67c1f95d = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xf1, 0x9f, 0x78, 0xe3, 0xf1, 0x26, 0xbb, 0xab, 0x5d, 0x82, 0xc8, 0x21, 0x18, 0x1f,
	0x16, 0xb3, 0x87, 0x1c, 0x92, 0xe3, 0xde, 0x92, 0x85, 0xed, 0xb1, 0x38, 0x94, 0x42, 0x4f, 0x75,
	0xec, 0x69, 0x10, 0x95, 0x25, 0x57, 0x56, 0x52, 0xfa, 0x79, 0xfb, 0x45, 0x8a, 0xe4, 0xd8, 0x71,
	0x21, 0x37, 0xbd, 0xb1, 0x46, 0xf3, 0xf3, 0x9b, 0x07, 0xdf, 0x1a, 0x59, 0x55, 0xc8, 0x19, 0xaa,
	0x65, 0xad, 0xa4, 0x96, 0x24, 0x28, 0x90, 0xf3, 0x5c, 0x25, 0x77, 0x10, 0xdd, 0xb2, 0xe2, 0x39,
	0xc3, 0x97, 0x23, 0x36, 0x9a, 0x10, 0xf0, 0x45, 0x5e, 0x21, 0x75, 0x62, 0x27, 0x0d, 0x33, 0x7b,
	0x26, 0x73, 0x18, 0x9f, 0x72, 0xc5, 0x50, 0xe7, 0x9c, 0xba, 0xb1, 0x97, 0x86, 0x59, 0xaf, 0xc9,
	0x0c, 0x82, 0x57, 0x26, 0x50, 0xbd, 0x51, 0xcf, 0x76, 0x9c, 0x55, 0xf2, 0x0f, 0x66, 0x3b, 0x2d,
	0x15, 0x96, 0x1b, 0xa9, 0x35, 0xc7, 0xad, 0xe4, 0x1c, 0x0b, 0xcd, 0xa4, 0x20, 0x7f, 0x60, 0xf4,
	0xc4, 0x90, 0x97, 0xd4, 0x89, 0xbd, 0x34, 0x5a, 0xfd, 0x5a, 0xb6, 0x20, 0xcb, 0xe1, 0xf5, 0xac,
	0xbd, 0x92, 0xbc, 0x3b, 0xf0, 0x75, 0x58, 0x27, 0x53, 0x70, 0x59, 0x49, 0xc7, 0x76, 0x94, 0xcb,
	0xca, 0xab, 0xb8, 0xbf, 0x7b, 0x24, 0x37, 0x76, 0xd2, 0x68, 0x35, 0xed, 0x26, 0xdc, 0xdb, 0x6a,
	0x87, 0x48, 0x28, 0x7c, 0x39, 0x31, 0xa1, 0xf3, 0x03, 0x5a, 0xf6, 0x49, 0xd6, 0x49, 0xb2, 0x86,
	0xa8, 0x90, 0x55, 0x2d, 0x1b, 0x66, 0x88, 0xa9, 0x6f, 0x41, 0x7f, 0x74, 0xcf, 0x6c, 0xcd, 0x27,
	0x81, 0x42, 0x67, 0xc3, 0x5b, 0x24, 0x86, 0xa8, 0xc4, 0xa6, 0x50, 0xac, 0xb6, 0x4d, 0x23, 0x4b,
	0x34, 0x2c, 0x19, 0xaf, 0x54, 0xae, 0x99, 0x38, 0xd0, 0xc0, 0xce, 0x3b, 0xab, 0xe4, 0x11, 0x82,
	0x16, 0xed, 0xea, 0xef, 0x98, 0x2e, 0x3c, 0x98, 0x27, 0xdd, 0xd6, 0xe1, 0x56, 0x19, 0xfc, 0x42,
	0x1e, 0x85, 0xee, 0xad, 0xef, 0x24, 0xf9, 0x0e, 0xde, 0x51, 0x71, 0xea, 0xdb, 0xaa, 0x39, 0x26,
	0xff, 0x21, 0xec, 0xa9, 0x3f, 0xad, 0xb3, 0x1d, 0x74, 0x59, 0xe7, 0x02, 0xa0, 0x46, 0x55, 0x60,
	0x6b, 0x8b, 0x6b, 0x31, 0x07, 0x95, 0xd5, 0x0d, 0x84, 0xbb, 0x2e, 0x48, 0xe4, 0x2f, 0xf8, 0x26,
	0x3a, 0xe4, 0x67, 0xe7, 0xcc, 0x20, 0x48, 0xf3, 0xc5, 0xb5, 0xbd, 0x5e, 0x62, 0xb0, 0x99, 0x3c,
	0x44, 0x7d, 0x24, 0xeb, 0xfd, 0x3e, 0xb0, 0xa9, 0x5c, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0xad,
	0x82, 0xeb, 0x50, 0xa8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SommelierClient is the client API for Sommelier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SommelierClient interface {
	// Pick implements pick.
	Pick(ctx context.Context, in *PickRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error)
}

type sommelierClient struct {
	cc *grpc.ClientConn
}

func NewSommelierClient(cc *grpc.ClientConn) SommelierClient {
	return &sommelierClient{cc}
}

func (c *sommelierClient) Pick(ctx context.Context, in *PickRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error) {
	out := new(StoredBottleCollection)
	err := c.cc.Invoke(ctx, "/cellar.Sommelier/Pick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SommelierServer is the server API for Sommelier service.
type SommelierServer interface {
	// Pick implements pick.
	Pick(context.Context, *PickRequest) (*StoredBottleCollection, error)
}

func RegisterSommelierServer(s *grpc.Server, srv SommelierServer) {
	s.RegisterService(&_Sommelier_serviceDesc, srv)
}

func _Sommelier_Pick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SommelierServer).Pick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cellar.Sommelier/Pick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SommelierServer).Pick(ctx, req.(*PickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sommelier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cellar.Sommelier",
	HandlerType: (*SommelierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pick",
			Handler:    _Sommelier_Pick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sommelier.proto",
}
