package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type AccelerometerData struct {
	Timestamp                  uint32
	TimestampMs                uint16
	SampleTimeOffset           uint16
	AccelX                     uint16
	AccelY                     uint16
	AccelZ                     uint16
	CalibratedAccelX           float32
	CalibratedAccelY           float32
	CalibratedAccelZ           float32
	CompressedCalibratedAccelX int16
	CompressedCalibratedAccelY int16
	CompressedCalibratedAccelZ int16
}

func (obj *AccelerometerData) MsgNumber() uint16 {
	return uint16(MesgNum_AccelerometerData)
}

func (obj *AccelerometerData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForAccelerometerData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the timestamp.", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "SampleTimeOffset", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Each time in the array describes the time at which the accelerometer sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in accel_x and accel_y and accel_z", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "AccelX", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "AccelY", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "AccelZ", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
	5:   {BaseType: messages.Float32, FieldId: 5, Name: "CalibratedAccelX", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "g", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated accel reading", Products: "", Example: ""},
	6:   {BaseType: messages.Float32, FieldId: 6, Name: "CalibratedAccelY", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "g", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated accel reading", Products: "", Example: ""},
	7:   {BaseType: messages.Float32, FieldId: 7, Name: "CalibratedAccelZ", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "g", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated accel reading", Products: "", Example: ""},
	8:   {BaseType: messages.Sint16, FieldId: 8, Name: "CompressedCalibratedAccelX", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mG", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated accel reading", Products: "", Example: ""},
	9:   {BaseType: messages.Sint16, FieldId: 9, Name: "CompressedCalibratedAccelY", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mG", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated accel reading", Products: "", Example: ""},
	10:  {BaseType: messages.Sint16, FieldId: 10, Name: "CompressedCalibratedAccelZ", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mG", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated accel reading", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_AccelerometerData),
		fieldMapForAccelerometerData,
		func() (registration.IFitMessage, error) {
			return &AccelerometerData{}, nil
		},
	)
}
