package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type BikeProfile struct {
	MessageIndex             MessageIndex
	Name                     string
	Sport                    Sport
	SubSport                 SubSport
	Odometer                 uint32
	BikeSpdAntId             uint16
	BikeCadAntId             uint16
	BikeSpdcadAntId          uint16
	BikePowerAntId           uint16
	CustomWheelsize          uint16
	AutoWheelsize            uint16
	BikeWeight               uint16
	PowerCalFactor           uint16
	AutoWheelCal             bool
	AutoPowerZero            bool
	Id                       uint8
	SpdEnabled               bool
	CadEnabled               bool
	SpdcadEnabled            bool
	PowerEnabled             bool
	CrankLength              uint8
	Enabled                  bool
	BikeSpdAntIdTransType    uint8
	BikeCadAntIdTransType    uint8
	BikeSpdcadAntIdTransType uint8
	BikePowerAntIdTransType  uint8
	OdometerRollover         uint8
	FrontGearNum             uint8
	FrontGear                uint8
	RearGearNum              uint8
	RearGear                 uint8
	ShimanoDi2Enabled        bool
}

func (obj *BikeProfile) MsgNumber() uint16 {
	return uint16(MesgNum_BikeProfile)
}

func (obj *BikeProfile) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForBikeProfile = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.String, FieldId: 0, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "16"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "Odometer", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16z, FieldId: 4, Name: "BikeSpdAntId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint16z, FieldId: 5, Name: "BikeCadAntId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint16z, FieldId: 6, Name: "BikeSpdcadAntId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint16z, FieldId: 7, Name: "BikePowerAntId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint16, FieldId: 8, Name: "CustomWheelsize", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "AutoWheelsize", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint16, FieldId: 10, Name: "BikeWeight", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "kg", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Uint16, FieldId: 11, Name: "PowerCalFactor", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	12:  {BaseType: messages.Bool, FieldId: 12, Name: "AutoWheelCal", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	13:  {BaseType: messages.Bool, FieldId: 13, Name: "AutoPowerZero", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	14:  {BaseType: messages.Uint8, FieldId: 14, Name: "Id", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	15:  {BaseType: messages.Bool, FieldId: 15, Name: "SpdEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	16:  {BaseType: messages.Bool, FieldId: 16, Name: "CadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "16", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	17:  {BaseType: messages.Bool, FieldId: 17, Name: "SpdcadEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "17", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	18:  {BaseType: messages.Bool, FieldId: 18, Name: "PowerEnabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "18", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	19:  {BaseType: messages.Uint8, FieldId: 19, Name: "CrankLength", IsArray: false, ArrayLength: 0, Components: "", Scale: 2, Offset: -110, Units: "mm", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	20:  {BaseType: messages.Bool, FieldId: 20, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "20", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	21:  {BaseType: messages.Uint8z, FieldId: 21, Name: "BikeSpdAntIdTransType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	22:  {BaseType: messages.Uint8z, FieldId: 22, Name: "BikeCadAntIdTransType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	23:  {BaseType: messages.Uint8z, FieldId: 23, Name: "BikeSpdcadAntIdTransType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	24:  {BaseType: messages.Uint8z, FieldId: 24, Name: "BikePowerAntIdTransType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "24", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	37:  {BaseType: messages.Uint8, FieldId: 37, Name: "OdometerRollover", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "37", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Rollover counter that can be used to extend the odometer", Products: "", Example: "1"},
	38:  {BaseType: messages.Uint8z, FieldId: 38, Name: "FrontGearNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "38", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of front gears", Products: "", Example: "1"},
	39:  {BaseType: messages.Uint8z, FieldId: 39, Name: "FrontGear", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "39", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of teeth on each gear 0 is innermost", Products: "", Example: "1"},
	40:  {BaseType: messages.Uint8z, FieldId: 40, Name: "RearGearNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "40", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of rear gears", Products: "", Example: "1"},
	41:  {BaseType: messages.Uint8z, FieldId: 41, Name: "RearGear", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "41", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of teeth on each gear 0 is innermost", Products: "", Example: "1"},
	44:  {BaseType: messages.Bool, FieldId: 44, Name: "ShimanoDi2Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "44", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_BikeProfile),
		fieldMapForBikeProfile,
		func() (registration.IFitMessage, error) {
			return &BikeProfile{}, nil
		},
	)
}
