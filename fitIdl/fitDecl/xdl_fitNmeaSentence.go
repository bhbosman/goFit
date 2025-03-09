package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type NmeaSentence struct {
	Timestamp   uint32
	TimestampMs uint16
	Sentence    string
}

func (obj *NmeaSentence) MsgNumber() uint16 {
	return uint16(MesgNum_NmeaSentence)
}

func (obj *NmeaSentence) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForNmeaSentence = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timestamp message was output", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fractional part of timestamp, added to timestamp", Products: "", Example: "1"},
	1:   {BaseType: messages.String, FieldId: 1, Name: "Sentence", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "NMEA sentence", Products: "", Example: "83"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_NmeaSentence),
		fieldMapForNmeaSentence,
		func() (registration.IFitMessage, error) {
			return &NmeaSentence{}, nil
		},
	)
}
