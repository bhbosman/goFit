package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type WatchfaceSettings struct {
	MessageIndex MessageIndex
	Mode         WatchfaceMode
}

func (obj *WatchfaceSettings) MsgNumber() uint16 {
	return uint16(MesgNum_WatchfaceSettings)
}

func (obj *WatchfaceSettings) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForWatchfaceSettings = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "Mode", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_WatchfaceSettings),
		fieldMapForWatchfaceSettings,
		func() (registration.IFitMessage, error) {
			return &WatchfaceSettings{}, nil
		},
	)
}
