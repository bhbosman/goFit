package fitDecl

import strconv "strconv"

type CardioExerciseName uint16

const (
	CardioExerciseName_BobAndWeaveCircle         CardioExerciseName = 0
	CardioExerciseName_WeightedBobAndWeaveCircle CardioExerciseName = 1
	CardioExerciseName_CardioCoreCrawl           CardioExerciseName = 2
	CardioExerciseName_WeightedCardioCoreCrawl   CardioExerciseName = 3
	CardioExerciseName_DoubleUnder               CardioExerciseName = 4
	CardioExerciseName_WeightedDoubleUnder       CardioExerciseName = 5
	CardioExerciseName_JumpRope                  CardioExerciseName = 6
	CardioExerciseName_WeightedJumpRope          CardioExerciseName = 7
	CardioExerciseName_JumpRopeCrossover         CardioExerciseName = 8
	CardioExerciseName_WeightedJumpRopeCrossover CardioExerciseName = 9
	CardioExerciseName_JumpRopeJog               CardioExerciseName = 10
	CardioExerciseName_WeightedJumpRopeJog       CardioExerciseName = 11
	CardioExerciseName_JumpingJacks              CardioExerciseName = 12
	CardioExerciseName_WeightedJumpingJacks      CardioExerciseName = 13
	CardioExerciseName_SkiMoguls                 CardioExerciseName = 14
	CardioExerciseName_WeightedSkiMoguls         CardioExerciseName = 15
	CardioExerciseName_SplitJacks                CardioExerciseName = 16
	CardioExerciseName_WeightedSplitJacks        CardioExerciseName = 17
	CardioExerciseName_SquatJacks                CardioExerciseName = 18
	CardioExerciseName_WeightedSquatJacks        CardioExerciseName = 19
	CardioExerciseName_TripleUnder               CardioExerciseName = 20
	CardioExerciseName_WeightedTripleUnder       CardioExerciseName = 21
	CardioExerciseName_Invalid                   CardioExerciseName = 65535
)

var CardioExerciseNameNames = map[CardioExerciseName]string{
	CardioExerciseName_BobAndWeaveCircle:         "BobAndWeaveCircle",
	CardioExerciseName_WeightedBobAndWeaveCircle: "WeightedBobAndWeaveCircle",
	CardioExerciseName_CardioCoreCrawl:           "CardioCoreCrawl",
	CardioExerciseName_WeightedCardioCoreCrawl:   "WeightedCardioCoreCrawl",
	CardioExerciseName_DoubleUnder:               "DoubleUnder",
	CardioExerciseName_WeightedDoubleUnder:       "WeightedDoubleUnder",
	CardioExerciseName_JumpRope:                  "JumpRope",
	CardioExerciseName_WeightedJumpRope:          "WeightedJumpRope",
	CardioExerciseName_JumpRopeCrossover:         "JumpRopeCrossover",
	CardioExerciseName_WeightedJumpRopeCrossover: "WeightedJumpRopeCrossover",
	CardioExerciseName_JumpRopeJog:               "JumpRopeJog",
	CardioExerciseName_WeightedJumpRopeJog:       "WeightedJumpRopeJog",
	CardioExerciseName_JumpingJacks:              "JumpingJacks",
	CardioExerciseName_WeightedJumpingJacks:      "WeightedJumpingJacks",
	CardioExerciseName_SkiMoguls:                 "SkiMoguls",
	CardioExerciseName_WeightedSkiMoguls:         "WeightedSkiMoguls",
	CardioExerciseName_SplitJacks:                "SplitJacks",
	CardioExerciseName_WeightedSplitJacks:        "WeightedSplitJacks",
	CardioExerciseName_SquatJacks:                "SquatJacks",
	CardioExerciseName_WeightedSquatJacks:        "WeightedSquatJacks",
	CardioExerciseName_TripleUnder:               "TripleUnder",
	CardioExerciseName_WeightedTripleUnder:       "WeightedTripleUnder",
}

func (k CardioExerciseName) String() string {
	if uint(k) < uint(len(CardioExerciseNameNames)) {
		if v, ok := CardioExerciseNameNames[k]; ok {
			return v
		}
	}
	return "CardioExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var CardioExerciseNameValues = map[string]CardioExerciseName{
	"BobAndWeaveCircle":         CardioExerciseName_BobAndWeaveCircle,
	"WeightedBobAndWeaveCircle": CardioExerciseName_WeightedBobAndWeaveCircle,
	"CardioCoreCrawl":           CardioExerciseName_CardioCoreCrawl,
	"WeightedCardioCoreCrawl":   CardioExerciseName_WeightedCardioCoreCrawl,
	"DoubleUnder":               CardioExerciseName_DoubleUnder,
	"WeightedDoubleUnder":       CardioExerciseName_WeightedDoubleUnder,
	"JumpRope":                  CardioExerciseName_JumpRope,
	"WeightedJumpRope":          CardioExerciseName_WeightedJumpRope,
	"JumpRopeCrossover":         CardioExerciseName_JumpRopeCrossover,
	"WeightedJumpRopeCrossover": CardioExerciseName_WeightedJumpRopeCrossover,
	"JumpRopeJog":               CardioExerciseName_JumpRopeJog,
	"WeightedJumpRopeJog":       CardioExerciseName_WeightedJumpRopeJog,
	"JumpingJacks":              CardioExerciseName_JumpingJacks,
	"WeightedJumpingJacks":      CardioExerciseName_WeightedJumpingJacks,
	"SkiMoguls":                 CardioExerciseName_SkiMoguls,
	"WeightedSkiMoguls":         CardioExerciseName_WeightedSkiMoguls,
	"SplitJacks":                CardioExerciseName_SplitJacks,
	"WeightedSplitJacks":        CardioExerciseName_WeightedSplitJacks,
	"SquatJacks":                CardioExerciseName_SquatJacks,
	"WeightedSquatJacks":        CardioExerciseName_WeightedSquatJacks,
	"TripleUnder":               CardioExerciseName_TripleUnder,
	"WeightedTripleUnder":       CardioExerciseName_WeightedTripleUnder,
}

func CardioExerciseNameFromString(s string) CardioExerciseName {
	if v, ok := CardioExerciseNameValues[s]; ok {
		return v
	}
	return CardioExerciseName(CardioExerciseName_Invalid)
}
