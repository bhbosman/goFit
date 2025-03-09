package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type HsaEvent struct {
	Timestamp uint32
	EventId   uint8
}

func (obj *HsaEvent) MsgNumber() uint16 {
	return uint16(MesgNum_HsaEvent)
}

func (obj *HsaEvent) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaEvent = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint8, FieldId: 0, Name: "EventId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Event ID. Health SDK use only", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaEvent),
		fieldMapForHsaEvent,
		func() (registration.IFitMessage, error) {
			return &HsaEvent{}, nil
		},
	)
}
