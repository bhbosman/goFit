package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DiveAlarm struct {
	MessageIndex     MessageIndex
	Depth            uint32
	Time             int32
	Enabled          bool
	AlarmType        DiveAlarmType
	Sound            Tone
	DiveTypes        SubSport
	Id               uint32
	PopupEnabled     bool
	TriggerOnDescent bool
	TriggerOnAscent  bool
	Repeating        bool
	Speed            int32
}

func (obj *DiveAlarm) MsgNumber() uint16 {
	return uint16(MesgNum_DiveAlarm)
}

func (obj *DiveAlarm) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDiveAlarm = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Index of the alarm", Products: "", Example: ""},
	0:   {BaseType: messages.Uint32, FieldId: 0, Name: "Depth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Depth setting (m) for depth type alarms", Products: "", Example: ""},
	1:   {BaseType: messages.Sint32, FieldId: 1, Name: "Time", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time setting (s) for time type alarms", Products: "", Example: ""},
	2:   {BaseType: messages.Bool, FieldId: 2, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Enablement flag", Products: "", Example: ""},
	3:   {BaseType: messages.Enum, FieldId: 3, Name: "AlarmType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Alarm type setting", Products: "", Example: ""},
	4:   {BaseType: messages.Enum, FieldId: 4, Name: "Sound", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Tone and Vibe setting for the alarm", Products: "", Example: ""},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "DiveTypes", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Dive types the alarm will trigger on", Products: "", Example: ""},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "Id", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Alarm ID", Products: "", Example: ""},
	7:   {BaseType: messages.Bool, FieldId: 7, Name: "PopupEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Show a visible pop-up for this alarm", Products: "", Example: ""},
	8:   {BaseType: messages.Bool, FieldId: 8, Name: "TriggerOnDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Trigger the alarm on descent", Products: "", Example: ""},
	9:   {BaseType: messages.Bool, FieldId: 9, Name: "TriggerOnAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Trigger the alarm on ascent", Products: "", Example: ""},
	10:  {BaseType: messages.Bool, FieldId: 10, Name: "Repeating", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Repeat alarm each time threshold is crossed?", Products: "", Example: ""},
	11:  {BaseType: messages.Sint32, FieldId: 11, Name: "Speed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "mps", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Ascent/descent rate (mps) setting for speed type alarms", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DiveAlarm),
		fieldMapForDiveAlarm,
		func() (registration.IFitMessage, error) {
			return &DiveAlarm{}, nil
		},
	)
}
