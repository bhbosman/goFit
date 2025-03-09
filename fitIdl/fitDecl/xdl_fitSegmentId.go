package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SegmentId struct {
	Name                  string
	Uuid                  string
	Sport                 Sport
	Enabled               bool
	UserProfilePrimaryKey uint32
	DeviceId              uint32
	DefaultRaceLeader     uint8
	DeleteStatus          SegmentDeleteStatus
	SelectionType         SegmentSelectionType
}

func (obj *SegmentId) MsgNumber() uint16 {
	return uint16(MesgNum_SegmentId)
}

func (obj *SegmentId) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSegmentId = map[uint16]registration.DefinedFields{
	0: {BaseType: messages.String, FieldId: 0, Name: "Name", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Friendly name assigned to segment", Products: "", Example: "1"},
	1: {BaseType: messages.String, FieldId: 1, Name: "Uuid", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "UUID of the segment", Products: "", Example: "1"},
	2: {BaseType: messages.Enum, FieldId: 2, Name: "Sport", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Sport associated with the segment", Products: "", Example: "1"},
	3: {BaseType: messages.Bool, FieldId: 3, Name: "Enabled", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Segment enabled for evaluation", Products: "", Example: "1"},
	4: {BaseType: messages.Uint32, FieldId: 4, Name: "UserProfilePrimaryKey", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Primary key of the user that created the segment", Products: "", Example: "1"},
	5: {BaseType: messages.Uint32, FieldId: 5, Name: "DeviceId", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "ID of the device that created the segment", Products: "", Example: "1"},
	6: {BaseType: messages.Uint8, FieldId: 6, Name: "DefaultRaceLeader", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Index for the Leader Board entry selected as the default race participant", Products: "", Example: "1"},
	7: {BaseType: messages.Enum, FieldId: 7, Name: "DeleteStatus", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates if any segments should be deleted", Products: "", Example: "1"},
	8: {BaseType: messages.Enum, FieldId: 8, Name: "SelectionType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates how the segment was selected to be sent to the device", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SegmentId),
		fieldMapForSegmentId,
		func() (registration.IFitMessage, error) {
			return &SegmentId{}, nil
		},
	)
}
