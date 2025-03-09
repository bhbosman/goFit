package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type GyroscopeData struct {
	Timestamp        uint32
	TimestampMs      uint16
	SampleTimeOffset uint16
	GyroX            uint16
	GyroY            uint16
	GyroZ            uint16
	CalibratedGyroX  float32
	CalibratedGyroY  float32
	CalibratedGyroZ  float32
}

func (obj *GyroscopeData) MsgNumber() uint16 {
	return uint16(MesgNum_GyroscopeData)
}

func (obj *GyroscopeData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForGyroscopeData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the timestamp.", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "SampleTimeOffset", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Each time in the array describes the time at which the gyro sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in gyro_x and gyro_y and gyro_z", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "GyroX", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "GyroY", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "GyroZ", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.", Products: "", Example: ""},
	5:   {BaseType: messages.Float32, FieldId: 5, Name: "CalibratedGyroX", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "deg/s", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated gyro reading", Products: "", Example: ""},
	6:   {BaseType: messages.Float32, FieldId: 6, Name: "CalibratedGyroY", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "deg/s", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated gyro reading", Products: "", Example: ""},
	7:   {BaseType: messages.Float32, FieldId: 7, Name: "CalibratedGyroZ", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "deg/s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibrated gyro reading", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_GyroscopeData),
		fieldMapForGyroscopeData,
		func() (registration.IFitMessage, error) {
			return &GyroscopeData{}, nil
		},
	)
}
