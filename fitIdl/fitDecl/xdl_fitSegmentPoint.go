package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SegmentPoint struct {
	MessageIndex     MessageIndex
	PositionLat      int32
	PositionLong     int32
	Distance         uint32
	Altitude         uint16
	LeaderTime       uint32
	EnhancedAltitude uint32
}

func (obj *SegmentPoint) MsgNumber() uint16 {
	return uint16(MesgNum_SegmentPoint)
}

func (obj *SegmentPoint) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSegmentPoint = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Sint32, FieldId: 1, Name: "PositionLat", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Sint32, FieldId: 2, Name: "PositionLong", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "semicircles", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "Distance", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "m", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Accumulated distance along the segment at the described point", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint16, FieldId: 4, Name: "Altitude", IsArray: false, ArrayLength: 0, Components: "enhanced_altitude", Scale: 5, Offset: 500, Units: "m", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Accumulated altitude along the segment at the described point", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "LeaderTime", IsArray: true, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "EnhancedAltitude", IsArray: false, ArrayLength: 0, Components: "", Scale: 5, Offset: 500, Units: "m", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Accumulated altitude along the segment at the described point", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SegmentPoint),
		fieldMapForSegmentPoint,
		func() (registration.IFitMessage, error) {
			return &SegmentPoint{}, nil
		},
	)
}
