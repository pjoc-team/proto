// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: callback.proto

package pay

import (
	context "context"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HttpCallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel    string             `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Account    string             `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	OrderId    string             `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	HttpMethod string             `protobuf:"bytes,4,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	Body       *httpbody.HttpBody `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HttpCallbackRequest) Reset() {
	*x = HttpCallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_callback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpCallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpCallbackRequest) ProtoMessage() {}

func (x *HttpCallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_callback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpCallbackRequest.ProtoReflect.Descriptor instead.
func (*HttpCallbackRequest) Descriptor() ([]byte, []int) {
	return file_callback_proto_rawDescGZIP(), []int{0}
}

func (x *HttpCallbackRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *HttpCallbackRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *HttpCallbackRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *HttpCallbackRequest) GetHttpMethod() string {
	if x != nil {
		return x.HttpMethod
	}
	return ""
}

func (x *HttpCallbackRequest) GetBody() *httpbody.HttpBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type HttpCallbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body   []byte            `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Header map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HttpCallbackResponse) Reset() {
	*x = HttpCallbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_callback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpCallbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpCallbackResponse) ProtoMessage() {}

func (x *HttpCallbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_callback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpCallbackResponse.ProtoReflect.Descriptor instead.
func (*HttpCallbackResponse) Descriptor() ([]byte, []int) {
	return file_callback_proto_rawDescGZIP(), []int{1}
}

func (x *HttpCallbackResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *HttpCallbackResponse) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_callback_proto protoreflect.FileDescriptor

var file_callback_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x61, 0x79, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01,
	0x0a, 0x13, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0xa4, 0x01, 0x0a, 0x14, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3d, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xb8, 0x04, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x83, 0x02, 0x0a, 0x0d, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x79, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xbf, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0xb8, 0x01, 0x12, 0x3f, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f,
	0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x5a, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x5a, 0x1f, 0x12,
	0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x30, 0x01,
	0x12, 0x9e, 0x02, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x79, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0xd9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xd2, 0x01, 0x22, 0x41, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x5a, 0x26, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x7b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x5a, 0x38,
	0x22, 0x30, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x5a, 0x25, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x7b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x30,
	0x01, 0x42, 0x16, 0x0a, 0x14, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x6a, 0x6f, 0x63, 0x2e, 0x70, 0x61,
	0x79, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_callback_proto_rawDescOnce sync.Once
	file_callback_proto_rawDescData = file_callback_proto_rawDesc
)

func file_callback_proto_rawDescGZIP() []byte {
	file_callback_proto_rawDescOnce.Do(func() {
		file_callback_proto_rawDescData = protoimpl.X.CompressGZIP(file_callback_proto_rawDescData)
	})
	return file_callback_proto_rawDescData
}

var file_callback_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_callback_proto_goTypes = []interface{}{
	(*HttpCallbackRequest)(nil),  // 0: pay.HttpCallbackRequest
	(*HttpCallbackResponse)(nil), // 1: pay.HttpCallbackResponse
	nil,                          // 2: pay.HttpCallbackResponse.HeaderEntry
	(*httpbody.HttpBody)(nil),    // 3: google.api.HttpBody
}
var file_callback_proto_depIdxs = []int32{
	3, // 0: pay.HttpCallbackRequest.body:type_name -> google.api.HttpBody
	2, // 1: pay.HttpCallbackResponse.header:type_name -> pay.HttpCallbackResponse.HeaderEntry
	0, // 2: pay.ChannelCallback.CallbackByGet:input_type -> pay.HttpCallbackRequest
	0, // 3: pay.ChannelCallback.CallbackByPost:input_type -> pay.HttpCallbackRequest
	3, // 4: pay.ChannelCallback.CallbackByGet:output_type -> google.api.HttpBody
	3, // 5: pay.ChannelCallback.CallbackByPost:output_type -> google.api.HttpBody
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_callback_proto_init() }
func file_callback_proto_init() {
	if File_callback_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_callback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpCallbackRequest); i {
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
		file_callback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpCallbackResponse); i {
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
			RawDescriptor: file_callback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_callback_proto_goTypes,
		DependencyIndexes: file_callback_proto_depIdxs,
		MessageInfos:      file_callback_proto_msgTypes,
	}.Build()
	File_callback_proto = out.File
	file_callback_proto_rawDesc = nil
	file_callback_proto_goTypes = nil
	file_callback_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChannelCallbackClient is the client API for ChannelCallback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChannelCallbackClient interface {
	CallbackByGet(ctx context.Context, in *HttpCallbackRequest, opts ...grpc.CallOption) (ChannelCallback_CallbackByGetClient, error)
	CallbackByPost(ctx context.Context, in *HttpCallbackRequest, opts ...grpc.CallOption) (ChannelCallback_CallbackByPostClient, error)
}

type channelCallbackClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelCallbackClient(cc grpc.ClientConnInterface) ChannelCallbackClient {
	return &channelCallbackClient{cc}
}

func (c *channelCallbackClient) CallbackByGet(ctx context.Context, in *HttpCallbackRequest, opts ...grpc.CallOption) (ChannelCallback_CallbackByGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChannelCallback_serviceDesc.Streams[0], "/pay.ChannelCallback/CallbackByGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &channelCallbackCallbackByGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChannelCallback_CallbackByGetClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type channelCallbackCallbackByGetClient struct {
	grpc.ClientStream
}

func (x *channelCallbackCallbackByGetClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *channelCallbackClient) CallbackByPost(ctx context.Context, in *HttpCallbackRequest, opts ...grpc.CallOption) (ChannelCallback_CallbackByPostClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChannelCallback_serviceDesc.Streams[1], "/pay.ChannelCallback/CallbackByPost", opts...)
	if err != nil {
		return nil, err
	}
	x := &channelCallbackCallbackByPostClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChannelCallback_CallbackByPostClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type channelCallbackCallbackByPostClient struct {
	grpc.ClientStream
}

func (x *channelCallbackCallbackByPostClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChannelCallbackServer is the server API for ChannelCallback service.
type ChannelCallbackServer interface {
	CallbackByGet(*HttpCallbackRequest, ChannelCallback_CallbackByGetServer) error
	CallbackByPost(*HttpCallbackRequest, ChannelCallback_CallbackByPostServer) error
}

// UnimplementedChannelCallbackServer can be embedded to have forward compatible implementations.
type UnimplementedChannelCallbackServer struct {
}

func (*UnimplementedChannelCallbackServer) CallbackByGet(*HttpCallbackRequest, ChannelCallback_CallbackByGetServer) error {
	return status.Errorf(codes.Unimplemented, "method CallbackByGet not implemented")
}
func (*UnimplementedChannelCallbackServer) CallbackByPost(*HttpCallbackRequest, ChannelCallback_CallbackByPostServer) error {
	return status.Errorf(codes.Unimplemented, "method CallbackByPost not implemented")
}

func RegisterChannelCallbackServer(s *grpc.Server, srv ChannelCallbackServer) {
	s.RegisterService(&_ChannelCallback_serviceDesc, srv)
}

func _ChannelCallback_CallbackByGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HttpCallbackRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChannelCallbackServer).CallbackByGet(m, &channelCallbackCallbackByGetServer{stream})
}

type ChannelCallback_CallbackByGetServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type channelCallbackCallbackByGetServer struct {
	grpc.ServerStream
}

func (x *channelCallbackCallbackByGetServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

func _ChannelCallback_CallbackByPost_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HttpCallbackRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChannelCallbackServer).CallbackByPost(m, &channelCallbackCallbackByPostServer{stream})
}

type ChannelCallback_CallbackByPostServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type channelCallbackCallbackByPostServer struct {
	grpc.ServerStream
}

func (x *channelCallbackCallbackByPostServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

var _ChannelCallback_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pay.ChannelCallback",
	HandlerType: (*ChannelCallbackServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CallbackByGet",
			Handler:       _ChannelCallback_CallbackByGet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CallbackByPost",
			Handler:       _ChannelCallback_CallbackByPost_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "callback.proto",
}
