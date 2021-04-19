// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spire/plugin/types/jwtkey.proto

package types

import (
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

type JWTKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The PKIX encoded public key.
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The key identifier.
	KeyId string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// When the key expires (seconds since Unix epoch). If zero, the key does
	// not expire.
	ExpiresAt int64 `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *JWTKey) Reset() {
	*x = JWTKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_types_jwtkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTKey) ProtoMessage() {}

func (x *JWTKey) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_types_jwtkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTKey.ProtoReflect.Descriptor instead.
func (*JWTKey) Descriptor() ([]byte, []int) {
	return file_spire_plugin_types_jwtkey_proto_rawDescGZIP(), []int{0}
}

func (x *JWTKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *JWTKey) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *JWTKey) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

var File_spire_plugin_types_jwtkey_proto protoreflect.FileDescriptor

var file_spire_plugin_types_jwtkey_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x6a, 0x77, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x06, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x15,
	0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_plugin_types_jwtkey_proto_rawDescOnce sync.Once
	file_spire_plugin_types_jwtkey_proto_rawDescData = file_spire_plugin_types_jwtkey_proto_rawDesc
)

func file_spire_plugin_types_jwtkey_proto_rawDescGZIP() []byte {
	file_spire_plugin_types_jwtkey_proto_rawDescOnce.Do(func() {
		file_spire_plugin_types_jwtkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_plugin_types_jwtkey_proto_rawDescData)
	})
	return file_spire_plugin_types_jwtkey_proto_rawDescData
}

var file_spire_plugin_types_jwtkey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spire_plugin_types_jwtkey_proto_goTypes = []interface{}{
	(*JWTKey)(nil), // 0: spire.plugin.types.JWTKey
}
var file_spire_plugin_types_jwtkey_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spire_plugin_types_jwtkey_proto_init() }
func file_spire_plugin_types_jwtkey_proto_init() {
	if File_spire_plugin_types_jwtkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_plugin_types_jwtkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTKey); i {
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
			RawDescriptor: file_spire_plugin_types_jwtkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_plugin_types_jwtkey_proto_goTypes,
		DependencyIndexes: file_spire_plugin_types_jwtkey_proto_depIdxs,
		MessageInfos:      file_spire_plugin_types_jwtkey_proto_msgTypes,
	}.Build()
	File_spire_plugin_types_jwtkey_proto = out.File
	file_spire_plugin_types_jwtkey_proto_rawDesc = nil
	file_spire_plugin_types_jwtkey_proto_goTypes = nil
	file_spire_plugin_types_jwtkey_proto_depIdxs = nil
}