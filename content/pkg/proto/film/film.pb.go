// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: film.proto

package film

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ContentID) Reset() {
	*x = ContentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentID) ProtoMessage() {}

func (x *ContentID) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentID.ProtoReflect.Descriptor instead.
func (*ContentID) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{0}
}

func (x *ContentID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{2}
}

func (x *Role) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PersonRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=Person,proto3" json:"Person,omitempty"`
	Role   *Role   `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *PersonRole) Reset() {
	*x = PersonRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRole) ProtoMessage() {}

func (x *PersonRole) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRole.ProtoReflect.Descriptor instead.
func (*PersonRole) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{3}
}

func (x *PersonRole) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *PersonRole) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{4}
}

func (x *Genre) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Selection) Reset() {
	*x = Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selection) ProtoMessage() {}

func (x *Selection) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selection.ProtoReflect.Descriptor instead.
func (*Selection) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{5}
}

func (x *Selection) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Selection) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{6}
}

func (x *Country) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title        string        `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string        `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Rating       float64       `protobuf:"fixed64,4,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Year         int32         `protobuf:"varint,5,opt,name=Year,proto3" json:"Year,omitempty"`
	IsFree       bool          `protobuf:"varint,6,opt,name=IsFree,proto3" json:"IsFree,omitempty"`
	AgeLimit     int32         `protobuf:"varint,7,opt,name=AgeLimit,proto3" json:"AgeLimit,omitempty"`
	TrailerURL   string        `protobuf:"bytes,8,opt,name=TrailerURL,proto3" json:"TrailerURL,omitempty"`
	PreviewURL   string        `protobuf:"bytes,9,opt,name=PreviewURL,proto3" json:"PreviewURL,omitempty"`
	Type         string        `protobuf:"bytes,10,opt,name=Type,proto3" json:"Type,omitempty"`
	PersonsRoles []*PersonRole `protobuf:"bytes,11,rep,name=PersonsRoles,proto3" json:"PersonsRoles,omitempty"`
	Genres       []*Genre      `protobuf:"bytes,12,rep,name=Genres,proto3" json:"Genres,omitempty"`
	Selections   []*Selection  `protobuf:"bytes,13,rep,name=Selections,proto3" json:"Selections,omitempty"`
	Countries    []*Country    `protobuf:"bytes,14,rep,name=Countries,proto3" json:"Countries,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{7}
}

func (x *Content) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Content) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Content) GetIsFree() bool {
	if x != nil {
		return x.IsFree
	}
	return false
}

func (x *Content) GetAgeLimit() int32 {
	if x != nil {
		return x.AgeLimit
	}
	return 0
}

func (x *Content) GetTrailerURL() string {
	if x != nil {
		return x.TrailerURL
	}
	return ""
}

func (x *Content) GetPreviewURL() string {
	if x != nil {
		return x.PreviewURL
	}
	return ""
}

func (x *Content) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Content) GetPersonsRoles() []*PersonRole {
	if x != nil {
		return x.PersonsRoles
	}
	return nil
}

func (x *Content) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Content) GetSelections() []*Selection {
	if x != nil {
		return x.Selections
	}
	return nil
}

func (x *Content) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ContentURL string   `protobuf:"bytes,2,opt,name=ContentURL,proto3" json:"ContentURL,omitempty"`
	Content    *Content `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{8}
}

func (x *Film) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Film) GetContentURL() string {
	if x != nil {
		return x.ContentURL
	}
	return ""
}

func (x *Film) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_film_proto protoreflect.FileDescriptor

var file_film_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x6d, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x2c, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x69, 0x6c, 0x6d,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22,
	0x2b, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x09,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x2d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbe,
	0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65,
	0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x73, 0x46, 0x72, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x49, 0x73, 0x46, 0x72, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x41, 0x67, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x55,
	0x52, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52, 0x4c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55,
	0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0c,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x06,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x5f, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x27, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x32, 0x3e, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x66, 0x69, 0x6c, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_film_proto_rawDescOnce sync.Once
	file_film_proto_rawDescData = file_film_proto_rawDesc
)

func file_film_proto_rawDescGZIP() []byte {
	file_film_proto_rawDescOnce.Do(func() {
		file_film_proto_rawDescData = protoimpl.X.CompressGZIP(file_film_proto_rawDescData)
	})
	return file_film_proto_rawDescData
}

var file_film_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_film_proto_goTypes = []interface{}{
	(*ContentID)(nil),  // 0: film.ContentID
	(*Person)(nil),     // 1: film.Person
	(*Role)(nil),       // 2: film.Role
	(*PersonRole)(nil), // 3: film.PersonRole
	(*Genre)(nil),      // 4: film.Genre
	(*Selection)(nil),  // 5: film.Selection
	(*Country)(nil),    // 6: film.Country
	(*Content)(nil),    // 7: film.Content
	(*Film)(nil),       // 8: film.Film
}
var file_film_proto_depIdxs = []int32{
	1, // 0: film.PersonRole.Person:type_name -> film.Person
	2, // 1: film.PersonRole.Role:type_name -> film.Role
	3, // 2: film.Content.PersonsRoles:type_name -> film.PersonRole
	4, // 3: film.Content.Genres:type_name -> film.Genre
	5, // 4: film.Content.Selections:type_name -> film.Selection
	6, // 5: film.Content.Countries:type_name -> film.Country
	7, // 6: film.Film.Content:type_name -> film.Content
	0, // 7: film.FilmService.GetByContentID:input_type -> film.ContentID
	8, // 8: film.FilmService.GetByContentID:output_type -> film.Film
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_film_proto_init() }
func file_film_proto_init() {
	if File_film_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_film_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentID); i {
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
		file_film_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_film_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_film_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRole); i {
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
		file_film_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_film_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selection); i {
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
		file_film_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_film_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_film_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
			RawDescriptor: file_film_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_film_proto_goTypes,
		DependencyIndexes: file_film_proto_depIdxs,
		MessageInfos:      file_film_proto_msgTypes,
	}.Build()
	File_film_proto = out.File
	file_film_proto_rawDesc = nil
	file_film_proto_goTypes = nil
	file_film_proto_depIdxs = nil
}
