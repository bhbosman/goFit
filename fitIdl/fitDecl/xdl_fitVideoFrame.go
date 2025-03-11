package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type VideoFrame struct {
	Timestamp   uint32
	TimestampMs uint16
	FrameNumber uint32
}

func (obj *VideoFrame) MsgNumber() uint16 {
	return uint16(MesgNum_VideoFrame)
}

func (obj *VideoFrame) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForVideoFrame = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the timestamp.", Products: "", Example: ""},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "FrameNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Number of the frame that the timestamp and timestamp_ms correlate to", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_VideoFrame),
		fieldMapForVideoFrame,
		func() (registration.IFitMessage, error) {
			return &VideoFrame{}, nil
		},
	)
}
