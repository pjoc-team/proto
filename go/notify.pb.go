// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notify.proto

package pay

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
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

type HttpNotifyRequest struct {
	Channel              string             `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Account              string             `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	OrderId              string             `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	HttpMethod           string             `protobuf:"bytes,4,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	Body                 *httpbody.HttpBody `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpNotifyRequest) Reset()         { *m = HttpNotifyRequest{} }
func (m *HttpNotifyRequest) String() string { return proto.CompactTextString(m) }
func (*HttpNotifyRequest) ProtoMessage()    {}
func (*HttpNotifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba76cc4ebe272d4, []int{0}
}

func (m *HttpNotifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpNotifyRequest.Unmarshal(m, b)
}
func (m *HttpNotifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpNotifyRequest.Marshal(b, m, deterministic)
}
func (m *HttpNotifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpNotifyRequest.Merge(m, src)
}
func (m *HttpNotifyRequest) XXX_Size() int {
	return xxx_messageInfo_HttpNotifyRequest.Size(m)
}
func (m *HttpNotifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpNotifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HttpNotifyRequest proto.InternalMessageInfo

func (m *HttpNotifyRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *HttpNotifyRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *HttpNotifyRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *HttpNotifyRequest) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *HttpNotifyRequest) GetBody() *httpbody.HttpBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type HttpNotifyResponse struct {
	Body                 []byte            `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Header               map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HttpNotifyResponse) Reset()         { *m = HttpNotifyResponse{} }
func (m *HttpNotifyResponse) String() string { return proto.CompactTextString(m) }
func (*HttpNotifyResponse) ProtoMessage()    {}
func (*HttpNotifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba76cc4ebe272d4, []int{1}
}

func (m *HttpNotifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpNotifyResponse.Unmarshal(m, b)
}
func (m *HttpNotifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpNotifyResponse.Marshal(b, m, deterministic)
}
func (m *HttpNotifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpNotifyResponse.Merge(m, src)
}
func (m *HttpNotifyResponse) XXX_Size() int {
	return xxx_messageInfo_HttpNotifyResponse.Size(m)
}
func (m *HttpNotifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpNotifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HttpNotifyResponse proto.InternalMessageInfo

func (m *HttpNotifyResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpNotifyResponse) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpNotifyRequest)(nil), "pay.HttpNotifyRequest")
	proto.RegisterType((*HttpNotifyResponse)(nil), "pay.HttpNotifyResponse")
	proto.RegisterMapType((map[string]string)(nil), "pay.HttpNotifyResponse.HeaderEntry")
}

func init() { proto.RegisterFile("notify.proto", fileDescriptor_aba76cc4ebe272d4) }

var fileDescriptor_aba76cc4ebe272d4 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x66, 0x92, 0x76, 0xd5, 0xc9, 0x1e, 0x74, 0x58, 0x4a, 0xba, 0x16, 0xba, 0xa4, 0x1e, 0xf6,
	0x34, 0x83, 0x2b, 0x88, 0xd6, 0x83, 0xb8, 0x20, 0xad, 0x07, 0x45, 0xf6, 0xe8, 0xa5, 0xcc, 0x26,
	0xd3, 0x26, 0xba, 0x99, 0x37, 0x26, 0x93, 0x42, 0x58, 0xf6, 0x22, 0xfe, 0x00, 0xc1, 0x3f, 0x20,
	0xde, 0xfd, 0x23, 0xde, 0xc4, 0xbf, 0xe0, 0x0f, 0x91, 0xcc, 0xbc, 0x85, 0x40, 0x8b, 0x50, 0xbc,
	0xbd, 0x6f, 0xbe, 0xf7, 0xcd, 0x7c, 0xef, 0x7b, 0x09, 0x1d, 0x6a, 0xb0, 0xc5, 0x79, 0xcb, 0x4d,
	0x05, 0x16, 0x58, 0x68, 0x64, 0x3b, 0xde, 0xbf, 0x00, 0xb8, 0x58, 0x29, 0x21, 0x4d, 0x21, 0x72,
	0x6b, 0xcd, 0x12, 0x32, 0xe4, 0xc7, 0x07, 0x3d, 0x4a, 0x6a, 0x0d, 0x56, 0xda, 0x02, 0x74, 0x8d,
	0xec, 0x21, 0xb2, 0x0e, 0x2d, 0x9b, 0x73, 0x61, 0x8b, 0x52, 0xd5, 0x56, 0x96, 0x06, 0x1b, 0x86,
	0x29, 0x94, 0x25, 0x68, 0x8f, 0x92, 0x1f, 0x84, 0xde, 0x3b, 0xb5, 0xd6, 0xbc, 0x71, 0x0e, 0x16,
	0xea, 0x63, 0xa3, 0x6a, 0xcb, 0x62, 0x7a, 0x2b, 0xcd, 0xa5, 0xd6, 0x6a, 0x15, 0x93, 0x09, 0x99,
	0xde, 0x59, 0x6c, 0x61, 0xc7, 0xc8, 0x34, 0x85, 0x46, 0xdb, 0x38, 0xf0, 0x0c, 0x42, 0xb6, 0x4f,
	0x6f, 0x43, 0x95, 0xa9, 0xea, 0xac, 0xc8, 0xe2, 0xd0, 0x53, 0x0e, 0xbf, 0xca, 0xd8, 0x21, 0x8d,
	0xba, 0x19, 0xce, 0x4a, 0x65, 0x73, 0xc8, 0xe2, 0x1d, 0xc7, 0xd2, 0xee, 0xe8, 0xb5, 0x3b, 0x61,
	0x53, 0xba, 0xd3, 0x0d, 0x18, 0xef, 0x4e, 0xc8, 0x34, 0x9a, 0x8d, 0xb8, 0x9f, 0x81, 0x4b, 0x53,
	0xf0, 0xce, 0xdc, 0x1c, 0xb2, 0x76, 0xe1, 0x3a, 0x92, 0x6f, 0x84, 0xb2, 0xbe, 0xdf, 0xda, 0x80,
	0xae, 0x15, 0x63, 0x78, 0x41, 0xe7, 0x76, 0xe8, 0x5b, 0xd9, 0x33, 0x3a, 0xc8, 0x95, 0xcc, 0x54,
	0x15, 0x07, 0x93, 0x70, 0x1a, 0xcd, 0x8e, 0xb8, 0x91, 0x2d, 0xbf, 0x2a, 0xe6, 0xa7, 0xae, 0xeb,
	0xa5, 0xb6, 0x55, 0xbb, 0x40, 0xc9, 0xf8, 0x29, 0x8d, 0x7a, 0xc7, 0xec, 0x2e, 0x0d, 0x3f, 0xa8,
	0x16, 0xc3, 0xe8, 0x4a, 0x36, 0xa2, 0xbb, 0x97, 0x72, 0xd5, 0x28, 0x8c, 0xc1, 0x83, 0xe3, 0xe0,
	0x09, 0x99, 0x7d, 0x0f, 0xe9, 0xc0, 0xbf, 0xc0, 0xbe, 0x04, 0x74, 0xe8, 0xcb, 0x79, 0xfb, 0x16,
	0x6a, 0xcb, 0xf6, 0xae, 0x78, 0x70, 0x81, 0x8f, 0xaf, 0x1d, 0x39, 0xf9, 0x45, 0x3e, 0xfd, 0xfe,
	0xf3, 0x35, 0xf8, 0x49, 0x92, 0xe7, 0xe2, 0xf2, 0xa1, 0xf0, 0x5f, 0x89, 0xc0, 0x55, 0x88, 0x35,
	0x16, 0x1b, 0x81, 0x2b, 0x10, 0x6b, 0x2c, 0x36, 0xc2, 0x25, 0x2f, 0xd6, 0xdb, 0x85, 0x6c, 0x8e,
	0x5d, 0x24, 0xef, 0x1e, 0x24, 0x07, 0xff, 0xba, 0x08, 0xbb, 0x1e, 0x27, 0xfc, 0x66, 0xcf, 0xa1,
	0xee, 0x28, 0xb9, 0xdf, 0xd3, 0x5d, 0x6f, 0x81, 0x7d, 0x26, 0x34, 0xda, 0x46, 0x72, 0xa2, 0x6e,
	0x9a, 0xc8, 0x89, 0x0b, 0xe4, 0x05, 0xfb, 0xdf, 0x3c, 0xe6, 0x7b, 0x74, 0x64, 0x9a, 0x25, 0x37,
	0xef, 0x21, 0x75, 0x06, 0x50, 0xbb, 0x1c, 0xb8, 0xdf, 0xe2, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd9, 0x60, 0x63, 0xff, 0x93, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifyClient interface {
	NotifyByPost(ctx context.Context, in *HttpNotifyRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	NotifyByGet(ctx context.Context, in *HttpNotifyRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type notifyClient struct {
	cc *grpc.ClientConn
}

func NewNotifyClient(cc *grpc.ClientConn) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) NotifyByPost(ctx context.Context, in *HttpNotifyRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/pay.Notify/NotifyByPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) NotifyByGet(ctx context.Context, in *HttpNotifyRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/pay.Notify/NotifyByGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServer is the server API for Notify service.
type NotifyServer interface {
	NotifyByPost(context.Context, *HttpNotifyRequest) (*httpbody.HttpBody, error)
	NotifyByGet(context.Context, *HttpNotifyRequest) (*httpbody.HttpBody, error)
}

// UnimplementedNotifyServer can be embedded to have forward compatible implementations.
type UnimplementedNotifyServer struct {
}

func (*UnimplementedNotifyServer) NotifyByPost(ctx context.Context, req *HttpNotifyRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyByPost not implemented")
}
func (*UnimplementedNotifyServer) NotifyByGet(ctx context.Context, req *HttpNotifyRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyByGet not implemented")
}

func RegisterNotifyServer(s *grpc.Server, srv NotifyServer) {
	s.RegisterService(&_Notify_serviceDesc, srv)
}

func _Notify_NotifyByPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).NotifyByPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay.Notify/NotifyByPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).NotifyByPost(ctx, req.(*HttpNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_NotifyByGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).NotifyByGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay.Notify/NotifyByGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).NotifyByGet(ctx, req.(*HttpNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pay.Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyByPost",
			Handler:    _Notify_NotifyByPost_Handler,
		},
		{
			MethodName: "NotifyByGet",
			Handler:    _Notify_NotifyByGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify.proto",
}