package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type ChronoShotSession struct {
	Timestamp         uint32
	MinSpeed          uint32
	MaxSpeed          uint32
	AvgSpeed          uint32
	ShotCount         uint16
	ProjectileType    ProjectileType
	GrainWeight       uint32
	StandardDeviation uint32
}

func (obj *ChronoShotSession) MsgNumber() uint16 {
	return uint16(MesgNum_ChronoShotSession)
}

func (obj *ChronoShotSession) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForChronoShotSession = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	0:   {BaseType: messages.Uint32, FieldId: 0, Name: "MinSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	1:   {BaseType: messages.Uint32, FieldId: 1, Name: "MaxSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	2:   {BaseType: messages.Uint32, FieldId: 2, Name: "AvgSpeed", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	3:   {BaseType: messages.Uint16, FieldId: 3, Name: "ShotCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	4:   {BaseType: messages.Enum, FieldId: 4, Name: "ProjectileType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	5:   {BaseType: messages.Uint32, FieldId: 5, Name: "GrainWeight", IsArray: false, ArrayLength: 0, Components: "", Scale: 10, Offset: 1, Units: "gr", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
	6:   {BaseType: messages.Uint32, FieldId: 6, Name: "StandardDeviation", IsArray: false, ArrayLength: 0, Components: "", Scale: 1000, Offset: 1, Units: "m/s", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_ChronoShotSession),
		fieldMapForChronoShotSession,
		func() (registration.IFitMessage, error) {
			return &ChronoShotSession{}, nil
		},
	)
}
