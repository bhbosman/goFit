package fitDecl

import strconv "strconv"

type AutoSyncFrequency byte

const (
	AutoSyncFrequency_Never        AutoSyncFrequency = 0
	AutoSyncFrequency_Occasionally AutoSyncFrequency = 1
	AutoSyncFrequency_Frequent     AutoSyncFrequency = 2
	AutoSyncFrequency_OnceADay     AutoSyncFrequency = 3
	AutoSyncFrequency_Remote       AutoSyncFrequency = 4
	AutoSyncFrequency_Invalid      AutoSyncFrequency = 255
)

var AutoSyncFrequencyNames = map[AutoSyncFrequency]string{
	AutoSyncFrequency_Never:        "Never",
	AutoSyncFrequency_Occasionally: "Occasionally",
	AutoSyncFrequency_Frequent:     "Frequent",
	AutoSyncFrequency_OnceADay:     "OnceADay",
	AutoSyncFrequency_Remote:       "Remote",
}

func (k AutoSyncFrequency) String() string {
	if uint(k) < uint(len(AutoSyncFrequencyNames)) {
		if v, ok := AutoSyncFrequencyNames[k]; ok {
			return v
		}
	}
	return "AutoSyncFrequency(" + strconv.Itoa(int(k)) + ")"
}

var AutoSyncFrequencyValues = map[string]AutoSyncFrequency{
	"Never":        AutoSyncFrequency_Never,
	"Occasionally": AutoSyncFrequency_Occasionally,
	"Frequent":     AutoSyncFrequency_Frequent,
	"OnceADay":     AutoSyncFrequency_OnceADay,
	"Remote":       AutoSyncFrequency_Remote,
}

func AutoSyncFrequencyFromString(s string) AutoSyncFrequency {
	if v, ok := AutoSyncFrequencyValues[s]; ok {
		return v
	}
	return AutoSyncFrequency(AutoSyncFrequency_Invalid)
}
