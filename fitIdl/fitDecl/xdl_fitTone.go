package fitDecl

import strconv "strconv"

type Tone byte

const (
	Tone_Off            Tone = 0
	Tone_Tone           Tone = 1
	Tone_Vibrate        Tone = 2
	Tone_ToneAndVibrate Tone = 3
	Tone_Invalid        Tone = 255
)

var ToneNames = map[Tone]string{
	Tone_Off:            "Off",
	Tone_Tone:           "Tone",
	Tone_Vibrate:        "Vibrate",
	Tone_ToneAndVibrate: "ToneAndVibrate",
}

func (k Tone) String() string {
	if uint(k) < uint(len(ToneNames)) {
		if v, ok := ToneNames[k]; ok {
			return v
		}
	}
	return "Tone(" + strconv.Itoa(int(k)) + ")"
}

var ToneValues = map[string]Tone{
	"Off":            Tone_Off,
	"Tone":           Tone_Tone,
	"Vibrate":        Tone_Vibrate,
	"ToneAndVibrate": Tone_ToneAndVibrate,
}

func ToneFromString(s string) Tone {
	if v, ok := ToneValues[s]; ok {
		return v
	}
	return Tone(Tone_Invalid)
}
