package fitDecl

import strconv "strconv"

type GoalSource byte

const (
	GoalSource_Auto      GoalSource = 0
	GoalSource_Community GoalSource = 1
	GoalSource_User      GoalSource = 2
	GoalSource_Invalid   GoalSource = 255
)

var GoalSourceNames = map[GoalSource]string{
	GoalSource_Auto:      "Auto",
	GoalSource_Community: "Community",
	GoalSource_User:      "User",
}

func (k GoalSource) String() string {
	if uint(k) < uint(len(GoalSourceNames)) {
		if v, ok := GoalSourceNames[k]; ok {
			return v
		}
	}
	return "GoalSource(" + strconv.Itoa(int(k)) + ")"
}

var GoalSourceValues = map[string]GoalSource{
	"Auto":      GoalSource_Auto,
	"Community": GoalSource_Community,
	"User":      GoalSource_User,
}

func GoalSourceFromString(s string) GoalSource {
	if v, ok := GoalSourceValues[s]; ok {
		return v
	}
	return GoalSource(GoalSource_Invalid)
}
