package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type HsaConfigurationData struct {
	Timestamp uint32
	Data      byte
	DataSize  uint8
}

func (obj *HsaConfigurationData) MsgNumber() uint16 {
	return uint16(MesgNum_HsaConfigurationData)
}

func (obj *HsaConfigurationData) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForHsaConfigurationData = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Encoded configuration data", Products: "", Example: ""},
	0:   {BaseType: messages.Byte, FieldId: 0, Name: "Data", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Encoded configuration data. Health SDK use only", Products: "", Example: ""},
	1:   {BaseType: messages.Uint8, FieldId: 1, Name: "DataSize", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Size in bytes of data field", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_HsaConfigurationData),
		fieldMapForHsaConfigurationData,
		func() (registration.IFitMessage, error) {
			return &HsaConfigurationData{}, nil
		},
	)
}
