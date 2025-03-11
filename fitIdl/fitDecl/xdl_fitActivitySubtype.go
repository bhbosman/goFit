package fitDecl

import strconv "strconv"

type ActivitySubtype byte

const (
	ActivitySubtype_Generic       ActivitySubtype = 0
	ActivitySubtype_Treadmill     ActivitySubtype = 1
	ActivitySubtype_Street        ActivitySubtype = 2
	ActivitySubtype_Trail         ActivitySubtype = 3
	ActivitySubtype_Track         ActivitySubtype = 4
	ActivitySubtype_Spin          ActivitySubtype = 5
	ActivitySubtype_IndoorCycling ActivitySubtype = 6
	ActivitySubtype_Road          ActivitySubtype = 7
	ActivitySubtype_Mountain      ActivitySubtype = 8
	ActivitySubtype_Downhill      ActivitySubtype = 9
	ActivitySubtype_Recumbent     ActivitySubtype = 10
	ActivitySubtype_Cyclocross    ActivitySubtype = 11
	ActivitySubtype_HandCycling   ActivitySubtype = 12
	ActivitySubtype_TrackCycling  ActivitySubtype = 13
	ActivitySubtype_IndoorRowing  ActivitySubtype = 14
	ActivitySubtype_Elliptical    ActivitySubtype = 15
	ActivitySubtype_StairClimbing ActivitySubtype = 16
	ActivitySubtype_LapSwimming   ActivitySubtype = 17
	ActivitySubtype_OpenWater     ActivitySubtype = 18
	ActivitySubtype_All           ActivitySubtype = 254
	ActivitySubtype_Invalid       ActivitySubtype = 255
)

var ActivitySubtypeNames = map[ActivitySubtype]string{
	ActivitySubtype_Generic:       "Generic",
	ActivitySubtype_Treadmill:     "Treadmill",
	ActivitySubtype_Street:        "Street",
	ActivitySubtype_Trail:         "Trail",
	ActivitySubtype_Track:         "Track",
	ActivitySubtype_Spin:          "Spin",
	ActivitySubtype_IndoorCycling: "IndoorCycling",
	ActivitySubtype_Road:          "Road",
	ActivitySubtype_Mountain:      "Mountain",
	ActivitySubtype_Downhill:      "Downhill",
	ActivitySubtype_Recumbent:     "Recumbent",
	ActivitySubtype_Cyclocross:    "Cyclocross",
	ActivitySubtype_HandCycling:   "HandCycling",
	ActivitySubtype_TrackCycling:  "TrackCycling",
	ActivitySubtype_IndoorRowing:  "IndoorRowing",
	ActivitySubtype_Elliptical:    "Elliptical",
	ActivitySubtype_StairClimbing: "StairClimbing",
	ActivitySubtype_LapSwimming:   "LapSwimming",
	ActivitySubtype_OpenWater:     "OpenWater",
	ActivitySubtype_All:           "All",
}

func (k ActivitySubtype) String() string {
	if uint(k) < uint(len(ActivitySubtypeNames)) {
		if v, ok := ActivitySubtypeNames[k]; ok {
			return v
		}
	}
	return "ActivitySubtype(" + strconv.Itoa(int(k)) + ")"
}

var ActivitySubtypeValues = map[string]ActivitySubtype{
	"Generic":       ActivitySubtype_Generic,
	"Treadmill":     ActivitySubtype_Treadmill,
	"Street":        ActivitySubtype_Street,
	"Trail":         ActivitySubtype_Trail,
	"Track":         ActivitySubtype_Track,
	"Spin":          ActivitySubtype_Spin,
	"IndoorCycling": ActivitySubtype_IndoorCycling,
	"Road":          ActivitySubtype_Road,
	"Mountain":      ActivitySubtype_Mountain,
	"Downhill":      ActivitySubtype_Downhill,
	"Recumbent":     ActivitySubtype_Recumbent,
	"Cyclocross":    ActivitySubtype_Cyclocross,
	"HandCycling":   ActivitySubtype_HandCycling,
	"TrackCycling":  ActivitySubtype_TrackCycling,
	"IndoorRowing":  ActivitySubtype_IndoorRowing,
	"Elliptical":    ActivitySubtype_Elliptical,
	"StairClimbing": ActivitySubtype_StairClimbing,
	"LapSwimming":   ActivitySubtype_LapSwimming,
	"OpenWater":     ActivitySubtype_OpenWater,
	"All":           ActivitySubtype_All,
}

func ActivitySubtypeFromString(s string) ActivitySubtype {
	if v, ok := ActivitySubtypeValues[s]; ok {
		return v
	}
	return ActivitySubtype(ActivitySubtype_Invalid)
}
