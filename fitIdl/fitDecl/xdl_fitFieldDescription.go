package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type FieldDescription struct {
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	FitBaseTypeId         FitBaseType
	FieldName             string
	Array                 uint8
	Components            string
	Scale                 uint8
	Offset                int8
	Units                 string
	Bits                  string
	Accumulate            string
	FitBaseUnitId         FitBaseUnit
	NativeMesgNum         MesgNum
	NativeFieldNum        uint8
}

func (obj *FieldDescription) MsgNumber() uint16 {
	return uint16(MesgNum_FieldDescription)
}

func (obj *FieldDescription) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForFieldDescription = map[uint16]registration.DefinedFields{
	0:  {BaseType: messages.Uint8, FieldId: 0, Name: "DeveloperDataIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:  {BaseType: messages.Uint8, FieldId: 1, Name: "FieldDefinitionNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:  {BaseType: messages.Enum, FieldId: 2, Name: "FitBaseTypeId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:  {BaseType: messages.String, FieldId: 3, Name: "FieldName", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "64"},
	4:  {BaseType: messages.Uint8, FieldId: 4, Name: "Array", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "0"},
	5:  {BaseType: messages.String, FieldId: 5, Name: "Components", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "0"},
	6:  {BaseType: messages.Uint8, FieldId: 6, Name: "Scale", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:  {BaseType: messages.Sint8, FieldId: 7, Name: "Offset", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:  {BaseType: messages.String, FieldId: 8, Name: "Units", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	9:  {BaseType: messages.String, FieldId: 9, Name: "Bits", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "0"},
	10: {BaseType: messages.String, FieldId: 10, Name: "Accumulate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "0"},
	13: {BaseType: messages.Enum, FieldId: 13, Name: "FitBaseUnitId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14: {BaseType: messages.Enum, FieldId: 14, Name: "NativeMesgNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	15: {BaseType: messages.Uint8, FieldId: 15, Name: "NativeFieldNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_FieldDescription),
		fieldMapForFieldDescription,
		func() (registration.IFitMessage, error) {
			return &FieldDescription{}, nil
		},
	)
}
