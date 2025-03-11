package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type HrmProfile struct {
	MessageIndex      MessageIndex
	Enabled           bool
	HrmAntId          uint16
	LogHrv            bool
	HrmAntIdTransType uint8
}

func (obj *HrmProfile) MsgNumber() uint16 {
	return uint16(MesgNum_HrmProfile)
}

func (obj *HrmProfile) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHrmProfile = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Bool, FieldId: 0, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16z, FieldId: 1, Name: "HrmAntId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Bool, FieldId: 2, Name: "LogHrv", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint8z, FieldId: 3, Name: "HrmAntIdTransType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HrmProfile),
		fieldMapForHrmProfile,
		func() (registration.IFitMessage, error) {
			return &HrmProfile{}, nil
		},
	)
}
