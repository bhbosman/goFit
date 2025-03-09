package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaWristTemperatureData struct {
	Timestamp          uint32
	ProcessingInterval uint16
	Value              uint16
}

func (obj *HsaWristTemperatureData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaWristTemperatureData)
}

func (obj *HsaWristTemperatureData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaWristTemperatureData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "Value", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "degC", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Wrist temperature reading", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaWristTemperatureData),
		fieldMapForHsaWristTemperatureData,
		func() (registration.IFitMessage, error) {
			return &HsaWristTemperatureData{}, nil
		},
	)
}
