package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ExerciseTitle struct {
	MessageIndex     MessageIndex
	ExerciseCategory ExerciseCategory
	ExerciseName     uint16
	WktStepName      string
}

func (obj *ExerciseTitle) MsgNumber() uint16 {
	return uint16(MesgNum_ExerciseTitle)
}

func (obj *ExerciseTitle) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForExerciseTitle = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "ExerciseCategory", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "ExerciseName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.String, FieldId: 2, Name: "WktStepName", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "200"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ExerciseTitle),
		fieldMapForExerciseTitle,
		func() (registration.IFitMessage, error) {
			return &ExerciseTitle{}, nil
		},
	)
}
