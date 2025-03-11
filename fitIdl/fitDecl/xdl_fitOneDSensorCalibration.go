package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type OneDSensorCalibration struct {
	Timestamp          uint32
	SensorType         SensorType
	CalibrationDivisor uint32
	LevelShift         uint32
	OffsetCal          int32
}

func (obj *OneDSensorCalibration) MsgNumber() uint16 {
	return uint16(MesgNum_OneDSensorCalibration)
}

func (obj *OneDSensorCalibration) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForOneDSensorCalibration = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "SensorType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates which sensor the calibration is for", Products: "", Example: ""},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "CalibrationDivisor", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Calibration factor divisor", Products: "", Example: ""},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "LevelShift", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Level shift value used to shift the ADC value back into range", Products: "", Example: ""},
	4:   {BaseType: messages.Sint32, FieldId: 4, Name: "OffsetCal", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Internal Calibration factor", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_OneDSensorCalibration),
		fieldMapForOneDSensorCalibration,
		func() (registration.IFitMessage, error) {
			return &OneDSensorCalibration{}, nil
		},
	)
}
