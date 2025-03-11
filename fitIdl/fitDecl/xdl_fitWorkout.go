package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Workout struct {
	MessageIndex   MessageIndex
	Sport          Sport
	Capabilities   WorkoutCapabilities
	NumValidSteps  uint16
	WktName        string
	SubSport       SubSport
	PoolLength     uint16
	PoolLengthUnit DisplayMeasure
	WktDescription string
}

func (obj *Workout) MsgNumber() uint16 {
	return uint16(MesgNum_Workout)
}

func (obj *Workout) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWorkout = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Enum, FieldId: 4, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "Capabilities", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint16, FieldId: 6, Name: "NumValidSteps", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "number of valid steps", Products: "", Example: "1"},
	8:   {BaseType: messages.String, FieldId: 8, Name: "WktName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	11:  {BaseType: messages.Enum, FieldId: 11, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Uint16, FieldId: 14, Name: "PoolLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	15:  {BaseType: messages.Enum, FieldId: 15, Name: "PoolLengthUnit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	17:  {BaseType: messages.String, FieldId: 17, Name: "WktDescription", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Description of the workout", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Workout),
		fieldMapForWorkout,
		func() (registration.IFitMessage, error) {
			return &Workout{}, nil
		},
	)
}
