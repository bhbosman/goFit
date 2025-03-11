package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HrvValue struct {
	Timestamp uint32
	Value     uint16
}

func (obj *HrvValue) MsgNumber() uint16 {
	return uint16(MesgNum_HrvValue)
}

func (obj *HrvValue) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHrvValue = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "Value", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "5 minute RMSSD", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HrvValue),
		fieldMapForHrvValue,
		func() (registration.IFitMessage, error) {
			return &HrvValue{}, nil
		},
	)
}
