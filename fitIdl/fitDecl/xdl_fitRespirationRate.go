package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type RespirationRate struct {
	Timestamp       uint32
	RespirationRate int16
}

func (obj *RespirationRate) MsgNumber() uint16 {
	return uint16(MesgNum_RespirationRate)
}

func (obj *RespirationRate) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForRespirationRate = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Sint16, FieldId: 0, Name: "RespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "breaths/min", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Breaths * 100 /min, -300 indicates invalid, -200 indicates large motion, -100 indicates off wrist", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_RespirationRate),
		fieldMapForRespirationRate,
		func() (registration.IFitMessage, error) {
			return &RespirationRate{}, nil
		},
	)
}
