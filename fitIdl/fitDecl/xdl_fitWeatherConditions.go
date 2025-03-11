package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type WeatherConditions struct {
	Timestamp                uint32
	WeatherReport            WeatherReport
	Temperature              int8
	Condition                WeatherStatus
	WindDirection            uint16
	WindSpeed                uint16
	PrecipitationProbability uint8
	TemperatureFeelsLike     int8
	RelativeHumidity         uint8
	Location                 string
	ObservedAtTime           uint32
	ObservedLocationLat      int32
	ObservedLocationLong     int32
	DayOfWeek                DayOfWeek
	HighTemperature          int8
	LowTemperature           int8
}

func (obj *WeatherConditions) MsgNumber() uint16 {
	return uint16(MesgNum_WeatherConditions)
}

func (obj *WeatherConditions) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWeatherConditions = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "time of update for current conditions, else forecast time", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "WeatherReport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Current or forecast", Products: "", Example: "1"},
	1:   {BaseType: messages.Sint8, FieldId: 1, Name: "Temperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "Condition", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Corresponds to GSC Response weatherIcon field", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "WindDirection", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "degrees", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "WindSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint8, FieldId: 5, Name: "PrecipitationProbability", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "range 0-100", Products: "", Example: "1"},
	6:   {BaseType: messages.Sint8, FieldId: 6, Name: "TemperatureFeelsLike", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Heat Index if GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint8, FieldId: 7, Name: "RelativeHumidity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.String, FieldId: 8, Name: "Location", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "string corresponding to GCS response location string", Products: "", Example: "64"},
	9:   {BaseType: messages.Uint32, FieldId: 9, Name: "ObservedAtTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Sint32, FieldId: 10, Name: "ObservedLocationLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Sint32, FieldId: 11, Name: "ObservedLocationLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Enum, FieldId: 12, Name: "DayOfWeek", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	13:  {BaseType: messages.Sint8, FieldId: 13, Name: "HighTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Sint8, FieldId: 14, Name: "LowTemperature", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "C", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_WeatherConditions),
		fieldMapForWeatherConditions,
		func() (registration.IFitMessage, error) {
			return &WeatherConditions{}, nil
		},
	)
}
