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
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x19, 0x5b, 0x2a, 0x04, 0x0a, 0x12, 0xc4, 0xea, 0xac, 0x82, 0x75, 0x55, 0x6c, 0x52,
	0xeb, 0xae, 0x2b, 0x7f, 0x41, 0x10, 0x37, 0xf6, 0x01, 0x86, 0x4c, 0x26, 0xa6, 0xd1, 0x4c, 0x12,
	0x92, 0x3b, 0x63, 0xe7, 0x11, 0xdc, 0xfa, 0x7a, 0x2e, 0x7c, 0x15, 0x31, 0x53, 0x70, 0xe1, 0xea,
	0xde, 0x7b, 0xee, 0x39, 0x07, 0x3e, 0x34, 0x8e, 0xef, 0x5c, 0x29, 0x19, 0xa8, 0x0f, 0x0e, 0x1c,
	0x1e, 0x71, 0xaf, 0x0b, 0x5f, 0xe6, 0xe7, 0xe9, 0x14, 0x73, 0x25, 0xed, 0x7c, 0xe7, 0x60, 0xce,
	0x83, 0x76, 0x36, 0x32, 0x6e, 0xad, 0x03, 0x9e, 0xf6, 0x3e, 0x95, 0x4f, 0x94, 0x73, 0xca, 0x48,
	0x16, 0xbc, 0x60, 0x11, 0x38, 0x34, 0xbb, 0xc7, 0xcd, 0x77, 0xf6, 0x79, 0xfd, 0x95, 0xe1, 0x8f,
	0x0c, 0x4d, 0x9e, 0x3a, 0xf2, 0xc0, 0xcb, 0x40, 0xee, 0xb7, 0xbc, 0xf6, 0x46, 0x92, 0xb5, 0x0c,
	0xad, 0x16, 0x12, 0x9f, 0x4d, 0xc9, 0x9d, 0x8c, 0x22, 0xe8, 0xd4, 0x8e, 0x10, 0x29, 0x40, 0x46,
	0x28, 0xc8, 0xac, 0xfa, 0x53, 0x67, 0xa7, 0xb7, 0x68, 0xfc, 0x68, 0x74, 0xed, 0xda, 0xdf, 0x9c,
	0x92, 0x1d, 0x3e, 0xd9, 0x00, 0xf8, 0xb8, 0x62, 0x4c, 0x69, 0xd8, 0x34, 0x25, 0x15, 0xae, 0x66,
	0x6f, 0x46, 0xd7, 0x8b, 0x36, 0x9f, 0xf4, 0x73, 0x1e, 0x93, 0xf3, 0xaa, 0xe3, 0xb6, 0x92, 0x5b,
	0x1a, 0x9a, 0xe5, 0xe0, 0x82, 0x2e, 0x66, 0xc3, 0x6c, 0x6f, 0x30, 0x5c, 0x1e, 0x70, 0xef, 0x8d,
	0x16, 0x89, 0x82, 0xbd, 0x46, 0x67, 0x57, 0xff, 0x94, 0xe7, 0x29, 0xda, 0xaf, 0xe4, 0x0b, 0x6f,
	0x0c, 0xe0, 0x63, 0x7c, 0x84, 0x0e, 0x73, 0x4c, 0x7b, 0x4e, 0x1a, 0xbc, 0xa0, 0xeb, 0xc4, 0x59,
	0x8e, 0x12, 0xe8, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa2, 0xd6, 0x96, 0x48, 0x01,
	0x00, 0x00,
}
