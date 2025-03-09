package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type WorkoutSession struct {
	MessageIndex   MessageIndex
	Sport          Sport
	SubSport       SubSport
	NumValidSteps  uint16
	FirstStepIndex uint16
	PoolLength     uint16
	PoolLengthUnit DisplayMeasure
}

func (obj *WorkoutSession) MsgNumber() uint16 {
	return uint16(MesgNum_WorkoutSession)
}

func (obj *WorkoutSession) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWorkoutSession = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "NumValidSteps", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "FirstStepIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "PoolLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "PoolLengthUnit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_WorkoutSession),
		fieldMapForWorkoutSession,
		func() (registration.IFitMessage, error) {
			return &WorkoutSession{}, nil
		},
	)
}
