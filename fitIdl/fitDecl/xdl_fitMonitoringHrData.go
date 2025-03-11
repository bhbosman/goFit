package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type MonitoringHrData struct {
	Timestamp                  uint32
	RestingHeartRate           uint8
	CurrentDayRestingHeartRate uint8
}

func (obj *MonitoringHrData) MsgNumber() uint16 {
	return uint16(MesgNum_MonitoringHrData)
}

func (obj *MonitoringHrData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMonitoringHrData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Must align to logging interval, for example, time must be 00:00:00 for daily log.", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint8, FieldId: 0, Name: "RestingHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "7-day rolling average", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "CurrentDayRestingHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "RHR for today only. (Feeds into 7-day average)", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_MonitoringHrData),
		fieldMapForMonitoringHrData,
		func() (registration.IFitMessage, error) {
			return &MonitoringHrData{}, nil
		},
	)
}
