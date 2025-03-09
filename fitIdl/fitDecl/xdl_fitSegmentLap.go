package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SegmentLap struct {
	MessageIndex                MessageIndex
	Timestamp                   uint32
	Event                       Event
	EventType                   EventType
	StartTime                   uint32
	StartPositionLat            int32
	StartPositionLong           int32
	EndPositionLat              int32
	EndPositionLong             int32
	TotalElapsedTime            uint32
	TotalTimerTime              uint32
	TotalDistance               uint32
	TotalCalories               uint16
	TotalFatCalories            uint16
	AvgSpeed                    uint16
	MaxSpeed                    uint16
	AvgHeartRate                uint8
	MaxHeartRate                uint8
	AvgCadence                  uint8
	MaxCadence                  uint8
	AvgPower                    uint16
	MaxPower                    uint16
	TotalAscent                 uint16
	TotalDescent                uint16
	Sport                       Sport
	EventGroup                  uint8
	NecLat                      int32
	NecLong                     int32
	SwcLat                      int32
	SwcLong                     int32
	Name                        string
	NormalizedPower             uint16
	LeftRightBalance            LeftRightBalance100
	SubSport                    SubSport
	TotalWork                   uint32
	AvgAltitude                 uint16
	MaxAltitude                 uint16
	GpsAccuracy                 uint8
	AvgGrade                    int16
	AvgPosGrade                 int16
	AvgNegGrade                 int16
	MaxPosGrade                 int16
	MaxNegGrade                 int16
	AvgTemperature              int8
	MaxTemperature              int8
	TotalMovingTime             uint32
	AvgPosVerticalSpeed         int16
	AvgNegVerticalSpeed         int16
	MaxPosVerticalSpeed         int16
	MaxNegVerticalSpeed         int16
	TimeInHrZone                uint32
	TimeInSpeedZone             uint32
	TimeInCadenceZone           uint32
	TimeInPowerZone             uint32
	RepetitionNum               uint16
	MinAltitude                 uint16
	MinHeartRate                uint8
	ActiveTime                  uint32
	WktStepIndex                MessageIndex
	SportEvent                  SportEvent
	AvgLeftTorqueEffectiveness  uint8
	AvgRightTorqueEffectiveness uint8
	AvgLeftPedalSmoothness      uint8
	AvgRightPedalSmoothness     uint8
	AvgCombinedPedalSmoothness  uint8
	Status                      SegmentLapStatus
	Uuid                        string
	AvgFractionalCadence        uint8
	MaxFractionalCadence        uint8
	TotalFractionalCycles       uint8
	FrontGearShiftCount         uint16
	RearGearShiftCount          uint16
	TimeStanding                uint32
	StandCount                  uint16
	AvgLeftPco                  int8
	AvgRightPco                 int8
	AvgLeftPowerPhase           uint8
	AvgLeftPowerPhasePeak       uint8
	AvgRightPowerPhase          uint8
	AvgRightPowerPhasePeak      uint8
	AvgPowerPosition            uint16
	MaxPowerPosition            uint16
	AvgCadencePosition          uint8
	MaxCadencePosition          uint8
	Manufacturer                Manufacturer
	TotalGrit                   float32
	TotalFlow                   float32
	AvgGrit                     float32
	AvgFlow                     float32
	TotalFractionalAscent       uint8
	TotalFractionalDescent      uint8
	EnhancedAvgAltitude         uint32
	EnhancedMaxAltitude         uint32
	EnhancedMinAltitude         uint32
}

func (obj *SegmentLap) MsgNumber() uint16 {
	return uint16(MesgNum_SegmentLap)
}

func (obj *SegmentLap) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSegmentLap = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Lap end time.", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Event", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "EventType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "StartTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Sint32, FieldId: 3, Name: "StartPositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Sint32, FieldId: 4, Name: "StartPositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Sint32, FieldId: 5, Name: "EndPositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Sint32, FieldId: 6, Name: "EndPositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint32, FieldId: 7, Name: "TotalElapsedTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time (includes pauses)", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint32, FieldId: 8, Name: "TotalTimerTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timer Time (excludes pauses)", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint32, FieldId: 9, Name: "TotalDistance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint16, FieldId: 11, Name: "TotalCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Uint16, FieldId: 12, Name: "TotalFatCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "If New Leaf", Products: "", Example: "1"},
	13:  {BaseType: messages.Uint16, FieldId: 13, Name: "AvgSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Uint16, FieldId: 14, Name: "MaxSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	15:  {BaseType: messages.Uint8, FieldId: 15, Name: "AvgHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	16:  {BaseType: messages.Uint8, FieldId: 16, Name: "MaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	17:  {BaseType: messages.Uint8, FieldId: 17, Name: "AvgCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time", Products: "", Example: "1"},
	18:  {BaseType: messages.Uint8, FieldId: 18, Name: "MaxCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "18", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	19:  {BaseType: messages.Uint16, FieldId: 19, Name: "AvgPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time", Products: "", Example: "1"},
	20:  {BaseType: messages.Uint16, FieldId: 20, Name: "MaxPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "20", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	21:  {BaseType: messages.Uint16, FieldId: 21, Name: "TotalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	22:  {BaseType: messages.Uint16, FieldId: 22, Name: "TotalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	23:  {BaseType: messages.Enum, FieldId: 23, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	24:  {BaseType: messages.Uint8, FieldId: 24, Name: "EventGroup", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	25:  {BaseType: messages.Sint32, FieldId: 25, Name: "NecLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "North east corner latitude.", Products: "", Example: "1"},
	26:  {BaseType: messages.Sint32, FieldId: 26, Name: "NecLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "26", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "North east corner longitude.", Products: "", Example: "1"},
	27:  {BaseType: messages.Sint32, FieldId: 27, Name: "SwcLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "27", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "South west corner latitude.", Products: "", Example: "1"},
	28:  {BaseType: messages.Sint32, FieldId: 28, Name: "SwcLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "28", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "South west corner latitude.", Products: "", Example: "1"},
	29:  {BaseType: messages.String, FieldId: 29, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "29", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	30:  {BaseType: messages.Uint16, FieldId: 30, Name: "NormalizedPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "30", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	31:  {BaseType: messages.Enum, FieldId: 31, Name: "LeftRightBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "31", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	32:  {BaseType: messages.Enum, FieldId: 32, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	33:  {BaseType: messages.Uint32, FieldId: 33, Name: "TotalWork", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "J", Bits: "33", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	34:  {BaseType: messages.Uint16, FieldId: 34, Name: "AvgAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "34", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	35:  {BaseType: messages.Uint16, FieldId: 35, Name: "MaxAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_max_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "35", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	36:  {BaseType: messages.Uint8, FieldId: 36, Name: "GpsAccuracy", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "36", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	37:  {BaseType: messages.Sint16, FieldId: 37, Name: "AvgGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "37", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	38:  {BaseType: messages.Sint16, FieldId: 38, Name: "AvgPosGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "38", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	39:  {BaseType: messages.Sint16, FieldId: 39, Name: "AvgNegGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "39", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	40:  {BaseType: messages.Sint16, FieldId: 40, Name: "MaxPosGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "40", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	41:  {BaseType: messages.Sint16, FieldId: 41, Name: "MaxNegGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "41", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	42:  {BaseType: messages.Sint8, FieldId: 42, Name: "AvgTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "42", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	43:  {BaseType: messages.Sint8, FieldId: 43, Name: "MaxTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "43", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	44:  {BaseType: messages.Uint32, FieldId: 44, Name: "TotalMovingTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "44", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	45:  {BaseType: messages.Sint16, FieldId: 45, Name: "AvgPosVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "45", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	46:  {BaseType: messages.Sint16, FieldId: 46, Name: "AvgNegVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "46", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	47:  {BaseType: messages.Sint16, FieldId: 47, Name: "MaxPosVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "47", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	48:  {BaseType: messages.Sint16, FieldId: 48, Name: "MaxNegVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "48", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	49:  {BaseType: messages.Uint32, FieldId: 49, Name: "TimeInHrZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "49", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	50:  {BaseType: messages.Uint32, FieldId: 50, Name: "TimeInSpeedZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "50", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	51:  {BaseType: messages.Uint32, FieldId: 51, Name: "TimeInCadenceZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "51", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	52:  {BaseType: messages.Uint32, FieldId: 52, Name: "TimeInPowerZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "52", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	53:  {BaseType: messages.Uint16, FieldId: 53, Name: "RepetitionNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "53", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	54:  {BaseType: messages.Uint16, FieldId: 54, Name: "MinAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_min_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "54", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	55:  {BaseType: messages.Uint8, FieldId: 55, Name: "MinHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "55", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	56:  {BaseType: messages.Uint32, FieldId: 56, Name: "ActiveTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "56", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	57:  {BaseType: messages.Enum, FieldId: 57, Name: "WktStepIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "57", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	58:  {BaseType: messages.Enum, FieldId: 58, Name: "SportEvent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "58", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	59:  {BaseType: messages.Uint8, FieldId: 59, Name: "AvgLeftTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "59", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	60:  {BaseType: messages.Uint8, FieldId: 60, Name: "AvgRightTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "60", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	61:  {BaseType: messages.Uint8, FieldId: 61, Name: "AvgLeftPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "61", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	62:  {BaseType: messages.Uint8, FieldId: 62, Name: "AvgRightPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "62", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	63:  {BaseType: messages.Uint8, FieldId: 63, Name: "AvgCombinedPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "63", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	64:  {BaseType: messages.Enum, FieldId: 64, Name: "Status", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "64", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	65:  {BaseType: messages.String, FieldId: 65, Name: "Uuid", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "65", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "33"},
	66:  {BaseType: messages.Uint8, FieldId: 66, Name: "AvgFractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "66", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the avg_cadence", Products: "", Example: "1"},
	67:  {BaseType: messages.Uint8, FieldId: 67, Name: "MaxFractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "67", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the max_cadence", Products: "", Example: "1"},
	68:  {BaseType: messages.Uint8, FieldId: 68, Name: "TotalFractionalCycles", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "cycles", Bits: "68", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the total_cycles", Products: "", Example: "1"},
	69:  {BaseType: messages.Uint16, FieldId: 69, Name: "FrontGearShiftCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "69", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	70:  {BaseType: messages.Uint16, FieldId: 70, Name: "RearGearShiftCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "70", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	71:  {BaseType: messages.Uint32, FieldId: 71, Name: "TimeStanding", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "71", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total time spent in the standing position", Products: "", Example: ""},
	72:  {BaseType: messages.Uint16, FieldId: 72, Name: "StandCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "72", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of transitions to the standing state", Products: "", Example: ""},
	73:  {BaseType: messages.Sint8, FieldId: 73, Name: "AvgLeftPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "73", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left platform center offset", Products: "", Example: ""},
	74:  {BaseType: messages.Sint8, FieldId: 74, Name: "AvgRightPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "74", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right platform center offset", Products: "", Example: ""},
	75:  {BaseType: messages.Uint8, FieldId: 75, Name: "AvgLeftPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "75", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	76:  {BaseType: messages.Uint8, FieldId: 76, Name: "AvgLeftPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "76", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	77:  {BaseType: messages.Uint8, FieldId: 77, Name: "AvgRightPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "77", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	78:  {BaseType: messages.Uint8, FieldId: 78, Name: "AvgRightPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "78", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	79:  {BaseType: messages.Uint16, FieldId: 79, Name: "AvgPowerPosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "79", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average power by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	80:  {BaseType: messages.Uint16, FieldId: 80, Name: "MaxPowerPosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "80", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum power by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	81:  {BaseType: messages.Uint8, FieldId: 81, Name: "AvgCadencePosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "81", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average cadence by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	82:  {BaseType: messages.Uint8, FieldId: 82, Name: "MaxCadencePosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "82", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum cadence by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	83:  {BaseType: messages.Enum, FieldId: 83, Name: "Manufacturer", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "83", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Manufacturer that produced the segment", Products: "", Example: ""},
	84:  {BaseType: messages.Float32, FieldId: 84, Name: "TotalGrit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kGrit", Bits: "84", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	85:  {BaseType: messages.Float32, FieldId: 85, Name: "TotalFlow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Flow", Bits: "85", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	86:  {BaseType: messages.Float32, FieldId: 86, Name: "AvgGrit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kGrit", Bits: "86", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	87:  {BaseType: messages.Float32, FieldId: 87, Name: "AvgFlow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Flow", Bits: "87", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	89:  {BaseType: messages.Uint8, FieldId: 89, Name: "TotalFractionalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "89", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of total_ascent", Products: "", Example: ""},
	90:  {BaseType: messages.Uint8, FieldId: 90, Name: "TotalFractionalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "90", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of total_descent", Products: "", Example: ""},
	91:  {BaseType: messages.Uint32, FieldId: 91, Name: "EnhancedAvgAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "91", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	92:  {BaseType: messages.Uint32, FieldId: 92, Name: "EnhancedMaxAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "92", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	93:  {BaseType: messages.Uint32, FieldId: 93, Name: "EnhancedMinAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "93", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SegmentLap),
		fieldMapForSegmentLap,
		func() (registration.IFitMessage, error) {
			return &SegmentLap{}, nil
		},
	)
}
