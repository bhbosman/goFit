package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type TimestampCorrelation struct {
	Timestamp                 uint32
	FractionalTimestamp       uint16
	SystemTimestamp           uint32
	FractionalSystemTimestamp uint16
	LocalTimestamp            LocalDateTime
	TimestampMs               uint16
	SystemTimestampMs         uint16
}

func (obj *TimestampCorrelation) MsgNumber() uint16 {
	return uint16(MesgNum_TimestampCorrelation)
}

func (obj *TimestampCorrelation) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForTimestampCorrelation = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of UTC timestamp at the time the system timestamp was recorded.", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "FractionalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 32768, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fractional part of the UTC timestamp at the time the system timestamp was recorded.", Products: "", Example: ""},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "SystemTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the system timestamp", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "FractionalSystemTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 32768, Offset: 1, Units: "s", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fractional part of the system timestamp", Products: "", Example: ""},
	3:   {BaseType: messages.Enum, FieldId: 3, Name: "LocalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "timestamp epoch expressed in local time used to convert timestamps to local time", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the UTC timestamp at the time the system timestamp was recorded.", Products: "", Example: ""},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "SystemTimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the system timestamp", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_TimestampCorrelation),
		fieldMapForTimestampCorrelation,
		func() (registration.IFitMessage, error) {
			return &TimestampCorrelation{}, nil
		},
	)
}
