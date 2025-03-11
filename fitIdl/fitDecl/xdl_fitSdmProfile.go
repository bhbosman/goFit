package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SdmProfile struct {
	MessageIndex      MessageIndex
	Enabled           bool
	SdmAntId          uint16
	SdmCalFactor      uint16
	Odometer          uint32
	SpeedSource       bool
	SdmAntIdTransType uint8
	OdometerRollover  uint8
}

func (obj *SdmProfile) MsgNumber() uint16 {
	return uint16(MesgNum_SdmProfile)
}

func (obj *SdmProfile) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSdmProfile = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Bool, FieldId: 0, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16z, FieldId: 1, Name: "SdmAntId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint16, FieldId: 2, Name: "SdmCalFactor", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "%", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "Odometer", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Bool, FieldId: 4, Name: "SpeedSource", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use footpod for speed source instead of GPS", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint8z, FieldId: 5, Name: "SdmAntIdTransType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint8, FieldId: 7, Name: "OdometerRollover", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Rollover counter that can be used to extend the odometer", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SdmProfile),
		fieldMapForSdmProfile,
		func() (registration.IFitMessage, error) {
			return &SdmProfile{}, nil
		},
	)
}
