package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Course struct {
	Sport        Sport
	Name         string
	Capabilities CourseCapabilities
	SubSport     SubSport
}

func (obj *Course) MsgNumber() uint16 {
	return uint16(MesgNum_Course)
}

func (obj *Course) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForCourse = map[uint16]registration.DefinedFields{
	4: {BaseType: messages.Enum, FieldId: 4, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5: {BaseType: messages.String, FieldId: 5, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	6: {BaseType: messages.Enum, FieldId: 6, Name: "Capabilities", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7: {BaseType: messages.Enum, FieldId: 7, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Course),
		fieldMapForCourse,
		func() (registration.IFitMessage, error) {
			return &Course{}, nil
		},
	)
}
