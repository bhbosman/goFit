package fitDecl

import strconv "strconv"

type SportEvent byte

const (
	SportEvent_Uncategorized  SportEvent = 0
	SportEvent_Geocaching     SportEvent = 1
	SportEvent_Fitness        SportEvent = 2
	SportEvent_Recreation     SportEvent = 3
	SportEvent_Race           SportEvent = 4
	SportEvent_SpecialEvent   SportEvent = 5
	SportEvent_Training       SportEvent = 6
	SportEvent_Transportation SportEvent = 7
	SportEvent_Touring        SportEvent = 8
	SportEvent_Invalid        SportEvent = 255
)

var SportEventNames = map[SportEvent]string{
	SportEvent_Uncategorized:  "Uncategorized",
	SportEvent_Geocaching:     "Geocaching",
	SportEvent_Fitness:        "Fitness",
	SportEvent_Recreation:     "Recreation",
	SportEvent_Race:           "Race",
	SportEvent_SpecialEvent:   "SpecialEvent",
	SportEvent_Training:       "Training",
	SportEvent_Transportation: "Transportation",
	SportEvent_Touring:        "Touring",
}

func (k SportEvent) String() string {
	if uint(k) < uint(len(SportEventNames)) {
		if v, ok := SportEventNames[k]; ok {
			return v
		}
	}
	return "SportEvent(" + strconv.Itoa(int(k)) + ")"
}

var SportEventValues = map[string]SportEvent{
	"Uncategorized":  SportEvent_Uncategorized,
	"Geocaching":     SportEvent_Geocaching,
	"Fitness":        SportEvent_Fitness,
	"Recreation":     SportEvent_Recreation,
	"Race":           SportEvent_Race,
	"SpecialEvent":   SportEvent_SpecialEvent,
	"Training":       SportEvent_Training,
	"Transportation": SportEvent_Transportation,
	"Touring":        SportEvent_Touring,
}

func SportEventFromString(s string) SportEvent {
	if v, ok := SportEventValues[s]; ok {
		return v
	}
	return SportEvent(SportEvent_Invalid)
}
