// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type AddressRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressRequest) Reset()         { *m = AddressRequest{} }
func (m *AddressRequest) String() string { return proto.CompactTextString(m) }
func (*AddressRequest) ProtoMessage()    {}
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *AddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressRequest.Unmarshal(m, b)
}
func (m *AddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressRequest.Marshal(b, m, deterministic)
}
func (m *AddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressRequest.Merge(m, src)
}
func (m *AddressRequest) XXX_Size() int {
	return xxx_messageInfo_AddressRequest.Size(m)
}
func (m *AddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressRequest proto.InternalMessageInfo

func (m *AddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressRequest) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type AddressResponse struct {
	Balance              map[string]string `protobuf:"bytes,1,rep,name=balance,proto3" json:"balance,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TransactionsCount    string            `protobuf:"bytes,2,opt,name=transactions_count,json=transactionsCount,proto3" json:"transactions_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddressResponse) Reset()         { *m = AddressResponse{} }
func (m *AddressResponse) String() string { return proto.CompactTextString(m) }
func (*AddressResponse) ProtoMessage()    {}
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *AddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressResponse.Unmarshal(m, b)
}
func (m *AddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressResponse.Marshal(b, m, deterministic)
}
func (m *AddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressResponse.Merge(m, src)
}
func (m *AddressResponse) XXX_Size() int {
	return xxx_messageInfo_AddressResponse.Size(m)
}
func (m *AddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressResponse proto.InternalMessageInfo

func (m *AddressResponse) GetBalance() map[string]string {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *AddressResponse) GetTransactionsCount() string {
	if m != nil {
		return m.TransactionsCount
	}
	return ""
}

type SubscribeRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SubscribeResponse struct {
	Query                string                     `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Data                 *_struct.Struct            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Events               []*SubscribeResponse_Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SubscribeResponse) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SubscribeResponse) GetEvents() []*SubscribeResponse_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type SubscribeResponse_Event struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Events               []string `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeResponse_Event) Reset()         { *m = SubscribeResponse_Event{} }
func (m *SubscribeResponse_Event) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse_Event) ProtoMessage()    {}
func (*SubscribeResponse_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3, 0}
}

func (m *SubscribeResponse_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse_Event.Unmarshal(m, b)
}
func (m *SubscribeResponse_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse_Event.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse_Event.Merge(m, src)
}
func (m *SubscribeResponse_Event) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse_Event.Size(m)
}
func (m *SubscribeResponse_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse_Event.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse_Event proto.InternalMessageInfo

func (m *SubscribeResponse_Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SubscribeResponse_Event) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressRequest)(nil), "api_pb.AddressRequest")
	proto.RegisterType((*AddressResponse)(nil), "api_pb.AddressResponse")
	proto.RegisterMapType((map[string]string)(nil), "api_pb.AddressResponse.BalanceEntry")
	proto.RegisterType((*SubscribeRequest)(nil), "api_pb.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "api_pb.SubscribeResponse")
	proto.RegisterType((*SubscribeResponse_Event)(nil), "api_pb.SubscribeResponse.Event")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x36, 0x91, 0x27, 0x11, 0x34, 0xa3, 0x28, 0x75, 0x0d, 0x12, 0x91, 0xc5, 0x21,
	0x12, 0xc2, 0x86, 0x70, 0x00, 0xf5, 0x80, 0xd4, 0x40, 0x10, 0x37, 0x24, 0x47, 0xea, 0xb5, 0x5a,
	0xdb, 0xdb, 0x64, 0x55, 0xb3, 0xeb, 0x7a, 0xd7, 0x81, 0x08, 0x71, 0xe1, 0x17, 0xe0, 0x73, 0x38,
	0xf1, 0x0b, 0xfc, 0x02, 0x1f, 0x82, 0xbc, 0xde, 0x8d, 0xa2, 0x26, 0xbd, 0x79, 0xf6, 0xbd, 0x79,
	0xb3, 0xf3, 0xde, 0x1a, 0x5c, 0x52, 0xb0, 0xb0, 0x28, 0x85, 0x12, 0xd8, 0x21, 0x05, 0xbb, 0x2a,
	0x12, 0xff, 0xf1, 0x52, 0x88, 0x65, 0x4e, 0x23, 0x52, 0xb0, 0x88, 0x70, 0x2e, 0x14, 0x51, 0x4c,
	0x70, 0xd9, 0xb0, 0xb6, 0xa8, 0xae, 0x92, 0xea, 0x3a, 0x92, 0xaa, 0xac, 0x52, 0x65, 0xd0, 0xb3,
	0x9d, 0xde, 0x95, 0x52, 0x45, 0x22, 0xb2, 0x8d, 0x81, 0x1e, 0xdd, 0x6d, 0xa4, 0x9f, 0x0b, 0x65,
	0xc0, 0x60, 0x06, 0x0f, 0x2e, 0xb2, 0xac, 0xa4, 0x52, 0xc6, 0xf4, 0xb6, 0xa2, 0x52, 0xa1, 0x07,
	0x5d, 0xd2, 0x9c, 0x78, 0xce, 0xd8, 0x99, 0xb8, 0xb1, 0x2d, 0x71, 0x04, 0x9d, 0x15, 0x65, 0xcb,
	0x95, 0xf2, 0x5a, 0x63, 0x67, 0x72, 0x14, 0x9b, 0x2a, 0xf8, 0xed, 0xc0, 0xc3, 0xad, 0x88, 0x2c,
	0x04, 0x97, 0x14, 0xdf, 0x42, 0x37, 0x21, 0x39, 0xe1, 0x29, 0xf5, 0x9c, 0x71, 0x7b, 0xd2, 0x9b,
	0x3e, 0x0d, 0x9b, 0x2d, 0xc3, 0x3b, 0xcc, 0x70, 0xd6, 0xd0, 0xe6, 0x5c, 0x95, 0x9b, 0xd8, 0x36,
	0xe1, 0x73, 0x40, 0x55, 0x12, 0x2e, 0x49, 0xaa, 0x3d, 0xb8, 0x4a, 0x45, 0xc5, 0x9b, 0xb9, 0x6e,
	0x3c, 0xd8, 0x45, 0xde, 0xd5, 0x80, 0x7f, 0x0e, 0xfd, 0x5d, 0x1d, 0x3c, 0x81, 0xf6, 0x0d, 0xdd,
	0x98, 0x05, 0xea, 0x4f, 0x1c, 0xc2, 0xf1, 0x9a, 0xe4, 0x15, 0x35, 0x1a, 0x4d, 0x71, 0xde, 0x7a,
	0xe3, 0x04, 0x13, 0x38, 0x59, 0x54, 0x89, 0x4c, 0x4b, 0x96, 0x50, 0x6b, 0xc2, 0x10, 0x8e, 0x6f,
	0x2b, 0x5a, 0x5a, 0x85, 0xa6, 0x08, 0xfe, 0x38, 0x30, 0xd8, 0xa1, 0x9a, 0x55, 0x0f, 0x72, 0xf1,
	0x19, 0x1c, 0x65, 0x44, 0x11, 0x3d, 0xae, 0x37, 0x3d, 0x0d, 0x9b, 0x10, 0x42, 0x1b, 0x42, 0xb8,
	0xd0, 0xe9, 0xc5, 0x9a, 0x84, 0xaf, 0xa1, 0x43, 0xd7, 0x94, 0x2b, 0xe9, 0xb5, 0xb5, 0x59, 0x4f,
	0xac, 0x59, 0x7b, 0xd3, 0xc2, 0x79, 0xcd, 0x8b, 0x0d, 0xdd, 0x7f, 0x09, 0xc7, 0xfa, 0xe0, 0xc0,
	0xc2, 0xa3, 0xad, 0x66, 0x6b, 0xdc, 0x9e, 0xb8, 0xb6, 0x65, 0xfa, 0xab, 0x05, 0x83, 0x59, 0x2e,
	0xd2, 0x9b, 0x74, 0x45, 0x18, 0x5f, 0xd0, 0x72, 0xcd, 0x52, 0x8a, 0x97, 0xd0, 0xbd, 0xb0, 0x31,
	0xef, 0x25, 0xa5, 0x3d, 0xf1, 0x4f, 0xef, 0x49, 0x30, 0xf0, 0x7f, 0xfc, 0xfd, 0xf7, 0xb3, 0x35,
	0x44, 0x8c, 0xcc, 0x4b, 0x89, 0xbe, 0x99, 0x8f, 0xef, 0xf8, 0x09, 0xfa, 0xef, 0xc5, 0x17, 0x9e,
	0x0b, 0x92, 0x7d, 0x60, 0x39, 0xc5, 0xd1, 0x9e, 0x11, 0xf3, 0xfa, 0x35, 0xfa, 0x43, 0x7b, 0x5e,
	0xff, 0x16, 0x1f, 0x95, 0x2a, 0x66, 0x22, 0xdb, 0x04, 0x03, 0xad, 0xdc, 0x43, 0x37, 0xba, 0x66,
	0x39, 0x0d, 0xd5, 0x57, 0x85, 0x97, 0xe0, 0x6e, 0x4d, 0x41, 0xef, 0x80, 0x4f, 0xcd, 0x65, 0xcf,
	0xee, 0x75, 0x30, 0x40, 0x2d, 0xda, 0x47, 0x88, 0xa4, 0xc5, 0x5e, 0x38, 0x49, 0x47, 0x5f, 0xe8,
	0xd5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x37, 0x17, 0x38, 0x98, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockchainServiceClient is the client API for BlockchainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainServiceClient interface {
	Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressResponse, error)
	DownloadFile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (BlockchainService_SubscribeClient, error)
}

type blockchainServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainServiceClient(cc *grpc.ClientConn) BlockchainServiceClient {
	return &blockchainServiceClient{cc}
}

func (c *blockchainServiceClient) Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressResponse, error) {
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, "/api_pb.BlockchainService/Address", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) DownloadFile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/api_pb.BlockchainService/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (BlockchainService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlockchainService_serviceDesc.Streams[0], "/api_pb.BlockchainService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockchainService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type blockchainServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockchainServiceServer is the server API for BlockchainService service.
type BlockchainServiceServer interface {
	Address(context.Context, *AddressRequest) (*AddressResponse, error)
	DownloadFile(context.Context, *empty.Empty) (*httpbody.HttpBody, error)
	Subscribe(*SubscribeRequest, BlockchainService_SubscribeServer) error
}

// UnimplementedBlockchainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockchainServiceServer struct {
}

func (*UnimplementedBlockchainServiceServer) Address(ctx context.Context, req *AddressRequest) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Address not implemented")
}
func (*UnimplementedBlockchainServiceServer) DownloadFile(ctx context.Context, req *empty.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (*UnimplementedBlockchainServiceServer) Subscribe(req *SubscribeRequest, srv BlockchainService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterBlockchainServiceServer(s *grpc.Server, srv BlockchainServiceServer) {
	s.RegisterService(&_BlockchainService_serviceDesc, srv)
}

func _BlockchainService_Address_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).Address(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_pb.BlockchainService/Address",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).Address(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_pb.BlockchainService/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).DownloadFile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockchainServiceServer).Subscribe(m, &blockchainServiceSubscribeServer{stream})
}

type BlockchainService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type blockchainServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlockchainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api_pb.BlockchainService",
	HandlerType: (*BlockchainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Address",
			Handler:    _BlockchainService_Address_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _BlockchainService_DownloadFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _BlockchainService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
