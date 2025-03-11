package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type Monitoring struct {
	Timestamp                    uint32
	DeviceIndex                  DeviceIndex
	Calories                     uint16
	Distance                     uint32
	ActiveTime                   uint32
	ActivityType                 ActivityType
	ActivitySubtype              ActivitySubtype
	ActivityLevel                ActivityLevel
	Distance16                   uint16
	Cycles16                     uint16
	ActiveTime16                 uint16
	LocalTimestamp               LocalDateTime
	Temperature                  int16
	TemperatureMin               int16
	TemperatureMax               int16
	ActivityTime                 uint16
	ActiveCalories               uint16
	CurrentActivityTypeIntensity byte
	TimestampMin8                uint8
	Timestamp16                  uint16
	HeartRate                    uint8
	Intensity                    uint8
	DurationMin                  uint16
	Duration                     uint32
	Ascent                       uint32
	Descent                      uint32
	ModerateActivityMinutes      uint16
	VigorousActivityMinutes      uint16
}

func (obj *Monitoring) MsgNumber() uint16 {
	return uint16(MesgNum_Monitoring)
}

func (obj *Monitoring) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMonitoring = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Must align to logging interval, for example, time must be 00:00:00 for daily log.", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "DeviceIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Associates this data to device_info message. Not required for file with single device (sensor).", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "Calories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Accumulated total calories. Maintained by MonitoringReader for each activity_type. See SDK documentation", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "Distance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "ActiveTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "ActivityType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Enum, FieldId: 6, Name: "ActivitySubtype", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "ActivityLevel", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:   {BaseType: messages.Uint16, FieldId: 8, Name: "Distance16", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "100 * m", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "Cycles16", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "2 * cycles (steps)", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint16, FieldId: 10, Name: "ActiveTime16", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Enum, FieldId: 11, Name: "LocalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Must align to logging interval, for example, time must be 00:00:00 for daily log.", Products: "", Example: "1"},
	12:  {BaseType: messages.Sint16, FieldId: 12, Name: "Temperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Avg temperature during the logging interval ended at timestamp", Products: "", Example: ""},
	14:  {BaseType: messages.Sint16, FieldId: 14, Name: "TemperatureMin", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Min temperature during the logging interval ended at timestamp", Products: "", Example: ""},
	15:  {BaseType: messages.Sint16, FieldId: 15, Name: "TemperatureMax", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "C", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Max temperature during the logging interval ended at timestamp", Products: "", Example: ""},
	16:  {BaseType: messages.Uint16, FieldId: 16, Name: "ActivityTime", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "minutes", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indexed using minute_activity_level enum", Products: "", Example: ""},
	19:  {BaseType: messages.Uint16, FieldId: 19, Name: "ActiveCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	24:  {BaseType: messages.Byte, FieldId: 24, Name: "CurrentActivityTypeIntensity", IsArray: false, ArrayLength: 0, Components: "activity_type,intensity", Scale: 1, Offset: 1, Units: "", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates single type / intensity for duration since last monitoring message.", Products: "", Example: ""},
	25:  {BaseType: messages.Uint8, FieldId: 25, Name: "TimestampMin8", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "min", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	26:  {BaseType: messages.Uint16, FieldId: 26, Name: "Timestamp16", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "26", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	27:  {BaseType: messages.Uint8, FieldId: 27, Name: "HeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "27", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	28:  {BaseType: messages.Uint8, FieldId: 28, Name: "Intensity", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "", Bits: "28", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	29:  {BaseType: messages.Uint16, FieldId: 29, Name: "DurationMin", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "min", Bits: "29", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	30:  {BaseType: messages.Uint32, FieldId: 30, Name: "Duration", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "30", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	31:  {BaseType: messages.Uint32, FieldId: 31, Name: "Ascent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "31", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	32:  {BaseType: messages.Uint32, FieldId: 32, Name: "Descent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	33:  {BaseType: messages.Uint16, FieldId: 33, Name: "ModerateActivityMinutes", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "minutes", Bits: "33", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	34:  {BaseType: messages.Uint16, FieldId: 34, Name: "VigorousActivityMinutes", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "minutes", Bits: "34", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Monitoring),
		fieldMapForMonitoring,
		func() (registration.IFitMessage, error) {
			return &Monitoring{}, nil
		},
	)
}
