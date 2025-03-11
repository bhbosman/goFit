package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DeviceSettings struct {
	ActiveTimeZone                      uint8
	UtcOffset                           uint32
	TimeOffset                          uint32
	TimeMode                            TimeMode
	TimeZoneOffset                      int8
	BacklightMode                       BacklightMode
	ActivityTrackerEnabled              bool
	ClockTime                           uint32
	PagesEnabled                        uint16
	MoveAlertEnabled                    bool
	DateMode                            DateMode
	DisplayOrientation                  DisplayOrientation
	MountingSide                        Side
	DefaultPage                         uint16
	AutosyncMinSteps                    uint16
	AutosyncMinTime                     uint16
	LactateThresholdAutodetectEnabled   bool
	BleAutoUploadEnabled                bool
	AutoSyncFrequency                   AutoSyncFrequency
	AutoActivityDetect                  AutoActivityDetect
	NumberOfScreens                     uint8
	SmartNotificationDisplayOrientation DisplayOrientation
	TapInterface                        Switch
	TapSensitivity                      TapSensitivity
}

func (obj *DeviceSettings) MsgNumber() uint16 {
	return uint16(MesgNum_DeviceSettings)
}

func (obj *DeviceSettings) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDeviceSettings = map[uint16]registration.DefinedFields{
	0:   {BaseType: messages.Uint8, FieldId: 0, Name: "ActiveTimeZone", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Index into time zone arrays.", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "UtcOffset", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Offset from system time. Required to convert timestamp from system time to UTC.", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "TimeOffset", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Offset from system time.", Products: "", Example: "2"},
	4:   {BaseType: messages.Enum, FieldId: 4, Name: "TimeMode", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Display mode for the time", Products: "", Example: "2"},
	5:   {BaseType: messages.Sint8, FieldId: 5, Name: "TimeZoneOffset", IsArray: true, ArrayLength: 0, Components: "", Scale: 4, Offset: 1, Units: "hr", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "timezone offset in 1/4 hour increments", Products: "", Example: "2"},
	12:  {BaseType: messages.Enum, FieldId: 12, Name: "BacklightMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Mode for backlight", Products: "", Example: "1"},
	36:  {BaseType: messages.Bool, FieldId: 36, Name: "ActivityTrackerEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "36", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Enabled state of the activity tracker functionality", Products: "", Example: "1"},
	39:  {BaseType: messages.Uint32, FieldId: 39, Name: "ClockTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "39", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "UTC timestamp used to set the devices clock and date", Products: "", Example: "1"},
	40:  {BaseType: messages.Uint16, FieldId: 40, Name: "PagesEnabled", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "40", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Bitfield to configure enabled screens for each supported loop", Products: "", Example: "1"},
	46:  {BaseType: messages.Bool, FieldId: 46, Name: "MoveAlertEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "46", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Enabled state of the move alert", Products: "", Example: "1"},
	47:  {BaseType: messages.Enum, FieldId: 47, Name: "DateMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "47", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Display mode for the date", Products: "", Example: "1"},
	55:  {BaseType: messages.Enum, FieldId: 55, Name: "DisplayOrientation", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "55", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	56:  {BaseType: messages.Enum, FieldId: 56, Name: "MountingSide", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "56", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	57:  {BaseType: messages.Uint16, FieldId: 57, Name: "DefaultPage", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "57", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Bitfield to indicate one page as default for each supported loop", Products: "", Example: "1"},
	58:  {BaseType: messages.Uint16, FieldId: 58, Name: "AutosyncMinSteps", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "steps", Bits: "58", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Minimum steps before an autosync can occur", Products: "", Example: "1"},
	59:  {BaseType: messages.Uint16, FieldId: 59, Name: "AutosyncMinTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "minutes", Bits: "59", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Minimum minutes before an autosync can occur", Products: "", Example: "1"},
	80:  {BaseType: messages.Bool, FieldId: 80, Name: "LactateThresholdAutodetectEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "80", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Enable auto-detect setting for the lactate threshold feature.", Products: "", Example: ""},
	86:  {BaseType: messages.Bool, FieldId: 86, Name: "BleAutoUploadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "86", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Automatically upload using BLE", Products: "", Example: ""},
	89:  {BaseType: messages.Enum, FieldId: 89, Name: "AutoSyncFrequency", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "89", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Helps to conserve battery by changing modes", Products: "", Example: ""},
	90:  {BaseType: messages.Enum, FieldId: 90, Name: "AutoActivityDetect", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "90", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Allows setting specific activities auto-activity detect enabled/disabled settings", Products: "", Example: ""},
	94:  {BaseType: messages.Uint8, FieldId: 94, Name: "NumberOfScreens", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "94", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of screens configured to display", Products: "", Example: ""},
	95:  {BaseType: messages.Enum, FieldId: 95, Name: "SmartNotificationDisplayOrientation", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "95", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Smart Notification display orientation", Products: "", Example: ""},
	134: {BaseType: messages.Enum, FieldId: 134, Name: "TapInterface", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "134", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	174: {BaseType: messages.Enum, FieldId: 174, Name: "TapSensitivity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "174", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Used to hold the tap threshold setting", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DeviceSettings),
		fieldMapForDeviceSettings,
		func() (registration.IFitMessage, error) {
			return &DeviceSettings{}, nil
		},
	)
}
