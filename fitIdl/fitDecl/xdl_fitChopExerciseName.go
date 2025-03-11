package fitDecl

import strconv "strconv"

type ChopExerciseName uint16

const (
	ChopExerciseName_CablePullThrough                   ChopExerciseName = 0
	ChopExerciseName_CableRotationalLift                ChopExerciseName = 1
	ChopExerciseName_CableWoodchop                      ChopExerciseName = 2
	ChopExerciseName_CrossChopToKnee                    ChopExerciseName = 3
	ChopExerciseName_WeightedCrossChopToKnee            ChopExerciseName = 4
	ChopExerciseName_DumbbellChop                       ChopExerciseName = 5
	ChopExerciseName_HalfKneelingRotation               ChopExerciseName = 6
	ChopExerciseName_WeightedHalfKneelingRotation       ChopExerciseName = 7
	ChopExerciseName_HalfKneelingRotationalChop         ChopExerciseName = 8
	ChopExerciseName_HalfKneelingRotationalReverseChop  ChopExerciseName = 9
	ChopExerciseName_HalfKneelingStabilityChop          ChopExerciseName = 10
	ChopExerciseName_HalfKneelingStabilityReverseChop   ChopExerciseName = 11
	ChopExerciseName_KneelingRotationalChop             ChopExerciseName = 12
	ChopExerciseName_KneelingRotationalReverseChop      ChopExerciseName = 13
	ChopExerciseName_KneelingStabilityChop              ChopExerciseName = 14
	ChopExerciseName_KneelingWoodchopper                ChopExerciseName = 15
	ChopExerciseName_MedicineBallWoodChops              ChopExerciseName = 16
	ChopExerciseName_PowerSquatChops                    ChopExerciseName = 17
	ChopExerciseName_WeightedPowerSquatChops            ChopExerciseName = 18
	ChopExerciseName_StandingRotationalChop             ChopExerciseName = 19
	ChopExerciseName_StandingSplitRotationalChop        ChopExerciseName = 20
	ChopExerciseName_StandingSplitRotationalReverseChop ChopExerciseName = 21
	ChopExerciseName_StandingStabilityReverseChop       ChopExerciseName = 22
	ChopExerciseName_Invalid                            ChopExerciseName = 65535
)

var ChopExerciseNameNames = map[ChopExerciseName]string{
	ChopExerciseName_CablePullThrough:                   "CablePullThrough",
	ChopExerciseName_CableRotationalLift:                "CableRotationalLift",
	ChopExerciseName_CableWoodchop:                      "CableWoodchop",
	ChopExerciseName_CrossChopToKnee:                    "CrossChopToKnee",
	ChopExerciseName_WeightedCrossChopToKnee:            "WeightedCrossChopToKnee",
	ChopExerciseName_DumbbellChop:                       "DumbbellChop",
	ChopExerciseName_HalfKneelingRotation:               "HalfKneelingRotation",
	ChopExerciseName_WeightedHalfKneelingRotation:       "WeightedHalfKneelingRotation",
	ChopExerciseName_HalfKneelingRotationalChop:         "HalfKneelingRotationalChop",
	ChopExerciseName_HalfKneelingRotationalReverseChop:  "HalfKneelingRotationalReverseChop",
	ChopExerciseName_HalfKneelingStabilityChop:          "HalfKneelingStabilityChop",
	ChopExerciseName_HalfKneelingStabilityReverseChop:   "HalfKneelingStabilityReverseChop",
	ChopExerciseName_KneelingRotationalChop:             "KneelingRotationalChop",
	ChopExerciseName_KneelingRotationalReverseChop:      "KneelingRotationalReverseChop",
	ChopExerciseName_KneelingStabilityChop:              "KneelingStabilityChop",
	ChopExerciseName_KneelingWoodchopper:                "KneelingWoodchopper",
	ChopExerciseName_MedicineBallWoodChops:              "MedicineBallWoodChops",
	ChopExerciseName_PowerSquatChops:                    "PowerSquatChops",
	ChopExerciseName_WeightedPowerSquatChops:            "WeightedPowerSquatChops",
	ChopExerciseName_StandingRotationalChop:             "StandingRotationalChop",
	ChopExerciseName_StandingSplitRotationalChop:        "StandingSplitRotationalChop",
	ChopExerciseName_StandingSplitRotationalReverseChop: "StandingSplitRotationalReverseChop",
	ChopExerciseName_StandingStabilityReverseChop:       "StandingStabilityReverseChop",
}

func (k ChopExerciseName) String() string {
	if uint(k) < uint(len(ChopExerciseNameNames)) {
		if v, ok := ChopExerciseNameNames[k]; ok {
			return v
		}
	}
	return "ChopExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var ChopExerciseNameValues = map[string]ChopExerciseName{
	"CablePullThrough":                   ChopExerciseName_CablePullThrough,
	"CableRotationalLift":                ChopExerciseName_CableRotationalLift,
	"CableWoodchop":                      ChopExerciseName_CableWoodchop,
	"CrossChopToKnee":                    ChopExerciseName_CrossChopToKnee,
	"WeightedCrossChopToKnee":            ChopExerciseName_WeightedCrossChopToKnee,
	"DumbbellChop":                       ChopExerciseName_DumbbellChop,
	"HalfKneelingRotation":               ChopExerciseName_HalfKneelingRotation,
	"WeightedHalfKneelingRotation":       ChopExerciseName_WeightedHalfKneelingRotation,
	"HalfKneelingRotationalChop":         ChopExerciseName_HalfKneelingRotationalChop,
	"HalfKneelingRotationalReverseChop":  ChopExerciseName_HalfKneelingRotationalReverseChop,
	"HalfKneelingStabilityChop":          ChopExerciseName_HalfKneelingStabilityChop,
	"HalfKneelingStabilityReverseChop":   ChopExerciseName_HalfKneelingStabilityReverseChop,
	"KneelingRotationalChop":             ChopExerciseName_KneelingRotationalChop,
	"KneelingRotationalReverseChop":      ChopExerciseName_KneelingRotationalReverseChop,
	"KneelingStabilityChop":              ChopExerciseName_KneelingStabilityChop,
	"KneelingWoodchopper":                ChopExerciseName_KneelingWoodchopper,
	"MedicineBallWoodChops":              ChopExerciseName_MedicineBallWoodChops,
	"PowerSquatChops":                    ChopExerciseName_PowerSquatChops,
	"WeightedPowerSquatChops":            ChopExerciseName_WeightedPowerSquatChops,
	"StandingRotationalChop":             ChopExerciseName_StandingRotationalChop,
	"StandingSplitRotationalChop":        ChopExerciseName_StandingSplitRotationalChop,
	"StandingSplitRotationalReverseChop": ChopExerciseName_StandingSplitRotationalReverseChop,
	"StandingStabilityReverseChop":       ChopExerciseName_StandingStabilityReverseChop,
}

func ChopExerciseNameFromString(s string) ChopExerciseName {
	if v, ok := ChopExerciseNameValues[s]; ok {
		return v
	}
	return ChopExerciseName(ChopExerciseName_Invalid)
}
