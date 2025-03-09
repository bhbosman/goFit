package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Connectivity struct {
	BluetoothEnabled            bool
	BluetoothLeEnabled          bool
	AntEnabled                  bool
	Name                        string
	LiveTrackingEnabled         bool
	WeatherConditionsEnabled    bool
	WeatherAlertsEnabled        bool
	AutoActivityUploadEnabled   bool
	CourseDownloadEnabled       bool
	WorkoutDownloadEnabled      bool
	GpsEphemerisDownloadEnabled bool
	IncidentDetectionEnabled    bool
	GrouptrackEnabled           bool
}

func (obj *Connectivity) MsgNumber() uint16 {
	return uint16(MesgNum_Connectivity)
}

func (obj *Connectivity) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForConnectivity = map[uint16]registration.DefinedFields{
	0:  {BaseType: messages.Bool, FieldId: 0, Name: "BluetoothEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use Bluetooth for connectivity features", Products: "", Example: "1"},
	1:  {BaseType: messages.Bool, FieldId: 1, Name: "BluetoothLeEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use Bluetooth Low Energy for connectivity features", Products: "", Example: "1"},
	2:  {BaseType: messages.Bool, FieldId: 2, Name: "AntEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use ANT for connectivity features", Products: "", Example: "1"},
	3:  {BaseType: messages.String, FieldId: 3, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "24"},
	4:  {BaseType: messages.Bool, FieldId: 4, Name: "LiveTrackingEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:  {BaseType: messages.Bool, FieldId: 5, Name: "WeatherConditionsEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:  {BaseType: messages.Bool, FieldId: 6, Name: "WeatherAlertsEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:  {BaseType: messages.Bool, FieldId: 7, Name: "AutoActivityUploadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:  {BaseType: messages.Bool, FieldId: 8, Name: "CourseDownloadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:  {BaseType: messages.Bool, FieldId: 9, Name: "WorkoutDownloadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10: {BaseType: messages.Bool, FieldId: 10, Name: "GpsEphemerisDownloadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11: {BaseType: messages.Bool, FieldId: 11, Name: "IncidentDetectionEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12: {BaseType: messages.Bool, FieldId: 12, Name: "GrouptrackEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Connectivity),
		fieldMapForConnectivity,
		func() (registration.IFitMessage, error) {
			return &Connectivity{}, nil
		},
	)
}
