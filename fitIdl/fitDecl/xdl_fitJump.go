package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Jump struct {
	Timestamp     uint32
	Distance      float32
	Height        float32
	Rotations     uint8
	HangTime      float32
	Score         float32
	PositionLat   int32
	PositionLong  int32
	Speed         uint16
	EnhancedSpeed uint32
}

func (obj *Jump) MsgNumber() uint16 {
	return uint16(MesgNum_Jump)
}

func (obj *Jump) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForJump = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Float32, FieldId: 0, Name: "Distance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Float32, FieldId: 1, Name: "Height", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Uint8, FieldId: 2, Name: "Rotations", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Float32, FieldId: 3, Name: "HangTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Float32, FieldId: 4, Name: "Score", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "A score for a jump calculated based on hang time, rotations, and distance.", Products: "", Example: ""},
	5:   {BaseType: messages.Sint32, FieldId: 5, Name: "PositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Sint32, FieldId: 6, Name: "PositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	7:   {BaseType: messages.Uint16, FieldId: 7, Name: "Speed", IsArray: false, ArrayLength: 0, Components: "enhanced_speed", Scale: 1000, Offset: 1, Units: "m/s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:   {BaseType: messages.Uint32, FieldId: 8, Name: "EnhancedSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Jump),
		fieldMapForJump,
		func() (registration.IFitMessage, error) {
			return &Jump{}, nil
		},
	)
}
