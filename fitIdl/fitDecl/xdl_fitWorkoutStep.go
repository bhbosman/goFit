package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type WorkoutStep struct {
	MessageIndex        MessageIndex
	WktStepName         string
	DurationType        WktStepDuration
	TargetType          WktStepTarget
	Intensity           Intensity
	Notes               string
	Equipment           WorkoutEquipment
	ExerciseCategory    ExerciseCategory
	ExerciseName        uint16
	ExerciseWeight      uint16
	WeightDisplayUnit   FitBaseUnit
	SecondaryTargetType WktStepTarget
}

func (obj *WorkoutStep) MsgNumber() uint16 {
	return uint16(MesgNum_WorkoutStep)
}

func (obj *WorkoutStep) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWorkoutStep = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.String, FieldId: 0, Name: "WktStepName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "DurationType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Enum, FieldId: 3, Name: "TargetType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "Intensity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.String, FieldId: 8, Name: "Notes", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "50"},
	9:   {BaseType: messages.Enum, FieldId: 9, Name: "Equipment", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Enum, FieldId: 10, Name: "ExerciseCategory", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint16, FieldId: 11, Name: "ExerciseName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	12:  {BaseType: messages.Uint16, FieldId: 12, Name: "ExerciseWeight", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "kg", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	13:  {BaseType: messages.Enum, FieldId: 13, Name: "WeightDisplayUnit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	19:  {BaseType: messages.Enum, FieldId: 19, Name: "SecondaryTargetType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_WorkoutStep),
		fieldMapForWorkoutStep,
		func() (registration.IFitMessage, error) {
			return &WorkoutStep{}, nil
		},
	)
}
