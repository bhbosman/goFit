package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Capabilities struct {
	Languages             uint8
	Sports                SportBits0
	WorkoutsSupported     WorkoutCapabilities
	ConnectivitySupported ConnectivityCapabilities
}

func (obj *Capabilities) MsgNumber() uint16 {
	return uint16(MesgNum_Capabilities)
}

func (obj *Capabilities) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForCapabilities = map[uint16]registration.DefinedFields{
	0:  {BaseType: messages.Uint8z, FieldId: 0, Name: "Languages", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use language_bits_x types where x is index of array.", Products: "", Example: "4"},
	1:  {BaseType: messages.Enum, FieldId: 1, Name: "Sports", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Use sport_bits_x types where x is index of array.", Products: "", Example: "1"},
	21: {BaseType: messages.Enum, FieldId: 21, Name: "WorkoutsSupported", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	23: {BaseType: messages.Enum, FieldId: 23, Name: "ConnectivitySupported", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "23", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Capabilities),
		fieldMapForCapabilities,
		func() (registration.IFitMessage, error) {
			return &Capabilities{}, nil
		},
	)
}
