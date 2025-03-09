package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type MetZone struct {
	MessageIndex MessageIndex
	HighBpm      uint8
	Calories     uint16
	FatCalories  uint8
}

func (obj *MetZone) MsgNumber() uint16 {
	return uint16(MesgNum_MetZone)
}

func (obj *MetZone) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMetZone = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "HighBpm", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "Calories", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "kcal / min", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "FatCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "kcal / min", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_MetZone),
		fieldMapForMetZone,
		func() (registration.IFitMessage, error) {
			return &MetZone{}, nil
		},
	)
}
