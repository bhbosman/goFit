package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type AadAccelFeatures struct {
	Timestamp          uint32
	Time               uint16
	EnergyTotal        uint32
	ZeroCrossCnt       uint16
	Instance           uint8
	TimeAboveThreshold uint16
}

func (obj *AadAccelFeatures) MsgNumber() uint16 {
	return uint16(MesgNum_AadAccelFeatures)
}

func (obj *AadAccelFeatures) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForAadAccelFeatures = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "Time", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time interval length in seconds", Products: "", Example: ""},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "EnergyTotal", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total accelerometer energy in the interval", Products: "", Example: ""},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "ZeroCrossCnt", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Count of zero crossings", Products: "", Example: ""},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "Instance", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Instance ID of zero crossing algorithm", Products: "", Example: ""},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "TimeAboveThreshold", IsArray: false, ArrayLength: 0, Components: "", Scale: 25, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total accelerometer time above threshold in the interval", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_AadAccelFeatures),
		fieldMapForAadAccelFeatures,
		func() (registration.IFitMessage, error) {
			return &AadAccelFeatures{}, nil
		},
	)
}
