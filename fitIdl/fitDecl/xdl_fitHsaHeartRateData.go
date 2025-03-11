package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaHeartRateData struct {
	Timestamp          uint32
	ProcessingInterval uint16
	Status             uint8
	HeartRate          uint8
}

func (obj *HsaHeartRateData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaHeartRateData)
}

func (obj *HsaHeartRateData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaHeartRateData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "ProcessingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Processing interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "Status", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Status of measurements in buffer - 0 indicates SEARCHING 1 indicates LOCKED", Products: "", Example: ""},
	2:   {BaseType: messages.Uint8, FieldId: 2, Name: "HeartRate", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Beats / min. Blank: 0", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaHeartRateData),
		fieldMapForHsaHeartRateData,
		func() (registration.IFitMessage, error) {
			return &HsaHeartRateData{}, nil
		},
	)
}
