package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SplitSummary struct {
	MessageIndex    MessageIndex
	SplitType       SplitType
	NumSplits       uint16
	TotalTimerTime  uint32
	TotalDistance   uint32
	AvgSpeed        uint32
	MaxSpeed        uint32
	TotalAscent     uint16
	TotalDescent    uint16
	AvgHeartRate    uint8
	MaxHeartRate    uint8
	AvgVertSpeed    int32
	TotalCalories   uint32
	TotalMovingTime uint32
}

func (obj *SplitSummary) MsgNumber() uint16 {
	return uint16(MesgNum_SplitSummary)
}

func (obj *SplitSummary) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSplitSummary = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "SplitType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "NumSplits", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "TotalTimerTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "TotalDistance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "AvgSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	7:   {BaseType: messages.Uint32, FieldId: 7, Name: "MaxSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:   {BaseType: messages.Uint16, FieldId: 8, Name: "TotalAscent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "TotalDescent", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "m", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	10:  {BaseType: messages.Uint8, FieldId: 10, Name: "AvgHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	11:  {BaseType: messages.Uint8, FieldId: 11, Name: "MaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	12:  {BaseType: messages.Sint32, FieldId: 12, Name: "AvgVertSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	13:  {BaseType: messages.Uint32, FieldId: 13, Name: "TotalCalories", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "kcal", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	77:  {BaseType: messages.Uint32, FieldId: 77, Name: "TotalMovingTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "77", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SplitSummary),
		fieldMapForSplitSummary,
		func() (registration.IFitMessage, error) {
			return &SplitSummary{}, nil
		},
	)
}
