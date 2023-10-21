// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: proto/api/rbac/permission.proto

package rbac

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

type Action int32

const (
	Action_UNKNOWN Action = 0
	Action_CONNECT Action = 1
	Action_READ    Action = 2
	Action_EDIT    Action = 3
	Action_ADMIN   Action = 4
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "UNKNOWN",
		1: "CONNECT",
		2: "READ",
		3: "EDIT",
		4: "ADMIN",
	}
	Action_value = map[string]int32{
		"UNKNOWN": 0,
		"CONNECT": 1,
		"READ":    2,
		"EDIT":    3,
		"ADMIN":   4,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_api_rbac_permission_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_proto_api_rbac_permission_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_api_rbac_permission_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Database *string  `protobuf:"bytes,2,opt,name=database,proto3,oneof" json:"database,omitempty"`
	Schemas  []string `protobuf:"bytes,3,rep,name=schemas,proto3" json:"schemas,omitempty"`
	Tables   []string `protobuf:"bytes,4,rep,name=tables,proto3" json:"tables,omitempty"`
	Action   *Action  `protobuf:"varint,5,opt,name=action,proto3,enum=rbac.Action,oneof" json:"action,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_rbac_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_rbac_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_proto_api_rbac_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetDatabase() string {
	if x != nil && x.Database != nil {
		return *x.Database
	}
	return ""
}

func (x *Permission) GetSchemas() []string {
	if x != nil {
		return x.Schemas
	}
	return nil
}

func (x *Permission) GetTables() []string {
	if x != nil {
		return x.Tables
	}
	return nil
}

func (x *Permission) GetAction() Action {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return Action_UNKNOWN
}

var File_proto_api_rbac_permission_proto protoreflect.FileDescriptor

var file_proto_api_rbac_permission_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x62, 0x61, 0x63,
	0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x72, 0x62, 0x61, 0x63, 0x22, 0xb6, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x29, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x72, 0x62, 0x61, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2a, 0x41, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x45, 0x44, 0x49, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x04, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x75, 0x74, 0x6e, 0x61, 0x6d, 0x31, 0x32, 0x30, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_rbac_permission_proto_rawDescOnce sync.Once
	file_proto_api_rbac_permission_proto_rawDescData = file_proto_api_rbac_permission_proto_rawDesc
)

func file_proto_api_rbac_permission_proto_rawDescGZIP() []byte {
	file_proto_api_rbac_permission_proto_rawDescOnce.Do(func() {
		file_proto_api_rbac_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_rbac_permission_proto_rawDescData)
	})
	return file_proto_api_rbac_permission_proto_rawDescData
}

var file_proto_api_rbac_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_api_rbac_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_api_rbac_permission_proto_goTypes = []interface{}{
	(Action)(0),        // 0: rbac.Action
	(*Permission)(nil), // 1: rbac.Permission
}
var file_proto_api_rbac_permission_proto_depIdxs = []int32{
	0, // 0: rbac.Permission.action:type_name -> rbac.Action
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_api_rbac_permission_proto_init() }
func file_proto_api_rbac_permission_proto_init() {
	if File_proto_api_rbac_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_rbac_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
	file_proto_api_rbac_permission_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_rbac_permission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_api_rbac_permission_proto_goTypes,
		DependencyIndexes: file_proto_api_rbac_permission_proto_depIdxs,
		EnumInfos:         file_proto_api_rbac_permission_proto_enumTypes,
		MessageInfos:      file_proto_api_rbac_permission_proto_msgTypes,
	}.Build()
	File_proto_api_rbac_permission_proto = out.File
	file_proto_api_rbac_permission_proto_rawDesc = nil
	file_proto_api_rbac_permission_proto_goTypes = nil
	file_proto_api_rbac_permission_proto_depIdxs = nil
}