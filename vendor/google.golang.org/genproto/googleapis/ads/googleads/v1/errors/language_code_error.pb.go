// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/language_code_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing language code errors.
type LanguageCodeErrorEnum_LanguageCodeError int32

const (
	// Enum unspecified.
	LanguageCodeErrorEnum_UNSPECIFIED LanguageCodeErrorEnum_LanguageCodeError = 0
	// The received error code is not known in this version.
	LanguageCodeErrorEnum_UNKNOWN LanguageCodeErrorEnum_LanguageCodeError = 1
	// The input language code is not recognized.
	LanguageCodeErrorEnum_LANGUAGE_CODE_NOT_FOUND LanguageCodeErrorEnum_LanguageCodeError = 2
	// The language is not allowed to use.
	LanguageCodeErrorEnum_INVALID_LANGUAGE_CODE LanguageCodeErrorEnum_LanguageCodeError = 3
)

var LanguageCodeErrorEnum_LanguageCodeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LANGUAGE_CODE_NOT_FOUND",
	3: "INVALID_LANGUAGE_CODE",
}
var LanguageCodeErrorEnum_LanguageCodeError_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"LANGUAGE_CODE_NOT_FOUND": 2,
	"INVALID_LANGUAGE_CODE":   3,
}

func (x LanguageCodeErrorEnum_LanguageCodeError) String() string {
	return proto.EnumName(LanguageCodeErrorEnum_LanguageCodeError_name, int32(x))
}
func (LanguageCodeErrorEnum_LanguageCodeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_language_code_error_66c9745952dd8d08, []int{0, 0}
}

// Container for enum describing language code errors.
type LanguageCodeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LanguageCodeErrorEnum) Reset()         { *m = LanguageCodeErrorEnum{} }
func (m *LanguageCodeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*LanguageCodeErrorEnum) ProtoMessage()    {}
func (*LanguageCodeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_language_code_error_66c9745952dd8d08, []int{0}
}
func (m *LanguageCodeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LanguageCodeErrorEnum.Unmarshal(m, b)
}
func (m *LanguageCodeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LanguageCodeErrorEnum.Marshal(b, m, deterministic)
}
func (dst *LanguageCodeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LanguageCodeErrorEnum.Merge(dst, src)
}
func (m *LanguageCodeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_LanguageCodeErrorEnum.Size(m)
}
func (m *LanguageCodeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LanguageCodeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LanguageCodeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LanguageCodeErrorEnum)(nil), "google.ads.googleads.v1.errors.LanguageCodeErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.LanguageCodeErrorEnum_LanguageCodeError", LanguageCodeErrorEnum_LanguageCodeError_name, LanguageCodeErrorEnum_LanguageCodeError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/language_code_error.proto", fileDescriptor_language_code_error_66c9745952dd8d08)
}

var fileDescriptor_language_code_error_66c9745952dd8d08 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x4d, 0x0a, 0x0a, 0xd3, 0x85, 0x31, 0x50, 0xc5, 0x3f, 0x74, 0x91, 0x03, 0x4c, 0x08,
	0x6e, 0x64, 0x5c, 0x4d, 0x9b, 0x34, 0x04, 0xcb, 0xa4, 0xa0, 0x89, 0x20, 0x81, 0x10, 0x3b, 0x61,
	0x08, 0xb4, 0xf3, 0x4a, 0xa6, 0xed, 0x01, 0x3c, 0x8a, 0x4b, 0x8f, 0xe2, 0x51, 0xdc, 0x79, 0x03,
	0x49, 0xa6, 0x0d, 0x48, 0xd1, 0x55, 0x3e, 0x5e, 0x7e, 0xdf, 0x9b, 0xf7, 0x7d, 0xe8, 0x4e, 0x00,
	0x88, 0x45, 0xe9, 0x16, 0x5c, 0xb9, 0x5a, 0x36, 0x6a, 0xeb, 0xb9, 0x65, 0x5d, 0x43, 0xad, 0xdc,
	0x45, 0x21, 0xc5, 0xa6, 0x10, 0x65, 0x3e, 0x07, 0x5e, 0xe6, 0xed, 0x10, 0xaf, 0x6a, 0x58, 0x83,
	0x3d, 0xd4, 0x38, 0x2e, 0xb8, 0xc2, 0x9d, 0x13, 0x6f, 0x3d, 0xac, 0x9d, 0x57, 0x37, 0xfb, 0xcd,
	0xab, 0xca, 0x2d, 0xa4, 0x84, 0x75, 0xb1, 0xae, 0x40, 0x2a, 0xed, 0x76, 0xde, 0x0c, 0x34, 0x98,
	0xee, 0x76, 0x8f, 0x81, 0x97, 0x41, 0x63, 0x0a, 0xe4, 0x66, 0xe9, 0x54, 0xe8, 0xec, 0xe0, 0x87,
	0x7d, 0x8a, 0xfa, 0x09, 0x7b, 0x9c, 0x05, 0xe3, 0x68, 0x12, 0x05, 0xbe, 0x75, 0x64, 0xf7, 0xd1,
	0x49, 0xc2, 0x1e, 0x58, 0xfc, 0xcc, 0x2c, 0xc3, 0xbe, 0x46, 0x17, 0x53, 0xca, 0xc2, 0x84, 0x86,
	0x41, 0x3e, 0x8e, 0xfd, 0x20, 0x67, 0xf1, 0x53, 0x3e, 0x89, 0x13, 0xe6, 0x5b, 0xa6, 0x7d, 0x89,
	0x06, 0x11, 0x4b, 0xe9, 0x34, 0xf2, 0xf3, 0x5f, 0x90, 0xd5, 0x1b, 0x7d, 0x1b, 0xc8, 0x99, 0xc3,
	0x12, 0xff, 0x9f, 0x64, 0x74, 0x7e, 0x70, 0xcf, 0xac, 0xc9, 0x30, 0x33, 0x5e, 0xfc, 0x9d, 0x53,
	0x40, 0xd3, 0x13, 0x86, 0x5a, 0xb8, 0xa2, 0x94, 0x6d, 0xc2, 0x7d, 0x9b, 0xab, 0x4a, 0xfd, 0x55,
	0xee, 0xbd, 0xfe, 0xbc, 0x9b, 0xbd, 0x90, 0xd2, 0x0f, 0x73, 0x18, 0xea, 0x65, 0x94, 0x2b, 0xac,
	0x65, 0xa3, 0x52, 0x0f, 0xb7, 0x4f, 0xaa, 0xcf, 0x3d, 0x90, 0x51, 0xae, 0xb2, 0x0e, 0xc8, 0x52,
	0x2f, 0xd3, 0xc0, 0x97, 0xe9, 0xe8, 0x29, 0x21, 0x94, 0x2b, 0x42, 0x3a, 0x84, 0x90, 0xd4, 0x23,
	0x44, 0x43, 0xaf, 0xc7, 0xed, 0x75, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0xef, 0x6d,
	0x3e, 0xf9, 0x01, 0x00, 0x00,
}
