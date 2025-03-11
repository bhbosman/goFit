package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HrvStatusSummary struct {
	Timestamp             uint32
	WeeklyAverage         uint16
	LastNightAverage      uint16
	LastNight5MinHigh     uint16
	BaselineLowUpper      uint16
	BaselineBalancedLower uint16
	BaselineBalancedUpper uint16
	Status                HrvStatus
}

func (obj *HrvStatusSummary) MsgNumber() uint16 {
	return uint16(MesgNum_HrvStatusSummary)
}

func (obj *HrvStatusSummary) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHrvStatusSummary = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "WeeklyAverage", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "7 day RMSSD average over sleep", Products: "", Example: ""},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "LastNightAverage", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Last night RMSSD average over sleep", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "LastNight5MinHigh", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "5 minute high RMSSD value over sleep", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "BaselineLowUpper", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "3 week baseline, upper boundary of low HRV status", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "BaselineBalancedLower", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "3 week baseline, lower boundary of balanced HRV status", Products: "", Example: ""},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "BaselineBalancedUpper", IsArray: false, ArrayLength: 0, Components: "", Scale: 128, Offset: 1, Units: "ms", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "3 week baseline, upper boundary of balanced HRV status", Products: "", Example: ""},
	6:   {BaseType: messages.Enum, FieldId: 6, Name: "Status", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HrvStatusSummary),
		fieldMapForHrvStatusSummary,
		func() (registration.IFitMessage, error) {
			return &HrvStatusSummary{}, nil
		},
	)
}
