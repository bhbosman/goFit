package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type Lap struct {
	MessageIndex                  MessageIndex
	Timestamp                     uint32
	Event                         Event
	EventType                     EventType
	StartTime                     uint32
	StartPositionLat              int32
	StartPositionLong             int32
	EndPositionLat                int32
	EndPositionLong               int32
	TotalElapsedTime              uint32
	TotalTimerTime                uint32
	TotalDistance                 uint32
	TotalCalories                 uint16
	TotalFatCalories              uint16
	AvgSpeed                      uint16
	MaxSpeed                      uint16
	AvgHeartRate                  uint8
	MaxHeartRate                  uint8
	AvgPower                      uint16
	MaxPower                      uint16
	TotalAscent                   uint16
	TotalDescent                  uint16
	Intensity                     Intensity
	LapTrigger                    LapTrigger
	Sport                         Sport
	EventGroup                    uint8
	NumLengths                    uint16
	NormalizedPower               uint16
	LeftRightBalance              LeftRightBalance100
	FirstLengthIndex              uint16
	AvgStrokeDistance             uint16
	SwimStroke                    SwimStroke
	SubSport                      SubSport
	NumActiveLengths              uint16
	TotalWork                     uint32
	AvgAltitude                   uint16
	MaxAltitude                   uint16
	GpsAccuracy                   uint8
	AvgGrade                      int16
	AvgPosGrade                   int16
	AvgNegGrade                   int16
	MaxPosGrade                   int16
	MaxNegGrade                   int16
	AvgTemperature                int8
	MaxTemperature                int8
	TotalMovingTime               uint32
	AvgPosVerticalSpeed           int16
	AvgNegVerticalSpeed           int16
	MaxPosVerticalSpeed           int16
	MaxNegVerticalSpeed           int16
	TimeInHrZone                  uint32
	TimeInSpeedZone               uint32
	TimeInCadenceZone             uint32
	TimeInPowerZone               uint32
	RepetitionNum                 uint16
	MinAltitude                   uint16
	MinHeartRate                  uint8
	WktStepIndex                  MessageIndex
	OpponentScore                 uint16
	StrokeCount                   uint16
	ZoneCount                     uint16
	AvgVerticalOscillation        uint16
	AvgStanceTimePercent          uint16
	AvgStanceTime                 uint16
	AvgFractionalCadence          uint8
	MaxFractionalCadence          uint8
	TotalFractionalCycles         uint8
	PlayerScore                   uint16
	AvgTotalHemoglobinConc        uint16
	MinTotalHemoglobinConc        uint16
	MaxTotalHemoglobinConc        uint16
	AvgSaturatedHemoglobinPercent uint16
	MinSaturatedHemoglobinPercent uint16
	MaxSaturatedHemoglobinPercent uint16
	AvgLeftTorqueEffectiveness    uint8
	AvgRightTorqueEffectiveness   uint8
	AvgLeftPedalSmoothness        uint8
	AvgRightPedalSmoothness       uint8
	AvgCombinedPedalSmoothness    uint8
	TimeStanding                  uint32
	StandCount                    uint16
	AvgLeftPco                    int8
	AvgRightPco                   int8
	AvgLeftPowerPhase             uint8
	AvgLeftPowerPhasePeak         uint8
	AvgRightPowerPhase            uint8
	AvgRightPowerPhasePeak        uint8
	AvgPowerPosition              uint16
	MaxPowerPosition              uint16
	AvgCadencePosition            uint8
	MaxCadencePosition            uint8
	EnhancedAvgSpeed              uint32
	EnhancedMaxSpeed              uint32
	EnhancedAvgAltitude           uint32
	EnhancedMinAltitude           uint32
	EnhancedMaxAltitude           uint32
	AvgLevMotorPower              uint16
	MaxLevMotorPower              uint16
	LevBatteryConsumption         uint8
	AvgVerticalRatio              uint16
	AvgStanceTimeBalance          uint16
	AvgStepLength                 uint16
	AvgVam                        uint16
	AvgDepth                      uint32
	MaxDepth                      uint32
	MinTemperature                int8
	EnhancedAvgRespirationRate    uint16
	EnhancedMaxRespirationRate    uint16
	AvgRespirationRate            uint8
	MaxRespirationRate            uint8
	TotalGrit                     float32
	TotalFlow                     float32
	JumpCount                     uint16
	AvgGrit                       float32
	AvgFlow                       float32
	TotalFractionalAscent         uint8
	TotalFractionalDescent        uint8
	AvgCoreTemperature            uint16
	MinCoreTemperature            uint16
	MaxCoreTemperature            uint16
}

func (obj *Lap) MsgNumber() uint16 {
	return uint16(MesgNum_Lap)
}

func (obj *Lap) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForLap = map[uint16]registration.DefinedFields{
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
	13:  {BaseType: messages.Uint16, FieldId: 13, Name: "AvgSpeed", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_speed", Scale: 1000, Offset: 1, Units: "m/s", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Uint16, FieldId: 14, Name: "MaxSpeed", IsArray: false, ArrayLength: 0, Components: "enhanced_max_speed", Scale: 1000, Offset: 1, Units: "m/s", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	15:  {BaseType: messages.Uint8, FieldId: 15, Name: "AvgHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	16:  {BaseType: messages.Uint8, FieldId: 16, Name: "MaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	19:  {BaseType: messages.Uint16, FieldId: 19, Name: "AvgPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time", Products: "", Example: "1"},
	20:  {BaseType: messages.Uint16, FieldId: 20, Name: "MaxPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "20", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	21:  {BaseType: messages.Uint16, FieldId: 21, Name: "TotalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	22:  {BaseType: messages.Uint16, FieldId: 22, Name: "TotalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	23:  {BaseType: messages.Enum, FieldId: 23, Name: "Intensity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	24:  {BaseType: messages.Enum, FieldId: 24, Name: "LapTrigger", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	25:  {BaseType: messages.Enum, FieldId: 25, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	26:  {BaseType: messages.Uint8, FieldId: 26, Name: "EventGroup", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "26", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	32:  {BaseType: messages.Uint16, FieldId: 32, Name: "NumLengths", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "lengths", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "# of lengths of swim pool", Products: "", Example: "1"},
	33:  {BaseType: messages.Uint16, FieldId: 33, Name: "NormalizedPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "33", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	34:  {BaseType: messages.Enum, FieldId: 34, Name: "LeftRightBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "34", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	35:  {BaseType: messages.Uint16, FieldId: 35, Name: "FirstLengthIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "35", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	37:  {BaseType: messages.Uint16, FieldId: 37, Name: "AvgStrokeDistance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "37", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	38:  {BaseType: messages.Enum, FieldId: 38, Name: "SwimStroke", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "38", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	39:  {BaseType: messages.Enum, FieldId: 39, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "39", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	40:  {BaseType: messages.Uint16, FieldId: 40, Name: "NumActiveLengths", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "lengths", Bits: "40", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "# of active lengths of swim pool", Products: "", Example: "1"},
	41:  {BaseType: messages.Uint32, FieldId: 41, Name: "TotalWork", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "J", Bits: "41", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	42:  {BaseType: messages.Uint16, FieldId: 42, Name: "AvgAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "42", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	43:  {BaseType: messages.Uint16, FieldId: 43, Name: "MaxAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_max_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "43", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	44:  {BaseType: messages.Uint8, FieldId: 44, Name: "GpsAccuracy", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "44", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	45:  {BaseType: messages.Sint16, FieldId: 45, Name: "AvgGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "45", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	46:  {BaseType: messages.Sint16, FieldId: 46, Name: "AvgPosGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "46", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	47:  {BaseType: messages.Sint16, FieldId: 47, Name: "AvgNegGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "47", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	48:  {BaseType: messages.Sint16, FieldId: 48, Name: "MaxPosGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "48", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	49:  {BaseType: messages.Sint16, FieldId: 49, Name: "MaxNegGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "49", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	50:  {BaseType: messages.Sint8, FieldId: 50, Name: "AvgTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "50", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	51:  {BaseType: messages.Sint8, FieldId: 51, Name: "MaxTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "51", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	52:  {BaseType: messages.Uint32, FieldId: 52, Name: "TotalMovingTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "52", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	53:  {BaseType: messages.Sint16, FieldId: 53, Name: "AvgPosVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "53", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	54:  {BaseType: messages.Sint16, FieldId: 54, Name: "AvgNegVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "54", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	55:  {BaseType: messages.Sint16, FieldId: 55, Name: "MaxPosVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "55", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	56:  {BaseType: messages.Sint16, FieldId: 56, Name: "MaxNegVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "56", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	57:  {BaseType: messages.Uint32, FieldId: 57, Name: "TimeInHrZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "57", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	58:  {BaseType: messages.Uint32, FieldId: 58, Name: "TimeInSpeedZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "58", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	59:  {BaseType: messages.Uint32, FieldId: 59, Name: "TimeInCadenceZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "59", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	60:  {BaseType: messages.Uint32, FieldId: 60, Name: "TimeInPowerZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "60", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	61:  {BaseType: messages.Uint16, FieldId: 61, Name: "RepetitionNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "61", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	62:  {BaseType: messages.Uint16, FieldId: 62, Name: "MinAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_min_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "62", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	63:  {BaseType: messages.Uint8, FieldId: 63, Name: "MinHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "63", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	71:  {BaseType: messages.Enum, FieldId: 71, Name: "WktStepIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "71", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	74:  {BaseType: messages.Uint16, FieldId: 74, Name: "OpponentScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "74", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	75:  {BaseType: messages.Uint16, FieldId: 75, Name: "StrokeCount", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "75", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "stroke_type enum used as the index", Products: "", Example: "1"},
	76:  {BaseType: messages.Uint16, FieldId: 76, Name: "ZoneCount", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "76", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "zone number used as the index", Products: "", Example: "1"},
	77:  {BaseType: messages.Uint16, FieldId: 77, Name: "AvgVerticalOscillation", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mm", Bits: "77", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	78:  {BaseType: messages.Uint16, FieldId: 78, Name: "AvgStanceTimePercent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "78", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	79:  {BaseType: messages.Uint16, FieldId: 79, Name: "AvgStanceTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "ms", Bits: "79", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	80:  {BaseType: messages.Uint8, FieldId: 80, Name: "AvgFractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "80", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the avg_cadence", Products: "", Example: "1"},
	81:  {BaseType: messages.Uint8, FieldId: 81, Name: "MaxFractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "81", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the max_cadence", Products: "", Example: "1"},
	82:  {BaseType: messages.Uint8, FieldId: 82, Name: "TotalFractionalCycles", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "cycles", Bits: "82", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the total_cycles", Products: "", Example: "1"},
	83:  {BaseType: messages.Uint16, FieldId: 83, Name: "PlayerScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "83", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	84:  {BaseType: messages.Uint16, FieldId: 84, Name: "AvgTotalHemoglobinConc", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "84", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Avg saturated and unsaturated hemoglobin", Products: "", Example: "1"},
	85:  {BaseType: messages.Uint16, FieldId: 85, Name: "MinTotalHemoglobinConc", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "85", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min saturated and unsaturated hemoglobin", Products: "", Example: "1"},
	86:  {BaseType: messages.Uint16, FieldId: 86, Name: "MaxTotalHemoglobinConc", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "86", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max saturated and unsaturated hemoglobin", Products: "", Example: "1"},
	87:  {BaseType: messages.Uint16, FieldId: 87, Name: "AvgSaturatedHemoglobinPercent", IsArray: true, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "87", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Avg percentage of hemoglobin saturated with oxygen", Products: "", Example: "1"},
	88:  {BaseType: messages.Uint16, FieldId: 88, Name: "MinSaturatedHemoglobinPercent", IsArray: true, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "88", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min percentage of hemoglobin saturated with oxygen", Products: "", Example: "1"},
	89:  {BaseType: messages.Uint16, FieldId: 89, Name: "MaxSaturatedHemoglobinPercent", IsArray: true, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "89", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max percentage of hemoglobin saturated with oxygen", Products: "", Example: "1"},
	91:  {BaseType: messages.Uint8, FieldId: 91, Name: "AvgLeftTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "91", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	92:  {BaseType: messages.Uint8, FieldId: 92, Name: "AvgRightTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "92", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	93:  {BaseType: messages.Uint8, FieldId: 93, Name: "AvgLeftPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "93", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	94:  {BaseType: messages.Uint8, FieldId: 94, Name: "AvgRightPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "94", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	95:  {BaseType: messages.Uint8, FieldId: 95, Name: "AvgCombinedPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "95", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	98:  {BaseType: messages.Uint32, FieldId: 98, Name: "TimeStanding", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "98", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total time spent in the standing position", Products: "", Example: ""},
	99:  {BaseType: messages.Uint16, FieldId: 99, Name: "StandCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "99", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of transitions to the standing state", Products: "", Example: ""},
	100: {BaseType: messages.Sint8, FieldId: 100, Name: "AvgLeftPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "100", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left platform center offset", Products: "", Example: ""},
	101: {BaseType: messages.Sint8, FieldId: 101, Name: "AvgRightPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "101", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right platform center offset", Products: "", Example: ""},
	102: {BaseType: messages.Uint8, FieldId: 102, Name: "AvgLeftPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "102", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	103: {BaseType: messages.Uint8, FieldId: 103, Name: "AvgLeftPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "103", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	104: {BaseType: messages.Uint8, FieldId: 104, Name: "AvgRightPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "104", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	105: {BaseType: messages.Uint8, FieldId: 105, Name: "AvgRightPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "105", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	106: {BaseType: messages.Uint16, FieldId: 106, Name: "AvgPowerPosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "106", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average power by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	107: {BaseType: messages.Uint16, FieldId: 107, Name: "MaxPowerPosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "107", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum power by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	108: {BaseType: messages.Uint8, FieldId: 108, Name: "AvgCadencePosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "108", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average cadence by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	109: {BaseType: messages.Uint8, FieldId: 109, Name: "MaxCadencePosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "109", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum cadence by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	110: {BaseType: messages.Uint32, FieldId: 110, Name: "EnhancedAvgSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "110", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	111: {BaseType: messages.Uint32, FieldId: 111, Name: "EnhancedMaxSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "111", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	112: {BaseType: messages.Uint32, FieldId: 112, Name: "EnhancedAvgAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "112", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	113: {BaseType: messages.Uint32, FieldId: 113, Name: "EnhancedMinAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "113", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	114: {BaseType: messages.Uint32, FieldId: 114, Name: "EnhancedMaxAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "114", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	115: {BaseType: messages.Uint16, FieldId: 115, Name: "AvgLevMotorPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "115", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev average motor power during lap", Products: "", Example: ""},
	116: {BaseType: messages.Uint16, FieldId: 116, Name: "MaxLevMotorPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "116", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev maximum motor power during lap", Products: "", Example: ""},
	117: {BaseType: messages.Uint8, FieldId: 117, Name: "LevBatteryConsumption", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "117", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev battery consumption during lap", Products: "", Example: ""},
	118: {BaseType: messages.Uint16, FieldId: 118, Name: "AvgVerticalRatio", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "118", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	119: {BaseType: messages.Uint16, FieldId: 119, Name: "AvgStanceTimeBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "119", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	120: {BaseType: messages.Uint16, FieldId: 120, Name: "AvgStepLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mm", Bits: "120", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	121: {BaseType: messages.Uint16, FieldId: 121, Name: "AvgVam", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "121", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	122: {BaseType: messages.Uint32, FieldId: 122, Name: "AvgDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "122", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	123: {BaseType: messages.Uint32, FieldId: 123, Name: "MaxDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "123", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	124: {BaseType: messages.Sint8, FieldId: 124, Name: "MinTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "124", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	136: {BaseType: messages.Uint16, FieldId: 136, Name: "EnhancedAvgRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "136", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	137: {BaseType: messages.Uint16, FieldId: 137, Name: "EnhancedMaxRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "137", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	147: {BaseType: messages.Uint8, FieldId: 147, Name: "AvgRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "147", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	148: {BaseType: messages.Uint8, FieldId: 148, Name: "MaxRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_max_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "148", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	149: {BaseType: messages.Float32, FieldId: 149, Name: "TotalGrit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kGrit", Bits: "149", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	150: {BaseType: messages.Float32, FieldId: 150, Name: "TotalFlow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Flow", Bits: "150", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	151: {BaseType: messages.Uint16, FieldId: 151, Name: "JumpCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "151", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	153: {BaseType: messages.Float32, FieldId: 153, Name: "AvgGrit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kGrit", Bits: "153", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	154: {BaseType: messages.Float32, FieldId: 154, Name: "AvgFlow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Flow", Bits: "154", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	156: {BaseType: messages.Uint8, FieldId: 156, Name: "TotalFractionalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "156", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of total_ascent", Products: "", Example: ""},
	157: {BaseType: messages.Uint8, FieldId: 157, Name: "TotalFractionalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "157", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of total_descent", Products: "", Example: ""},
	158: {BaseType: messages.Uint16, FieldId: 158, Name: "AvgCoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "158", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	159: {BaseType: messages.Uint16, FieldId: 159, Name: "MinCoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "159", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	160: {BaseType: messages.Uint16, FieldId: 160, Name: "MaxCoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "160", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Lap),
		fieldMapForLap,
		func() (registration.IFitMessage, error) {
			return &Lap{}, nil
		},
	)
}
