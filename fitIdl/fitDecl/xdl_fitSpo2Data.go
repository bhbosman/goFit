package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Spo2Data struct {
	Timestamp         uint32
	ReadingSpo2       uint8
	ReadingConfidence uint8
	Mode              Spo2MeasurementType
}

func (obj *Spo2Data) MsgNumber() uint16 {
	return uint16(MesgNum_Spo2Data)
}

func (obj *Spo2Data) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSpo2Data = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint8, FieldId: 0, Name: "ReadingSpo2", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "ReadingConfidence", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "Mode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Mode when data was captured", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Spo2Data),
		fieldMapForSpo2Data,
		func() (registration.IFitMessage, error) {
			return &Spo2Data{}, nil
		},
	)
}
