package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type TankUpdate struct {
	Timestamp uint32
	Sensor    AntChannelId
	Pressure  uint16
}

func (obj *TankUpdate) MsgNumber() uint16 {
	return uint16(MesgNum_TankUpdate)
}

func (obj *TankUpdate) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForTankUpdate = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Sensor", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "Pressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "bar", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_TankUpdate),
		fieldMapForTankUpdate,
		func() (registration.IFitMessage, error) {
			return &TankUpdate{}, nil
		},
	)
}
