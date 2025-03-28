package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type MesgCapabilities struct {
	MessageIndex MessageIndex
	File         File
	MesgNum      MesgNum
	CountType    MesgCount
}

func (obj *MesgCapabilities) MsgNumber() uint16 {
	return uint16(MesgNum_MesgCapabilities)
}

func (obj *MesgCapabilities) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMesgCapabilities = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "File", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "MesgNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "CountType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_MesgCapabilities),
		fieldMapForMesgCapabilities,
		func() (registration.IFitMessage, error) {
			return &MesgCapabilities{}, nil
		},
	)
}
