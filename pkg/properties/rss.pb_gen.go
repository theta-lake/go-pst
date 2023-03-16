package properties

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *RSS) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "27136431":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssChannel")
					return
				}
				z.PostRssChannel = nil
			} else {
				if z.PostRssChannel == nil {
					z.PostRssChannel = new(string)
				}
				*z.PostRssChannel, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "PostRssChannel")
					return
				}
			}
		case "27136031":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssChannelLink")
					return
				}
				z.PostRssChannelLink = nil
			} else {
				if z.PostRssChannelLink == nil {
					z.PostRssChannelLink = new(string)
				}
				*z.PostRssChannelLink, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "PostRssChannelLink")
					return
				}
			}
		case "27136331":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemGuid")
					return
				}
				z.PostRssItemGuid = nil
			} else {
				if z.PostRssItemGuid == nil {
					z.PostRssItemGuid = new(string)
				}
				*z.PostRssItemGuid, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemGuid")
					return
				}
			}
		case "2713623":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemHash")
					return
				}
				z.PostRssItemHash = nil
			} else {
				if z.PostRssItemHash == nil {
					z.PostRssItemHash = new(int32)
				}
				*z.PostRssItemHash, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemHash")
					return
				}
			}
		case "27136131":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemLink")
					return
				}
				z.PostRssItemLink = nil
			} else {
				if z.PostRssItemLink == nil {
					z.PostRssItemLink = new(string)
				}
				*z.PostRssItemLink, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemLink")
					return
				}
			}
		case "27136531":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemXml")
					return
				}
				z.PostRssItemXml = nil
			} else {
				if z.PostRssItemXml == nil {
					z.PostRssItemXml = new(string)
				}
				*z.PostRssItemXml, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemXml")
					return
				}
			}
		case "27136631":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "PostRssSubscription")
					return
				}
				z.PostRssSubscription = nil
			} else {
				if z.PostRssSubscription == nil {
					z.PostRssSubscription = new(string)
				}
				*z.PostRssSubscription, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "PostRssSubscription")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RSS) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(7)
	var zb0001Mask uint8 /* 7 bits */
	if z.PostRssChannel == nil {
		zb0001Len--
		zb0001Mask |= 0x1
	}
	if z.PostRssChannelLink == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if z.PostRssItemGuid == nil {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if z.PostRssItemHash == nil {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if z.PostRssItemLink == nil {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	if z.PostRssItemXml == nil {
		zb0001Len--
		zb0001Mask |= 0x20
	}
	if z.PostRssSubscription == nil {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	if (zb0001Mask & 0x1) == 0 { // if not empty
		// write "27136431"
		err = en.Append(0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x34, 0x33, 0x31)
		if err != nil {
			return
		}
		if z.PostRssChannel == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteString(*z.PostRssChannel)
			if err != nil {
				err = msgp.WrapError(err, "PostRssChannel")
				return
			}
		}
	}
	if (zb0001Mask & 0x2) == 0 { // if not empty
		// write "27136031"
		err = en.Append(0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x30, 0x33, 0x31)
		if err != nil {
			return
		}
		if z.PostRssChannelLink == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteString(*z.PostRssChannelLink)
			if err != nil {
				err = msgp.WrapError(err, "PostRssChannelLink")
				return
			}
		}
	}
	if (zb0001Mask & 0x4) == 0 { // if not empty
		// write "27136331"
		err = en.Append(0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x33, 0x33, 0x31)
		if err != nil {
			return
		}
		if z.PostRssItemGuid == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteString(*z.PostRssItemGuid)
			if err != nil {
				err = msgp.WrapError(err, "PostRssItemGuid")
				return
			}
		}
	}
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// write "2713623"
		err = en.Append(0xa7, 0x32, 0x37, 0x31, 0x33, 0x36, 0x32, 0x33)
		if err != nil {
			return
		}
		if z.PostRssItemHash == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteInt32(*z.PostRssItemHash)
			if err != nil {
				err = msgp.WrapError(err, "PostRssItemHash")
				return
			}
		}
	}
	if (zb0001Mask & 0x10) == 0 { // if not empty
		// write "27136131"
		err = en.Append(0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x31, 0x33, 0x31)
		if err != nil {
			return
		}
		if z.PostRssItemLink == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteString(*z.PostRssItemLink)
			if err != nil {
				err = msgp.WrapError(err, "PostRssItemLink")
				return
			}
		}
	}
	if (zb0001Mask & 0x20) == 0 { // if not empty
		// write "27136531"
		err = en.Append(0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x35, 0x33, 0x31)
		if err != nil {
			return
		}
		if z.PostRssItemXml == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteString(*z.PostRssItemXml)
			if err != nil {
				err = msgp.WrapError(err, "PostRssItemXml")
				return
			}
		}
	}
	if (zb0001Mask & 0x40) == 0 { // if not empty
		// write "27136631"
		err = en.Append(0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x36, 0x33, 0x31)
		if err != nil {
			return
		}
		if z.PostRssSubscription == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteString(*z.PostRssSubscription)
			if err != nil {
				err = msgp.WrapError(err, "PostRssSubscription")
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RSS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(7)
	var zb0001Mask uint8 /* 7 bits */
	if z.PostRssChannel == nil {
		zb0001Len--
		zb0001Mask |= 0x1
	}
	if z.PostRssChannelLink == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if z.PostRssItemGuid == nil {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if z.PostRssItemHash == nil {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if z.PostRssItemLink == nil {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	if z.PostRssItemXml == nil {
		zb0001Len--
		zb0001Mask |= 0x20
	}
	if z.PostRssSubscription == nil {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	if (zb0001Mask & 0x1) == 0 { // if not empty
		// string "27136431"
		o = append(o, 0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x34, 0x33, 0x31)
		if z.PostRssChannel == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendString(o, *z.PostRssChannel)
		}
	}
	if (zb0001Mask & 0x2) == 0 { // if not empty
		// string "27136031"
		o = append(o, 0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x30, 0x33, 0x31)
		if z.PostRssChannelLink == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendString(o, *z.PostRssChannelLink)
		}
	}
	if (zb0001Mask & 0x4) == 0 { // if not empty
		// string "27136331"
		o = append(o, 0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x33, 0x33, 0x31)
		if z.PostRssItemGuid == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendString(o, *z.PostRssItemGuid)
		}
	}
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// string "2713623"
		o = append(o, 0xa7, 0x32, 0x37, 0x31, 0x33, 0x36, 0x32, 0x33)
		if z.PostRssItemHash == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendInt32(o, *z.PostRssItemHash)
		}
	}
	if (zb0001Mask & 0x10) == 0 { // if not empty
		// string "27136131"
		o = append(o, 0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x31, 0x33, 0x31)
		if z.PostRssItemLink == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendString(o, *z.PostRssItemLink)
		}
	}
	if (zb0001Mask & 0x20) == 0 { // if not empty
		// string "27136531"
		o = append(o, 0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x35, 0x33, 0x31)
		if z.PostRssItemXml == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendString(o, *z.PostRssItemXml)
		}
	}
	if (zb0001Mask & 0x40) == 0 { // if not empty
		// string "27136631"
		o = append(o, 0xa8, 0x32, 0x37, 0x31, 0x33, 0x36, 0x36, 0x33, 0x31)
		if z.PostRssSubscription == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendString(o, *z.PostRssSubscription)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RSS) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "27136431":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssChannel = nil
			} else {
				if z.PostRssChannel == nil {
					z.PostRssChannel = new(string)
				}
				*z.PostRssChannel, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssChannel")
					return
				}
			}
		case "27136031":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssChannelLink = nil
			} else {
				if z.PostRssChannelLink == nil {
					z.PostRssChannelLink = new(string)
				}
				*z.PostRssChannelLink, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssChannelLink")
					return
				}
			}
		case "27136331":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssItemGuid = nil
			} else {
				if z.PostRssItemGuid == nil {
					z.PostRssItemGuid = new(string)
				}
				*z.PostRssItemGuid, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemGuid")
					return
				}
			}
		case "2713623":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssItemHash = nil
			} else {
				if z.PostRssItemHash == nil {
					z.PostRssItemHash = new(int32)
				}
				*z.PostRssItemHash, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemHash")
					return
				}
			}
		case "27136131":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssItemLink = nil
			} else {
				if z.PostRssItemLink == nil {
					z.PostRssItemLink = new(string)
				}
				*z.PostRssItemLink, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemLink")
					return
				}
			}
		case "27136531":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssItemXml = nil
			} else {
				if z.PostRssItemXml == nil {
					z.PostRssItemXml = new(string)
				}
				*z.PostRssItemXml, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssItemXml")
					return
				}
			}
		case "27136631":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PostRssSubscription = nil
			} else {
				if z.PostRssSubscription == nil {
					z.PostRssSubscription = new(string)
				}
				*z.PostRssSubscription, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PostRssSubscription")
					return
				}
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
func (z *RSS) Msgsize() (s int) {
	s = 1 + 9
	if z.PostRssChannel == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.PostRssChannel)
	}
	s += 9
	if z.PostRssChannelLink == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.PostRssChannelLink)
	}
	s += 9
	if z.PostRssItemGuid == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.PostRssItemGuid)
	}
	s += 8
	if z.PostRssItemHash == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	s += 9
	if z.PostRssItemLink == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.PostRssItemLink)
	}
	s += 9
	if z.PostRssItemXml == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.PostRssItemXml)
	}
	s += 9
	if z.PostRssSubscription == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.PostRssSubscription)
	}
	return
}