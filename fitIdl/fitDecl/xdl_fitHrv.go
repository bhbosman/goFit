package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Hrv struct {
	Time uint16
}

func (obj *Hrv) MsgNumber() uint16 {
	return uint16(MesgNum_Hrv)
}

func (obj *Hrv) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHrv = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Uint16, FieldId: 0, Name: "Time", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Time between beats", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Hrv),
		fieldMapForHrv,
		func() (registration.IFitMessage, error) {
			return &Hrv{}, nil
		},
	)
}
