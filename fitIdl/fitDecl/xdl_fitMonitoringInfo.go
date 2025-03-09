package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type MonitoringInfo struct {
	Timestamp            uint32
	LocalTimestamp       LocalDateTime
	ActivityType         ActivityType
	CyclesToDistance     uint16
	CyclesToCalories     uint16
	RestingMetabolicRate uint16
}

func (obj *MonitoringInfo) MsgNumber() uint16 {
	return uint16(MesgNum_MonitoringInfo)
}

func (obj *MonitoringInfo) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMonitoringInfo = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "LocalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "ActivityType", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "CyclesToDistance", IsArray: true, ArrayLength: 0, Components: "", Scale: 5000, Offset: 1, Units: "m/cycle", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indexed by activity_type", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "CyclesToCalories", IsArray: true, ArrayLength: 0, Components: "", Scale: 5000, Offset: 1, Units: "kcal/cycle", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indexed by activity_type", Products: "", Example: ""},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "RestingMetabolicRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal / day", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_MonitoringInfo),
		fieldMapForMonitoringInfo,
		func() (registration.IFitMessage, error) {
			return &MonitoringInfo{}, nil
		},
	)
}
