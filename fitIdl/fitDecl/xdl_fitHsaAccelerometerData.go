package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaAccelerometerData struct {
	Timestamp        uint32
	TimestampMs      uint16
	SamplingInterval uint16
	AccelX           int16
	AccelY           int16
	AccelZ           int16
	Timestamp32k     uint32
}

func (obj *HsaAccelerometerData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaAccelerometerData)
}

func (obj *HsaAccelerometerData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaAccelerometerData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond resolution of the timestamp", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "SamplingInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Sampling Interval in Milliseconds", Products: "", Example: ""},
	2:   {BaseType: messages.Sint16, FieldId: 2, Name: "AccelX", IsArray: true, ArrayLength: 0, Components: "", Scale: 1.024, Offset: 1, Units: "mG", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "X-Axis Measurement", Products: "", Example: ""},
	3:   {BaseType: messages.Sint16, FieldId: 3, Name: "AccelY", IsArray: true, ArrayLength: 0, Components: "", Scale: 1.024, Offset: 1, Units: "mG", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Y-Axis Measurement", Products: "", Example: ""},
	4:   {BaseType: messages.Sint16, FieldId: 4, Name: "AccelZ", IsArray: true, ArrayLength: 0, Components: "", Scale: 1.024, Offset: 1, Units: "mG", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Z-Axis Measurement", Products: "", Example: ""},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "Timestamp32k", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "32 kHz timestamp", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaAccelerometerData),
		fieldMapForHsaAccelerometerData,
		func() (registration.IFitMessage, error) {
			return &HsaAccelerometerData{}, nil
		},
	)
}
