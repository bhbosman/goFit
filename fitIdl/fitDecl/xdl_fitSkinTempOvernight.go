package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type SkinTempOvernight struct {
	Timestamp            uint32
	LocalTimestamp       LocalDateTime
	AverageDeviation     float32
	Average7DayDeviation float32
	NightlyValue         float32
}

func (obj *SkinTempOvernight) MsgNumber() uint16 {
	return uint16(MesgNum_SkinTempOvernight)
}

func (obj *SkinTempOvernight) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSkinTempOvernight = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "LocalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Float32, FieldId: 1, Name: "AverageDeviation", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The average overnight deviation from baseline temperature in degrees C", Products: "", Example: ""},
	2:   {BaseType: messages.Float32, FieldId: 2, Name: "Average7DayDeviation", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The average 7 day overnight deviation from baseline temperature in degrees C", Products: "", Example: ""},
	4:   {BaseType: messages.Float32, FieldId: 4, Name: "NightlyValue", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Final overnight temperature value", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SkinTempOvernight),
		fieldMapForSkinTempOvernight,
		func() (registration.IFitMessage, error) {
			return &SkinTempOvernight{}, nil
		},
	)
}
