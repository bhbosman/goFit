package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DiveGas struct {
	MessageIndex  MessageIndex
	HeliumContent uint8
	OxygenContent uint8
	Status        DiveGasStatus
	Mode          DiveGasMode
}

func (obj *DiveGas) MsgNumber() uint16 {
	return uint16(MesgNum_DiveGas)
}

func (obj *DiveGas) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDiveGas = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint8, FieldId: 0, Name: "HeliumContent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "OxygenContent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "Status", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Enum, FieldId: 3, Name: "Mode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DiveGas),
		fieldMapForDiveGas,
		func() (registration.IFitMessage, error) {
			return &DiveGas{}, nil
		},
	)
}
