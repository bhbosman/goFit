package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type OhrSettings struct {
	Timestamp uint32
	Enabled   Switch
}

func (obj *OhrSettings) MsgNumber() uint16 {
	return uint16(MesgNum_OhrSettings)
}

func (obj *OhrSettings) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForOhrSettings = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_OhrSettings),
		fieldMapForOhrSettings,
		func() (registration.IFitMessage, error) {
			return &OhrSettings{}, nil
		},
	)
}
