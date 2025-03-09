package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaStressData struct {
	Timestamp          uint32
	ProcessingInterval uint16
	StressLevel        int8
}

func (obj *HsaStressData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaStressData)
}

func (obj *HsaStressData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaStressData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Sint8, FieldId: 1, Name: "StressLevel", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Stress Level: [0,100] Off wrist: -1 Excess motion: -2 Not enough data: -3 Recovering from exercise: -4 Unidentified: -5 Blank: -16", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaStressData),
		fieldMapForHsaStressData,
		func() (registration.IFitMessage, error) {
			return &HsaStressData{}, nil
		},
	)
}
