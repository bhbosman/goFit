package fitDecl

import strconv "strconv"

type RowExerciseName uint16

const (
	RowExerciseName_BarbellStraightLegDeadliftToRow            RowExerciseName = 0
	RowExerciseName_CableRowStanding                           RowExerciseName = 1
	RowExerciseName_DumbbellRow                                RowExerciseName = 2
	RowExerciseName_ElevatedFeetInvertedRow                    RowExerciseName = 3
	RowExerciseName_WeightedElevatedFeetInvertedRow            RowExerciseName = 4
	RowExerciseName_FacePull                                   RowExerciseName = 5
	RowExerciseName_FacePullWithExternalRotation               RowExerciseName = 6
	RowExerciseName_InvertedRowWithFeetOnSwissBall             RowExerciseName = 7
	RowExerciseName_WeightedInvertedRowWithFeetOnSwissBall     RowExerciseName = 8
	RowExerciseName_KettlebellRow                              RowExerciseName = 9
	RowExerciseName_ModifiedInvertedRow                        RowExerciseName = 10
	RowExerciseName_WeightedModifiedInvertedRow                RowExerciseName = 11
	RowExerciseName_NeutralGripAlternatingDumbbellRow          RowExerciseName = 12
	RowExerciseName_OneArmBentOverRow                          RowExerciseName = 13
	RowExerciseName_OneLeggedDumbbellRow                       RowExerciseName = 14
	RowExerciseName_RenegadeRow                                RowExerciseName = 15
	RowExerciseName_ReverseGripBarbellRow                      RowExerciseName = 16
	RowExerciseName_RopeHandleCableRow                         RowExerciseName = 17
	RowExerciseName_SeatedCableRow                             RowExerciseName = 18
	RowExerciseName_SeatedDumbbellRow                          RowExerciseName = 19
	RowExerciseName_SingleArmCableRow                          RowExerciseName = 20
	RowExerciseName_SingleArmCableRowAndRotation               RowExerciseName = 21
	RowExerciseName_SingleArmInvertedRow                       RowExerciseName = 22
	RowExerciseName_WeightedSingleArmInvertedRow               RowExerciseName = 23
	RowExerciseName_SingleArmNeutralGripDumbbellRow            RowExerciseName = 24
	RowExerciseName_SingleArmNeutralGripDumbbellRowAndRotation RowExerciseName = 25
	RowExerciseName_SuspendedInvertedRow                       RowExerciseName = 26
	RowExerciseName_WeightedSuspendedInvertedRow               RowExerciseName = 27
	RowExerciseName_TBarRow                                    RowExerciseName = 28
	RowExerciseName_TowelGripInvertedRow                       RowExerciseName = 29
	RowExerciseName_WeightedTowelGripInvertedRow               RowExerciseName = 30
	RowExerciseName_UnderhandGripCableRow                      RowExerciseName = 31
	RowExerciseName_VGripCableRow                              RowExerciseName = 32
	RowExerciseName_WideGripSeatedCableRow                     RowExerciseName = 33
	RowExerciseName_Invalid                                    RowExerciseName = 65535
)

var RowExerciseNameNames = map[RowExerciseName]string{
	RowExerciseName_BarbellStraightLegDeadliftToRow:            "BarbellStraightLegDeadliftToRow",
	RowExerciseName_CableRowStanding:                           "CableRowStanding",
	RowExerciseName_DumbbellRow:                                "DumbbellRow",
	RowExerciseName_ElevatedFeetInvertedRow:                    "ElevatedFeetInvertedRow",
	RowExerciseName_WeightedElevatedFeetInvertedRow:            "WeightedElevatedFeetInvertedRow",
	RowExerciseName_FacePull:                                   "FacePull",
	RowExerciseName_FacePullWithExternalRotation:               "FacePullWithExternalRotation",
	RowExerciseName_InvertedRowWithFeetOnSwissBall:             "InvertedRowWithFeetOnSwissBall",
	RowExerciseName_WeightedInvertedRowWithFeetOnSwissBall:     "WeightedInvertedRowWithFeetOnSwissBall",
	RowExerciseName_KettlebellRow:                              "KettlebellRow",
	RowExerciseName_ModifiedInvertedRow:                        "ModifiedInvertedRow",
	RowExerciseName_WeightedModifiedInvertedRow:                "WeightedModifiedInvertedRow",
	RowExerciseName_NeutralGripAlternatingDumbbellRow:          "NeutralGripAlternatingDumbbellRow",
	RowExerciseName_OneArmBentOverRow:                          "OneArmBentOverRow",
	RowExerciseName_OneLeggedDumbbellRow:                       "OneLeggedDumbbellRow",
	RowExerciseName_RenegadeRow:                                "RenegadeRow",
	RowExerciseName_ReverseGripBarbellRow:                      "ReverseGripBarbellRow",
	RowExerciseName_RopeHandleCableRow:                         "RopeHandleCableRow",
	RowExerciseName_SeatedCableRow:                             "SeatedCableRow",
	RowExerciseName_SeatedDumbbellRow:                          "SeatedDumbbellRow",
	RowExerciseName_SingleArmCableRow:                          "SingleArmCableRow",
	RowExerciseName_SingleArmCableRowAndRotation:               "SingleArmCableRowAndRotation",
	RowExerciseName_SingleArmInvertedRow:                       "SingleArmInvertedRow",
	RowExerciseName_WeightedSingleArmInvertedRow:               "WeightedSingleArmInvertedRow",
	RowExerciseName_SingleArmNeutralGripDumbbellRow:            "SingleArmNeutralGripDumbbellRow",
	RowExerciseName_SingleArmNeutralGripDumbbellRowAndRotation: "SingleArmNeutralGripDumbbellRowAndRotation",
	RowExerciseName_SuspendedInvertedRow:                       "SuspendedInvertedRow",
	RowExerciseName_WeightedSuspendedInvertedRow:               "WeightedSuspendedInvertedRow",
	RowExerciseName_TBarRow:                                    "TBarRow",
	RowExerciseName_TowelGripInvertedRow:                       "TowelGripInvertedRow",
	RowExerciseName_WeightedTowelGripInvertedRow:               "WeightedTowelGripInvertedRow",
	RowExerciseName_UnderhandGripCableRow:                      "UnderhandGripCableRow",
	RowExerciseName_VGripCableRow:                              "VGripCableRow",
	RowExerciseName_WideGripSeatedCableRow:                     "WideGripSeatedCableRow",
}

func (k RowExerciseName) String() string {
	if uint(k) < uint(len(RowExerciseNameNames)) {
		if v, ok := RowExerciseNameNames[k]; ok {
			return v
		}
	}
	return "RowExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var RowExerciseNameValues = map[string]RowExerciseName{
	"BarbellStraightLegDeadliftToRow":            RowExerciseName_BarbellStraightLegDeadliftToRow,
	"CableRowStanding":                           RowExerciseName_CableRowStanding,
	"DumbbellRow":                                RowExerciseName_DumbbellRow,
	"ElevatedFeetInvertedRow":                    RowExerciseName_ElevatedFeetInvertedRow,
	"WeightedElevatedFeetInvertedRow":            RowExerciseName_WeightedElevatedFeetInvertedRow,
	"FacePull":                                   RowExerciseName_FacePull,
	"FacePullWithExternalRotation":               RowExerciseName_FacePullWithExternalRotation,
	"InvertedRowWithFeetOnSwissBall":             RowExerciseName_InvertedRowWithFeetOnSwissBall,
	"WeightedInvertedRowWithFeetOnSwissBall":     RowExerciseName_WeightedInvertedRowWithFeetOnSwissBall,
	"KettlebellRow":                              RowExerciseName_KettlebellRow,
	"ModifiedInvertedRow":                        RowExerciseName_ModifiedInvertedRow,
	"WeightedModifiedInvertedRow":                RowExerciseName_WeightedModifiedInvertedRow,
	"NeutralGripAlternatingDumbbellRow":          RowExerciseName_NeutralGripAlternatingDumbbellRow,
	"OneArmBentOverRow":                          RowExerciseName_OneArmBentOverRow,
	"OneLeggedDumbbellRow":                       RowExerciseName_OneLeggedDumbbellRow,
	"RenegadeRow":                                RowExerciseName_RenegadeRow,
	"ReverseGripBarbellRow":                      RowExerciseName_ReverseGripBarbellRow,
	"RopeHandleCableRow":                         RowExerciseName_RopeHandleCableRow,
	"SeatedCableRow":                             RowExerciseName_SeatedCableRow,
	"SeatedDumbbellRow":                          RowExerciseName_SeatedDumbbellRow,
	"SingleArmCableRow":                          RowExerciseName_SingleArmCableRow,
	"SingleArmCableRowAndRotation":               RowExerciseName_SingleArmCableRowAndRotation,
	"SingleArmInvertedRow":                       RowExerciseName_SingleArmInvertedRow,
	"WeightedSingleArmInvertedRow":               RowExerciseName_WeightedSingleArmInvertedRow,
	"SingleArmNeutralGripDumbbellRow":            RowExerciseName_SingleArmNeutralGripDumbbellRow,
	"SingleArmNeutralGripDumbbellRowAndRotation": RowExerciseName_SingleArmNeutralGripDumbbellRowAndRotation,
	"SuspendedInvertedRow":                       RowExerciseName_SuspendedInvertedRow,
	"WeightedSuspendedInvertedRow":               RowExerciseName_WeightedSuspendedInvertedRow,
	"TBarRow":                                    RowExerciseName_TBarRow,
	"TowelGripInvertedRow":                       RowExerciseName_TowelGripInvertedRow,
	"WeightedTowelGripInvertedRow":               RowExerciseName_WeightedTowelGripInvertedRow,
	"UnderhandGripCableRow":                      RowExerciseName_UnderhandGripCableRow,
	"VGripCableRow":                              RowExerciseName_VGripCableRow,
	"WideGripSeatedCableRow":                     RowExerciseName_WideGripSeatedCableRow,
}

func RowExerciseNameFromString(s string) RowExerciseName {
	if v, ok := RowExerciseNameValues[s]; ok {
		return v
	}
	return RowExerciseName(RowExerciseName_Invalid)
}
