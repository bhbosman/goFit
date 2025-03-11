package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ExdDataFieldConfiguration struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptCount uint8
	DisplayType  ExdDisplayType
	Title        string
}

func (obj *ExdDataFieldConfiguration) MsgNumber() uint16 {
	return uint16(MesgNum_ExdDataFieldConfiguration)
}

func (obj *ExdDataFieldConfiguration) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForExdDataFieldConfiguration = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Uint8, FieldId: 0, Name: "ScreenIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1: {BaseType: messages.Byte, FieldId: 1, Name: "ConceptField", IsArray: false, ArrayLength: 0, Components: "field_id,concept_count", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2: {BaseType: messages.Uint8, FieldId: 2, Name: "FieldId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3: {BaseType: messages.Uint8, FieldId: 3, Name: "ConceptCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4: {BaseType: messages.Enum, FieldId: 4, Name: "DisplayType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5: {BaseType: messages.String, FieldId: 5, Name: "Title", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ExdDataFieldConfiguration),
		fieldMapForExdDataFieldConfiguration,
		func() (registration.IFitMessage, error) {
			return &ExdDataFieldConfiguration{}, nil
		},
	)
}
