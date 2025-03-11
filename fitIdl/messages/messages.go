package messages

import (
	"encoding/binary"
	"math"
)

type MessageDefinition struct {
	Reserved            byte
	ByteOrder           binary.ByteOrder
	GlobalMessageNumber uint16
	NumberOfFields      byte
	Fields              []MessageEntry
	NumberOfDevFields   byte
	DevFields           []MessageEntry
	MsgSize             uint16
}

type MessageEntry struct {
	Id       byte
	Length   byte
	BaseType BaseType
}
type BaseType byte

const (
	Enum    BaseType = 0x00
	Sint8   BaseType = 0x01 // 2’s complement format
	Uint8   BaseType = 0x02
	Sint16  BaseType = 0x83 // 2’s complement format
	Uint16  BaseType = 0x84
	Sint32  BaseType = 0x85 // 2’s complement format
	Uint32  BaseType = 0x86
	String  BaseType = 0x07 // Null terminated string encoded in UTF-8 format: 0x00
	Float32 BaseType = 0x88
	Float64 BaseType = 0x89
	Uint8z  BaseType = 0x0A
	Uint16z BaseType = 0x8B
	Uint32z BaseType = 0x8C
	Byte    BaseType = 0x0D // Array of bytes. Field is invalid if all bytes are invalid.
	Sint64  BaseType = 0x8E // 2’s complement format
	Uint64  BaseType = 0x8F
	Uint64z BaseType = 0x90
	Bool    BaseType = 0xA0
)

const (
	EnumInvalid    byte   = math.MaxUint8  // 0xFF
	Sint8Invalid   int8   = math.MaxInt8   // 0x7F
	Uint8Invalid   uint8  = math.MaxUint8  // 0xFF
	Sint16Invalid  int16  = math.MaxInt16  // 0x7FFF
	Uint16Invalid  uint16 = math.MaxUint16 // 0xFFFF
	Sint32Invalid  int32  = math.MaxInt32  // 0x7FFFFFFF
	Uint32Invalid  uint32 = math.MaxUint32 // 0xFFFFFFFF
	StringInvalid  string = ""             // We use empty string to represent an invalid string in Go. However, it will be converted automatically into an utf8 null-terminated string "\x00" by the Value Marshaler.
	Float32Invalid uint32 = math.MaxUint32 // 0xFFFFFFFF. math.Float32frombits(0xFFFFFFFF) produces float64 NaN which is uncomparable. Can only check in its integer form e.g. math.Float32bits(float32value) == Float32Invalid.
	Float64Invalid uint64 = math.MaxUint64 // 0xFFFFFFFFFFFFFFFF. math.Float64frombits(0xFFFFFFFFFFFFFFFF) produces float64 NaN which is uncomparable. Can only check in its integer form e.g. math.Float64bits(float64value) == Float64Invalid.
	Uint8zInvalid  uint8  = 0              // 0x00
	Uint16zInvalid uint16 = 0              // 0x0000
	Uint32zInvalid uint32 = 0              // 0x00000000
	ByteInvalid    byte   = math.MaxUint8  // 0xFF
	Sint64Invalid  int64  = math.MaxInt64  // 0x7FFFFFFFFFFFFFFF
	Uint64Invalid  uint64 = math.MaxUint64 // 0xFFFFFFFFFFFFFFFF
	Uint64zInvalid uint64 = 0              // 0x0000000000000000
)
