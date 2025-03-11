package fitDecl

import strconv "strconv"

type DeviceIndex uint8

const (
	DeviceIndex_Creator DeviceIndex = 0
	DeviceIndex_Invalid DeviceIndex = 255
)

var DeviceIndexNames = map[DeviceIndex]string{
	DeviceIndex_Creator: "Creator",
}

func (k DeviceIndex) String() string {
	if uint(k) < uint(len(DeviceIndexNames)) {
		if v, ok := DeviceIndexNames[k]; ok {
			return v
		}
	}
	return "DeviceIndex(" + strconv.Itoa(int(k)) + ")"
}

var DeviceIndexValues = map[string]DeviceIndex{
	"Creator": DeviceIndex_Creator,
}

func DeviceIndexFromString(s string) DeviceIndex {
	if v, ok := DeviceIndexValues[s]; ok {
		return v
	}
	return DeviceIndex(DeviceIndex_Invalid)
}
