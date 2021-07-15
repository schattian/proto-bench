// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: structdef-go-v2.proto

package protobench

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

type TypeV2 int32

const (
	TypeV2_TYPEV2_UNSPECIFIED TypeV2 = 0
	TypeV2_TYPEV2_R           TypeV2 = 1
	TypeV2_TYPEV2_S           TypeV2 = 2
)

// Enum value maps for TypeV2.
var (
	TypeV2_name = map[int32]string{
		0: "TYPEV2_UNSPECIFIED",
		1: "TYPEV2_R",
		2: "TYPEV2_S",
	}
	TypeV2_value = map[string]int32{
		"TYPEV2_UNSPECIFIED": 0,
		"TYPEV2_R":           1,
		"TYPEV2_S":           2,
	}
)

func (x TypeV2) Enum() *TypeV2 {
	p := new(TypeV2)
	*p = x
	return p
}

func (x TypeV2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeV2) Descriptor() protoreflect.EnumDescriptor {
	return file_structdef_go_v2_proto_enumTypes[0].Descriptor()
}

func (TypeV2) Type() protoreflect.EnumType {
	return &file_structdef_go_v2_proto_enumTypes[0]
}

func (x TypeV2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeV2.Descriptor instead.
func (TypeV2) EnumDescriptor() ([]byte, []int) {
	return file_structdef_go_v2_proto_rawDescGZIP(), []int{0}
}

type GoV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BirthDay int64   `protobuf:"varint,2,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Phone    string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Siblings int32   `protobuf:"varint,4,opt,name=siblings,proto3" json:"siblings,omitempty"`
	Spouse   bool    `protobuf:"varint,5,opt,name=spouse,proto3" json:"spouse,omitempty"`
	Money    float64 `protobuf:"fixed64,6,opt,name=money,proto3" json:"money,omitempty"`
	Type     TypeV2  `protobuf:"varint,7,opt,name=type,proto3,enum=protobench.TypeV2" json:"type,omitempty"`
	// Types that are assignable to Values:
	//	*GoV2_ValueS
	//	*GoV2_ValueI
	//	*GoV2_ValueD
	Values isGoV2_Values `protobuf_oneof:"values"`
}

func (x *GoV2) Reset() {
	*x = GoV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_structdef_go_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoV2) ProtoMessage() {}

func (x *GoV2) ProtoReflect() protoreflect.Message {
	mi := &file_structdef_go_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoV2.ProtoReflect.Descriptor instead.
func (*GoV2) Descriptor() ([]byte, []int) {
	return file_structdef_go_v2_proto_rawDescGZIP(), []int{0}
}

func (x *GoV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoV2) GetBirthDay() int64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

func (x *GoV2) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GoV2) GetSiblings() int32 {
	if x != nil {
		return x.Siblings
	}
	return 0
}

func (x *GoV2) GetSpouse() bool {
	if x != nil {
		return x.Spouse
	}
	return false
}

func (x *GoV2) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *GoV2) GetType() TypeV2 {
	if x != nil {
		return x.Type
	}
	return TypeV2_TYPEV2_UNSPECIFIED
}

func (m *GoV2) GetValues() isGoV2_Values {
	if m != nil {
		return m.Values
	}
	return nil
}

func (x *GoV2) GetValueS() string {
	if x, ok := x.GetValues().(*GoV2_ValueS); ok {
		return x.ValueS
	}
	return ""
}

func (x *GoV2) GetValueI() int32 {
	if x, ok := x.GetValues().(*GoV2_ValueI); ok {
		return x.ValueI
	}
	return 0
}

func (x *GoV2) GetValueD() float64 {
	if x, ok := x.GetValues().(*GoV2_ValueD); ok {
		return x.ValueD
	}
	return 0
}

type isGoV2_Values interface {
	isGoV2_Values()
}

type GoV2_ValueS struct {
	ValueS string `protobuf:"bytes,8,opt,name=value_s,json=valueS,proto3,oneof"`
}

type GoV2_ValueI struct {
	ValueI int32 `protobuf:"varint,9,opt,name=value_i,json=valueI,proto3,oneof"`
}

type GoV2_ValueD struct {
	ValueD float64 `protobuf:"fixed64,10,opt,name=value_d,json=valueD,proto3,oneof"`
}

func (*GoV2_ValueS) isGoV2_Values() {}

func (*GoV2_ValueI) isGoV2_Values() {}

func (*GoV2_ValueD) isGoV2_Values() {}

var File_structdef_go_v2_proto protoreflect.FileDescriptor

var file_structdef_go_v2_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x64, 0x65, 0x66, 0x2d, 0x67, 0x6f, 0x2d, 0x76,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x22, 0x99, 0x02, 0x0a, 0x04, 0x47, 0x6f, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x56, 0x32, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x12,
	0x19, 0x0a, 0x07, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a,
	0x3c, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x56, 0x32, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50,
	0x45, 0x56, 0x32, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x56, 0x32, 0x5f, 0x52, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x56, 0x32, 0x5f, 0x53, 0x10, 0x02, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78,
	0x73, 0x68, 0x74, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_structdef_go_v2_proto_rawDescOnce sync.Once
	file_structdef_go_v2_proto_rawDescData = file_structdef_go_v2_proto_rawDesc
)

func file_structdef_go_v2_proto_rawDescGZIP() []byte {
	file_structdef_go_v2_proto_rawDescOnce.Do(func() {
		file_structdef_go_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_structdef_go_v2_proto_rawDescData)
	})
	return file_structdef_go_v2_proto_rawDescData
}

var file_structdef_go_v2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_structdef_go_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_structdef_go_v2_proto_goTypes = []interface{}{
	(TypeV2)(0),  // 0: protobench.TypeV2
	(*GoV2)(nil), // 1: protobench.GoV2
}
var file_structdef_go_v2_proto_depIdxs = []int32{
	0, // 0: protobench.GoV2.type:type_name -> protobench.TypeV2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_structdef_go_v2_proto_init() }
func file_structdef_go_v2_proto_init() {
	if File_structdef_go_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_structdef_go_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoV2); i {
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
	file_structdef_go_v2_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GoV2_ValueS)(nil),
		(*GoV2_ValueI)(nil),
		(*GoV2_ValueD)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_structdef_go_v2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_structdef_go_v2_proto_goTypes,
		DependencyIndexes: file_structdef_go_v2_proto_depIdxs,
		EnumInfos:         file_structdef_go_v2_proto_enumTypes,
		MessageInfos:      file_structdef_go_v2_proto_msgTypes,
	}.Build()
	File_structdef_go_v2_proto = out.File
	file_structdef_go_v2_proto_rawDesc = nil
	file_structdef_go_v2_proto_goTypes = nil
	file_structdef_go_v2_proto_depIdxs = nil
}
