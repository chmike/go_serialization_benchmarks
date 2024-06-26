package msgpacktinylib

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *A) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Name"
	o = append(o, 0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "BirthDay"
	o = append(o, 0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	o = msgp.AppendTime(o, z.BirthDay)
	// string "Phone"
	o = append(o, 0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Phone)
	// string "Siblings"
	o = append(o, 0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendInt(o, z.Siblings)
	// string "Spouse"
	o = append(o, 0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Spouse)
	// string "Money"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Money)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "BirthDay":
			z.BirthDay, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BirthDay")
				return
			}
		case "Phone":
			z.Phone, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Phone")
				return
			}
		case "Siblings":
			z.Siblings, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Siblings")
				return
			}
		case "Spouse":
			z.Spouse, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Spouse")
				return
			}
		case "Money":
			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Money")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.TimeSize + 6 + msgp.StringPrefixSize + len(z.Phone) + 9 + msgp.IntSize + 7 + msgp.BoolSize + 6 + msgp.Float64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NoTimeA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Name"
	o = append(o, 0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "BirthDay"
	o = append(o, 0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	o = msgp.AppendInt64(o, z.BirthDay)
	// string "Phone"
	o = append(o, 0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Phone)
	// string "Siblings"
	o = append(o, 0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendInt(o, z.Siblings)
	// string "Spouse"
	o = append(o, 0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Spouse)
	// string "Money"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Money)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NoTimeA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "BirthDay":
			z.BirthDay, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BirthDay")
				return
			}
		case "Phone":
			z.Phone, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Phone")
				return
			}
		case "Siblings":
			z.Siblings, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Siblings")
				return
			}
		case "Spouse":
			z.Spouse, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Spouse")
				return
			}
		case "Money":
			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Money")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NoTimeA) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Phone) + 9 + msgp.IntSize + 7 + msgp.BoolSize + 6 + msgp.Float64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NoTimeNoStringNoFloatA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Name"
	o = append(o, 0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendBytes(o, z.Name)
	// string "BirthDay"
	o = append(o, 0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	o = msgp.AppendUint64(o, z.BirthDay)
	// string "Phone"
	o = append(o, 0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendBytes(o, z.Phone)
	// string "Siblings"
	o = append(o, 0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendUint32(o, z.Siblings)
	// string "Spouse"
	o = append(o, 0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Spouse)
	// string "Money"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendUint64(o, z.Money)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NoTimeNoStringNoFloatA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadBytesBytes(bts, z.Name)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "BirthDay":
			z.BirthDay, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BirthDay")
				return
			}
		case "Phone":
			z.Phone, bts, err = msgp.ReadBytesBytes(bts, z.Phone)
			if err != nil {
				err = msgp.WrapError(err, "Phone")
				return
			}
		case "Siblings":
			z.Siblings, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Siblings")
				return
			}
		case "Spouse":
			z.Spouse, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Spouse")
				return
			}
		case "Money":
			z.Money, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Money")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NoTimeNoStringNoFloatA) Msgsize() (s int) {
	s = 1 + 5 + msgp.BytesPrefixSize + len(z.Name) + 9 + msgp.Uint64Size + 6 + msgp.BytesPrefixSize + len(z.Phone) + 9 + msgp.Uint32Size + 7 + msgp.BoolSize + 6 + msgp.Uint64Size
	return
}
