package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Hr struct {
	Timestamp           uint32
	FractionalTimestamp uint16
	Time256             uint8
	FilteredBpm         uint8
	EventTimestamp      uint32
	EventTimestamp12    byte
}

func (obj *Hr) MsgNumber() uint16 {
	return uint16(MesgNum_Hr)
}

func (obj *Hr) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHr = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "FractionalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 32768, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "Time256", IsArray: false, ArrayLength: 0, Components: "fractional_timestamp", Scale: 256, Offset: 1, Units: "s", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint8, FieldId: 6, Name: "FilteredBpm", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint32, FieldId: 9, Name: "EventTimestamp", IsArray: true, ArrayLength: 0, Components: "", Scale: 1024, Offset: 1, Units: "s", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Byte, FieldId: 10, Name: "EventTimestamp12", IsArray: true, ArrayLength: 0, Components: "event_timestamp,event_timestamp,event_timestamp,event_timestamp,event_timestamp,event_timestamp,event_timestamp,event_timestamp,event_timestamp,event_timestamp", Scale: 0, Offset: 1, Units: "s", Bits: "10", Accumulate: "1,1,1,1,1,1,1,1,1,1", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Hr),
		fieldMapForHr,
		func() (registration.IFitMessage, error) {
			return &Hr{}, nil
		},
	)
}
