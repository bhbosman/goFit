package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SegmentFile struct {
	MessageIndex           MessageIndex
	FileUuid               string
	Enabled                bool
	UserProfilePrimaryKey  uint32
	LeaderType             SegmentLeaderboardType
	LeaderGroupPrimaryKey  uint32
	LeaderActivityId       uint32
	LeaderActivityIdString string
	DefaultRaceLeader      uint8
}

func (obj *SegmentFile) MsgNumber() uint16 {
	return uint16(MesgNum_SegmentFile)
}

func (obj *SegmentFile) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSegmentFile = map[uint16]registration.DefinedFields{
	254: {BaseType: messages.Enum, FieldId: 254, Name: "MessageIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "254", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.String, FieldId: 1, Name: "FileUuid", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "UUID of the segment file", Products: "", Example: "1"},
	3:   {BaseType: messages.Bool, FieldId: 3, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Enabled state of the segment file", Products: "", Example: "1"},
	4:   {BaseType: messages.Uint32, FieldId: 4, Name: "UserProfilePrimaryKey", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Primary key of the user that created the segment file", Products: "", Example: "1"},
	7:   {BaseType: messages.Enum, FieldId: 7, Name: "LeaderType", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Leader type of each leader in the segment file", Products: "", Example: "1"},
	8:   {BaseType: messages.Uint32, FieldId: 8, Name: "LeaderGroupPrimaryKey", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Group primary key of each leader in the segment file", Products: "", Example: "1"},
	9:   {BaseType: messages.Uint32, FieldId: 9, Name: "LeaderActivityId", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Activity ID of each leader in the segment file", Products: "", Example: "1"},
	10:  {BaseType: messages.String, FieldId: 10, Name: "LeaderActivityIdString", IsArray: true, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "String version of the activity ID of each leader in the segment file. 21 characters long for each ID, express in decimal", Products: "", Example: ""},
	11:  {BaseType: messages.Uint8, FieldId: 11, Name: "DefaultRaceLeader", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Index for the Leader Board entry selected as the default race participant", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SegmentFile),
		fieldMapForSegmentFile,
		func() (registration.IFitMessage, error) {
			return &SegmentFile{}, nil
		},
	)
}
