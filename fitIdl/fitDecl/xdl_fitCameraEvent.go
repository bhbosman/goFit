package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type CameraEvent struct {
	Timestamp         uint32
	TimestampMs       uint16
	CameraEventType   CameraEventType
	CameraFileUuid    string
	CameraOrientation CameraOrientationType
}

func (obj *CameraEvent) MsgNumber() uint16 {
	return uint16(MesgNum_CameraEvent)
}

func (obj *CameraEvent) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForCameraEvent = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp.", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the timestamp.", Products: "", Example: ""},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "CameraEventType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.String, FieldId: 2, Name: "CameraFileUuid", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Enum, FieldId: 3, Name: "CameraOrientation", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_CameraEvent),
		fieldMapForCameraEvent,
		func() (registration.IFitMessage, error) {
			return &CameraEvent{}, nil
		},
	)
}
