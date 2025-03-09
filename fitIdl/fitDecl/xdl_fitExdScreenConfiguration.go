package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ExdScreenConfiguration struct {
	ScreenIndex   uint8
	FieldCount    uint8
	Layout        ExdLayout
	ScreenEnabled bool
}

func (obj *ExdScreenConfiguration) MsgNumber() uint16 {
	return uint16(MesgNum_ExdScreenConfiguration)
}

func (obj *ExdScreenConfiguration) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForExdScreenConfiguration = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Uint8, FieldId: 0, Name: "ScreenIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1: {BaseType: messages.Uint8, FieldId: 1, Name: "FieldCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "number of fields in screen", Products: "", Example: "1"},
	2: {BaseType: messages.Enum, FieldId: 2, Name: "Layout", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3: {BaseType: messages.Bool, FieldId: 3, Name: "ScreenEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ExdScreenConfiguration),
		fieldMapForExdScreenConfiguration,
		func() (registration.IFitMessage, error) {
			return &ExdScreenConfiguration{}, nil
		},
	)
}
