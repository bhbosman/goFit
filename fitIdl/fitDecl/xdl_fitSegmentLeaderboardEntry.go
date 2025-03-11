package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SegmentLeaderboardEntry struct {
	MessageIndex     MessageIndex
	Name             string
	Type             SegmentLeaderboardType
	GroupPrimaryKey  uint32
	ActivityId       uint32
	SegmentTime      uint32
	ActivityIdString string
}

func (obj *SegmentLeaderboardEntry) MsgNumber() uint16 {
	return uint16(MesgNum_SegmentLeaderboardEntry)
}

func (obj *SegmentLeaderboardEntry) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSegmentLeaderboardEntry = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.String, FieldId: 0, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Friendly name assigned to leader", Products: "", Example: "1"},
	1:   {BaseType: messages.Enum, FieldId: 1, Name: "Type", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Leader classification", Products: "", Example: "1"},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "GroupPrimaryKey", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Primary user ID of this leader", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint32, FieldId: 3, Name: "ActivityId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "ID of the activity associated with this leader time", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "SegmentTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "s", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Segment Time (includes pauses)", Products: "", Example: "1"},
	5:   {BaseType: messages.String, FieldId: 5, Name: "ActivityIdString", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "String version of the activity_id. 21 characters long, express in decimal", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SegmentLeaderboardEntry),
		fieldMapForSegmentLeaderboardEntry,
		func() (registration.IFitMessage, error) {
			return &SegmentLeaderboardEntry{}, nil
		},
	)
}
