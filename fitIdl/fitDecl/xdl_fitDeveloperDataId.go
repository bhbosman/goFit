package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DeveloperDataId struct {
	DeveloperId        byte
	ApplicationId      byte
	ManufacturerId     Manufacturer
	DeveloperDataIndex uint8
	ApplicationVersion uint32
}

func (obj *DeveloperDataId) MsgNumber() uint16 {
	return uint16(MesgNum_DeveloperDataId)
}

func (obj *DeveloperDataId) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDeveloperDataId = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Byte, FieldId: 0, Name: "DeveloperId", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	1: {BaseType: messages.Byte, FieldId: 1, Name: "ApplicationId", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	2: {BaseType: messages.Enum, FieldId: 2, Name: "ManufacturerId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3: {BaseType: messages.Uint8, FieldId: 3, Name: "DeveloperDataIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4: {BaseType: messages.Uint32, FieldId: 4, Name: "ApplicationVersion", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DeveloperDataId),
		fieldMapForDeveloperDataId,
		func() (registration.IFitMessage, error) {
			return &DeveloperDataId{}, nil
		},
	)
}
