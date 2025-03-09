package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type HsaRespirationData struct {
	Timestamp          uint32
	ProcessingInterval uint16
	RespirationRate    int16
}

func (obj *HsaRespirationData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaRespirationData)
}

func (obj *HsaRespirationData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaRespirationData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Sint16, FieldId: 1, Name: "RespirationRate", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "breaths/min", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Breaths / min: [1,100] Invalid: 255 Excess motion: 254 Off wrist: 253 Not available: 252 Blank: 2.4", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaRespirationData),
		fieldMapForHsaRespirationData,
		func() (registration.IFitMessage, error) {
			return &HsaRespirationData{}, nil
		},
	)
}
