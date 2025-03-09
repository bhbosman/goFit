package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ExdDataConceptConfiguration struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptIndex uint8
	DataPage     uint8
	ConceptKey   uint8
	Scaling      uint8
	DataUnits    ExdDataUnits
	Qualifier    ExdQualifiers
	Descriptor   ExdDescriptors
	IsSigned     bool
}

func (obj *ExdDataConceptConfiguration) MsgNumber() uint16 {
	return uint16(MesgNum_ExdDataConceptConfiguration)
}

func (obj *ExdDataConceptConfiguration) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForExdDataConceptConfiguration = map[uint16]registration.DefinedFields{
	0:  {BaseType: messages.Uint8, FieldId: 0, Name: "ScreenIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:  {BaseType: messages.Byte, FieldId: 1, Name: "ConceptField", IsArray: false, ArrayLength: 0, Components: "field_id,concept_index", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:  {BaseType: messages.Uint8, FieldId: 2, Name: "FieldId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:  {BaseType: messages.Uint8, FieldId: 3, Name: "ConceptIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:  {BaseType: messages.Uint8, FieldId: 4, Name: "DataPage", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:  {BaseType: messages.Uint8, FieldId: 5, Name: "ConceptKey", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:  {BaseType: messages.Uint8, FieldId: 6, Name: "Scaling", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:  {BaseType: messages.Enum, FieldId: 8, Name: "DataUnits", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:  {BaseType: messages.Enum, FieldId: 9, Name: "Qualifier", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10: {BaseType: messages.Enum, FieldId: 10, Name: "Descriptor", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11: {BaseType: messages.Bool, FieldId: 11, Name: "IsSigned", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ExdDataConceptConfiguration),
		fieldMapForExdDataConceptConfiguration,
		func() (registration.IFitMessage, error) {
			return &ExdDataConceptConfiguration{}, nil
		},
	)
}
