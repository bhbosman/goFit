package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type FileCapabilities struct {
	MessageIndex MessageIndex
	Type         File
	Flags        FileFlags
	Directory    string
	MaxCount     uint16
	MaxSize      uint32
}

func (obj *FileCapabilities) MsgNumber() uint16 {
	return uint16(MesgNum_FileCapabilities)
}

func (obj *FileCapabilities) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForFileCapabilities = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Type", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "Flags", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.String, FieldId: 2, Name: "Directory", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "MaxCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "MaxSize", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bytes", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_FileCapabilities),
		fieldMapForFileCapabilities,
		func() (registration.IFitMessage, error) {
			return &FileCapabilities{}, nil
		},
	)
}
