package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Set struct {
	Timestamp         uint32
	Duration          uint32
	Repetitions       uint16
	Weight            uint16
	SetType           SetType
	StartTime         uint32
	Category          ExerciseCategory
	CategorySubtype   uint16
	WeightDisplayUnit FitBaseUnit
	MessageIndex      MessageIndex
	WktStepIndex      MessageIndex
}

func (obj *Set) MsgNumber() uint16 {
	return uint16(MesgNum_Set)
}

func (obj *Set) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSet = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Uint32, FieldId: 254, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timestamp of the set", Products: "", Example: ""},
	0:   {BaseType: messages.Uint32, FieldId: 0, Name: "Duration", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "Repetitions", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "# of repitions of the movement", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "Weight", IsArray: false, ArrayLength: 0, Components: "", Scale: 16, Offset: 1, Units: "kg", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Amount of weight applied for the set", Products: "", Example: ""},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "SetType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "StartTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Start time of the set", Products: "", Example: ""},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "Category", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:   {BaseType: messages.Uint16, FieldId: 8, Name: "CategorySubtype", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Based on the associated category, see [category]_exercise_names", Products: "", Example: ""},
	9:   {BaseType: messages.Enum, FieldId: 9, Name: "WeightDisplayUnit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Enum, FieldId: 10, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	11:  {BaseType: messages.Enum, FieldId: 11, Name: "WktStepIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Set),
		fieldMapForSet,
		func() (registration.IFitMessage, error) {
			return &Set{}, nil
		},
	)
}
