package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type FileId struct {
	Type         File
	Manufacturer Manufacturer
	SerialNumber uint32
	TimeCreated  uint32
	Number       uint16
	ProductName  string
}

func (obj *FileId) MsgNumber() uint16 {
	return uint16(MesgNum_FileId)
}

func (obj *FileId) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForFileId = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Enum, FieldId: 0, Name: "Type", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1: {BaseType: messages.Enum, FieldId: 1, Name: "Manufacturer", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3: {BaseType: messages.Uint32z, FieldId: 3, Name: "SerialNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4: {BaseType: messages.Uint32, FieldId: 4, Name: "TimeCreated", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Only set for files that are can be created/erased.", Products: "", Example: "1"},
	5: {BaseType: messages.Uint16, FieldId: 5, Name: "Number", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Only set for files that are not created/erased.", Products: "", Example: "1"},
	8: {BaseType: messages.String, FieldId: 8, Name: "ProductName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Optional free form string to indicate the devices name or model", Products: "", Example: "20"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_FileId),
		fieldMapForFileId,
		func() (registration.IFitMessage, error) {
			return &FileId{}, nil
		},
	)
}
