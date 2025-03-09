package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type BeatIntervals struct {
	Timestamp   uint32
	TimestampMs uint16
	Time        uint16
}

func (obj *BeatIntervals) MsgNumber() uint16 {
	return uint16(MesgNum_BeatIntervals)
}

func (obj *BeatIntervals) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForBeatIntervals = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Milliseconds past date_time", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "Time", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Array of millisecond times between beats", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_BeatIntervals),
		fieldMapForBeatIntervals,
		func() (registration.IFitMessage, error) {
			return &BeatIntervals{}, nil
		},
	)
}
