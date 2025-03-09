package fitDecl

import strconv "strconv"

type ShoulderStabilityExerciseName uint16

const (
	ShoulderStabilityExerciseName__90DegreeCableExternalRotation         ShoulderStabilityExerciseName = 0
	ShoulderStabilityExerciseName_BandExternalRotation                   ShoulderStabilityExerciseName = 1
	ShoulderStabilityExerciseName_BandInternalRotation                   ShoulderStabilityExerciseName = 2
	ShoulderStabilityExerciseName_BentArmLateralRaiseAndExternalRotation ShoulderStabilityExerciseName = 3
	ShoulderStabilityExerciseName_CableExternalRotation                  ShoulderStabilityExerciseName = 4
	ShoulderStabilityExerciseName_DumbbellFacePullWithExternalRotation   ShoulderStabilityExerciseName = 5
	ShoulderStabilityExerciseName_FloorIRaise                            ShoulderStabilityExerciseName = 6
	ShoulderStabilityExerciseName_WeightedFloorIRaise                    ShoulderStabilityExerciseName = 7
	ShoulderStabilityExerciseName_FloorTRaise                            ShoulderStabilityExerciseName = 8
	ShoulderStabilityExerciseName_WeightedFloorTRaise                    ShoulderStabilityExerciseName = 9
	ShoulderStabilityExerciseName_FloorYRaise                            ShoulderStabilityExerciseName = 10
	ShoulderStabilityExerciseName_WeightedFloorYRaise                    ShoulderStabilityExerciseName = 11
	ShoulderStabilityExerciseName_InclineIRaise                          ShoulderStabilityExerciseName = 12
	ShoulderStabilityExerciseName_WeightedInclineIRaise                  ShoulderStabilityExerciseName = 13
	ShoulderStabilityExerciseName_InclineLRaise                          ShoulderStabilityExerciseName = 14
	ShoulderStabilityExerciseName_WeightedInclineLRaise                  ShoulderStabilityExerciseName = 15
	ShoulderStabilityExerciseName_InclineTRaise                          ShoulderStabilityExerciseName = 16
	ShoulderStabilityExerciseName_WeightedInclineTRaise                  ShoulderStabilityExerciseName = 17
	ShoulderStabilityExerciseName_InclineWRaise                          ShoulderStabilityExerciseName = 18
	ShoulderStabilityExerciseName_WeightedInclineWRaise                  ShoulderStabilityExerciseName = 19
	ShoulderStabilityExerciseName_InclineYRaise                          ShoulderStabilityExerciseName = 20
	ShoulderStabilityExerciseName_WeightedInclineYRaise                  ShoulderStabilityExerciseName = 21
	ShoulderStabilityExerciseName_LyingExternalRotation                  ShoulderStabilityExerciseName = 22
	ShoulderStabilityExerciseName_SeatedDumbbellExternalRotation         ShoulderStabilityExerciseName = 23
	ShoulderStabilityExerciseName_StandingLRaise                         ShoulderStabilityExerciseName = 24
	ShoulderStabilityExerciseName_SwissBallIRaise                        ShoulderStabilityExerciseName = 25
	ShoulderStabilityExerciseName_WeightedSwissBallIRaise                ShoulderStabilityExerciseName = 26
	ShoulderStabilityExerciseName_SwissBallTRaise                        ShoulderStabilityExerciseName = 27
	ShoulderStabilityExerciseName_WeightedSwissBallTRaise                ShoulderStabilityExerciseName = 28
	ShoulderStabilityExerciseName_SwissBallWRaise                        ShoulderStabilityExerciseName = 29
	ShoulderStabilityExerciseName_WeightedSwissBallWRaise                ShoulderStabilityExerciseName = 30
	ShoulderStabilityExerciseName_SwissBallYRaise                        ShoulderStabilityExerciseName = 31
	ShoulderStabilityExerciseName_WeightedSwissBallYRaise                ShoulderStabilityExerciseName = 32
	ShoulderStabilityExerciseName_Invalid                                ShoulderStabilityExerciseName = 65535
)

var ShoulderStabilityExerciseNameNames = map[ShoulderStabilityExerciseName]string{
	ShoulderStabilityExerciseName__90DegreeCableExternalRotation:         "_90DegreeCableExternalRotation",
	ShoulderStabilityExerciseName_BandExternalRotation:                   "BandExternalRotation",
	ShoulderStabilityExerciseName_BandInternalRotation:                   "BandInternalRotation",
	ShoulderStabilityExerciseName_BentArmLateralRaiseAndExternalRotation: "BentArmLateralRaiseAndExternalRotation",
	ShoulderStabilityExerciseName_CableExternalRotation:                  "CableExternalRotation",
	ShoulderStabilityExerciseName_DumbbellFacePullWithExternalRotation:   "DumbbellFacePullWithExternalRotation",
	ShoulderStabilityExerciseName_FloorIRaise:                            "FloorIRaise",
	ShoulderStabilityExerciseName_WeightedFloorIRaise:                    "WeightedFloorIRaise",
	ShoulderStabilityExerciseName_FloorTRaise:                            "FloorTRaise",
	ShoulderStabilityExerciseName_WeightedFloorTRaise:                    "WeightedFloorTRaise",
	ShoulderStabilityExerciseName_FloorYRaise:                            "FloorYRaise",
	ShoulderStabilityExerciseName_WeightedFloorYRaise:                    "WeightedFloorYRaise",
	ShoulderStabilityExerciseName_InclineIRaise:                          "InclineIRaise",
	ShoulderStabilityExerciseName_WeightedInclineIRaise:                  "WeightedInclineIRaise",
	ShoulderStabilityExerciseName_InclineLRaise:                          "InclineLRaise",
	ShoulderStabilityExerciseName_WeightedInclineLRaise:                  "WeightedInclineLRaise",
	ShoulderStabilityExerciseName_InclineTRaise:                          "InclineTRaise",
	ShoulderStabilityExerciseName_WeightedInclineTRaise:                  "WeightedInclineTRaise",
	ShoulderStabilityExerciseName_InclineWRaise:                          "InclineWRaise",
	ShoulderStabilityExerciseName_WeightedInclineWRaise:                  "WeightedInclineWRaise",
	ShoulderStabilityExerciseName_InclineYRaise:                          "InclineYRaise",
	ShoulderStabilityExerciseName_WeightedInclineYRaise:                  "WeightedInclineYRaise",
	ShoulderStabilityExerciseName_LyingExternalRotation:                  "LyingExternalRotation",
	ShoulderStabilityExerciseName_SeatedDumbbellExternalRotation:         "SeatedDumbbellExternalRotation",
	ShoulderStabilityExerciseName_StandingLRaise:                         "StandingLRaise",
	ShoulderStabilityExerciseName_SwissBallIRaise:                        "SwissBallIRaise",
	ShoulderStabilityExerciseName_WeightedSwissBallIRaise:                "WeightedSwissBallIRaise",
	ShoulderStabilityExerciseName_SwissBallTRaise:                        "SwissBallTRaise",
	ShoulderStabilityExerciseName_WeightedSwissBallTRaise:                "WeightedSwissBallTRaise",
	ShoulderStabilityExerciseName_SwissBallWRaise:                        "SwissBallWRaise",
	ShoulderStabilityExerciseName_WeightedSwissBallWRaise:                "WeightedSwissBallWRaise",
	ShoulderStabilityExerciseName_SwissBallYRaise:                        "SwissBallYRaise",
	ShoulderStabilityExerciseName_WeightedSwissBallYRaise:                "WeightedSwissBallYRaise",
}

func (k ShoulderStabilityExerciseName) String() string {
	if uint(k) < uint(len(ShoulderStabilityExerciseNameNames)) {
		if v, ok := ShoulderStabilityExerciseNameNames[k]; ok {
			return v
		}
	}
	return "ShoulderStabilityExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var ShoulderStabilityExerciseNameValues = map[string]ShoulderStabilityExerciseName{
	"_90DegreeCableExternalRotation":         ShoulderStabilityExerciseName__90DegreeCableExternalRotation,
	"BandExternalRotation":                   ShoulderStabilityExerciseName_BandExternalRotation,
	"BandInternalRotation":                   ShoulderStabilityExerciseName_BandInternalRotation,
	"BentArmLateralRaiseAndExternalRotation": ShoulderStabilityExerciseName_BentArmLateralRaiseAndExternalRotation,
	"CableExternalRotation":                  ShoulderStabilityExerciseName_CableExternalRotation,
	"DumbbellFacePullWithExternalRotation":   ShoulderStabilityExerciseName_DumbbellFacePullWithExternalRotation,
	"FloorIRaise":                            ShoulderStabilityExerciseName_FloorIRaise,
	"WeightedFloorIRaise":                    ShoulderStabilityExerciseName_WeightedFloorIRaise,
	"FloorTRaise":                            ShoulderStabilityExerciseName_FloorTRaise,
	"WeightedFloorTRaise":                    ShoulderStabilityExerciseName_WeightedFloorTRaise,
	"FloorYRaise":                            ShoulderStabilityExerciseName_FloorYRaise,
	"WeightedFloorYRaise":                    ShoulderStabilityExerciseName_WeightedFloorYRaise,
	"InclineIRaise":                          ShoulderStabilityExerciseName_InclineIRaise,
	"WeightedInclineIRaise":                  ShoulderStabilityExerciseName_WeightedInclineIRaise,
	"InclineLRaise":                          ShoulderStabilityExerciseName_InclineLRaise,
	"WeightedInclineLRaise":                  ShoulderStabilityExerciseName_WeightedInclineLRaise,
	"InclineTRaise":                          ShoulderStabilityExerciseName_InclineTRaise,
	"WeightedInclineTRaise":                  ShoulderStabilityExerciseName_WeightedInclineTRaise,
	"InclineWRaise":                          ShoulderStabilityExerciseName_InclineWRaise,
	"WeightedInclineWRaise":                  ShoulderStabilityExerciseName_WeightedInclineWRaise,
	"InclineYRaise":                          ShoulderStabilityExerciseName_InclineYRaise,
	"WeightedInclineYRaise":                  ShoulderStabilityExerciseName_WeightedInclineYRaise,
	"LyingExternalRotation":                  ShoulderStabilityExerciseName_LyingExternalRotation,
	"SeatedDumbbellExternalRotation":         ShoulderStabilityExerciseName_SeatedDumbbellExternalRotation,
	"StandingLRaise":                         ShoulderStabilityExerciseName_StandingLRaise,
	"SwissBallIRaise":                        ShoulderStabilityExerciseName_SwissBallIRaise,
	"WeightedSwissBallIRaise":                ShoulderStabilityExerciseName_WeightedSwissBallIRaise,
	"SwissBallTRaise":                        ShoulderStabilityExerciseName_SwissBallTRaise,
	"WeightedSwissBallTRaise":                ShoulderStabilityExerciseName_WeightedSwissBallTRaise,
	"SwissBallWRaise":                        ShoulderStabilityExerciseName_SwissBallWRaise,
	"WeightedSwissBallWRaise":                ShoulderStabilityExerciseName_WeightedSwissBallWRaise,
	"SwissBallYRaise":                        ShoulderStabilityExerciseName_SwissBallYRaise,
	"WeightedSwissBallYRaise":                ShoulderStabilityExerciseName_WeightedSwissBallYRaise,
}

func ShoulderStabilityExerciseNameFromString(s string) ShoulderStabilityExerciseName {
	if v, ok := ShoulderStabilityExerciseNameValues[s]; ok {
		return v
	}
	return ShoulderStabilityExerciseName(ShoulderStabilityExerciseName_Invalid)
}
