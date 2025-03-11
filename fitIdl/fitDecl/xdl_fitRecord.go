package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Record struct {
	Timestamp                     uint32
	PositionLat                   int32
	PositionLong                  int32
	Altitude                      uint16
	HeartRate                     uint8
	Cadence                       uint8
	Distance                      uint32
	Speed                         uint16
	Power                         uint16
	CompressedSpeedDistance       byte
	Grade                         int16
	Resistance                    uint8
	TimeFromCourse                int32
	CycleLength                   uint8
	Temperature                   int8
	Speed1s                       uint8
	Cycles                        uint8
	TotalCycles                   uint32
	CompressedAccumulatedPower    uint16
	AccumulatedPower              uint32
	LeftRightBalance              LeftRightBalance
	GpsAccuracy                   uint8
	VerticalSpeed                 int16
	Calories                      uint16
	VerticalOscillation           uint16
	StanceTimePercent             uint16
	StanceTime                    uint16
	ActivityType                  ActivityType
	LeftTorqueEffectiveness       uint8
	RightTorqueEffectiveness      uint8
	LeftPedalSmoothness           uint8
	RightPedalSmoothness          uint8
	CombinedPedalSmoothness       uint8
	Time128                       uint8
	StrokeType                    StrokeType
	Zone                          uint8
	BallSpeed                     uint16
	Cadence256                    uint16
	FractionalCadence             uint8
	TotalHemoglobinConc           uint16
	TotalHemoglobinConcMin        uint16
	TotalHemoglobinConcMax        uint16
	SaturatedHemoglobinPercent    uint16
	SaturatedHemoglobinPercentMin uint16
	SaturatedHemoglobinPercentMax uint16
	DeviceIndex                   DeviceIndex
	LeftPco                       int8
	RightPco                      int8
	LeftPowerPhase                uint8
	LeftPowerPhasePeak            uint8
	RightPowerPhase               uint8
	RightPowerPhasePeak           uint8
	EnhancedSpeed                 uint32
	EnhancedAltitude              uint32
	BatterySoc                    uint8
	MotorPower                    uint16
	VerticalRatio                 uint16
	StanceTimeBalance             uint16
	StepLength                    uint16
	CycleLength16                 uint16
	AbsolutePressure              uint32
	Depth                         uint32
	NextStopDepth                 uint32
	NextStopTime                  uint32
	TimeToSurface                 uint32
	NdlTime                       uint32
	CnsLoad                       uint8
	N2Load                        uint16
	RespirationRate               uint8
	EnhancedRespirationRate       uint16
	Grit                          float32
	Flow                          float32
	CurrentStress                 uint16
	EbikeTravelRange              uint16
	EbikeBatteryLevel             uint8
	EbikeAssistMode               uint8
	EbikeAssistLevelPercent       uint8
	AirTimeRemaining              uint32
	PressureSac                   uint16
	VolumeSac                     uint16
	Rmv                           uint16
	AscentRate                    int32
	Po2                           uint8
	CoreTemperature               uint16
}

func (obj *Record) MsgNumber() uint16 {
	return uint16(MesgNum_Record)
}

func (obj *Record) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForRecord = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Sint32, FieldId: 0, Name: "PositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Sint32, FieldId: 1, Name: "PositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "Altitude", IsArray: false, ArrayLength: 0, Components: "enhanced_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "HeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint8, FieldId: 4, Name: "Cadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "Distance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint16, FieldId: 6, Name: "Speed", IsArray: false, ArrayLength: 0, Components: "enhanced_speed", Scale: 1000, Offset: 1, Units: "m/s", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint16, FieldId: 7, Name: "Power", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.Byte, FieldId: 8, Name: "CompressedSpeedDistance", IsArray: true, ArrayLength: 0, Components: "speed,distance", Scale: 0, Offset: 1, Units: "m/s,m", Bits: "8", Accumulate: "0,1", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Sint16, FieldId: 9, Name: "Grade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint8, FieldId: 10, Name: "Resistance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Relative. 0 is none 254 is Max.", Products: "", Example: "1"},
	11:  {BaseType: messages.Sint32, FieldId: 11, Name: "TimeFromCourse", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Uint8, FieldId: 12, Name: "CycleLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	13:  {BaseType: messages.Sint8, FieldId: 13, Name: "Temperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	17:  {BaseType: messages.Uint8, FieldId: 17, Name: "Speed1s", IsArray: true, ArrayLength: 0, Components: "", Scale: 16, Offset: 1, Units: "m/s", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Speed at 1s intervals. Timestamp field indicates time of last array element.", Products: "", Example: "5"},
	18:  {BaseType: messages.Uint8, FieldId: 18, Name: "Cycles", IsArray: false, ArrayLength: 0, Components: "total_cycles", Scale: 1, Offset: 1, Units: "cycles", Bits: "18", Accumulate: "1", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	19:  {BaseType: messages.Uint32, FieldId: 19, Name: "TotalCycles", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "cycles", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	28:  {BaseType: messages.Uint16, FieldId: 28, Name: "CompressedAccumulatedPower", IsArray: false, ArrayLength: 0, Components: "accumulated_power", Scale: 1, Offset: 1, Units: "watts", Bits: "28", Accumulate: "1", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	29:  {BaseType: messages.Uint32, FieldId: 29, Name: "AccumulatedPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "29", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	30:  {BaseType: messages.Enum, FieldId: 30, Name: "LeftRightBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "30", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	31:  {BaseType: messages.Uint8, FieldId: 31, Name: "GpsAccuracy", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "31", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	32:  {BaseType: messages.Sint16, FieldId: 32, Name: "VerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	33:  {BaseType: messages.Uint16, FieldId: 33, Name: "Calories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "33", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	39:  {BaseType: messages.Uint16, FieldId: 39, Name: "VerticalOscillation", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mm", Bits: "39", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	40:  {BaseType: messages.Uint16, FieldId: 40, Name: "StanceTimePercent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "40", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	41:  {BaseType: messages.Uint16, FieldId: 41, Name: "StanceTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "ms", Bits: "41", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	42:  {BaseType: messages.Enum, FieldId: 42, Name: "ActivityType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "42", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	43:  {BaseType: messages.Uint8, FieldId: 43, Name: "LeftTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "43", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	44:  {BaseType: messages.Uint8, FieldId: 44, Name: "RightTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "44", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	45:  {BaseType: messages.Uint8, FieldId: 45, Name: "LeftPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "45", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	46:  {BaseType: messages.Uint8, FieldId: 46, Name: "RightPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "46", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	47:  {BaseType: messages.Uint8, FieldId: 47, Name: "CombinedPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "47", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	48:  {BaseType: messages.Uint8, FieldId: 48, Name: "Time128", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "s", Bits: "48", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	49:  {BaseType: messages.Enum, FieldId: 49, Name: "StrokeType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "49", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	50:  {BaseType: messages.Uint8, FieldId: 50, Name: "Zone", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "50", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	51:  {BaseType: messages.Uint16, FieldId: 51, Name: "BallSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m/s", Bits: "51", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	52:  {BaseType: messages.Uint16, FieldId: 52, Name: "Cadence256", IsArray: false, ArrayLength: 0, Components: "", Scale: 256, Offset: 1, Units: "rpm", Bits: "52", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Log cadence and fractional cadence for backwards compatability", Products: "", Example: "1"},
	53:  {BaseType: messages.Uint8, FieldId: 53, Name: "FractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "53", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	54:  {BaseType: messages.Uint16, FieldId: 54, Name: "TotalHemoglobinConc", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "54", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total saturated and unsaturated hemoglobin", Products: "", Example: "1"},
	55:  {BaseType: messages.Uint16, FieldId: 55, Name: "TotalHemoglobinConcMin", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "55", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min saturated and unsaturated hemoglobin", Products: "", Example: "1"},
	56:  {BaseType: messages.Uint16, FieldId: 56, Name: "TotalHemoglobinConcMax", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "56", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max saturated and unsaturated hemoglobin", Products: "", Example: "1"},
	57:  {BaseType: messages.Uint16, FieldId: 57, Name: "SaturatedHemoglobinPercent", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "57", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Percentage of hemoglobin saturated with oxygen", Products: "", Example: "1"},
	58:  {BaseType: messages.Uint16, FieldId: 58, Name: "SaturatedHemoglobinPercentMin", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "58", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min percentage of hemoglobin saturated with oxygen", Products: "", Example: "1"},
	59:  {BaseType: messages.Uint16, FieldId: 59, Name: "SaturatedHemoglobinPercentMax", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "59", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max percentage of hemoglobin saturated with oxygen", Products: "", Example: "1"},
	62:  {BaseType: messages.Enum, FieldId: 62, Name: "DeviceIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "62", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	67:  {BaseType: messages.Sint8, FieldId: 67, Name: "LeftPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "67", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Left platform center offset", Products: "", Example: ""},
	68:  {BaseType: messages.Sint8, FieldId: 68, Name: "RightPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "68", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Right platform center offset", Products: "", Example: ""},
	69:  {BaseType: messages.Uint8, FieldId: 69, Name: "LeftPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "69", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Left power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	70:  {BaseType: messages.Uint8, FieldId: 70, Name: "LeftPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "70", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Left power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	71:  {BaseType: messages.Uint8, FieldId: 71, Name: "RightPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "71", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Right power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	72:  {BaseType: messages.Uint8, FieldId: 72, Name: "RightPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "72", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Right power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	73:  {BaseType: messages.Uint32, FieldId: 73, Name: "EnhancedSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "73", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	78:  {BaseType: messages.Uint32, FieldId: 78, Name: "EnhancedAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "78", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	81:  {BaseType: messages.Uint8, FieldId: 81, Name: "BatterySoc", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "81", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev battery state of charge", Products: "", Example: ""},
	82:  {BaseType: messages.Uint16, FieldId: 82, Name: "MotorPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "82", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev motor power", Products: "", Example: ""},
	83:  {BaseType: messages.Uint16, FieldId: 83, Name: "VerticalRatio", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "83", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	84:  {BaseType: messages.Uint16, FieldId: 84, Name: "StanceTimeBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "84", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	85:  {BaseType: messages.Uint16, FieldId: 85, Name: "StepLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mm", Bits: "85", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	87:  {BaseType: messages.Uint16, FieldId: 87, Name: "CycleLength16", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "87", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Supports larger cycle sizes needed for paddlesports. Max cycle size: 655.35", Products: "", Example: ""},
	91:  {BaseType: messages.Uint32, FieldId: 91, Name: "AbsolutePressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Pa", Bits: "91", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Includes atmospheric pressure", Products: "", Example: ""},
	92:  {BaseType: messages.Uint32, FieldId: 92, Name: "Depth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "92", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	93:  {BaseType: messages.Uint32, FieldId: 93, Name: "NextStopDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "93", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	94:  {BaseType: messages.Uint32, FieldId: 94, Name: "NextStopTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "94", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	95:  {BaseType: messages.Uint32, FieldId: 95, Name: "TimeToSurface", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "95", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	96:  {BaseType: messages.Uint32, FieldId: 96, Name: "NdlTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "96", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	97:  {BaseType: messages.Uint8, FieldId: 97, Name: "CnsLoad", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "97", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	98:  {BaseType: messages.Uint16, FieldId: 98, Name: "N2Load", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "98", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	99:  {BaseType: messages.Uint8, FieldId: 99, Name: "RespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_respiration_rate", Scale: 1, Offset: 1, Units: "s", Bits: "99", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	108: {BaseType: messages.Uint16, FieldId: 108, Name: "EnhancedRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "108", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	114: {BaseType: messages.Float32, FieldId: 114, Name: "Grit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "114", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	115: {BaseType: messages.Float32, FieldId: 115, Name: "Flow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "115", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	116: {BaseType: messages.Uint16, FieldId: 116, Name: "CurrentStress", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "", Bits: "116", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Current Stress value", Products: "", Example: ""},
	117: {BaseType: messages.Uint16, FieldId: 117, Name: "EbikeTravelRange", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "km", Bits: "117", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	118: {BaseType: messages.Uint8, FieldId: 118, Name: "EbikeBatteryLevel", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "118", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	119: {BaseType: messages.Uint8, FieldId: 119, Name: "EbikeAssistMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "depends on sensor", Bits: "119", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	120: {BaseType: messages.Uint8, FieldId: 120, Name: "EbikeAssistLevelPercent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "120", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	123: {BaseType: messages.Uint32, FieldId: 123, Name: "AirTimeRemaining", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "123", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	124: {BaseType: messages.Uint16, FieldId: 124, Name: "PressureSac", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "bar/min", Bits: "124", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Pressure-based surface air consumption", Products: "", Example: ""},
	125: {BaseType: messages.Uint16, FieldId: 125, Name: "VolumeSac", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "L/min", Bits: "125", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Volumetric surface air consumption", Products: "", Example: ""},
	126: {BaseType: messages.Uint16, FieldId: 126, Name: "Rmv", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "L/min", Bits: "126", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Respiratory minute volume", Products: "", Example: ""},
	127: {BaseType: messages.Sint32, FieldId: 127, Name: "AscentRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "127", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	129: {BaseType: messages.Uint8, FieldId: 129, Name: "Po2", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "129", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Current partial pressure of oxygen", Products: "", Example: ""},
	139: {BaseType: messages.Uint16, FieldId: 139, Name: "CoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "139", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Record),
		fieldMapForRecord,
		func() (registration.IFitMessage, error) {
			return &Record{}, nil
		},
	)
}
