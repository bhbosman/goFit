package fitDecl

import strconv "strconv"

type AttitudeValidity uint16

const (
	AttitudeValidity_TrackAngleHeadingValid AttitudeValidity = 1
	AttitudeValidity_PitchValid             AttitudeValidity = 2
	AttitudeValidity_RollValid              AttitudeValidity = 4
	AttitudeValidity_LateralBodyAccelValid  AttitudeValidity = 8
	AttitudeValidity_NormalBodyAccelValid   AttitudeValidity = 16
	AttitudeValidity_TurnRateValid          AttitudeValidity = 32
	AttitudeValidity_HwFail                 AttitudeValidity = 64
	AttitudeValidity_MagInvalid             AttitudeValidity = 128
	AttitudeValidity_NoGps                  AttitudeValidity = 256
	AttitudeValidity_GpsInvalid             AttitudeValidity = 512
	AttitudeValidity_SolutionCoasting       AttitudeValidity = 1024
	AttitudeValidity_TrueTrackAngle         AttitudeValidity = 2048
	AttitudeValidity_MagneticHeading        AttitudeValidity = 4096
	AttitudeValidity_Invalid                AttitudeValidity = 65535
)

var AttitudeValidityNames = map[AttitudeValidity]string{
	AttitudeValidity_TrackAngleHeadingValid: "TrackAngleHeadingValid",
	AttitudeValidity_PitchValid:             "PitchValid",
	AttitudeValidity_RollValid:              "RollValid",
	AttitudeValidity_LateralBodyAccelValid:  "LateralBodyAccelValid",
	AttitudeValidity_NormalBodyAccelValid:   "NormalBodyAccelValid",
	AttitudeValidity_TurnRateValid:          "TurnRateValid",
	AttitudeValidity_HwFail:                 "HwFail",
	AttitudeValidity_MagInvalid:             "MagInvalid",
	AttitudeValidity_NoGps:                  "NoGps",
	AttitudeValidity_GpsInvalid:             "GpsInvalid",
	AttitudeValidity_SolutionCoasting:       "SolutionCoasting",
	AttitudeValidity_TrueTrackAngle:         "TrueTrackAngle",
	AttitudeValidity_MagneticHeading:        "MagneticHeading",
}

func (k AttitudeValidity) String() string {
	if uint(k) < uint(len(AttitudeValidityNames)) {
		if v, ok := AttitudeValidityNames[k]; ok {
			return v
		}
	}
	return "AttitudeValidity(" + strconv.Itoa(int(k)) + ")"
}

var AttitudeValidityValues = map[string]AttitudeValidity{
	"TrackAngleHeadingValid": AttitudeValidity_TrackAngleHeadingValid,
	"PitchValid":             AttitudeValidity_PitchValid,
	"RollValid":              AttitudeValidity_RollValid,
	"LateralBodyAccelValid":  AttitudeValidity_LateralBodyAccelValid,
	"NormalBodyAccelValid":   AttitudeValidity_NormalBodyAccelValid,
	"TurnRateValid":          AttitudeValidity_TurnRateValid,
	"HwFail":                 AttitudeValidity_HwFail,
	"MagInvalid":             AttitudeValidity_MagInvalid,
	"NoGps":                  AttitudeValidity_NoGps,
	"GpsInvalid":             AttitudeValidity_GpsInvalid,
	"SolutionCoasting":       AttitudeValidity_SolutionCoasting,
	"TrueTrackAngle":         AttitudeValidity_TrueTrackAngle,
	"MagneticHeading":        AttitudeValidity_MagneticHeading,
}

func AttitudeValidityFromString(s string) AttitudeValidity {
	if v, ok := AttitudeValidityValues[s]; ok {
		return v
	}
	return AttitudeValidity(AttitudeValidity_Invalid)
}
