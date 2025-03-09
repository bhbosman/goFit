package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ClimbPro struct {
	Timestamp     uint32
	PositionLat   int32
	PositionLong  int32
	ClimbProEvent ClimbProEvent
	ClimbNumber   uint16
	ClimbCategory uint8
	CurrentDist   float32
}

func (obj *ClimbPro) MsgNumber() uint16 {
	return uint16(MesgNum_ClimbPro)
}

func (obj *ClimbPro) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForClimbPro = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Sint32, FieldId: 0, Name: "PositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Sint32, FieldId: 1, Name: "PositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "ClimbProEvent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "ClimbNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Uint8, FieldId: 4, Name: "ClimbCategory", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:   {BaseType: messages.Float32, FieldId: 5, Name: "CurrentDist", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ClimbPro),
		fieldMapForClimbPro,
		func() (registration.IFitMessage, error) {
			return &ClimbPro{}, nil
		},
	)
}
