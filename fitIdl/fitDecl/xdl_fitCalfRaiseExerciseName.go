package fitDecl

import strconv "strconv"

type CalfRaiseExerciseName uint16

const (
	CalfRaiseExerciseName__3WayCalfRaise                     CalfRaiseExerciseName = 0
	CalfRaiseExerciseName__3WayWeightedCalfRaise             CalfRaiseExerciseName = 1
	CalfRaiseExerciseName__3WaySingleLegCalfRaise            CalfRaiseExerciseName = 2
	CalfRaiseExerciseName__3WayWeightedSingleLegCalfRaise    CalfRaiseExerciseName = 3
	CalfRaiseExerciseName_DonkeyCalfRaise                    CalfRaiseExerciseName = 4
	CalfRaiseExerciseName_WeightedDonkeyCalfRaise            CalfRaiseExerciseName = 5
	CalfRaiseExerciseName_SeatedCalfRaise                    CalfRaiseExerciseName = 6
	CalfRaiseExerciseName_WeightedSeatedCalfRaise            CalfRaiseExerciseName = 7
	CalfRaiseExerciseName_SeatedDumbbellToeRaise             CalfRaiseExerciseName = 8
	CalfRaiseExerciseName_SingleLegBentKneeCalfRaise         CalfRaiseExerciseName = 9
	CalfRaiseExerciseName_WeightedSingleLegBentKneeCalfRaise CalfRaiseExerciseName = 10
	CalfRaiseExerciseName_SingleLegDeclinePushUp             CalfRaiseExerciseName = 11
	CalfRaiseExerciseName_SingleLegDonkeyCalfRaise           CalfRaiseExerciseName = 12
	CalfRaiseExerciseName_WeightedSingleLegDonkeyCalfRaise   CalfRaiseExerciseName = 13
	CalfRaiseExerciseName_SingleLegHipRaiseWithKneeHold      CalfRaiseExerciseName = 14
	CalfRaiseExerciseName_SingleLegStandingCalfRaise         CalfRaiseExerciseName = 15
	CalfRaiseExerciseName_SingleLegStandingDumbbellCalfRaise CalfRaiseExerciseName = 16
	CalfRaiseExerciseName_StandingBarbellCalfRaise           CalfRaiseExerciseName = 17
	CalfRaiseExerciseName_StandingCalfRaise                  CalfRaiseExerciseName = 18
	CalfRaiseExerciseName_WeightedStandingCalfRaise          CalfRaiseExerciseName = 19
	CalfRaiseExerciseName_StandingDumbbellCalfRaise          CalfRaiseExerciseName = 20
	CalfRaiseExerciseName_Invalid                            CalfRaiseExerciseName = 65535
)

var CalfRaiseExerciseNameNames = map[CalfRaiseExerciseName]string{
	CalfRaiseExerciseName__3WayCalfRaise:                     "_3WayCalfRaise",
	CalfRaiseExerciseName__3WayWeightedCalfRaise:             "_3WayWeightedCalfRaise",
	CalfRaiseExerciseName__3WaySingleLegCalfRaise:            "_3WaySingleLegCalfRaise",
	CalfRaiseExerciseName__3WayWeightedSingleLegCalfRaise:    "_3WayWeightedSingleLegCalfRaise",
	CalfRaiseExerciseName_DonkeyCalfRaise:                    "DonkeyCalfRaise",
	CalfRaiseExerciseName_WeightedDonkeyCalfRaise:            "WeightedDonkeyCalfRaise",
	CalfRaiseExerciseName_SeatedCalfRaise:                    "SeatedCalfRaise",
	CalfRaiseExerciseName_WeightedSeatedCalfRaise:            "WeightedSeatedCalfRaise",
	CalfRaiseExerciseName_SeatedDumbbellToeRaise:             "SeatedDumbbellToeRaise",
	CalfRaiseExerciseName_SingleLegBentKneeCalfRaise:         "SingleLegBentKneeCalfRaise",
	CalfRaiseExerciseName_WeightedSingleLegBentKneeCalfRaise: "WeightedSingleLegBentKneeCalfRaise",
	CalfRaiseExerciseName_SingleLegDeclinePushUp:             "SingleLegDeclinePushUp",
	CalfRaiseExerciseName_SingleLegDonkeyCalfRaise:           "SingleLegDonkeyCalfRaise",
	CalfRaiseExerciseName_WeightedSingleLegDonkeyCalfRaise:   "WeightedSingleLegDonkeyCalfRaise",
	CalfRaiseExerciseName_SingleLegHipRaiseWithKneeHold:      "SingleLegHipRaiseWithKneeHold",
	CalfRaiseExerciseName_SingleLegStandingCalfRaise:         "SingleLegStandingCalfRaise",
	CalfRaiseExerciseName_SingleLegStandingDumbbellCalfRaise: "SingleLegStandingDumbbellCalfRaise",
	CalfRaiseExerciseName_StandingBarbellCalfRaise:           "StandingBarbellCalfRaise",
	CalfRaiseExerciseName_StandingCalfRaise:                  "StandingCalfRaise",
	CalfRaiseExerciseName_WeightedStandingCalfRaise:          "WeightedStandingCalfRaise",
	CalfRaiseExerciseName_StandingDumbbellCalfRaise:          "StandingDumbbellCalfRaise",
}

func (k CalfRaiseExerciseName) String() string {
	if uint(k) < uint(len(CalfRaiseExerciseNameNames)) {
		if v, ok := CalfRaiseExerciseNameNames[k]; ok {
			return v
		}
	}
	return "CalfRaiseExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var CalfRaiseExerciseNameValues = map[string]CalfRaiseExerciseName{
	"_3WayCalfRaise":                     CalfRaiseExerciseName__3WayCalfRaise,
	"_3WayWeightedCalfRaise":             CalfRaiseExerciseName__3WayWeightedCalfRaise,
	"_3WaySingleLegCalfRaise":            CalfRaiseExerciseName__3WaySingleLegCalfRaise,
	"_3WayWeightedSingleLegCalfRaise":    CalfRaiseExerciseName__3WayWeightedSingleLegCalfRaise,
	"DonkeyCalfRaise":                    CalfRaiseExerciseName_DonkeyCalfRaise,
	"WeightedDonkeyCalfRaise":            CalfRaiseExerciseName_WeightedDonkeyCalfRaise,
	"SeatedCalfRaise":                    CalfRaiseExerciseName_SeatedCalfRaise,
	"WeightedSeatedCalfRaise":            CalfRaiseExerciseName_WeightedSeatedCalfRaise,
	"SeatedDumbbellToeRaise":             CalfRaiseExerciseName_SeatedDumbbellToeRaise,
	"SingleLegBentKneeCalfRaise":         CalfRaiseExerciseName_SingleLegBentKneeCalfRaise,
	"WeightedSingleLegBentKneeCalfRaise": CalfRaiseExerciseName_WeightedSingleLegBentKneeCalfRaise,
	"SingleLegDeclinePushUp":             CalfRaiseExerciseName_SingleLegDeclinePushUp,
	"SingleLegDonkeyCalfRaise":           CalfRaiseExerciseName_SingleLegDonkeyCalfRaise,
	"WeightedSingleLegDonkeyCalfRaise":   CalfRaiseExerciseName_WeightedSingleLegDonkeyCalfRaise,
	"SingleLegHipRaiseWithKneeHold":      CalfRaiseExerciseName_SingleLegHipRaiseWithKneeHold,
	"SingleLegStandingCalfRaise":         CalfRaiseExerciseName_SingleLegStandingCalfRaise,
	"SingleLegStandingDumbbellCalfRaise": CalfRaiseExerciseName_SingleLegStandingDumbbellCalfRaise,
	"StandingBarbellCalfRaise":           CalfRaiseExerciseName_StandingBarbellCalfRaise,
	"StandingCalfRaise":                  CalfRaiseExerciseName_StandingCalfRaise,
	"WeightedStandingCalfRaise":          CalfRaiseExerciseName_WeightedStandingCalfRaise,
	"StandingDumbbellCalfRaise":          CalfRaiseExerciseName_StandingDumbbellCalfRaise,
}

func CalfRaiseExerciseNameFromString(s string) CalfRaiseExerciseName {
	if v, ok := CalfRaiseExerciseNameValues[s]; ok {
		return v
	}
	return CalfRaiseExerciseName(CalfRaiseExerciseName_Invalid)
}
