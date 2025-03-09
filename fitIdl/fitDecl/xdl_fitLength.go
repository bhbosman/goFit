package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Length struct {
	MessageIndex               MessageIndex
	Timestamp                  uint32
	Event                      Event
	EventType                  EventType
	StartTime                  uint32
	TotalElapsedTime           uint32
	TotalTimerTime             uint32
	TotalStrokes               uint16
	AvgSpeed                   uint16
	SwimStroke                 SwimStroke
	AvgSwimmingCadence         uint8
	EventGroup                 uint8
	TotalCalories              uint16
	LengthType                 LengthType
	PlayerScore                uint16
	OpponentScore              uint16
	StrokeCount                uint16
	ZoneCount                  uint16
	EnhancedAvgRespirationRate uint16
	EnhancedMaxRespirationRate uint16
	AvgRespirationRate         uint8
	MaxRespirationRate         uint8
}

func (obj *Length) MsgNumber() uint16 {
	return uint16(MesgNum_Length)
}

func (obj *Length) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForLength = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Event", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "EventType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "StartTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "TotalElapsedTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "TotalTimerTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "TotalStrokes", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "strokes", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint16, FieldId: 6, Name: "AvgSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "SwimStroke", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "swim_stroke", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint8, FieldId: 9, Name: "AvgSwimmingCadence", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "strokes/min", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint8, FieldId: 10, Name: "EventGroup", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint16, FieldId: 11, Name: "TotalCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Enum, FieldId: 12, Name: "LengthType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	18:  {BaseType: messages.Uint16, FieldId: 18, Name: "PlayerScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "18", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	19:  {BaseType: messages.Uint16, FieldId: 19, Name: "OpponentScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	20:  {BaseType: messages.Uint16, FieldId: 20, Name: "StrokeCount", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "20", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "stroke_type enum used as the index", Products: "", Example: "1"},
	21:  {BaseType: messages.Uint16, FieldId: 21, Name: "ZoneCount", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "counts", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "zone number used as the index", Products: "", Example: "1"},
	22:  {BaseType: messages.Uint16, FieldId: 22, Name: "EnhancedAvgRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	23:  {BaseType: messages.Uint16, FieldId: 23, Name: "EnhancedMaxRespirationRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "Breaths/min", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	24:  {BaseType: messages.Uint8, FieldId: 24, Name: "AvgRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_avg_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	25:  {BaseType: messages.Uint8, FieldId: 25, Name: "MaxRespirationRate", IsArray: false, ArrayLength: 0, Components: "enhanced_max_respiration_rate", Scale: 1, Offset: 1, Units: "", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Length),
		fieldMapForLength,
		func() (registration.IFitMessage, error) {
			return &Length{}, nil
		},
	)
}
