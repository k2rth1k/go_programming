// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/data-structures/data-structures.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RearrangeArrayAlternately_Test_Case struct {
	Length               int64    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Array                []int64  `protobuf:"varint,2,rep,packed,name=array,proto3" json:"array,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RearrangeArrayAlternately_Test_Case) Reset()         { *m = RearrangeArrayAlternately_Test_Case{} }
func (m *RearrangeArrayAlternately_Test_Case) String() string { return proto.CompactTextString(m) }
func (*RearrangeArrayAlternately_Test_Case) ProtoMessage()    {}
func (*RearrangeArrayAlternately_Test_Case) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{0}
}

func (m *RearrangeArrayAlternately_Test_Case) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RearrangeArrayAlternately_Test_Case.Unmarshal(m, b)
}
func (m *RearrangeArrayAlternately_Test_Case) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RearrangeArrayAlternately_Test_Case.Marshal(b, m, deterministic)
}
func (m *RearrangeArrayAlternately_Test_Case) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RearrangeArrayAlternately_Test_Case.Merge(m, src)
}
func (m *RearrangeArrayAlternately_Test_Case) XXX_Size() int {
	return xxx_messageInfo_RearrangeArrayAlternately_Test_Case.Size(m)
}
func (m *RearrangeArrayAlternately_Test_Case) XXX_DiscardUnknown() {
	xxx_messageInfo_RearrangeArrayAlternately_Test_Case.DiscardUnknown(m)
}

var xxx_messageInfo_RearrangeArrayAlternately_Test_Case proto.InternalMessageInfo

func (m *RearrangeArrayAlternately_Test_Case) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *RearrangeArrayAlternately_Test_Case) GetArray() []int64 {
	if m != nil {
		return m.Array
	}
	return nil
}

type RearrangeArrayAlternatelyRequest struct {
	Array                []int64  `protobuf:"varint,1,rep,packed,name=Array,proto3" json:"Array,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RearrangeArrayAlternatelyRequest) Reset()         { *m = RearrangeArrayAlternatelyRequest{} }
func (m *RearrangeArrayAlternatelyRequest) String() string { return proto.CompactTextString(m) }
func (*RearrangeArrayAlternatelyRequest) ProtoMessage()    {}
func (*RearrangeArrayAlternatelyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{1}
}

func (m *RearrangeArrayAlternatelyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RearrangeArrayAlternatelyRequest.Unmarshal(m, b)
}
func (m *RearrangeArrayAlternatelyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RearrangeArrayAlternatelyRequest.Marshal(b, m, deterministic)
}
func (m *RearrangeArrayAlternatelyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RearrangeArrayAlternatelyRequest.Merge(m, src)
}
func (m *RearrangeArrayAlternatelyRequest) XXX_Size() int {
	return xxx_messageInfo_RearrangeArrayAlternatelyRequest.Size(m)
}
func (m *RearrangeArrayAlternatelyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RearrangeArrayAlternatelyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RearrangeArrayAlternatelyRequest proto.InternalMessageInfo

func (m *RearrangeArrayAlternatelyRequest) GetArray() []int64 {
	if m != nil {
		return m.Array
	}
	return nil
}

type RearrangeArrayAlternatelyResponse struct {
	Array                []int64  `protobuf:"varint,1,rep,packed,name=array,proto3" json:"array,omitempty"`
	ArrangedArray        []int64  `protobuf:"varint,2,rep,packed,name=arranged_array,json=arrangedArray,proto3" json:"arranged_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RearrangeArrayAlternatelyResponse) Reset()         { *m = RearrangeArrayAlternatelyResponse{} }
func (m *RearrangeArrayAlternatelyResponse) String() string { return proto.CompactTextString(m) }
func (*RearrangeArrayAlternatelyResponse) ProtoMessage()    {}
func (*RearrangeArrayAlternatelyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{2}
}

func (m *RearrangeArrayAlternatelyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RearrangeArrayAlternatelyResponse.Unmarshal(m, b)
}
func (m *RearrangeArrayAlternatelyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RearrangeArrayAlternatelyResponse.Marshal(b, m, deterministic)
}
func (m *RearrangeArrayAlternatelyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RearrangeArrayAlternatelyResponse.Merge(m, src)
}
func (m *RearrangeArrayAlternatelyResponse) XXX_Size() int {
	return xxx_messageInfo_RearrangeArrayAlternatelyResponse.Size(m)
}
func (m *RearrangeArrayAlternatelyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RearrangeArrayAlternatelyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RearrangeArrayAlternatelyResponse proto.InternalMessageInfo

func (m *RearrangeArrayAlternatelyResponse) GetArray() []int64 {
	if m != nil {
		return m.Array
	}
	return nil
}

func (m *RearrangeArrayAlternatelyResponse) GetArrangedArray() []int64 {
	if m != nil {
		return m.ArrangedArray
	}
	return nil
}

type StreamClientSideRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamClientSideRequest) Reset()         { *m = StreamClientSideRequest{} }
func (m *StreamClientSideRequest) String() string { return proto.CompactTextString(m) }
func (*StreamClientSideRequest) ProtoMessage()    {}
func (*StreamClientSideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{3}
}

func (m *StreamClientSideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamClientSideRequest.Unmarshal(m, b)
}
func (m *StreamClientSideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamClientSideRequest.Marshal(b, m, deterministic)
}
func (m *StreamClientSideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamClientSideRequest.Merge(m, src)
}
func (m *StreamClientSideRequest) XXX_Size() int {
	return xxx_messageInfo_StreamClientSideRequest.Size(m)
}
func (m *StreamClientSideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamClientSideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamClientSideRequest proto.InternalMessageInfo

func (m *StreamClientSideRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamClientSideResponse struct {
	Sol                  string   `protobuf:"bytes,2,opt,name=sol,proto3" json:"sol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamClientSideResponse) Reset()         { *m = StreamClientSideResponse{} }
func (m *StreamClientSideResponse) String() string { return proto.CompactTextString(m) }
func (*StreamClientSideResponse) ProtoMessage()    {}
func (*StreamClientSideResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{4}
}

func (m *StreamClientSideResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamClientSideResponse.Unmarshal(m, b)
}
func (m *StreamClientSideResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamClientSideResponse.Marshal(b, m, deterministic)
}
func (m *StreamClientSideResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamClientSideResponse.Merge(m, src)
}
func (m *StreamClientSideResponse) XXX_Size() int {
	return xxx_messageInfo_StreamClientSideResponse.Size(m)
}
func (m *StreamClientSideResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamClientSideResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamClientSideResponse proto.InternalMessageInfo

func (m *StreamClientSideResponse) GetSol() string {
	if m != nil {
		return m.Sol
	}
	return ""
}

type ServerSideStreamRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerSideStreamRequest) Reset()         { *m = ServerSideStreamRequest{} }
func (m *ServerSideStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ServerSideStreamRequest) ProtoMessage()    {}
func (*ServerSideStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{5}
}

func (m *ServerSideStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerSideStreamRequest.Unmarshal(m, b)
}
func (m *ServerSideStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerSideStreamRequest.Marshal(b, m, deterministic)
}
func (m *ServerSideStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerSideStreamRequest.Merge(m, src)
}
func (m *ServerSideStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ServerSideStreamRequest.Size(m)
}
func (m *ServerSideStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerSideStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerSideStreamRequest proto.InternalMessageInfo

func (m *ServerSideStreamRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ServerSideStreamResponse struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerSideStreamResponse) Reset()         { *m = ServerSideStreamResponse{} }
func (m *ServerSideStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ServerSideStreamResponse) ProtoMessage()    {}
func (*ServerSideStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{6}
}

func (m *ServerSideStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerSideStreamResponse.Unmarshal(m, b)
}
func (m *ServerSideStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerSideStreamResponse.Marshal(b, m, deterministic)
}
func (m *ServerSideStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerSideStreamResponse.Merge(m, src)
}
func (m *ServerSideStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ServerSideStreamResponse.Size(m)
}
func (m *ServerSideStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerSideStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerSideStreamResponse proto.InternalMessageInfo

func (m *ServerSideStreamResponse) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type BidirectionalStreamRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidirectionalStreamRequest) Reset()         { *m = BidirectionalStreamRequest{} }
func (m *BidirectionalStreamRequest) String() string { return proto.CompactTextString(m) }
func (*BidirectionalStreamRequest) ProtoMessage()    {}
func (*BidirectionalStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{7}
}

func (m *BidirectionalStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalStreamRequest.Unmarshal(m, b)
}
func (m *BidirectionalStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalStreamRequest.Marshal(b, m, deterministic)
}
func (m *BidirectionalStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalStreamRequest.Merge(m, src)
}
func (m *BidirectionalStreamRequest) XXX_Size() int {
	return xxx_messageInfo_BidirectionalStreamRequest.Size(m)
}
func (m *BidirectionalStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalStreamRequest proto.InternalMessageInfo

func (m *BidirectionalStreamRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type BidirectionalStreamResponse struct {
	Numbers              int64    `protobuf:"varint,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidirectionalStreamResponse) Reset()         { *m = BidirectionalStreamResponse{} }
func (m *BidirectionalStreamResponse) String() string { return proto.CompactTextString(m) }
func (*BidirectionalStreamResponse) ProtoMessage()    {}
func (*BidirectionalStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d995446cfa84ab4, []int{8}
}

func (m *BidirectionalStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalStreamResponse.Unmarshal(m, b)
}
func (m *BidirectionalStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalStreamResponse.Marshal(b, m, deterministic)
}
func (m *BidirectionalStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalStreamResponse.Merge(m, src)
}
func (m *BidirectionalStreamResponse) XXX_Size() int {
	return xxx_messageInfo_BidirectionalStreamResponse.Size(m)
}
func (m *BidirectionalStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalStreamResponse proto.InternalMessageInfo

func (m *BidirectionalStreamResponse) GetNumbers() int64 {
	if m != nil {
		return m.Numbers
	}
	return 0
}

func init() {
	proto.RegisterType((*RearrangeArrayAlternately_Test_Case)(nil), "proto.RearrangeArrayAlternately_Test_Case")
	proto.RegisterType((*RearrangeArrayAlternatelyRequest)(nil), "proto.RearrangeArrayAlternatelyRequest")
	proto.RegisterType((*RearrangeArrayAlternatelyResponse)(nil), "proto.RearrangeArrayAlternatelyResponse")
	proto.RegisterType((*StreamClientSideRequest)(nil), "proto.StreamClientSideRequest")
	proto.RegisterType((*StreamClientSideResponse)(nil), "proto.StreamClientSideResponse")
	proto.RegisterType((*ServerSideStreamRequest)(nil), "proto.ServerSideStreamRequest")
	proto.RegisterType((*ServerSideStreamResponse)(nil), "proto.ServerSideStreamResponse")
	proto.RegisterType((*BidirectionalStreamRequest)(nil), "proto.BidirectionalStreamRequest")
	proto.RegisterType((*BidirectionalStreamResponse)(nil), "proto.BidirectionalStreamResponse")
}

func init() {
	proto.RegisterFile("pkg/proto/data-structures/data-structures.proto", fileDescriptor_5d995446cfa84ab4)
}

var fileDescriptor_5d995446cfa84ab4 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x0d, 0x4d, 0xc5, 0x48, 0x54, 0xd5, 0x0a, 0xa5, 0xc6, 0x20, 0x9a, 0x2e, 0x42,
	0x44, 0x55, 0x5b, 0xb7, 0xe9, 0xa1, 0x55, 0x6f, 0xa5, 0x3c, 0x41, 0xc2, 0x3d, 0x6c, 0x92, 0xc1,
	0x58, 0xd8, 0xeb, 0xb0, 0x3b, 0xae, 0x94, 0x2b, 0x0f, 0xc0, 0x85, 0x33, 0x77, 0xde, 0x87, 0x57,
	0xe0, 0x35, 0x90, 0x90, 0x77, 0xd7, 0x75, 0xe4, 0xc6, 0x6d, 0x4f, 0xed, 0x64, 0xe7, 0x9b, 0xff,
	0xf7, 0xce, 0xbf, 0x10, 0x2d, 0xbe, 0xc6, 0xd1, 0x42, 0xe7, 0x94, 0x47, 0x73, 0x49, 0xf2, 0xc8,
	0x90, 0x2e, 0x66, 0x54, 0x68, 0x34, 0xcd, 0xfa, 0xd8, 0x76, 0xf1, 0x4d, 0xfb, 0x27, 0xdc, 0x8b,
	0xf3, 0x3c, 0x4e, 0xd1, 0xa1, 0xd3, 0xe2, 0x73, 0x44, 0x49, 0x86, 0x86, 0x64, 0xb6, 0x70, 0x7d,
	0xe1, 0x2b, 0xdf, 0x20, 0x17, 0x49, 0x24, 0x95, 0xca, 0x49, 0x52, 0x92, 0x2b, 0x3f, 0x45, 0x8c,
	0xe1, 0xcd, 0x08, 0xa5, 0xd6, 0x52, 0xc5, 0x78, 0xa5, 0xb5, 0x5c, 0x5e, 0xa5, 0x84, 0x5a, 0x49,
	0xc2, 0x74, 0x39, 0xf9, 0x88, 0x86, 0x26, 0xd7, 0xd2, 0x20, 0xef, 0x41, 0x37, 0x45, 0x15, 0xd3,
	0x97, 0x80, 0xf5, 0xd9, 0xa0, 0x33, 0xf2, 0x15, 0x7f, 0x0e, 0x9b, 0x25, 0xbc, 0x0c, 0x36, 0xfa,
	0x9d, 0x41, 0x67, 0xe4, 0x0a, 0x71, 0x01, 0xfd, 0xd6, 0xa1, 0x23, 0xfc, 0x56, 0xa0, 0xa1, 0x92,
	0xb4, 0x47, 0x01, 0x73, 0xa4, 0x2d, 0xc4, 0x27, 0xd8, 0xbf, 0x87, 0x34, 0x8b, 0x5c, 0x19, 0xac,
	0x45, 0xd9, 0x8a, 0x28, 0x7f, 0x0b, 0xdb, 0x1e, 0x9c, 0x4f, 0x56, 0x3d, 0x3d, 0xab, 0x7e, 0x75,
	0x0a, 0x67, 0xb0, 0x3b, 0x26, 0x8d, 0x32, 0xbb, 0x4e, 0x13, 0x54, 0x34, 0x4e, 0xe6, 0x58, 0x59,
	0x0a, 0x60, 0x2b, 0x43, 0x63, 0x64, 0x8c, 0xf6, 0x2b, 0x9f, 0x8e, 0xaa, 0x52, 0x1c, 0x42, 0x70,
	0x17, 0xf2, 0x6e, 0x76, 0xa0, 0x63, 0xf2, 0x34, 0xd8, 0xb0, 0x44, 0xf9, 0xaf, 0x88, 0x60, 0x77,
	0x8c, 0xfa, 0x06, 0x75, 0xd9, 0xe7, 0xb8, 0x95, 0xaf, 0x9e, 0xe5, 0x85, 0x22, 0x7f, 0x8d, 0xae,
	0x10, 0x43, 0x08, 0xee, 0x02, 0x7e, 0x7c, 0x0f, 0xba, 0xaa, 0xc8, 0xa6, 0xa8, 0xab, 0x9b, 0x77,
	0x95, 0x18, 0x42, 0xf8, 0x3e, 0x99, 0x27, 0x1a, 0x67, 0xe5, 0x3a, 0x65, 0xfa, 0x18, 0x9d, 0x73,
	0x78, 0xb9, 0x96, 0xf1, 0x52, 0x01, 0x6c, 0xb9, 0xe1, 0xc6, 0x63, 0x55, 0x39, 0xfc, 0xf7, 0x04,
	0xb6, 0x3f, 0x48, 0x92, 0xe3, 0xdb, 0x10, 0xf2, 0xdf, 0x0c, 0x5e, 0xb4, 0xae, 0x8a, 0xbf, 0x73,
	0xf1, 0x3a, 0x7e, 0x28, 0x06, 0xe1, 0xe0, 0xe1, 0x46, 0xe7, 0x4e, 0x5c, 0x7c, 0xff, 0xf3, 0xf7,
	0xe7, 0xc6, 0x50, 0x1c, 0x45, 0x37, 0xa7, 0xf6, 0x49, 0x4c, 0x8c, 0xa6, 0xea, 0x89, 0xb4, 0xe2,
	0x97, 0xec, 0x80, 0xff, 0x60, 0xb0, 0xd3, 0x5c, 0x1f, 0x7f, 0xed, 0x85, 0x5b, 0xc2, 0x10, 0xee,
	0xb5, 0x9e, 0x7b, 0x3f, 0xe7, 0xd6, 0xcf, 0xa9, 0x38, 0xac, 0xfd, 0xd4, 0x4f, 0x96, 0xd0, 0x50,
	0x34, 0xbb, 0x85, 0xdc, 0x90, 0x44, 0xc5, 0x97, 0xec, 0x60, 0xc0, 0x9c, 0xa1, 0xc6, 0xc2, 0x6b,
	0x43, 0xeb, 0xa3, 0x53, 0x1b, 0x6a, 0x49, 0xca, 0x23, 0x0c, 0x99, 0x06, 0xea, 0x0c, 0x9d, 0x30,
	0xfe, 0x8b, 0x41, 0x6f, 0x4d, 0x32, 0x12, 0x15, 0xf3, 0x7d, 0x2f, 0xdb, 0x1e, 0xb6, 0x50, 0xdc,
	0xd7, 0xd2, 0xbe, 0xbd, 0x86, 0xb9, 0xe9, 0x2a, 0x5d, 0x9e, 0xa2, 0xcc, 0xec, 0x75, 0x9d, 0xb0,
	0x69, 0xd7, 0x0a, 0x9c, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x65, 0x87, 0x2c, 0x3f, 0x25, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataStructuresClient is the client API for DataStructures service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataStructuresClient interface {
	RearrangeArrayAlternately(ctx context.Context, in *RearrangeArrayAlternatelyRequest, opts ...grpc.CallOption) (*RearrangeArrayAlternatelyResponse, error)
	StreamClientSide(ctx context.Context, opts ...grpc.CallOption) (DataStructures_StreamClientSideClient, error)
	ServerSideStream(ctx context.Context, in *ServerSideStreamRequest, opts ...grpc.CallOption) (DataStructures_ServerSideStreamClient, error)
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DataStructures_BidirectionalStreamingClient, error)
}

type dataStructuresClient struct {
	cc *grpc.ClientConn
}

func NewDataStructuresClient(cc *grpc.ClientConn) DataStructuresClient {
	return &dataStructuresClient{cc}
}

func (c *dataStructuresClient) RearrangeArrayAlternately(ctx context.Context, in *RearrangeArrayAlternatelyRequest, opts ...grpc.CallOption) (*RearrangeArrayAlternatelyResponse, error) {
	out := new(RearrangeArrayAlternatelyResponse)
	err := c.cc.Invoke(ctx, "/proto.DataStructures/RearrangeArrayAlternately", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStructuresClient) StreamClientSide(ctx context.Context, opts ...grpc.CallOption) (DataStructures_StreamClientSideClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataStructures_serviceDesc.Streams[0], "/proto.DataStructures/StreamClientSide", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataStructuresStreamClientSideClient{stream}
	return x, nil
}

type DataStructures_StreamClientSideClient interface {
	Send(*StreamClientSideRequest) error
	CloseAndRecv() (*StreamClientSideResponse, error)
	grpc.ClientStream
}

type dataStructuresStreamClientSideClient struct {
	grpc.ClientStream
}

func (x *dataStructuresStreamClientSideClient) Send(m *StreamClientSideRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataStructuresStreamClientSideClient) CloseAndRecv() (*StreamClientSideResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamClientSideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataStructuresClient) ServerSideStream(ctx context.Context, in *ServerSideStreamRequest, opts ...grpc.CallOption) (DataStructures_ServerSideStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataStructures_serviceDesc.Streams[1], "/proto.DataStructures/ServerSideStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataStructuresServerSideStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataStructures_ServerSideStreamClient interface {
	Recv() (*ServerSideStreamResponse, error)
	grpc.ClientStream
}

type dataStructuresServerSideStreamClient struct {
	grpc.ClientStream
}

func (x *dataStructuresServerSideStreamClient) Recv() (*ServerSideStreamResponse, error) {
	m := new(ServerSideStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataStructuresClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DataStructures_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataStructures_serviceDesc.Streams[2], "/proto.DataStructures/BidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataStructuresBidirectionalStreamingClient{stream}
	return x, nil
}

type DataStructures_BidirectionalStreamingClient interface {
	Send(*BidirectionalStreamRequest) error
	Recv() (*BidirectionalStreamResponse, error)
	grpc.ClientStream
}

type dataStructuresBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *dataStructuresBidirectionalStreamingClient) Send(m *BidirectionalStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataStructuresBidirectionalStreamingClient) Recv() (*BidirectionalStreamResponse, error) {
	m := new(BidirectionalStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataStructuresServer is the server API for DataStructures service.
type DataStructuresServer interface {
	RearrangeArrayAlternately(context.Context, *RearrangeArrayAlternatelyRequest) (*RearrangeArrayAlternatelyResponse, error)
	StreamClientSide(DataStructures_StreamClientSideServer) error
	ServerSideStream(*ServerSideStreamRequest, DataStructures_ServerSideStreamServer) error
	BidirectionalStreaming(DataStructures_BidirectionalStreamingServer) error
}

// UnimplementedDataStructuresServer can be embedded to have forward compatible implementations.
type UnimplementedDataStructuresServer struct {
}

func (*UnimplementedDataStructuresServer) RearrangeArrayAlternately(ctx context.Context, req *RearrangeArrayAlternatelyRequest) (*RearrangeArrayAlternatelyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RearrangeArrayAlternately not implemented")
}
func (*UnimplementedDataStructuresServer) StreamClientSide(srv DataStructures_StreamClientSideServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientSide not implemented")
}
func (*UnimplementedDataStructuresServer) ServerSideStream(req *ServerSideStreamRequest, srv DataStructures_ServerSideStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStream not implemented")
}
func (*UnimplementedDataStructuresServer) BidirectionalStreaming(srv DataStructures_BidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreaming not implemented")
}

func RegisterDataStructuresServer(s *grpc.Server, srv DataStructuresServer) {
	s.RegisterService(&_DataStructures_serviceDesc, srv)
}

func _DataStructures_RearrangeArrayAlternately_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RearrangeArrayAlternatelyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStructuresServer).RearrangeArrayAlternately(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DataStructures/RearrangeArrayAlternately",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStructuresServer).RearrangeArrayAlternately(ctx, req.(*RearrangeArrayAlternatelyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataStructures_StreamClientSide_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataStructuresServer).StreamClientSide(&dataStructuresStreamClientSideServer{stream})
}

type DataStructures_StreamClientSideServer interface {
	SendAndClose(*StreamClientSideResponse) error
	Recv() (*StreamClientSideRequest, error)
	grpc.ServerStream
}

type dataStructuresStreamClientSideServer struct {
	grpc.ServerStream
}

func (x *dataStructuresStreamClientSideServer) SendAndClose(m *StreamClientSideResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataStructuresStreamClientSideServer) Recv() (*StreamClientSideRequest, error) {
	m := new(StreamClientSideRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataStructures_ServerSideStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerSideStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataStructuresServer).ServerSideStream(m, &dataStructuresServerSideStreamServer{stream})
}

type DataStructures_ServerSideStreamServer interface {
	Send(*ServerSideStreamResponse) error
	grpc.ServerStream
}

type dataStructuresServerSideStreamServer struct {
	grpc.ServerStream
}

func (x *dataStructuresServerSideStreamServer) Send(m *ServerSideStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataStructures_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataStructuresServer).BidirectionalStreaming(&dataStructuresBidirectionalStreamingServer{stream})
}

type DataStructures_BidirectionalStreamingServer interface {
	Send(*BidirectionalStreamResponse) error
	Recv() (*BidirectionalStreamRequest, error)
	grpc.ServerStream
}

type dataStructuresBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *dataStructuresBidirectionalStreamingServer) Send(m *BidirectionalStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataStructuresBidirectionalStreamingServer) Recv() (*BidirectionalStreamRequest, error) {
	m := new(BidirectionalStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataStructures_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DataStructures",
	HandlerType: (*DataStructuresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RearrangeArrayAlternately",
			Handler:    _DataStructures_RearrangeArrayAlternately_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientSide",
			Handler:       _DataStructures_StreamClientSide_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerSideStream",
			Handler:       _DataStructures_ServerSideStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _DataStructures_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/data-structures/data-structures.proto",
}
