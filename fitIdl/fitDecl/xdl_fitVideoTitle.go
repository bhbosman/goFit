package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type VideoTitle struct {
	MessageIndex MessageIndex
	MessageCount uint16
	Text         string
}

func (obj *VideoTitle) MsgNumber() uint16 {
	return uint16(MesgNum_VideoTitle)
}

func (obj *VideoTitle) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForVideoTitle = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Long titles will be split into multiple parts", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "MessageCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total number of title parts", Products: "", Example: "1"},
	1:   {BaseType: messages.String, FieldId: 1, Name: "Text", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "80"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_VideoTitle),
		fieldMapForVideoTitle,
		func() (registration.IFitMessage, error) {
			return &VideoTitle{}, nil
		},
	)
}
