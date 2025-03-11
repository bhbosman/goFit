package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SlaveDevice struct {
	Manufacturer Manufacturer
}

func (obj *SlaveDevice) MsgNumber() uint16 {
	return uint16(MesgNum_SlaveDevice)
}

func (obj *SlaveDevice) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSlaveDevice = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Enum, FieldId: 0, Name: "Manufacturer", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SlaveDevice),
		fieldMapForSlaveDevice,
		func() (registration.IFitMessage, error) {
			return &SlaveDevice{}, nil
		},
	)
}
