package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type Video struct {
	Url             string
	HostingProvider string
	Duration        uint32
}

func (obj *Video) MsgNumber() uint16 {
	return uint16(MesgNum_Video)
}

func (obj *Video) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForVideo = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.String, FieldId: 0, Name: "Url", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1: {BaseType: messages.String, FieldId: 1, Name: "HostingProvider", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2: {BaseType: messages.Uint32, FieldId: 2, Name: "Duration", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Playback time of video", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_Video),
		fieldMapForVideo,
		func() (registration.IFitMessage, error) {
			return &Video{}, nil
		},
	)
}
