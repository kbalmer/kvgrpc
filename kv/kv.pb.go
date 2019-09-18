// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv.proto

package kv

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Request struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	External             bool     `protobuf:"varint,2,opt,name=external,proto3" json:"external,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Request) GetExternal() bool {
	if m != nil {
		return m.External
	}
	return false
}

func (m *Request) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ValueReturn struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueReturn) Reset()         { *m = ValueReturn{} }
func (m *ValueReturn) String() string { return proto.CompactTextString(m) }
func (*ValueReturn) ProtoMessage()    {}
func (*ValueReturn) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{2}
}

func (m *ValueReturn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueReturn.Unmarshal(m, b)
}
func (m *ValueReturn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueReturn.Marshal(b, m, deterministic)
}
func (m *ValueReturn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueReturn.Merge(m, src)
}
func (m *ValueReturn) XXX_Size() int {
	return xxx_messageInfo_ValueReturn.Size(m)
}
func (m *ValueReturn) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueReturn.DiscardUnknown(m)
}

var xxx_messageInfo_ValueReturn proto.InternalMessageInfo

func (m *ValueReturn) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PostRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	External             bool     `protobuf:"varint,3,opt,name=external,proto3" json:"external,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{3}
}

func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (m *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(m, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PostRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PostRequest) GetExternal() bool {
	if m != nil {
		return m.External
	}
	return false
}

func (m *PostRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "kv.Empty")
	proto.RegisterType((*Request)(nil), "kv.Request")
	proto.RegisterType((*ValueReturn)(nil), "kv.ValueReturn")
	proto.RegisterType((*PostRequest)(nil), "kv.PostRequest")
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor_2216fe83c9c12408) }

var fileDescriptor_2216fe83c9c12408 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x51, 0x4b, 0x84, 0x40,
	0x10, 0xc7, 0x59, 0xd7, 0xf3, 0x74, 0x0c, 0x8a, 0xa1, 0x07, 0x91, 0x08, 0xd9, 0xeb, 0xc1, 0x27,
	0x89, 0xfa, 0x0a, 0x45, 0x41, 0x3d, 0xc4, 0x42, 0xf7, 0x6e, 0x30, 0xd0, 0xa1, 0x9e, 0xa6, 0xa3,
	0x74, 0x1f, 0xac, 0xef, 0x17, 0xbb, 0xc5, 0xb9, 0x57, 0xd4, 0xdb, 0xfc, 0xff, 0x3b, 0x33, 0xff,
	0x1f, 0x3b, 0x10, 0x56, 0x53, 0xd1, 0xf5, 0x2d, 0xb7, 0xe8, 0x55, 0x93, 0x5a, 0xc2, 0xe2, 0xb6,
	0xe9, 0x78, 0xa7, 0x9e, 0x61, 0xa9, 0xe9, 0x6d, 0xa4, 0x81, 0xf1, 0x04, 0x64, 0x45, 0xbb, 0x44,
	0x64, 0x22, 0x8f, 0xb4, 0x29, 0x31, 0x85, 0x90, 0xde, 0x99, 0xfa, 0x6d, 0x59, 0x27, 0x5e, 0x26,
	0xf2, 0x50, 0xef, 0x35, 0x9e, 0x41, 0xc4, 0x9b, 0x86, 0x06, 0x2e, 0x9b, 0x2e, 0x91, 0x99, 0xc8,
	0xa5, 0x9e, 0x0d, 0xb5, 0x82, 0x78, 0x5d, 0xd6, 0x23, 0x69, 0xe2, 0xb1, 0xdf, 0xe2, 0x29, 0x2c,
	0x26, 0x23, 0xed, 0xf2, 0x23, 0xfd, 0x25, 0x54, 0x0b, 0xf1, 0x53, 0x3b, 0xf0, 0xdf, 0xf9, 0xfb,
	0x31, 0xcf, 0x19, 0x3b, 0xa0, 0x92, 0xff, 0x51, 0xf9, 0x3f, 0xa8, 0xae, 0x3e, 0x04, 0x78, 0x0f,
	0x6b, 0x5c, 0x81, 0xbc, 0x23, 0xc6, 0xb8, 0xa8, 0xa6, 0xe2, 0x3b, 0x3c, 0x3d, 0x36, 0xc2, 0x45,
	0x56, 0xe0, 0x1b, 0x38, 0xb4, 0x0f, 0x0e, 0x66, 0x1a, 0x19, 0xc3, 0x7e, 0x1e, 0x66, 0x10, 0xdc,
	0x50, 0x4d, 0x4c, 0x87, 0xbb, 0x9c, 0x8e, 0x0b, 0xf0, 0x1f, 0x37, 0x03, 0xe3, 0x6c, 0xfd, 0x4a,
	0xba, 0x14, 0x78, 0x0e, 0xc1, 0x3d, 0x95, 0x35, 0xbf, 0xba, 0x7d, 0x73, 0xf9, 0x12, 0xd8, 0xc3,
	0x5d, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x63, 0x1d, 0x81, 0x5c, 0xc4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	// Get method
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ValueReturn, error)
	// Post method
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Empty, error)
	// Delete method
	Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	// List method
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (KV_ListClient, error)
	// Health method
	Health(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type kVClient struct {
	cc *grpc.ClientConn
}

func NewKVClient(cc *grpc.ClientConn) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ValueReturn, error) {
	out := new(ValueReturn)
	err := c.cc.Invoke(ctx, "/kv.KV/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kv.KV/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kv.KV/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (KV_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KV_serviceDesc.Streams[0], "/kv.KV/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_ListClient interface {
	Recv() (*ValueReturn, error)
	grpc.ClientStream
}

type kVListClient struct {
	grpc.ClientStream
}

func (x *kVListClient) Recv() (*ValueReturn, error) {
	m := new(ValueReturn)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) Health(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kv.KV/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	// Get method
	Get(context.Context, *Request) (*ValueReturn, error)
	// Post method
	Post(context.Context, *PostRequest) (*Empty, error)
	// Delete method
	Delete(context.Context, *Request) (*Empty, error)
	// List method
	List(*Empty, KV_ListServer) error
	// Health method
	Health(context.Context, *Empty) (*Empty, error)
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) Get(ctx context.Context, req *Request) (*ValueReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVServer) Post(ctx context.Context, req *PostRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (*UnimplementedKVServer) Delete(ctx context.Context, req *Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedKVServer) List(req *Empty, srv KV_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedKVServer) Health(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Delete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).List(m, &kVListServer{stream})
}

type KV_ListServer interface {
	Send(*ValueReturn) error
	grpc.ServerStream
}

type kVListServer struct {
	grpc.ServerStream
}

func (x *kVListServer) Send(m *ValueReturn) error {
	return x.ServerStream.SendMsg(m)
}

func _KV_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Health(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kv.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KV_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _KV_Post_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KV_Delete_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _KV_Health_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _KV_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kv.proto",
}