package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ChronoShotData struct {
	Timestamp uint32
	ShotSpeed uint32
	ShotNum   uint16
}

func (obj *ChronoShotData) MsgNumber() uint16 {
	return uint16(MesgNum_ChronoShotData)
}

func (obj *ChronoShotData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForChronoShotData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint32, FieldId: 0, Name: "ShotSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "ShotNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ChronoShotData),
		fieldMapForChronoShotData,
		func() (registration.IFitMessage, error) {
			return &ChronoShotData{}, nil
		},
	)
}
