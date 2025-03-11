package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type TimeInZone struct {
	Timestamp                uint32
	ReferenceMesg            MesgNum
	ReferenceIndex           MessageIndex
	TimeInHrZone             uint32
	TimeInSpeedZone          uint32
	TimeInCadenceZone        uint32
	TimeInPowerZone          uint32
	HrZoneHighBoundary       uint8
	SpeedZoneHighBoundary    uint16
	CadenceZoneHighBondary   uint8
	PowerZoneHighBoundary    uint16
	HrCalcType               HrZoneCalc
	MaxHeartRate             uint8
	RestingHeartRate         uint8
	ThresholdHeartRate       uint8
	PwrCalcType              PwrZoneCalc
	FunctionalThresholdPower uint16
}

func (obj *TimeInZone) MsgNumber() uint16 {
	return uint16(MesgNum_TimeInZone)
}

func (obj *TimeInZone) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForTimeInZone = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "ReferenceMesg", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "ReferenceIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "TimeInHrZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "TimeInSpeedZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "TimeInCadenceZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "TimeInPowerZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Uint8, FieldId: 6, Name: "HrZoneHighBoundary", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	7:   {BaseType: messages.Uint16, FieldId: 7, Name: "SpeedZoneHighBoundary", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:   {BaseType: messages.Uint8, FieldId: 8, Name: "CadenceZoneHighBondary", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "PowerZoneHighBoundary", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	10:  {BaseType: messages.Enum, FieldId: 10, Name: "HrCalcType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	11:  {BaseType: messages.Uint8, FieldId: 11, Name: "MaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	12:  {BaseType: messages.Uint8, FieldId: 12, Name: "RestingHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	13:  {BaseType: messages.Uint8, FieldId: 13, Name: "ThresholdHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	14:  {BaseType: messages.Enum, FieldId: 14, Name: "PwrCalcType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	15:  {BaseType: messages.Uint16, FieldId: 15, Name: "FunctionalThresholdPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_TimeInZone),
		fieldMapForTimeInZone,
		func() (registration.IFitMessage, error) {
			return &TimeInZone{}, nil
		},
	)
}
