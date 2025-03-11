package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type CadenceZone struct {
	MessageIndex MessageIndex
	HighValue    uint8
	Name         string
}

func (obj *CadenceZone) MsgNumber() uint16 {
	return uint16(MesgNum_CadenceZone)
}

func (obj *CadenceZone) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForCadenceZone = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint8, FieldId: 0, Name: "HighValue", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.String, FieldId: 1, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_CadenceZone),
		fieldMapForCadenceZone,
		func() (registration.IFitMessage, error) {
			return &CadenceZone{}, nil
		},
	)
}
