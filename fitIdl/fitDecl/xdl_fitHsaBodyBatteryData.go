package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaBodyBatteryData struct {
	Timestamp          uint32
	ProcessingInterval uint16
	Level              int8
	Charged            int16
	Uncharged          int16
}

func (obj *HsaBodyBatteryData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaBodyBatteryData)
}

func (obj *HsaBodyBatteryData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaBodyBatteryData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Sint8, FieldId: 1, Name: "Level", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Body battery level: [0,100] Blank: -16", Products: "", Example: ""},
	2:   {BaseType: messages.Sint16, FieldId: 2, Name: "Charged", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Body battery charged value", Products: "", Example: ""},
	3:   {BaseType: messages.Sint16, FieldId: 3, Name: "Uncharged", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Body battery uncharged value", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaBodyBatteryData),
		fieldMapForHsaBodyBatteryData,
		func() (registration.IFitMessage, error) {
			return &HsaBodyBatteryData{}, nil
		},
	)
}
