package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type UserProfile struct {
	MessageIndex               MessageIndex
	FriendlyName               string
	Gender                     Gender
	Age                        uint8
	Height                     uint8
	Weight                     uint16
	Language                   Language
	ElevSetting                DisplayMeasure
	WeightSetting              DisplayMeasure
	RestingHeartRate           uint8
	DefaultMaxRunningHeartRate uint8
	DefaultMaxBikingHeartRate  uint8
	DefaultMaxHeartRate        uint8
	HrSetting                  DisplayHeart
	SpeedSetting               DisplayMeasure
	DistSetting                DisplayMeasure
	PowerSetting               DisplayPower
	ActivityClass              ActivityClass
	PositionSetting            DisplayPosition
	TemperatureSetting         DisplayMeasure
	LocalId                    UserLocalId
	GlobalId                   byte
	WakeTime                   LocaltimeIntoDay
	SleepTime                  LocaltimeIntoDay
	HeightSetting              DisplayMeasure
	UserRunningStepLength      uint16
	UserWalkingStepLength      uint16
	DepthSetting               DisplayMeasure
	DiveCount                  uint32
}

func (obj *UserProfile) MsgNumber() uint16 {
	return uint16(MesgNum_UserProfile)
}

func (obj *UserProfile) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForUserProfile = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.String, FieldId: 0, Name: "FriendlyName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Used for Morning Report greeting", Products: "", Example: "16"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "Gender", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint8, FieldId: 2, Name: "Age", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "years", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "Height", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "Weight", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "kg", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Enum, FieldId: 5, Name: "Language", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Enum, FieldId: 6, Name: "ElevSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "WeightSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint8, FieldId: 8, Name: "RestingHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint8, FieldId: 9, Name: "DefaultMaxRunningHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint8, FieldId: 10, Name: "DefaultMaxBikingHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint8, FieldId: 11, Name: "DefaultMaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Enum, FieldId: 12, Name: "HrSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	13:  {BaseType: messages.Enum, FieldId: 13, Name: "SpeedSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Enum, FieldId: 14, Name: "DistSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	16:  {BaseType: messages.Enum, FieldId: 16, Name: "PowerSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	17:  {BaseType: messages.Enum, FieldId: 17, Name: "ActivityClass", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	18:  {BaseType: messages.Enum, FieldId: 18, Name: "PositionSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "18", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	21:  {BaseType: messages.Enum, FieldId: 21, Name: "TemperatureSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	22:  {BaseType: messages.Enum, FieldId: 22, Name: "LocalId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	23:  {BaseType: messages.Byte, FieldId: 23, Name: "GlobalId", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	28:  {BaseType: messages.Enum, FieldId: 28, Name: "WakeTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "28", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Typical wake time", Products: "", Example: ""},
	29:  {BaseType: messages.Enum, FieldId: 29, Name: "SleepTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "29", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Typical bed time", Products: "", Example: ""},
	30:  {BaseType: messages.Enum, FieldId: 30, Name: "HeightSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "30", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	31:  {BaseType: messages.Uint16, FieldId: 31, Name: "UserRunningStepLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "31", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "User defined running step length set to 0 for auto length", Products: "", Example: "1"},
	32:  {BaseType: messages.Uint16, FieldId: 32, Name: "UserWalkingStepLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "User defined walking step length set to 0 for auto length", Products: "", Example: "1"},
	47:  {BaseType: messages.Enum, FieldId: 47, Name: "DepthSetting", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "47", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	49:  {BaseType: messages.Uint32, FieldId: 49, Name: "DiveCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "49", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_UserProfile),
		fieldMapForUserProfile,
		func() (registration.IFitMessage, error) {
			return &UserProfile{}, nil
		},
	)
}
