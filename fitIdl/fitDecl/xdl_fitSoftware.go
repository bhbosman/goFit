package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Software struct {
	MessageIndex MessageIndex
	Version      uint16
	PartNumber   string
}

func (obj *Software) MsgNumber() uint16 {
	return uint16(MesgNum_Software)
}

func (obj *Software) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSoftware = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "Version", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.String, FieldId: 5, Name: "PartNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Software),
		fieldMapForSoftware,
		func() (registration.IFitMessage, error) {
			return &Software{}, nil
		},
	)
}
