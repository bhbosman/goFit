package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DiveSummary struct {
	Timestamp       uint32
	ReferenceMesg   MesgNum
	ReferenceIndex  MessageIndex
	AvgDepth        uint32
	MaxDepth        uint32
	SurfaceInterval uint32
	StartCns        uint8
	EndCns          uint8
	StartN2         uint16
	EndN2           uint16
	O2Toxicity      uint16
	DiveNumber      uint32
	BottomTime      uint32
	AvgPressureSac  uint16
	AvgVolumeSac    uint16
	AvgRmv          uint16
	DescentTime     uint32
	AscentTime      uint32
	AvgAscentRate   int32
	AvgDescentRate  uint32
	MaxAscentRate   uint32
	MaxDescentRate  uint32
	HangTime        uint32
}

func (obj *DiveSummary) MsgNumber() uint16 {
	return uint16(MesgNum_DiveSummary)
}

func (obj *DiveSummary) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDiveSummary = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "ReferenceMesg", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "ReferenceIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "AvgDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "MaxDepth", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "0 if above water", Products: "", Example: ""},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "SurfaceInterval", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time since end of last dive", Products: "", Example: ""},
	5:   {BaseType: messages.Uint8, FieldId: 5, Name: "StartCns", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Uint8, FieldId: 6, Name: "EndCns", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	7:   {BaseType: messages.Uint16, FieldId: 7, Name: "StartN2", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:   {BaseType: messages.Uint16, FieldId: 8, Name: "EndN2", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "percent", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "O2Toxicity", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "OTUs", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	10:  {BaseType: messages.Uint32, FieldId: 10, Name: "DiveNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	11:  {BaseType: messages.Uint32, FieldId: 11, Name: "BottomTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	12:  {BaseType: messages.Uint16, FieldId: 12, Name: "AvgPressureSac", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "bar/min", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average pressure-based surface air consumption", Products: "", Example: ""},
	13:  {BaseType: messages.Uint16, FieldId: 13, Name: "AvgVolumeSac", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "L/min", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average volumetric surface air consumption", Products: "", Example: ""},
	14:  {BaseType: messages.Uint16, FieldId: 14, Name: "AvgRmv", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "L/min", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average respiratory minute volume", Products: "", Example: ""},
	15:  {BaseType: messages.Uint32, FieldId: 15, Name: "DescentTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time to reach deepest level stop", Products: "", Example: ""},
	16:  {BaseType: messages.Uint32, FieldId: 16, Name: "AscentTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time after leaving bottom until reaching surface", Products: "", Example: ""},
	17:  {BaseType: messages.Sint32, FieldId: 17, Name: "AvgAscentRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average ascent rate, not including descents or stops", Products: "", Example: ""},
	22:  {BaseType: messages.Uint32, FieldId: 22, Name: "AvgDescentRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average descent rate, not including ascents or stops", Products: "", Example: ""},
	23:  {BaseType: messages.Uint32, FieldId: 23, Name: "MaxAscentRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum ascent rate", Products: "", Example: ""},
	24:  {BaseType: messages.Uint32, FieldId: 24, Name: "MaxDescentRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Maximum descent rate", Products: "", Example: ""},
	25:  {BaseType: messages.Uint32, FieldId: 25, Name: "HangTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time spent neither ascending nor descending", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DiveSummary),
		fieldMapForDiveSummary,
		func() (registration.IFitMessage, error) {
			return &DiveSummary{}, nil
		},
	)
}
