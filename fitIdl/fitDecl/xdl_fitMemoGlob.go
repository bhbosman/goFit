package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type MemoGlob struct {
	PartIndex   uint32
	Memo        byte
	MesgNum     MesgNum
	ParentIndex MessageIndex
	FieldNum    uint8
	Data        uint8
}

func (obj *MemoGlob) MsgNumber() uint16 {
	return uint16(MesgNum_MemoGlob)
}

func (obj *MemoGlob) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForMemoGlob = map[uint16]registration.DefinedFields{
	250: {BaseType: messages.Uint32, FieldId: 250, Name: "PartIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "250", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Sequence number of memo blocks", Products: "", Example: ""},
	0:   {BaseType: messages.Byte, FieldId: 0, Name: "Memo", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Deprecated. Use data field.", Products: "", Example: ""},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "MesgNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Message Number of the parent message", Products: "", Example: ""},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "ParentIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Index of mesg that this glob is associated with.", Products: "", Example: ""},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "FieldNum", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Field within the parent that this glob is associated with", Products: "", Example: ""},
	4:   {BaseType: messages.Uint8z, FieldId: 4, Name: "Data", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Block of utf8 bytes. Note, mutltibyte characters may be split across adjoining memo_glob messages.", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_MemoGlob),
		fieldMapForMemoGlob,
		func() (registration.IFitMessage, error) {
			return &MemoGlob{}, nil
		},
	)
}
