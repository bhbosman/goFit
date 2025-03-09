package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type MaxMetData struct {
	UpdateTime     uint32
	Vo2Max         uint16
	Sport          Sport
	SubSport       SubSport
	MaxMetCategory MaxMetCategory
	CalibratedData bool
	HrSource       MaxMetHeartRateSource
	SpeedSource    MaxMetSpeedSource
}

func (obj *MaxMetData) MsgNumber() uint16 {
	return uint16(MesgNum_MaxMetData)
}

func (obj *MaxMetData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMaxMetData = map[uint16]registration.DefinedFields{
	0:  {BaseType: messages.Uint32, FieldId: 0, Name: "UpdateTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time maxMET and vo2 were calculated", Products: "", Example: ""},
	2:  {BaseType: messages.Uint16, FieldId: 2, Name: "Vo2Max", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "mL/kg/min", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:  {BaseType: messages.Enum, FieldId: 5, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:  {BaseType: messages.Enum, FieldId: 6, Name: "SubSport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	8:  {BaseType: messages.Enum, FieldId: 8, Name: "MaxMetCategory", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	9:  {BaseType: messages.Bool, FieldId: 9, Name: "CalibratedData", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates if calibrated data was used in the calculation", Products: "", Example: ""},
	12: {BaseType: messages.Enum, FieldId: 12, Name: "HrSource", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "12", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates if the estimate was obtained using a chest strap or wrist heart rate", Products: "", Example: ""},
	13: {BaseType: messages.Enum, FieldId: 13, Name: "SpeedSource", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "13", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indidcates if the estimate was obtained using onboard GPS or connected GPS", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_MaxMetData),
		fieldMapForMaxMetData,
		func() (registration.IFitMessage, error) {
			return &MaxMetData{}, nil
		},
	)
}
