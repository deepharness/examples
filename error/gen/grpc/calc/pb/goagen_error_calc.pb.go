// Code generated with goa v3.16.0, DO NOT EDIT.
//
// calc protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/error/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: goagen_error_calc.proto

package calcpb

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

type DivideDivByZeroError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// division by zero leads to infinity.
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
}

func (x *DivideDivByZeroError) Reset() {
	*x = DivideDivByZeroError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_error_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideDivByZeroError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideDivByZeroError) ProtoMessage() {}

func (x *DivideDivByZeroError) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_error_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideDivByZeroError.ProtoReflect.Descriptor instead.
func (*DivideDivByZeroError) Descriptor() ([]byte, []int) {
	return file_goagen_error_calc_proto_rawDescGZIP(), []int{0}
}

func (x *DivideDivByZeroError) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

type DivideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dividend int32 `protobuf:"zigzag32,1,opt,name=dividend,proto3" json:"dividend,omitempty"`
	Divisor  int32 `protobuf:"zigzag32,2,opt,name=divisor,proto3" json:"divisor,omitempty"`
}

func (x *DivideRequest) Reset() {
	*x = DivideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_error_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideRequest) ProtoMessage() {}

func (x *DivideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_error_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideRequest.ProtoReflect.Descriptor instead.
func (*DivideRequest) Descriptor() ([]byte, []int) {
	return file_goagen_error_calc_proto_rawDescGZIP(), []int{1}
}

func (x *DivideRequest) GetDividend() int32 {
	if x != nil {
		return x.Dividend
	}
	return 0
}

func (x *DivideRequest) GetDivisor() int32 {
	if x != nil {
		return x.Divisor
	}
	return 0
}

type DivideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quotient int32 `protobuf:"zigzag32,1,opt,name=quotient,proto3" json:"quotient,omitempty"`
	Reminder int32 `protobuf:"zigzag32,2,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *DivideResponse) Reset() {
	*x = DivideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_error_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideResponse) ProtoMessage() {}

func (x *DivideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_error_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideResponse.ProtoReflect.Descriptor instead.
func (*DivideResponse) Descriptor() ([]byte, []int) {
	return file_goagen_error_calc_proto_rawDescGZIP(), []int{2}
}

func (x *DivideResponse) GetQuotient() int32 {
	if x != nil {
		return x.Quotient
	}
	return 0
}

func (x *DivideResponse) GetReminder() int32 {
	if x != nil {
		return x.Reminder
	}
	return 0
}

var File_goagen_error_calc_proto protoreflect.FileDescriptor

var file_goagen_error_calc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22,
	0x31, 0x0a, 0x14, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x44, 0x69, 0x76, 0x42, 0x79, 0x5a, 0x65,
	0x72, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x45, 0x0a, 0x0d, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x69, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x07, 0x64, 0x69, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x22, 0x48, 0x0a, 0x0e, 0x44, 0x69, 0x76,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x6f, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x71,
	0x75, 0x6f, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x32, 0x3b, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x33, 0x0a, 0x06, 0x44,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x44, 0x69, 0x76,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_goagen_error_calc_proto_rawDescOnce sync.Once
	file_goagen_error_calc_proto_rawDescData = file_goagen_error_calc_proto_rawDesc
)

func file_goagen_error_calc_proto_rawDescGZIP() []byte {
	file_goagen_error_calc_proto_rawDescOnce.Do(func() {
		file_goagen_error_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_goagen_error_calc_proto_rawDescData)
	})
	return file_goagen_error_calc_proto_rawDescData
}

var file_goagen_error_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goagen_error_calc_proto_goTypes = []interface{}{
	(*DivideDivByZeroError)(nil), // 0: calc.DivideDivByZeroError
	(*DivideRequest)(nil),        // 1: calc.DivideRequest
	(*DivideResponse)(nil),       // 2: calc.DivideResponse
}
var file_goagen_error_calc_proto_depIdxs = []int32{
	1, // 0: calc.Calc.Divide:input_type -> calc.DivideRequest
	2, // 1: calc.Calc.Divide:output_type -> calc.DivideResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_error_calc_proto_init() }
func file_goagen_error_calc_proto_init() {
	if File_goagen_error_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goagen_error_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideDivByZeroError); i {
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
		file_goagen_error_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideRequest); i {
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
		file_goagen_error_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideResponse); i {
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
			RawDescriptor: file_goagen_error_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_error_calc_proto_goTypes,
		DependencyIndexes: file_goagen_error_calc_proto_depIdxs,
		MessageInfos:      file_goagen_error_calc_proto_msgTypes,
	}.Build()
	File_goagen_error_calc_proto = out.File
	file_goagen_error_calc_proto_rawDesc = nil
	file_goagen_error_calc_proto_goTypes = nil
	file_goagen_error_calc_proto_depIdxs = nil
}
