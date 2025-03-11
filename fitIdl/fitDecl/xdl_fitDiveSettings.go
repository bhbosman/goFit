package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DiveSettings struct {
	Timestamp                 uint32
	MessageIndex              MessageIndex
	Name                      string
	Model                     TissueModelType
	GfLow                     uint8
	GfHigh                    uint8
	WaterType                 WaterType
	WaterDensity              float32
	Po2Warn                   uint8
	Po2Critical               uint8
	Po2Deco                   uint8
	SafetyStopEnabled         bool
	BottomDepth               float32
	BottomTime                uint32
	ApneaCountdownEnabled     bool
	ApneaCountdownTime        uint32
	BacklightMode             DiveBacklightMode
	BacklightBrightness       uint8
	BacklightTimeout          BacklightTimeout
	RepeatDiveInterval        uint16
	SafetyStopTime            uint16
	HeartRateSourceType       SourceType
	TravelGas                 MessageIndex
	CcrLowSetpointSwitchMode  CcrSetpointSwitchMode
	CcrLowSetpoint            uint8
	CcrLowSetpointDepth       uint32
	CcrHighSetpointSwitchMode CcrSetpointSwitchMode
	CcrHighSetpoint           uint8
	CcrHighSetpointDepth      uint32
	GasConsumptionDisplay     GasConsumptionRateType
	UpKeyEnabled              bool
	DiveSounds                Tone
	LastStopMultiple          uint8
	NoFlyTimeMode             NoFlyTimeMode
}

func (obj *DiveSettings) MsgNumber() uint16 {
	return uint16(MesgNum_DiveSettings)
}

func (obj *DiveSettings) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDiveSettings = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.String, FieldId: 0, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "Model", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Uint8, FieldId: 2, Name: "GfLow", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "GfHigh", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Enum, FieldId: 4, Name: "WaterType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:   {BaseType: messages.Float32, FieldId: 5, Name: "WaterDensity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kg/m^3", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fresh water is usually 1000; salt water is usually 1025", Products: "", Example: ""},
	6:   {BaseType: messages.Uint8, FieldId: 6, Name: "Po2Warn", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Typically 1.40", Products: "", Example: ""},
	7:   {BaseType: messages.Uint8, FieldId: 7, Name: "Po2Critical", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Typically 1.60", Products: "", Example: ""},
	8:   {BaseType: messages.Uint8, FieldId: 8, Name: "Po2Deco", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	9:   {BaseType: messages.Bool, FieldId: 9, Name: "SafetyStopEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	10:  {BaseType: messages.Float32, FieldId: 10, Name: "BottomDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	11:  {BaseType: messages.Uint32, FieldId: 11, Name: "BottomTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	12:  {BaseType: messages.Bool, FieldId: 12, Name: "ApneaCountdownEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	13:  {BaseType: messages.Uint32, FieldId: 13, Name: "ApneaCountdownTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	14:  {BaseType: messages.Enum, FieldId: 14, Name: "BacklightMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	15:  {BaseType: messages.Uint8, FieldId: 15, Name: "BacklightBrightness", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	16:  {BaseType: messages.Enum, FieldId: 16, Name: "BacklightTimeout", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	17:  {BaseType: messages.Uint16, FieldId: 17, Name: "RepeatDiveInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time between surfacing and ending the activity", Products: "", Example: ""},
	18:  {BaseType: messages.Uint16, FieldId: 18, Name: "SafetyStopTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "18", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time at safety stop (if enabled)", Products: "", Example: ""},
	19:  {BaseType: messages.Enum, FieldId: 19, Name: "HeartRateSourceType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	21:  {BaseType: messages.Enum, FieldId: 21, Name: "TravelGas", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Index of travel dive_gas message", Products: "", Example: ""},
	22:  {BaseType: messages.Enum, FieldId: 22, Name: "CcrLowSetpointSwitchMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "If low PO2 should be switched to automatically", Products: "", Example: ""},
	23:  {BaseType: messages.Uint8, FieldId: 23, Name: "CcrLowSetpoint", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Target PO2 when using low setpoint", Products: "", Example: ""},
	24:  {BaseType: messages.Uint32, FieldId: 24, Name: "CcrLowSetpointDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Depth to switch to low setpoint in automatic mode", Products: "", Example: ""},
	25:  {BaseType: messages.Enum, FieldId: 25, Name: "CcrHighSetpointSwitchMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "If high PO2 should be switched to automatically", Products: "", Example: ""},
	26:  {BaseType: messages.Uint8, FieldId: 26, Name: "CcrHighSetpoint", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "percent", Bits: "26", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Target PO2 when using high setpoint", Products: "", Example: ""},
	27:  {BaseType: messages.Uint32, FieldId: 27, Name: "CcrHighSetpointDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "27", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Depth to switch to high setpoint in automatic mode", Products: "", Example: ""},
	29:  {BaseType: messages.Enum, FieldId: 29, Name: "GasConsumptionDisplay", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "29", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Type of gas consumption rate to display. Some values are only valid if tank volume is known.", Products: "", Example: ""},
	30:  {BaseType: messages.Bool, FieldId: 30, Name: "UpKeyEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "30", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates whether the up key is enabled during dives", Products: "", Example: ""},
	35:  {BaseType: messages.Enum, FieldId: 35, Name: "DiveSounds", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "35", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Sounds and vibration enabled or disabled in-dive", Products: "", Example: ""},
	36:  {BaseType: messages.Uint8, FieldId: 36, Name: "LastStopMultiple", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "", Bits: "36", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Usually 1.0/1.5/2.0 representing 3/4.5/6m or 10/15/20ft", Products: "", Example: ""},
	37:  {BaseType: messages.Enum, FieldId: 37, Name: "NoFlyTimeMode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "37", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates which guidelines to use for no-fly surface interval.", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DiveSettings),
		fieldMapForDiveSettings,
		func() (registration.IFitMessage, error) {
			return &DiveSettings{}, nil
		},
	)
}
