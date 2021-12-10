// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: def.proto

package def

import (
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_TYPE_UNSPECIFIED Type = 0
	Type_TYPE_R           Type = 1
	Type_TYPE_S           Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_R",
		2: "TYPE_S",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_R":           1,
		"TYPE_S":           2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_def_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_def_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BirthDay int64   `protobuf:"varint,2,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Phone    string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Siblings int32   `protobuf:"varint,4,opt,name=siblings,proto3" json:"siblings,omitempty"`
	Spouse   bool    `protobuf:"varint,5,opt,name=spouse,proto3" json:"spouse,omitempty"`
	Money    float64 `protobuf:"fixed64,6,opt,name=money,proto3" json:"money,omitempty"`
	Type     Type    `protobuf:"varint,7,opt,name=type,proto3,enum=protobench.Type" json:"type,omitempty"`
	// Types that are assignable to Values:
	//	*Message_ValueS
	//	*Message_ValueI
	//	*Message_ValueD
	Values isMessage_Values       `protobuf_oneof:"values"`
	Book   *Book                  `protobuf:"bytes,11,opt,name=book,proto3" json:"book,omitempty"`
	Ts     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Message) GetBirthDay() int64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

func (x *Message) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Message) GetSiblings() int32 {
	if x != nil {
		return x.Siblings
	}
	return 0
}

func (x *Message) GetSpouse() bool {
	if x != nil {
		return x.Spouse
	}
	return false
}

func (x *Message) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *Message) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (m *Message) GetValues() isMessage_Values {
	if m != nil {
		return m.Values
	}
	return nil
}

func (x *Message) GetValueS() string {
	if x, ok := x.GetValues().(*Message_ValueS); ok {
		return x.ValueS
	}
	return ""
}

func (x *Message) GetValueI() int32 {
	if x, ok := x.GetValues().(*Message_ValueI); ok {
		return x.ValueI
	}
	return 0
}

func (x *Message) GetValueD() float64 {
	if x, ok := x.GetValues().(*Message_ValueD); ok {
		return x.ValueD
	}
	return 0
}

func (x *Message) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *Message) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type isMessage_Values interface {
	isMessage_Values()
}

type Message_ValueS struct {
	ValueS string `protobuf:"bytes,8,opt,name=value_s,json=valueS,proto3,oneof"`
}

type Message_ValueI struct {
	ValueI int32 `protobuf:"varint,9,opt,name=value_i,json=valueI,proto3,oneof"`
}

type Message_ValueD struct {
	ValueD float64 `protobuf:"fixed64,10,opt,name=value_d,json=valueD,proto3,oneof"`
}

func (*Message_ValueS) isMessage_Values() {}

func (*Message_ValueI) isMessage_Values() {}

func (*Message_ValueD) isMessage_Values() {}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RandomReader string   `protobuf:"bytes,1,opt,name=random_reader,json=randomReader,proto3" json:"random_reader,omitempty"`
	Readers      []string `protobuf:"bytes,2,rep,name=readers,proto3" json:"readers,omitempty"`
	Author       *Author  `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetRandomReader() string {
	if x != nil {
		return x.RandomReader
	}
	return ""
}

func (x *Book) GetReaders() []string {
	if x != nil {
		return x.Readers
	}
	return nil
}

func (x *Book) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age       int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	RandomNum int64  `protobuf:"varint,4,opt,name=random_num,json=randomNum,proto3" json:"random_num,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{2}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Author) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Author) GetRandomNum() int64 {
	if x != nil {
		return x.RandomNum
	}
	return 0
}

var File_def_proto protoreflect.FileDescriptor

var file_def_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x12, 0x19, 0x0a, 0x07,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x34, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x90, 0xdf, 0x1f, 0x01, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x02, 0x74, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x71, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x22, 0x6a, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x2a, 0x34, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x10, 0x02, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x63, 0x68, 0x61, 0x74, 0x74, 0x69, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x64, 0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_def_proto_rawDescOnce sync.Once
	file_def_proto_rawDescData = file_def_proto_rawDesc
)

func file_def_proto_rawDescGZIP() []byte {
	file_def_proto_rawDescOnce.Do(func() {
		file_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_def_proto_rawDescData)
	})
	return file_def_proto_rawDescData
}

var file_def_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_def_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_def_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: protobench.Type
	(*Message)(nil),               // 1: protobench.Message
	(*Book)(nil),                  // 2: protobench.Book
	(*Author)(nil),                // 3: protobench.Author
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_def_proto_depIdxs = []int32{
	0, // 0: protobench.Message.type:type_name -> protobench.Type
	2, // 1: protobench.Message.book:type_name -> protobench.Book
	4, // 2: protobench.Message.ts:type_name -> google.protobuf.Timestamp
	3, // 3: protobench.Book.Author:type_name -> protobench.Author
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_def_proto_init() }
func file_def_proto_init() {
	if File_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
	file_def_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_ValueS)(nil),
		(*Message_ValueI)(nil),
		(*Message_ValueD)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_def_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_def_proto_goTypes,
		DependencyIndexes: file_def_proto_depIdxs,
		EnumInfos:         file_def_proto_enumTypes,
		MessageInfos:      file_def_proto_msgTypes,
	}.Build()
	File_def_proto = out.File
	file_def_proto_rawDesc = nil
	file_def_proto_goTypes = nil
	file_def_proto_depIdxs = nil
}
