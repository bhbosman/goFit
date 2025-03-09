package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type FileCreator struct {
	SoftwareVersion uint16
	HardwareVersion uint8
}

func (obj *FileCreator) MsgNumber() uint16 {
	return uint16(MesgNum_FileCreator)
}

func (obj *FileCreator) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForFileCreator = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.Uint16, FieldId: 0, Name: "SoftwareVersion", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1: {BaseType: messages.Uint8, FieldId: 1, Name: "HardwareVersion", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_FileCreator),
		fieldMapForFileCreator,
		func() (registration.IFitMessage, error) {
			return &FileCreator{}, nil
		},
	)
}
