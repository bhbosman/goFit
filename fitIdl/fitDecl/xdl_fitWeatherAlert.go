package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type WeatherAlert struct {
	Timestamp  uint32
	ReportId   string
	IssueTime  uint32
	ExpireTime uint32
	Severity   WeatherSeverity
	Type       WeatherSevereType
}

func (obj *WeatherAlert) MsgNumber() uint16 {
	return uint16(MesgNum_WeatherAlert)
}

func (obj *WeatherAlert) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWeatherAlert = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.String, FieldId: 0, Name: "ReportId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Unique identifier from GCS report ID string, length is 12", Products: "", Example: "12"},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "IssueTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time alert was issued", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "ExpireTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time alert expires", Products: "", Example: "1"},
	3:   {BaseType: messages.Enum, FieldId: 3, Name: "Severity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Warning, Watch, Advisory, Statement", Products: "", Example: "1"},
	4:   {BaseType: messages.Enum, FieldId: 4, Name: "Type", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Tornado, Severe Thunderstorm, etc.", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_WeatherAlert),
		fieldMapForWeatherAlert,
		func() (registration.IFitMessage, error) {
			return &WeatherAlert{}, nil
		},
	)
}
