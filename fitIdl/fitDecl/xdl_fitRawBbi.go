package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type RawBbi struct {
	Timestamp   uint32
	TimestampMs uint16
	Data        uint16
	Time        uint16
	Quality     uint8
	Gap         uint8
}

func (obj *RawBbi) MsgNumber() uint16 {
	return uint16(MesgNum_RawBbi)
}

func (obj *RawBbi) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForRawBbi = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond resolution of the timestamp", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "Data", IsArray: true, ArrayLength: 0, Components: "time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap,time,quality,gap", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "1 bit for gap indicator, 1 bit for quality indicator, and 14 bits for Beat-to-Beat interval values in whole-integer millisecond resolution", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "Time", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Array of millisecond times between beats", Products: "", Example: ""},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "Quality", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "1 = high confidence. 0 = low confidence. N/A when gap = 1", Products: "", Example: ""},
	4:   {BaseType: messages.Uint8, FieldId: 4, Name: "Gap", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "1 = gap (time represents ms gap length). 0 = BBI data", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_RawBbi),
		fieldMapForRawBbi,
		func() (registration.IFitMessage, error) {
			return &RawBbi{}, nil
		},
	)
}
