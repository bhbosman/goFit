package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type BloodPressure struct {
	Timestamp            uint32
	SystolicPressure     uint16
	DiastolicPressure    uint16
	MeanArterialPressure uint16
	Map3SampleMean       uint16
	MapMorningValues     uint16
	MapEveningValues     uint16
	HeartRate            uint8
	HeartRateType        HrType
	Status               BpStatus
	UserProfileIndex     MessageIndex
}

func (obj *BloodPressure) MsgNumber() uint16 {
	return uint16(MesgNum_BloodPressure)
}

func (obj *BloodPressure) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForBloodPressure = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "SystolicPressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mmHg", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "DiastolicPressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mmHg", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "MeanArterialPressure", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mmHg", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "Map3SampleMean", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mmHg", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "MapMorningValues", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mmHg", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "MapEveningValues", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "mmHg", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint8, FieldId: 6, Name: "HeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "bpm", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "HeartRateType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.Enum, FieldId: 8, Name: "Status", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	9:   {BaseType: messages.Enum, FieldId: 9, Name: "UserProfileIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Associates this blood pressure message to a user. This corresponds to the index of the user profile message in the blood pressure file.", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_BloodPressure),
		fieldMapForBloodPressure,
		func() (registration.IFitMessage, error) {
			return &BloodPressure{}, nil
		},
	)
}
