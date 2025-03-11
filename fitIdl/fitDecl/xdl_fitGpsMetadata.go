package fitDecl

import registration "github.com/bhbosman/goFit/fitIdl/registration"
import messages "github.com/bhbosman/goFit/fitIdl/messages"

type GpsMetadata struct {
	Timestamp        uint32
	TimestampMs      uint16
	PositionLat      int32
	PositionLong     int32
	EnhancedAltitude uint32
	EnhancedSpeed    uint32
	Heading          uint16
	UtcTimestamp     uint32
	Velocity         int16
}

func (obj *GpsMetadata) MsgNumber() uint16 {
	return uint16(MesgNum_GpsMetadata)
}

func (obj *GpsMetadata) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForGpsMetadata = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Whole second part of the timestamp.", Products: "", Example: ""},
	0:   {BaseType: messages.Uint16, FieldId: 0, Name: "TimestampMs", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "ms", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Millisecond part of the timestamp.", Products: "", Example: ""},
	1:   {BaseType: messages.Sint32, FieldId: 1, Name: "PositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Sint32, FieldId: 2, Name: "PositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "EnhancedAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "EnhancedSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "Heading", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "degrees", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "UtcTimestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Used to correlate UTC to system time if the timestamp of the message is in system time. This UTC time is derived from the GPS data.", Products: "", Example: ""},
	7:   {BaseType: messages.Sint16, FieldId: 7, Name: "Velocity", IsArray: true, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m/s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_GpsMetadata),
		fieldMapForGpsMetadata,
		func() (registration.IFitMessage, error) {
			return &GpsMetadata{}, nil
		},
	)
}
