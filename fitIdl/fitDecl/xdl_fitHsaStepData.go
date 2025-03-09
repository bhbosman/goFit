package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaStepData struct {
	Timestamp          uint32
	ProcessingInterval uint16
	Steps              uint32
}

func (obj *HsaStepData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaStepData)
}

func (obj *HsaStepData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaStepData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds. File start: 0xFFFFFFEF File stop: 0xFFFFFFEE", Products: "", Example: ""},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "Steps", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "steps", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total step sum", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaStepData),
		fieldMapForHsaStepData,
		func() (registration.IFitMessage, error) {
			return &HsaStepData{}, nil
		},
	)
}
