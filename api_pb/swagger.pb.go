// Code generated by protoc-gen-go. DO NOT EDIT.
// source: swagger.proto

package api_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/rpc/status"
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

func init() {
	proto.RegisterFile("swagger.proto", fileDescriptor_49635b75e059a131)
}

var fileDescriptor_49635b75e059a131 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x6a, 0x42, 0x31,
	0x10, 0x85, 0xf1, 0x07, 0x0b, 0x01, 0xa1, 0x84, 0x52, 0xdb, 0xbb, 0x0a, 0x75, 0x21, 0x84, 0x9a,
	0xa8, 0xdd, 0x89, 0x95, 0xfe, 0xd2, 0x42, 0xe9, 0xa6, 0x3e, 0x80, 0xe6, 0xc6, 0x18, 0x53, 0x73,
	0x93, 0x90, 0xe4, 0x5a, 0xdd, 0xf6, 0x11, 0xfa, 0x1c, 0x7d, 0xc6, 0x52, 0x8c, 0x42, 0x17, 0xdd,
	0x24, 0x33, 0x67, 0xce, 0x1c, 0x98, 0x0f, 0x34, 0xc3, 0x07, 0x93, 0x52, 0x78, 0xe2, 0xbc, 0x8d,
	0x16, 0x36, 0x98, 0x53, 0x53, 0x97, 0x67, 0x97, 0xa9, 0xe5, 0x5d, 0x29, 0x4c, 0xf7, 0xe0, 0xa0,
	0xd6, 0x45, 0x65, 0x4d, 0xa0, 0xcc, 0x18, 0x1b, 0x59, 0xaa, 0xf7, 0x5b, 0x59, 0x4b, 0x5a, 0x2b,
	0xb5, 0xa0, 0xde, 0x71, 0x1a, 0x22, 0x8b, 0xe5, 0x61, 0x70, 0xf7, 0x59, 0xfd, 0xba, 0xfd, 0xa9,
	0xc0, 0xef, 0x0a, 0x68, 0xbd, 0x6e, 0xd1, 0x33, 0xcb, 0x3d, 0x7a, 0xdc, 0xb0, 0xc2, 0x69, 0x81,
	0x26, 0xc2, 0xaf, 0x15, 0x17, 0xf0, 0xa9, 0x8d, 0x1e, 0x44, 0xe0, 0x5e, 0xa5, 0x74, 0x00, 0xa6,
	0x51, 0x84, 0x38, 0x45, 0x18, 0xcf, 0xff, 0x54, 0x8c, 0xc1, 0x88, 0xa1, 0xa5, 0x17, 0x8b, 0xeb,
	0x4e, 0xbb, 0x33, 0xd6, 0xca, 0xac, 0x46, 0x94, 0x8d, 0x01, 0x46, 0xb3, 0xfe, 0x6c, 0xf7, 0x0e,
	0x66, 0xe0, 0xe2, 0x1e, 0x34, 0x5f, 0xb4, 0x2a, 0xec, 0x7a, 0x17, 0x2d, 0xc5, 0x16, 0x9e, 0x2f,
	0x63, 0x74, 0x61, 0x48, 0xa9, 0x54, 0x71, 0x59, 0xe6, 0x84, 0xdb, 0x82, 0xae, 0xb4, 0x2a, 0x7a,
	0xeb, 0xac, 0xb5, 0xff, 0xbb, 0x21, 0x39, 0x6f, 0xb6, 0xcc, 0xcc, 0xc5, 0x86, 0xf8, 0x72, 0x50,
	0xeb, 0x93, 0x1e, 0xae, 0x57, 0xaa, 0xb5, 0xfa, 0xe0, 0x98, 0x39, 0xa7, 0x15, 0x4f, 0x87, 0xd2,
	0xf7, 0x60, 0xcd, 0xf0, 0x9f, 0xf2, 0xd6, 0x06, 0x47, 0x73, 0xb1, 0x60, 0xa5, 0x8e, 0xf0, 0x0c,
	0x9e, 0x82, 0x93, 0x0c, 0x92, 0x3d, 0x0a, 0xe2, 0x1d, 0x27, 0x93, 0x84, 0x22, 0x6f, 0x24, 0x16,
	0x57, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0x23, 0x13, 0xa0, 0x6b, 0x01, 0x00, 0x00,
}
