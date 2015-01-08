// Code generated by protoc-gen-go.
// source: structdef.proto
// DO NOT EDIT!

/*
Package goserbench is a generated protocol buffer package.

It is generated from these files:
	structdef.proto

It has these top-level messages:
	ProtoBufA
*/
package goserbench

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ProtoBufA struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	BirthDay         *int64   `protobuf:"varint,2,req,name=birthDay" json:"birthDay,omitempty"`
	Phone            *string  `protobuf:"bytes,3,req,name=phone" json:"phone,omitempty"`
	Siblings         *int32   `protobuf:"varint,4,req,name=siblings" json:"siblings,omitempty"`
	Spouse           *bool    `protobuf:"varint,5,req,name=spouse" json:"spouse,omitempty"`
	Money            *float64 `protobuf:"fixed64,6,req,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProtoBufA) Reset()         { *m = ProtoBufA{} }
func (m *ProtoBufA) String() string { return proto.CompactTextString(m) }
func (*ProtoBufA) ProtoMessage()    {}

func (m *ProtoBufA) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProtoBufA) GetBirthDay() int64 {
	if m != nil && m.BirthDay != nil {
		return *m.BirthDay
	}
	return 0
}

func (m *ProtoBufA) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

func (m *ProtoBufA) GetSiblings() int32 {
	if m != nil && m.Siblings != nil {
		return *m.Siblings
	}
	return 0
}

func (m *ProtoBufA) GetSpouse() bool {
	if m != nil && m.Spouse != nil {
		return *m.Spouse
	}
	return false
}

func (m *ProtoBufA) GetMoney() float64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func init() {
}
