package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaSpo2Data struct {
	Timestamp          uint32
	ProcessingInterval uint16
	ReadingSpo2        uint8
	Confidence         uint8
}

func (obj *HsaSpo2Data) MsgNumber() uint16 {
	return uint16(MesgNum_HsaSpo2Data)
}

func (obj *HsaSpo2Data) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaSpo2Data = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "ReadingSpo2", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "SpO2 Reading: [70,100] Blank: 240", Products: "", Example: ""},
	2:   {BaseType: messages.Uint8, FieldId: 2, Name: "Confidence", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "SpO2 Confidence: [0,254]", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaSpo2Data),
		fieldMapForHsaSpo2Data,
		func() (registration.IFitMessage, error) {
			return &HsaSpo2Data{}, nil
		},
	)
}
