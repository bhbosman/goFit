package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ZonesTarget struct {
	MaxHeartRate             uint8
	ThresholdHeartRate       uint8
	FunctionalThresholdPower uint16
	HrCalcType               HrZoneCalc
	PwrCalcType              PwrZoneCalc
}

func (obj *ZonesTarget) MsgNumber() uint16 {
	return uint16(MesgNum_ZonesTarget)
}

func (obj *ZonesTarget) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForZonesTarget = map[uint16]registration.DefinedFields{
	1: {BaseType: messages.Uint8, FieldId: 1, Name: "MaxHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2: {BaseType: messages.Uint8, FieldId: 2, Name: "ThresholdHeartRate", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3: {BaseType: messages.Uint16, FieldId: 3, Name: "FunctionalThresholdPower", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5: {BaseType: messages.Enum, FieldId: 5, Name: "HrCalcType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7: {BaseType: messages.Enum, FieldId: 7, Name: "PwrCalcType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ZonesTarget),
		fieldMapForZonesTarget,
		func() (registration.IFitMessage, error) {
			return &ZonesTarget{}, nil
		},
	)
}
