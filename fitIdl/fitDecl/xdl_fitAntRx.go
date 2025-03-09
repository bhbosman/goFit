package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type AntRx struct {
	Timestamp           uint32
	FractionalTimestamp uint16
	MesgId              byte
	MesgData            byte
	ChannelNumber       uint8
	Data                byte
}

func (obj *AntRx) MsgNumber() uint16 {
	return uint16(MesgNum_AntRx)
}

func (obj *AntRx) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForAntRx = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "FractionalTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 32768, Offset: 1, Units: "s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Byte, FieldId: 1, Name: "MesgId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Byte, FieldId: 2, Name: "MesgData", IsArray: true, ArrayLength: 0, Components: "channel_number,data,data,data,data,data,data,data,data", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "9"},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "ChannelNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	4:   {BaseType: messages.Byte, FieldId: 4, Name: "Data", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "8"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_AntRx),
		fieldMapForAntRx,
		func() (registration.IFitMessage, error) {
			return &AntRx{}, nil
		},
	)
}
