package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type SleepAssessment struct {
	CombinedAwakeScore       uint8
	AwakeTimeScore           uint8
	AwakeningsCountScore     uint8
	DeepSleepScore           uint8
	SleepDurationScore       uint8
	LightSleepScore          uint8
	OverallSleepScore        uint8
	SleepQualityScore        uint8
	SleepRecoveryScore       uint8
	RemSleepScore            uint8
	SleepRestlessnessScore   uint8
	AwakeningsCount          uint8
	InterruptionsScore       uint8
	AverageStressDuringSleep uint16
}

func (obj *SleepAssessment) MsgNumber() uint16 {
	return uint16(MesgNum_SleepAssessment)
}

func (obj *SleepAssessment) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForSleepAssessment = map[uint16]registration.DefinedFields{
	0:  {BaseType: messages.Uint8, FieldId: 0, Name: "CombinedAwakeScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Average of awake_time_score and awakenings_count_score. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	1:  {BaseType: messages.Uint8, FieldId: 1, Name: "AwakeTimeScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the total time spent awake between sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	2:  {BaseType: messages.Uint8, FieldId: 2, Name: "AwakeningsCountScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the number of awakenings that interrupt sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	3:  {BaseType: messages.Uint8, FieldId: 3, Name: "DeepSleepScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the amount of deep sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	4:  {BaseType: messages.Uint8, FieldId: 4, Name: "SleepDurationScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "4", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	5:  {BaseType: messages.Uint8, FieldId: 5, Name: "LightSleepScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the amount of light sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	6:  {BaseType: messages.Uint8, FieldId: 6, Name: "OverallSleepScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Total score that summarizes the overall quality of sleep, combining sleep duration and quality. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	7:  {BaseType: messages.Uint8, FieldId: 7, Name: "SleepQualityScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	8:  {BaseType: messages.Uint8, FieldId: 8, Name: "SleepRecoveryScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "8", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates stress and recovery during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	9:  {BaseType: messages.Uint8, FieldId: 9, Name: "RemSleepScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "9", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the amount of REM sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	10: {BaseType: messages.Uint8, FieldId: 10, Name: "SleepRestlessnessScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the amount of restlessness during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	11: {BaseType: messages.Uint8, FieldId: 11, Name: "AwakeningsCount", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "The number of awakenings during sleep.", Products: "", Example: ""},
	14: {BaseType: messages.Uint8, FieldId: 14, Name: "InterruptionsScore", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "14", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Score that evaluates the sleep interruptions. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.", Products: "", Example: ""},
	15: {BaseType: messages.Uint16, FieldId: 15, Name: "AverageStressDuringSleep", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "", Bits: "15", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Excludes stress during awake periods in the sleep window", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_SleepAssessment),
		fieldMapForSleepAssessment,
		func() (registration.IFitMessage, error) {
			return &SleepAssessment{}, nil
		},
	)
}
