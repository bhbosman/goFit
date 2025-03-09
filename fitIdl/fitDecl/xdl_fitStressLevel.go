package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type StressLevel struct {
	StressLevelValue int16
	StressLevelTime  uint32
}

func (obj *StressLevel) MsgNumber() uint16 {
	return uint16(MesgNum_StressLevel)
}

func (obj *StressLevel) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForStressLevel = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Sint16, FieldId: 0, Name: "StressLevelValue", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1: {BaseType: messages.Uint32, FieldId: 1, Name: "StressLevelTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time stress score was calculated", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_StressLevel),
		fieldMapForStressLevel,
		func() (registration.IFitMessage, error) {
			return &StressLevel{}, nil
		},
	)
}
