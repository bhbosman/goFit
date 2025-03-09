package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type AviationAttitude struct {
	Timestamp             uint32
	TimestampMs           uint16
	SystemTime            uint32
	Pitch                 int16
	Roll                  int16
	AccelLateral          int16
	AccelNormal           int16
	TurnRate              int16
	Stage                 AttitudeStage
	AttitudeStageComplete uint8
	Track                 uint16
	Validity              AttitudeValidity
}

func (obj *AviationAttitude) MsgNumber() uint16 {
	return uint16(MesgNum_AviationAttitude)
}

func (obj *AviationAttitude) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForAviationAttitude = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Timestamp message was output", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Fractional part of timestamp, added to timestamp", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "SystemTime", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "System time associated with sample expressed in ms.", Products: "", Example: "1"},
	2:   {BaseType: messages.Sint16, FieldId: 2, Name: "Pitch", IsArray: true, ArrayLength: 0, Components: "", Scale: 10430.38, Offset: 1, Units: "radians", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Range -PI/2 to +PI/2", Products: "", Example: "1"},
	3:   {BaseType: messages.Sint16, FieldId: 3, Name: "Roll", IsArray: true, ArrayLength: 0, Components: "", Scale: 10430.38, Offset: 1, Units: "radians", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Range -PI to +PI", Products: "", Example: "1"},
	4:   {BaseType: messages.Sint16, FieldId: 4, Name: "AccelLateral", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m/s^2", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Range -78.4 to +78.4 (-8 Gs to 8 Gs)", Products: "", Example: "1"},
	5:   {BaseType: messages.Sint16, FieldId: 5, Name: "AccelNormal", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m/s^2", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Range -78.4 to +78.4 (-8 Gs to 8 Gs)", Products: "", Example: "1"},
	6:   {BaseType: messages.Sint16, FieldId: 6, Name: "TurnRate", IsArray: true, ArrayLength: 0, Components: "", Scale: 1024, Offset: 1, Units: "radians/second", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "Stage", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint8, FieldId: 8, Name: "AttitudeStageComplete", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "%", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The percent complete of the current attitude stage. Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it. Range - 100", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint16, FieldId: 9, Name: "Track", IsArray: true, ArrayLength: 0, Components: "", Scale: 10430.38, Offset: 1, Units: "radians", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Track Angle/Heading Range 0 - 2pi", Products: "", Example: "1"},
	10:  {BaseType: messages.Enum, FieldId: 10, Name: "Validity", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_AviationAttitude),
		fieldMapForAviationAttitude,
		func() (registration.IFitMessage, error) {
			return &AviationAttitude{}, nil
		},
	)
}
