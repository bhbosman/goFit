package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Session struct {
	MessageIndex                  MessageIndex
	Timestamp                     uint32
	Event                         Event
	EventType                     EventType
	StartTime                     uint32
	StartPositionLat              int32
	StartPositionLong             int32
	Sport                         Sport
	SubSport                      SubSport
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
	TotalTrainingEffect           uint8
	FirstLapIndex                 uint16
	NumLaps                       uint16
	EventGroup                    uint8
	Trigger                       SessionTrigger
	NecLat                        int32
	NecLong                       int32
	SwcLat                        int32
	SwcLong                       int32
	NumLengths                    uint16
	NormalizedPower               uint16
	TrainingStressScore           uint16
	IntensityFactor               uint16
	LeftRightBalance              LeftRightBalance100
	EndPositionLat                int32
	EndPositionLong               int32
	AvgStrokeCount                uint32
	AvgStrokeDistance             uint16
	SwimStroke                    SwimStroke
	PoolLength                    uint16
	ThresholdPower                uint16
	PoolLengthUnit                DisplayMeasure
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
	MinHeartRate                  uint8
	TimeInHrZone                  uint32
	TimeInSpeedZone               uint32
	TimeInCadenceZone             uint32
	TimeInPowerZone               uint32
	AvgLapTime                    uint32
	BestLapIndex                  uint16
	MinAltitude                   uint16
	PlayerScore                   uint16
	OpponentScore                 uint16
	OpponentName                  string
	StrokeCount                   uint16
	ZoneCount                     uint16
	MaxBallSpeed                  uint16
	AvgBallSpeed                  uint16
	AvgVerticalOscillation        uint16
	AvgStanceTimePercent          uint16
	AvgStanceTime                 uint16
	AvgFractionalCadence          uint8
	MaxFractionalCadence          uint8
	TotalFractionalCycles         uint8
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
	SportProfileName              string
	SportIndex                    uint8
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
	TotalAnaerobicTrainingEffect  uint8
	AvgVam                        uint16
	AvgDepth                      uint32
	MaxDepth                      uint32
	SurfaceInterval               uint32
	StartCns                      uint8
	EndCns                        uint8
	StartN2                       uint16
	EndN2                         uint16
	AvgRespirationRate            uint8
	MaxRespirationRate            uint8
	MinRespirationRate            uint8
	MinTemperature                int8
	O2Toxicity                    uint16
	DiveNumber                    uint32
	TrainingLoadPeak              int32
	EnhancedAvgRespirationRate    uint16
	EnhancedMaxRespirationRate    uint16
	EnhancedMinRespirationRate    uint16
	TotalGrit                     float32
	TotalFlow                     float32
	JumpCount                     uint16
	AvgGrit                       float32
	AvgFlow                       float32
	WorkoutFeel                   uint8
	WorkoutRpe                    uint8
	AvgSpo2                       uint8
	AvgStress                     uint8
	SdrrHrv                       uint8
	RmssdHrv                      uint8
	TotalFractionalAscent         uint8
	TotalFractionalDescent        uint8
	AvgCoreTemperature            uint16
	MinCoreTemperature            uint16
	MaxCoreTemperature            uint16
}

func (obj *Session) MsgNumber() uint16 {
	return uint16(MesgNum_Session)
}

func (obj *Session) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSession = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Selected bit is set for the current session.", Products: "", Example: "1"},
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Sesson end time.", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Event", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "session", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "EventType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "stop", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "StartTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Sint32, FieldId: 3, Name: "StartPositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Sint32, FieldId: 4, Name: "StartPositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Enum, FieldId: 6, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint32, FieldId: 7, Name: "TotalElapsedTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time (includes pauses)", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint32, FieldId: 8, Name: "TotalTimerTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timer Time (excludes pauses)", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint32, FieldId: 9, Name: "TotalDistance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint16, FieldId: 11, Name: "TotalCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	13:  {BaseType: messages.Uint16, FieldId: 13, Name: "TotalFatCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Uint16, FieldId: 14, Name: "AvgSpeed", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_speed", Scale: 1000, Offset: 1, Units: "m/s", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "total_distance / total_timer_time", Products: "", Example: "1"},
	15:  {BaseType: messages.Uint16, FieldId: 15, Name: "MaxSpeed", IsArray: false, ArrayLength: 0, Components: "enhanced_max_speed", Scale: 1000, Offset: 1, Units: "m/s", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	16:  {BaseType: messages.Uint8, FieldId: 16, Name: "AvgHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "average heart rate (excludes pause time)", Products: "", Example: "1"},
	17:  {BaseType: messages.Uint8, FieldId: 17, Name: "MaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	20:  {BaseType: messages.Uint16, FieldId: 20, Name: "AvgPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "20", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time", Products: "", Example: "1"},
	21:  {BaseType: messages.Uint16, FieldId: 21, Name: "MaxPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	22:  {BaseType: messages.Uint16, FieldId: 22, Name: "TotalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	23:  {BaseType: messages.Uint16, FieldId: 23, Name: "TotalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	24:  {BaseType: messages.Uint8, FieldId: 24, Name: "TotalTrainingEffect", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	25:  {BaseType: messages.Uint16, FieldId: 25, Name: "FirstLapIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	26:  {BaseType: messages.Uint16, FieldId: 26, Name: "NumLaps", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "26", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	27:  {BaseType: messages.Uint8, FieldId: 27, Name: "EventGroup", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "27", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	28:  {BaseType: messages.Enum, FieldId: 28, Name: "Trigger", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "28", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	29:  {BaseType: messages.Sint32, FieldId: 29, Name: "NecLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "29", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "North east corner latitude", Products: "", Example: "1"},
	30:  {BaseType: messages.Sint32, FieldId: 30, Name: "NecLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "30", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "North east corner longitude", Products: "", Example: "1"},
	31:  {BaseType: messages.Sint32, FieldId: 31, Name: "SwcLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "31", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "South west corner latitude", Products: "", Example: "1"},
	32:  {BaseType: messages.Sint32, FieldId: 32, Name: "SwcLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "South west corner longitude", Products: "", Example: "1"},
	33:  {BaseType: messages.Uint16, FieldId: 33, Name: "NumLengths", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "lengths", Bits: "33", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "# of lengths of swim pool", Products: "", Example: "1"},
	34:  {BaseType: messages.Uint16, FieldId: 34, Name: "NormalizedPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "34", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	35:  {BaseType: messages.Uint16, FieldId: 35, Name: "TrainingStressScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "tss", Bits: "35", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	36:  {BaseType: messages.Uint16, FieldId: 36, Name: "IntensityFactor", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "if", Bits: "36", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	37:  {BaseType: messages.Enum, FieldId: 37, Name: "LeftRightBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "37", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	38:  {BaseType: messages.Sint32, FieldId: 38, Name: "EndPositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "38", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	39:  {BaseType: messages.Sint32, FieldId: 39, Name: "EndPositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "39", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	41:  {BaseType: messages.Uint32, FieldId: 41, Name: "AvgStrokeCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "strokes/lap", Bits: "41", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	42:  {BaseType: messages.Uint16, FieldId: 42, Name: "AvgStrokeDistance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "42", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	43:  {BaseType: messages.Enum, FieldId: 43, Name: "SwimStroke", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "swim_stroke", Bits: "43", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	44:  {BaseType: messages.Uint16, FieldId: 44, Name: "PoolLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "44", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	45:  {BaseType: messages.Uint16, FieldId: 45, Name: "ThresholdPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "45", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	46:  {BaseType: messages.Enum, FieldId: 46, Name: "PoolLengthUnit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "46", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	47:  {BaseType: messages.Uint16, FieldId: 47, Name: "NumActiveLengths", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "lengths", Bits: "47", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "# of active lengths of swim pool", Products: "", Example: "1"},
	48:  {BaseType: messages.Uint32, FieldId: 48, Name: "TotalWork", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "J", Bits: "48", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	49:  {BaseType: messages.Uint16, FieldId: 49, Name: "AvgAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "49", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	50:  {BaseType: messages.Uint16, FieldId: 50, Name: "MaxAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_max_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "50", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	51:  {BaseType: messages.Uint8, FieldId: 51, Name: "GpsAccuracy", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "51", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	52:  {BaseType: messages.Sint16, FieldId: 52, Name: "AvgGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "52", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	53:  {BaseType: messages.Sint16, FieldId: 53, Name: "AvgPosGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "53", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	54:  {BaseType: messages.Sint16, FieldId: 54, Name: "AvgNegGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "54", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	55:  {BaseType: messages.Sint16, FieldId: 55, Name: "MaxPosGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "55", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	56:  {BaseType: messages.Sint16, FieldId: 56, Name: "MaxNegGrade", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "56", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	57:  {BaseType: messages.Sint8, FieldId: 57, Name: "AvgTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "57", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	58:  {BaseType: messages.Sint8, FieldId: 58, Name: "MaxTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "58", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	59:  {BaseType: messages.Uint32, FieldId: 59, Name: "TotalMovingTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "59", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	60:  {BaseType: messages.Sint16, FieldId: 60, Name: "AvgPosVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "60", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	61:  {BaseType: messages.Sint16, FieldId: 61, Name: "AvgNegVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "61", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	62:  {BaseType: messages.Sint16, FieldId: 62, Name: "MaxPosVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "62", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	63:  {BaseType: messages.Sint16, FieldId: 63, Name: "MaxNegVerticalSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "63", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	64:  {BaseType: messages.Uint8, FieldId: 64, Name: "MinHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "64", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	65:  {BaseType: messages.Uint32, FieldId: 65, Name: "TimeInHrZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "65", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	66:  {BaseType: messages.Uint32, FieldId: 66, Name: "TimeInSpeedZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "66", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	67:  {BaseType: messages.Uint32, FieldId: 67, Name: "TimeInCadenceZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "67", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	68:  {BaseType: messages.Uint32, FieldId: 68, Name: "TimeInPowerZone", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "68", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	69:  {BaseType: messages.Uint32, FieldId: 69, Name: "AvgLapTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "69", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	70:  {BaseType: messages.Uint16, FieldId: 70, Name: "BestLapIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "70", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	71:  {BaseType: messages.Uint16, FieldId: 71, Name: "MinAltitude", IsArray: false, ArrayLength: 0, Components: "enhanced_min_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "71", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	82:  {BaseType: messages.Uint16, FieldId: 82, Name: "PlayerScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "82", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	83:  {BaseType: messages.Uint16, FieldId: 83, Name: "OpponentScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "83", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	84:  {BaseType: messages.String, FieldId: 84, Name: "OpponentName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "84", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	85:  {BaseType: messages.Uint16, FieldId: 85, Name: "StrokeCount", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "85", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "stroke_type enum used as the index", Products: "", Example: "1"},
	86:  {BaseType: messages.Uint16, FieldId: 86, Name: "ZoneCount", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "86", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "zone number used as the index", Products: "", Example: "1"},
	87:  {BaseType: messages.Uint16, FieldId: 87, Name: "MaxBallSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m/s", Bits: "87", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	88:  {BaseType: messages.Uint16, FieldId: 88, Name: "AvgBallSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m/s", Bits: "88", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	89:  {BaseType: messages.Uint16, FieldId: 89, Name: "AvgVerticalOscillation", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mm", Bits: "89", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	90:  {BaseType: messages.Uint16, FieldId: 90, Name: "AvgStanceTimePercent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "90", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	91:  {BaseType: messages.Uint16, FieldId: 91, Name: "AvgStanceTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "ms", Bits: "91", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	92:  {BaseType: messages.Uint8, FieldId: 92, Name: "AvgFractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "92", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the avg_cadence", Products: "", Example: "1"},
	93:  {BaseType: messages.Uint8, FieldId: 93, Name: "MaxFractionalCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "rpm", Bits: "93", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the max_cadence", Products: "", Example: "1"},
	94:  {BaseType: messages.Uint8, FieldId: 94, Name: "TotalFractionalCycles", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "cycles", Bits: "94", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of the total_cycles", Products: "", Example: "1"},
	95:  {BaseType: messages.Uint16, FieldId: 95, Name: "AvgTotalHemoglobinConc", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "95", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Avg saturated and unsaturated hemoglobin", Products: "", Example: ""},
	96:  {BaseType: messages.Uint16, FieldId: 96, Name: "MinTotalHemoglobinConc", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "96", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min saturated and unsaturated hemoglobin", Products: "", Example: ""},
	97:  {BaseType: messages.Uint16, FieldId: 97, Name: "MaxTotalHemoglobinConc", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "g/dL", Bits: "97", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max saturated and unsaturated hemoglobin", Products: "", Example: ""},
	98:  {BaseType: messages.Uint16, FieldId: 98, Name: "AvgSaturatedHemoglobinPercent", IsArray: true, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "98", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Avg percentage of hemoglobin saturated with oxygen", Products: "", Example: ""},
	99:  {BaseType: messages.Uint16, FieldId: 99, Name: "MinSaturatedHemoglobinPercent", IsArray: true, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "99", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min percentage of hemoglobin saturated with oxygen", Products: "", Example: ""},
	100: {BaseType: messages.Uint16, FieldId: 100, Name: "MaxSaturatedHemoglobinPercent", IsArray: true, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "100", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max percentage of hemoglobin saturated with oxygen", Products: "", Example: ""},
	101: {BaseType: messages.Uint8, FieldId: 101, Name: "AvgLeftTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "101", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	102: {BaseType: messages.Uint8, FieldId: 102, Name: "AvgRightTorqueEffectiveness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "102", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	103: {BaseType: messages.Uint8, FieldId: 103, Name: "AvgLeftPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "103", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	104: {BaseType: messages.Uint8, FieldId: 104, Name: "AvgRightPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "104", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	105: {BaseType: messages.Uint8, FieldId: 105, Name: "AvgCombinedPedalSmoothness", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "105", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	110: {BaseType: messages.String, FieldId: 110, Name: "SportProfileName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "110", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Sport name from associated sport mesg", Products: "", Example: "16"},
	111: {BaseType: messages.Uint8, FieldId: 111, Name: "SportIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "111", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	112: {BaseType: messages.Uint32, FieldId: 112, Name: "TimeStanding", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "112", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total time spend in the standing position", Products: "", Example: ""},
	113: {BaseType: messages.Uint16, FieldId: 113, Name: "StandCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "113", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of transitions to the standing state", Products: "", Example: ""},
	114: {BaseType: messages.Sint8, FieldId: 114, Name: "AvgLeftPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "114", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average platform center offset Left", Products: "", Example: ""},
	115: {BaseType: messages.Sint8, FieldId: 115, Name: "AvgRightPco", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mm", Bits: "115", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average platform center offset Right", Products: "", Example: ""},
	116: {BaseType: messages.Uint8, FieldId: 116, Name: "AvgLeftPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "116", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left power phase angles. Indexes defined by power_phase_type.", Products: "", Example: ""},
	117: {BaseType: messages.Uint8, FieldId: 117, Name: "AvgLeftPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "117", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average left power phase peak angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	118: {BaseType: messages.Uint8, FieldId: 118, Name: "AvgRightPowerPhase", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "118", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right power phase angles. Data value indexes defined by power_phase_type.", Products: "", Example: ""},
	119: {BaseType: messages.Uint8, FieldId: 119, Name: "AvgRightPowerPhasePeak", IsArray: true, ArrayLength: 0, Components: "", Scale: 0.7111111, Offset: 1, Units: "degrees", Bits: "119", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average right power phase peak angles data value indexes defined by power_phase_type.", Products: "", Example: ""},
	120: {BaseType: messages.Uint16, FieldId: 120, Name: "AvgPowerPosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "120", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average power by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	121: {BaseType: messages.Uint16, FieldId: 121, Name: "MaxPowerPosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "121", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum power by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	122: {BaseType: messages.Uint8, FieldId: 122, Name: "AvgCadencePosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "122", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average cadence by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	123: {BaseType: messages.Uint8, FieldId: 123, Name: "MaxCadencePosition", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "rpm", Bits: "123", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum cadence by position. Data value indexes defined by rider_position_type.", Products: "", Example: ""},
	124: {BaseType: messages.Uint32, FieldId: 124, Name: "EnhancedAvgSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "124", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "total_distance / total_timer_time", Products: "", Example: "1"},
	125: {BaseType: messages.Uint32, FieldId: 125, Name: "EnhancedMaxSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "125", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	126: {BaseType: messages.Uint32, FieldId: 126, Name: "EnhancedAvgAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "126", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	127: {BaseType: messages.Uint32, FieldId: 127, Name: "EnhancedMinAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "127", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	128: {BaseType: messages.Uint32, FieldId: 128, Name: "EnhancedMaxAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "128", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	129: {BaseType: messages.Uint16, FieldId: 129, Name: "AvgLevMotorPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "129", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev average motor power during session", Products: "", Example: ""},
	130: {BaseType: messages.Uint16, FieldId: 130, Name: "MaxLevMotorPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "watts", Bits: "130", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev maximum motor power during session", Products: "", Example: ""},
	131: {BaseType: messages.Uint8, FieldId: 131, Name: "LevBatteryConsumption", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: 1, Units: "percent", Bits: "131", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "lev battery consumption during session", Products: "", Example: ""},
	132: {BaseType: messages.Uint16, FieldId: 132, Name: "AvgVerticalRatio", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "132", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	133: {BaseType: messages.Uint16, FieldId: 133, Name: "AvgStanceTimeBalance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "133", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	134: {BaseType: messages.Uint16, FieldId: 134, Name: "AvgStepLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mm", Bits: "134", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	137: {BaseType: messages.Uint8, FieldId: 137, Name: "TotalAnaerobicTrainingEffect", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "", Bits: "137", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	139: {BaseType: messages.Uint16, FieldId: 139, Name: "AvgVam", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "139", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	140: {BaseType: messages.Uint32, FieldId: 140, Name: "AvgDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "140", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	141: {BaseType: messages.Uint32, FieldId: 141, Name: "MaxDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "141", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	142: {BaseType: messages.Uint32, FieldId: 142, Name: "SurfaceInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "142", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time since end of last dive", Products: "", Example: ""},
	143: {BaseType: messages.Uint8, FieldId: 143, Name: "StartCns", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "143", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	144: {BaseType: messages.Uint8, FieldId: 144, Name: "EndCns", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "144", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	145: {BaseType: messages.Uint16, FieldId: 145, Name: "StartN2", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "145", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	146: {BaseType: messages.Uint16, FieldId: 146, Name: "EndN2", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "146", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	147: {BaseType: messages.Uint8, FieldId: 147, Name: "AvgRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "147", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	148: {BaseType: messages.Uint8, FieldId: 148, Name: "MaxRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_max_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "148", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	149: {BaseType: messages.Uint8, FieldId: 149, Name: "MinRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_min_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "149", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	150: {BaseType: messages.Sint8, FieldId: 150, Name: "MinTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "150", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	155: {BaseType: messages.Uint16, FieldId: 155, Name: "O2Toxicity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "OTUs", Bits: "155", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	156: {BaseType: messages.Uint32, FieldId: 156, Name: "DiveNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "156", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	168: {BaseType: messages.Sint32, FieldId: 168, Name: "TrainingLoadPeak", IsArray: false, ArrayLength: 0, Components: "", Scale: 65536, Offset: 1, Units: "", Bits: "168", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	169: {BaseType: messages.Uint16, FieldId: 169, Name: "EnhancedAvgRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "169", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	170: {BaseType: messages.Uint16, FieldId: 170, Name: "EnhancedMaxRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "170", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	180: {BaseType: messages.Uint16, FieldId: 180, Name: "EnhancedMinRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "", Bits: "180", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	181: {BaseType: messages.Float32, FieldId: 181, Name: "TotalGrit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kGrit", Bits: "181", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	182: {BaseType: messages.Float32, FieldId: 182, Name: "TotalFlow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Flow", Bits: "182", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	183: {BaseType: messages.Uint16, FieldId: 183, Name: "JumpCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "183", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	186: {BaseType: messages.Float32, FieldId: 186, Name: "AvgGrit", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kGrit", Bits: "186", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.", Products: "", Example: ""},
	187: {BaseType: messages.Float32, FieldId: 187, Name: "AvgFlow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "Flow", Bits: "187", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.", Products: "", Example: ""},
	192: {BaseType: messages.Uint8, FieldId: 192, Name: "WorkoutFeel", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "192", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "A 0-100 scale representing how a user felt while performing a workout. Low values are considered feeling bad, while high values are good.", Products: "", Example: ""},
	193: {BaseType: messages.Uint8, FieldId: 193, Name: "WorkoutRpe", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "193", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Common Borg CR10 / 0-10 RPE scale, multiplied 10x.. Aggregate score for all workouts in a single session.", Products: "", Example: ""},
	194: {BaseType: messages.Uint8, FieldId: 194, Name: "AvgSpo2", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "194", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average SPO2 for the monitoring session", Products: "", Example: ""},
	195: {BaseType: messages.Uint8, FieldId: 195, Name: "AvgStress", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "195", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average stress for the monitoring session", Products: "", Example: ""},
	197: {BaseType: messages.Uint8, FieldId: 197, Name: "SdrrHrv", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mS", Bits: "197", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Standard deviation of R-R interval (SDRR) - Heart rate variability measure most useful for wellness users.", Products: "", Example: ""},
	198: {BaseType: messages.Uint8, FieldId: 198, Name: "RmssdHrv", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mS", Bits: "198", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Root mean square successive difference (RMSSD) - Heart rate variability measure most useful for athletes", Products: "", Example: ""},
	199: {BaseType: messages.Uint8, FieldId: 199, Name: "TotalFractionalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "199", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of total_ascent", Products: "", Example: ""},
	200: {BaseType: messages.Uint8, FieldId: 200, Name: "TotalFractionalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "200", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "fractional part of total_descent", Products: "", Example: ""},
	208: {BaseType: messages.Uint16, FieldId: 208, Name: "AvgCoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "208", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	209: {BaseType: messages.Uint16, FieldId: 209, Name: "MinCoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "209", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	210: {BaseType: messages.Uint16, FieldId: 210, Name: "MaxCoreTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "210", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Session),
		fieldMapForSession,
		func() (registration.IFitMessage, error) {
			return &Session{}, nil
		},
	)
}
