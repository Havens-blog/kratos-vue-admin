// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: sysuser_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type SysUserErrorReason int32

const (
	// 为某个枚举单独设置错误码
	SysUserErrorReason_USER_NOT_FOUND    SysUserErrorReason = 0
	SysUserErrorReason_CONTENT_MISSING   SysUserErrorReason = 1
	SysUserErrorReason_LOGIN_FAIL        SysUserErrorReason = 2
	SysUserErrorReason_CAPTCHA_INVALID   SysUserErrorReason = 3
	SysUserErrorReason_INTERNAL_ERR      SysUserErrorReason = 4
	SysUserErrorReason_CODE_NOT_MATCH    SysUserErrorReason = 5
	SysUserErrorReason_DATABASE_ERR      SysUserErrorReason = 6
	SysUserErrorReason_TENTCENT_API      SysUserErrorReason = 7
	SysUserErrorReason_BizError_API      SysUserErrorReason = 8
	SysUserErrorReason_ACCOUNT_FORBIDDEN SysUserErrorReason = 9
	SysUserErrorReason_ROLE_BIND_ACCOUNT SysUserErrorReason = 10
	SysUserErrorReason_ACCOUNT_EXISTED   SysUserErrorReason = 11
)

// Enum value maps for SysUserErrorReason.
var (
	SysUserErrorReason_name = map[int32]string{
		0:  "USER_NOT_FOUND",
		1:  "CONTENT_MISSING",
		2:  "LOGIN_FAIL",
		3:  "CAPTCHA_INVALID",
		4:  "INTERNAL_ERR",
		5:  "CODE_NOT_MATCH",
		6:  "DATABASE_ERR",
		7:  "TENTCENT_API",
		8:  "BizError_API",
		9:  "ACCOUNT_FORBIDDEN",
		10: "ROLE_BIND_ACCOUNT",
		11: "ACCOUNT_EXISTED",
	}
	SysUserErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":    0,
		"CONTENT_MISSING":   1,
		"LOGIN_FAIL":        2,
		"CAPTCHA_INVALID":   3,
		"INTERNAL_ERR":      4,
		"CODE_NOT_MATCH":    5,
		"DATABASE_ERR":      6,
		"TENTCENT_API":      7,
		"BizError_API":      8,
		"ACCOUNT_FORBIDDEN": 9,
		"ROLE_BIND_ACCOUNT": 10,
		"ACCOUNT_EXISTED":   11,
	}
)

func (x SysUserErrorReason) Enum() *SysUserErrorReason {
	p := new(SysUserErrorReason)
	*p = x
	return p
}

func (x SysUserErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SysUserErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_sysuser_error_proto_enumTypes[0].Descriptor()
}

func (SysUserErrorReason) Type() protoreflect.EnumType {
	return &file_sysuser_error_proto_enumTypes[0]
}

func (x SysUserErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SysUserErrorReason.Descriptor instead.
func (SysUserErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_sysuser_error_proto_rawDescGZIP(), []int{0}
}

var File_sysuser_error_proto protoreflect.FileDescriptor

var file_sysuser_error_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x73, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xcf, 0x02, 0x0a, 0x12, 0x53, 0x79, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x14, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x41,
	0x50, 0x54, 0x43, 0x48, 0x41, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x18, 0x0a,
	0x0e, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10,
	0x05, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x44, 0x41, 0x54, 0x41, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12,
	0x16, 0x0a, 0x0c, 0x54, 0x45, 0x4e, 0x54, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x50, 0x49, 0x10,
	0x07, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x16, 0x0a, 0x0c, 0x42, 0x69, 0x7a, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x41, 0x50, 0x49, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0xcc, 0x01, 0x12,
	0x1b, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49,
	0x44, 0x44, 0x45, 0x4e, 0x10, 0x09, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x1b, 0x0a, 0x11,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x10, 0x0a, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x19, 0x0a, 0x0f, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x1a, 0x04,
	0xa8, 0x45, 0xc8, 0x01, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x19, 0x5a, 0x17, 0x66, 0x65,
	0x6e, 0x67, 0x79, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sysuser_error_proto_rawDescOnce sync.Once
	file_sysuser_error_proto_rawDescData = file_sysuser_error_proto_rawDesc
)

func file_sysuser_error_proto_rawDescGZIP() []byte {
	file_sysuser_error_proto_rawDescOnce.Do(func() {
		file_sysuser_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_sysuser_error_proto_rawDescData)
	})
	return file_sysuser_error_proto_rawDescData
}

var file_sysuser_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sysuser_error_proto_goTypes = []interface{}{
	(SysUserErrorReason)(0), // 0: api.admin.v1.SysUserErrorReason
}
var file_sysuser_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sysuser_error_proto_init() }
func file_sysuser_error_proto_init() {
	if File_sysuser_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sysuser_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sysuser_error_proto_goTypes,
		DependencyIndexes: file_sysuser_error_proto_depIdxs,
		EnumInfos:         file_sysuser_error_proto_enumTypes,
	}.Build()
	File_sysuser_error_proto = out.File
	file_sysuser_error_proto_rawDesc = nil
	file_sysuser_error_proto_goTypes = nil
	file_sysuser_error_proto_depIdxs = nil
}
