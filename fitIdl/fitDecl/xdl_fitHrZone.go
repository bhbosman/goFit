package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HrZone struct {
	MessageIndex MessageIndex
	HighBpm      uint8
	Name         string
}

func (obj *HrZone) MsgNumber() uint16 {
	return uint16(MesgNum_HrZone)
}

func (obj *HrZone) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHrZone = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "HighBpm", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.String, FieldId: 2, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HrZone),
		fieldMapForHrZone,
		func() (registration.IFitMessage, error) {
			return &HrZone{}, nil
		},
	)
}
