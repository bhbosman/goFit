package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ObdiiData struct {
	Timestamp        uint32
	TimestampMs      uint16
	TimeOffset       uint16
	Pid              byte
	RawData          byte
	PidDataSize      uint8
	SystemTime       uint32
	StartTimestamp   uint32
	StartTimestampMs uint16
}

func (obj *ObdiiData) MsgNumber() uint16 {
	return uint16(MesgNum_ObdiiData)
}

func (obj *ObdiiData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForObdiiData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timestamp message was output", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fractional part of timestamp, added to timestamp", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "TimeOffset", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Offset of PID reading [i] from start_timestamp+start_timestamp_ms. Readings may span accross seconds.", Products: "", Example: ""},
	2:   {BaseType: messages.Byte, FieldId: 2, Name: "Pid", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Parameter ID", Products: "", Example: ""},
	3:   {BaseType: messages.Byte, FieldId: 3, Name: "RawData", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Raw parameter data", Products: "", Example: ""},
	4:   {BaseType: messages.Uint8, FieldId: 4, Name: "PidDataSize", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Optional, data size of PID[i]. If not specified refer to SAE J1979.", Products: "", Example: ""},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "SystemTime", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "System time associated with sample expressed in ms, can be used instead of time_offset. There will be a system_time value for each raw_data element. For multibyte pids the system_time is repeated.", Products: "", Example: ""},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "StartTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timestamp of first sample recorded in the message. Used with time_offset to generate time of each sample", Products: "", Example: ""},
	7:   {BaseType: messages.Uint16, FieldId: 7, Name: "StartTimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fractional part of start_timestamp", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ObdiiData),
		fieldMapForObdiiData,
		func() (registration.IFitMessage, error) {
			return &ObdiiData{}, nil
		},
	)
}
