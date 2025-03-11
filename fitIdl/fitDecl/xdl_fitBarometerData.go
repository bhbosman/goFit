package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type BarometerData struct {
	Timestamp        uint32
	TimestampMs      uint16
	SampleTimeOffset uint16
	BaroPres         uint32
}

func (obj *BarometerData) MsgNumber() uint16 {
	return uint16(MesgNum_BarometerData)
}

func (obj *BarometerData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForBarometerData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the timestamp.", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "SampleTimeOffset", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Each time in the array describes the time at which the barometer sample with the corrosponding index was taken. The samples may span across seconds. Array size must match the number of samples in baro_cal", Products: "", Example: ""},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "BaroPres", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Pa", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_BarometerData),
		fieldMapForBarometerData,
		func() (registration.IFitMessage, error) {
			return &BarometerData{}, nil
		},
	)
}
