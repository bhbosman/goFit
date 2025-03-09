package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type TankSummary struct {
	Timestamp     uint32
	Sensor        AntChannelId
	StartPressure uint16
	EndPressure   uint16
	VolumeUsed    uint32
}

func (obj *TankSummary) MsgNumber() uint16 {
	return uint16(MesgNum_TankSummary)
}

func (obj *TankSummary) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForTankSummary = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Sensor", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "StartPressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "bar", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "EndPressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "bar", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "VolumeUsed", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "L", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_TankSummary),
		fieldMapForTankSummary,
		func() (registration.IFitMessage, error) {
			return &TankSummary{}, nil
		},
	)
}
