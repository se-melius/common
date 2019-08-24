// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attributes.proto

package attrspb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Attribute struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f97a9220cd723d2e, []int{0}
}

func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute.Unmarshal(m, b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
}
func (m *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(m, src)
}
func (m *Attribute) XXX_Size() int {
	return xxx_messageInfo_Attribute.Size(m)
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attribute) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Attribute) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetAttributeRequest struct {
	Key                  string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Info                 *Attribute `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SetAttributeRequest) Reset()         { *m = SetAttributeRequest{} }
func (m *SetAttributeRequest) String() string { return proto.CompactTextString(m) }
func (*SetAttributeRequest) ProtoMessage()    {}
func (*SetAttributeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f97a9220cd723d2e, []int{1}
}

func (m *SetAttributeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAttributeRequest.Unmarshal(m, b)
}
func (m *SetAttributeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAttributeRequest.Marshal(b, m, deterministic)
}
func (m *SetAttributeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAttributeRequest.Merge(m, src)
}
func (m *SetAttributeRequest) XXX_Size() int {
	return xxx_messageInfo_SetAttributeRequest.Size(m)
}
func (m *SetAttributeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAttributeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetAttributeRequest proto.InternalMessageInfo

func (m *SetAttributeRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetAttributeRequest) GetInfo() *Attribute {
	if m != nil {
		return m.Info
	}
	return nil
}

type SetAttributeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetAttributeResponse) Reset()         { *m = SetAttributeResponse{} }
func (m *SetAttributeResponse) String() string { return proto.CompactTextString(m) }
func (*SetAttributeResponse) ProtoMessage()    {}
func (*SetAttributeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f97a9220cd723d2e, []int{2}
}

func (m *SetAttributeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAttributeResponse.Unmarshal(m, b)
}
func (m *SetAttributeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAttributeResponse.Marshal(b, m, deterministic)
}
func (m *SetAttributeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAttributeResponse.Merge(m, src)
}
func (m *SetAttributeResponse) XXX_Size() int {
	return xxx_messageInfo_SetAttributeResponse.Size(m)
}
func (m *SetAttributeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAttributeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetAttributeResponse proto.InternalMessageInfo

type GetAttributesRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttributesRequest) Reset()         { *m = GetAttributesRequest{} }
func (m *GetAttributesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAttributesRequest) ProtoMessage()    {}
func (*GetAttributesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f97a9220cd723d2e, []int{3}
}

func (m *GetAttributesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttributesRequest.Unmarshal(m, b)
}
func (m *GetAttributesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttributesRequest.Marshal(b, m, deterministic)
}
func (m *GetAttributesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttributesRequest.Merge(m, src)
}
func (m *GetAttributesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAttributesRequest.Size(m)
}
func (m *GetAttributesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttributesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttributesRequest proto.InternalMessageInfo

func (m *GetAttributesRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetAttributesResponse struct {
	InfoList             map[string]*Attribute `protobuf:"bytes,1,rep,name=infoList,proto3" json:"infoList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetAttributesResponse) Reset()         { *m = GetAttributesResponse{} }
func (m *GetAttributesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAttributesResponse) ProtoMessage()    {}
func (*GetAttributesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f97a9220cd723d2e, []int{4}
}

func (m *GetAttributesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttributesResponse.Unmarshal(m, b)
}
func (m *GetAttributesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttributesResponse.Marshal(b, m, deterministic)
}
func (m *GetAttributesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttributesResponse.Merge(m, src)
}
func (m *GetAttributesResponse) XXX_Size() int {
	return xxx_messageInfo_GetAttributesResponse.Size(m)
}
func (m *GetAttributesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttributesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttributesResponse proto.InternalMessageInfo

func (m *GetAttributesResponse) GetInfoList() map[string]*Attribute {
	if m != nil {
		return m.InfoList
	}
	return nil
}

func init() {
	proto.RegisterType((*Attribute)(nil), "attrspb.Attribute")
	proto.RegisterType((*SetAttributeRequest)(nil), "attrspb.SetAttributeRequest")
	proto.RegisterType((*SetAttributeResponse)(nil), "attrspb.SetAttributeResponse")
	proto.RegisterType((*GetAttributesRequest)(nil), "attrspb.GetAttributesRequest")
	proto.RegisterType((*GetAttributesResponse)(nil), "attrspb.GetAttributesResponse")
	proto.RegisterMapType((map[string]*Attribute)(nil), "attrspb.GetAttributesResponse.InfoListEntry")
}

func init() { proto.RegisterFile("attributes.proto", fileDescriptor_f97a9220cd723d2e) }

var fileDescriptor_f97a9220cd723d2e = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x4a, 0x02, 0x41,
	0x14, 0xc6, 0x59, 0xb5, 0x3f, 0x9e, 0x4d, 0x90, 0x93, 0xc5, 0x22, 0x1a, 0xb2, 0x17, 0x21, 0x11,
	0x2b, 0xd8, 0x4d, 0x78, 0xd7, 0x45, 0x54, 0x10, 0x08, 0xfa, 0x04, 0xb3, 0x76, 0xb2, 0xa1, 0x75,
	0x66, 0xdb, 0x19, 0x05, 0x6f, 0x7b, 0x85, 0xde, 0xa4, 0x57, 0xe9, 0xbe, 0xab, 0x1e, 0x24, 0x66,
	0x76, 0xdd, 0xb4, 0xd6, 0xee, 0x66, 0xce, 0x39, 0xdf, 0xf9, 0xcd, 0xf7, 0x31, 0x50, 0x67, 0x5a,
	0x27, 0x3c, 0x9c, 0x6b, 0x52, 0x41, 0x9c, 0x48, 0x2d, 0x71, 0xcf, 0x54, 0x54, 0x1c, 0x36, 0x5b,
	0x53, 0x29, 0xa7, 0x11, 0xf5, 0x58, 0xcc, 0x7b, 0x4c, 0x08, 0xa9, 0x99, 0xe6, 0x52, 0x64, 0x63,
	0xfe, 0x0c, 0xaa, 0x57, 0x2b, 0x29, 0x22, 0x54, 0x04, 0x9b, 0x91, 0xe7, 0x74, 0x9c, 0x6e, 0x75,
	0x64, 0xcf, 0xd8, 0x80, 0x9d, 0x88, 0x85, 0x14, 0x79, 0x25, 0x5b, 0x4c, 0x2f, 0xd8, 0x01, 0xf7,
	0x81, 0xd4, 0x24, 0xe1, 0xb1, 0x59, 0xe6, 0x95, 0x6d, 0x6f, 0xbd, 0x64, 0x74, 0x0b, 0x16, 0xcd,
	0xc9, 0xab, 0xa4, 0x3a, 0x7b, 0xf1, 0x87, 0x70, 0x38, 0x26, 0x9d, 0x13, 0x47, 0xf4, 0x32, 0x27,
	0xa5, 0xb1, 0x0e, 0xe5, 0x67, 0x5a, 0x66, 0x5c, 0x73, 0xc4, 0x53, 0xa8, 0x70, 0xf1, 0x28, 0x2d,
	0xd5, 0xed, 0x63, 0x90, 0xb9, 0x09, 0x7e, 0xa4, 0xb6, 0xef, 0x1f, 0x43, 0x63, 0x73, 0xa1, 0x8a,
	0xa5, 0x50, 0xe4, 0x77, 0xa1, 0x71, 0xb3, 0x56, 0x57, 0x7f, 0x48, 0xe5, 0x8c, 0xe4, 0xbf, 0x3b,
	0x70, 0xf4, 0x6b, 0x34, 0xdd, 0x81, 0xb7, 0xb0, 0x6f, 0x18, 0xf7, 0x5c, 0x69, 0x2b, 0x70, 0xfb,
	0xe7, 0xf9, 0x3b, 0x0a, 0x15, 0xc1, 0x5d, 0x36, 0x7e, 0x2d, 0x74, 0xb2, 0x1c, 0xe5, 0xea, 0xe6,
	0x10, 0x6a, 0x1b, 0xad, 0x02, 0xc3, 0xdd, 0x55, 0x5e, 0xdb, 0x1d, 0xa7, 0x03, 0x83, 0xd2, 0xa5,
	0xd3, 0xff, 0x74, 0xc0, 0x35, 0x1b, 0xc7, 0x94, 0x2c, 0xf8, 0x84, 0x90, 0xe0, 0x60, 0x3d, 0x06,
	0x6c, 0xe5, 0xf2, 0x82, 0xb8, 0x9b, 0xed, 0x2d, 0xdd, 0x2c, 0x3b, 0xef, 0xf5, 0xe3, 0xeb, 0xad,
	0x84, 0x7e, 0xcd, 0xfe, 0x19, 0x63, 0xa2, 0xa7, 0x48, 0x0f, 0x9c, 0x33, 0x7c, 0x82, 0xda, 0x86,
	0x71, 0x6c, 0x6f, 0x0b, 0x24, 0x05, 0x9d, 0xfc, 0x9f, 0x57, 0x11, 0x69, 0x6a, 0x49, 0xe1, 0xae,
	0xfd, 0x9e, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x30, 0x31, 0x0d, 0xd9, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoServiceClient interface {
	SetAttribute(ctx context.Context, in *SetAttributeRequest, opts ...grpc.CallOption) (*SetAttributeResponse, error)
	GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error)
}

type infoServiceClient struct {
	cc *grpc.ClientConn
}

func NewInfoServiceClient(cc *grpc.ClientConn) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) SetAttribute(ctx context.Context, in *SetAttributeRequest, opts ...grpc.CallOption) (*SetAttributeResponse, error) {
	out := new(SetAttributeResponse)
	err := c.cc.Invoke(ctx, "/attrspb.InfoService/SetAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error) {
	out := new(GetAttributesResponse)
	err := c.cc.Invoke(ctx, "/attrspb.InfoService/GetAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
type InfoServiceServer interface {
	SetAttribute(context.Context, *SetAttributeRequest) (*SetAttributeResponse, error)
	GetAttributes(context.Context, *GetAttributesRequest) (*GetAttributesResponse, error)
}

// UnimplementedInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (*UnimplementedInfoServiceServer) SetAttribute(ctx context.Context, req *SetAttributeRequest) (*SetAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAttribute not implemented")
}
func (*UnimplementedInfoServiceServer) GetAttributes(ctx context.Context, req *GetAttributesRequest) (*GetAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributes not implemented")
}

func RegisterInfoServiceServer(s *grpc.Server, srv InfoServiceServer) {
	s.RegisterService(&_InfoService_serviceDesc, srv)
}

func _InfoService_SetAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).SetAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attrspb.InfoService/SetAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).SetAttribute(ctx, req.(*SetAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attrspb.InfoService/GetAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetAttributes(ctx, req.(*GetAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "attrspb.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAttribute",
			Handler:    _InfoService_SetAttribute_Handler,
		},
		{
			MethodName: "GetAttributes",
			Handler:    _InfoService_GetAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attributes.proto",
}
