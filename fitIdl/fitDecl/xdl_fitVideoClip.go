package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type VideoClip struct {
	ClipNumber       uint16
	StartTimestamp   uint32
	StartTimestampMs uint16
	EndTimestamp     uint32
	EndTimestampMs   uint16
	ClipStart        uint32
	ClipEnd          uint32
}

func (obj *VideoClip) MsgNumber() uint16 {
	return uint16(MesgNum_VideoClip)
}

func (obj *VideoClip) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForVideoClip = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Uint16, FieldId: 0, Name: "ClipNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1: {BaseType: messages.Uint32, FieldId: 1, Name: "StartTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2: {BaseType: messages.Uint16, FieldId: 2, Name: "StartTimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3: {BaseType: messages.Uint32, FieldId: 3, Name: "EndTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4: {BaseType: messages.Uint16, FieldId: 4, Name: "EndTimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6: {BaseType: messages.Uint32, FieldId: 6, Name: "ClipStart", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Start of clip in video time", Products: "", Example: ""},
	7: {BaseType: messages.Uint32, FieldId: 7, Name: "ClipEnd", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "End of clip in video time", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_VideoClip),
		fieldMapForVideoClip,
		func() (registration.IFitMessage, error) {
			return &VideoClip{}, nil
		},
	)
}
