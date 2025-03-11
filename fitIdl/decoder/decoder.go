package decoder

import (
	"encoding/binary"
	"errors"
	"github.com/bhbosman/goFit/fitIdl/messages"
	"github.com/bhbosman/goFit/fitIdl/registration"
	"io"
	"time"
)

type IFitRecord interface {
	GlobalMessageNumber() uint16
}

type ValueKind byte

const (
	VkEnum ValueKind = iota
	VkSignedInt08
	VkUnsignedInt08
	VkSignedInt16
	VkUnsignedInt16
	VkSignedInt32
	VkUnsignedInt32
	VkString
	vkFloat32
	vkFloat64
	VkUnsignedInt08z
	VkUnsignedInt16z
	VkUnsignedInt32z
	vkByteArray
	VkSignedInt64
	VkUnsignedInt64
	VkUnsignedInt64z
)

type FitValue struct {
}

type UnknownFitRecord struct {
	globalMessageNumber uint16
	m                   map[uint16]FitValue
}

func (self *UnknownFitRecord) GlobalMessageNumber() uint16 {
	return self.globalMessageNumber
}

type decoder struct {
	m     map[byte]messages.MessageDefinition
	r     io.Reader
	bytes struct {
		buf  [4096]byte
		i, j int
		pos  int
	}
	tmp             [2048]byte
	headerSize      byte
	protocolVersion byte
	profileVersion  uint16
	dataSize        uint32
	dataTypeByte    [4]byte
	crc             uint16
}

func (d *decoder) readFull(p []byte) error {
	// Unread the overshot bytes, if any.

	for {
		n := copy(p, d.bytes.buf[d.bytes.i:d.bytes.j])
		p = p[n:]
		d.bytes.i += n
		d.bytes.pos += n
		if len(p) == 0 {
			break
		}
		if err := d.fill(); err != nil {
			return err
		}
	}
	return nil
}

func (d *decoder) fill() error {
	if d.bytes.i != d.bytes.j {
		panic("fit: fill called when unread bytes exist")
	}
	d.bytes.i = 0
	d.bytes.j = 0
	n, err := d.r.Read(d.bytes.buf[d.bytes.j:])
	d.bytes.j += n
	if n > 0 {
		return nil
	}
	if err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return err
}

type ReadDataFunc func(md messages.MessageDefinition, dh DataHeader, data [2048]byte) error

func (d *decoder) decode(r io.Reader, cb ReadDataFunc) error {
	d.r = r
	d.m = map[byte]messages.MessageDefinition{}
	if err := d.readHeader(); err != nil {
		return err
	}
	if err := d.readRecords(cb); err != nil {
		return err
	}
	return nil
}

func (d *decoder) readRecords(cb ReadDataFunc) error {
	offset := d.bytes.pos

	for (d.bytes.pos - offset) < int(d.dataSize) {
		dh, err := d.readDataHeader()
		if err != nil {
			return nil
		}
		switch dh.normal {
		case true:
			switch dh.normalData.messageType {
			case dataMessage:
				md, ok := d.m[dh.normalData.LocalMessageType]
				if !ok {
					return errors.New("unknown message type")
				}
				if err = d.readFull(d.tmp[:md.MsgSize]); err != nil {
					return err
				}
				//

				for _, field := range md.Fields {
					fieldData, ok := registration.GetField(md.GlobalMessageNumber, uint16(field.Id))
					if !ok {
						continue
					}

					println(md.GlobalMessageNumber, fieldData.FieldId)

				}

				if cb != nil {
					err = cb(md, dh, d.tmp)
					if err != nil {
						return err
					}
				}
				break
			case definitionMessage:
				md, err := d.readDefinitionMessage(dh)
				if err != nil {
					return nil
				}
				d.m[dh.normalData.LocalMessageType] = md
				break
			}
			break
		case false:
			break
		}
	}
	return nil
}

func (d *decoder) readHeader() error {
	if err := d.readFull(d.tmp[:1]); err != nil {
		return err
	}
	d.headerSize = d.tmp[0]
	if err := d.readFull(d.tmp[:d.headerSize-1]); err != nil {
		return err
	}
	d.protocolVersion = d.tmp[0]
	d.profileVersion = binary.LittleEndian.Uint16(d.tmp[1:3])
	d.dataSize = binary.LittleEndian.Uint32(d.tmp[3:7])
	copy(d.dataTypeByte[:], d.tmp[7:11])
	if d.headerSize >= 14 {
		d.crc = binary.LittleEndian.Uint16(d.tmp[11:13])
	}
	return nil
}

type dataHeaderMessageType byte

const (
	dataMessage       dataHeaderMessageType = 0
	definitionMessage                       = 64
)

type dataHeaderNormalData struct {
	messageType         dataHeaderMessageType
	MessageTypeSpecific bool
	LocalMessageType    byte
}
type DataHeader struct {
	normal     bool
	normalData dataHeaderNormalData
}

func (d *decoder) readDataHeader() (DataHeader, error) {
	if err := d.readFull(d.tmp[:1]); err != nil {
		return DataHeader{}, err
	}
	v := d.tmp[0]
	result := DataHeader{
		v&0x80 == 0,
		dataHeaderNormalData{
			dataHeaderMessageType(v & 0x40),
			v&0x20 != 0,
			v & 0x0f,
		},
	}
	return result, nil
}

func (d *decoder) readDefinitionMessage(dh DataHeader) (messages.MessageDefinition, error) {
	if err := d.readFull(d.tmp[:5]); err != nil {
		return messages.MessageDefinition{}, err
	}
	arch := d.tmp[1]
	var byteOrder binary.ByteOrder = binary.LittleEndian
	if arch == 1 {
		byteOrder = binary.BigEndian
	}

	md := messages.MessageDefinition{}
	md.Reserved = d.tmp[0]
	md.ByteOrder = byteOrder
	md.GlobalMessageNumber = byteOrder.Uint16(d.tmp[2:4])
	md.NumberOfFields = d.tmp[4]
	md.MsgSize = 0
	md.Fields = make([]messages.MessageEntry, md.NumberOfFields)
	for i := 0; i < int(md.NumberOfFields); i++ {
		if err := d.readFull(d.tmp[:3]); err != nil {
			return messages.MessageDefinition{}, err
		}
		md.Fields[i].Id = d.tmp[0]
		md.Fields[i].Length = d.tmp[1]
		md.Fields[i].BaseType = messages.BaseType(d.tmp[2])
		md.MsgSize += uint16(d.tmp[1])
	}
	if dh.normalData.MessageTypeSpecific {
		if err := d.readFull(d.tmp[:1]); err != nil {
			return messages.MessageDefinition{}, err
		}
		md.NumberOfDevFields = d.tmp[0]
		md.DevFields = make([]messages.MessageEntry, md.NumberOfDevFields)
		for i := 0; i < int(md.NumberOfDevFields); i++ {
			if err := d.readFull(d.tmp[:3]); err != nil {
				return messages.MessageDefinition{}, err
			}
			md.DevFields[i].Id = d.tmp[0]
			md.DevFields[i].Length = d.tmp[1]
			md.DevFields[i].BaseType = messages.BaseType(d.tmp[2])
			md.MsgSize += uint16(d.tmp[1])
		}
	}
	return md, nil
}

type Some[TData interface{}] struct {
	Data     TData
	assigned bool
}

type FileIdInformation struct {
	fileType      Some[byte]
	manufacturer  Some[byte]
	serial_number Some[uint32]
	time_created  Some[time.Time]
	number        Some[uint16]
	product_name  Some[string]
}

func (d *decoder) readBaseType01(index int) (Some[int8], error) {
	v := d.tmp[index]
	return Some[int8]{int8(v), int8(v) != 0x7F}, nil
}

func (d *decoder) readBaseType0A(index int) (Some[uint8], error) {
	v := d.tmp[index]
	return Some[uint8]{uint8(v), int8(v) != 0x00}, nil
}

func (d *decoder) readBaseType02(index int) (Some[int16], error) {
	v := d.tmp[index]
	return Some[int16]{int16(v), int16(v) != 0x7FFF}, nil
}

func (d *decoder) readBaseTypeNullTerminatedString(index int, length byte) (Some[string], error) {

	for idx := 0; idx < int(length); idx++ {
		if d.tmp[index+idx] == 0x00 {
			s := string(d.tmp[index : index+idx])
			return Some[string]{s, true}, nil

		}
	}

	return Some[string]{"", false}, nil
}

func (d *decoder) readBaseType8B(byteOrder binary.ByteOrder, index int) (Some[uint16], error) {
	v := byteOrder.Uint16(d.tmp[index:])
	return Some[uint16]{v, v != 0}, nil
}

func (d *decoder) readBaseType8C(byteOrder binary.ByteOrder, index int) (Some[uint32], error) {
	v := byteOrder.Uint32(d.tmp[index:])
	return Some[uint32]{v, v != 0}, nil
}

func (d *decoder) readBaseType83(byteOrder binary.ByteOrder, index int) (Some[int16], error) {
	v := byteOrder.Uint16(d.tmp[index:])
	return Some[int16]{int16(v), v != 0x7FFF}, nil
}

func (d *decoder) readBaseType84(byteOrder binary.ByteOrder, index int) (Some[uint16], error) {
	v := byteOrder.Uint16(d.tmp[index:])
	return Some[uint16]{v, v != 0xFFFF}, nil
}

func (d *decoder) readBaseType85(byteOrder binary.ByteOrder, index int) (Some[int32], error) {
	v := byteOrder.Uint32(d.tmp[index:])
	return Some[int32]{int32(v), v != 0x7FFFFFFF}, nil
}

func (d *decoder) readBaseType86(byteOrder binary.ByteOrder, index int) (Some[uint32], error) {
	v := byteOrder.Uint32(d.tmp[index:])
	return Some[uint32]{v, v != 0xFFFFFFFF}, nil
}

func (d *decoder) readBaseType0D(index int, length byte) (Some[[]byte], error) {
	arr := d.tmp[index : index+int(length)]
	return Some[[]byte]{arr, true}, nil
}

func Decode(r io.Reader, cb ReadDataFunc) error {
	var d decoder
	return d.decode(r, cb)
}
