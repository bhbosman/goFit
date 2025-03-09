package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type WeightScale struct {
	Timestamp         uint32
	Weight            Weight
	PercentFat        uint16
	PercentHydration  uint16
	VisceralFatMass   uint16
	BoneMass          uint16
	MuscleMass        uint16
	BasalMet          uint16
	PhysiqueRating    uint8
	ActiveMet         uint16
	MetabolicAge      uint8
	VisceralFatRating uint8
	UserProfileIndex  MessageIndex
	Bmi               uint16
}

func (obj *WeightScale) MsgNumber() uint16 {
	return uint16(MesgNum_WeightScale)
}

func (obj *WeightScale) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWeightScale = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Weight", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "kg", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "PercentFat", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "PercentHydration", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "%", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "VisceralFatMass", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "kg", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "BoneMass", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "kg", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "MuscleMass", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "kg", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint16, FieldId: 7, Name: "BasalMet", IsArray: false, ArrayLength: 0, Components: "", Scale: 4, Offset: 1, Units: "kcal/day", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint8, FieldId: 8, Name: "PhysiqueRating", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "ActiveMet", IsArray: false, ArrayLength: 0, Components: "", Scale: 4, Offset: 1, Units: "kcal/day", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "~4kJ per kcal, 0.25 allows max 16384 kcal", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint8, FieldId: 10, Name: "MetabolicAge", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "years", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint8, FieldId: 11, Name: "VisceralFatRating", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Enum, FieldId: 12, Name: "UserProfileIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Associates this weight scale message to a user. This corresponds to the index of the user profile message in the weight scale file.", Products: "", Example: "1"},
	13:  {BaseType: messages.Uint16, FieldId: 13, Name: "Bmi", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "kg/m^2", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_WeightScale),
		fieldMapForWeightScale,
		func() (registration.IFitMessage, error) {
			return &WeightScale{}, nil
		},
	)
}
