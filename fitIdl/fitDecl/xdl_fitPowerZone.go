package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type PowerZone struct {
	MessageIndex MessageIndex
	HighValue    uint16
	Name         string
}

func (obj *PowerZone) MsgNumber() uint16 {
	return uint16(MesgNum_PowerZone)
}

func (obj *PowerZone) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForPowerZone = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "HighValue", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.String, FieldId: 2, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_PowerZone),
		fieldMapForPowerZone,
		func() (registration.IFitMessage, error) {
			return &PowerZone{}, nil
		},
	)
}
