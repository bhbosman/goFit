package fitDecl

import strconv "strconv"

type AntChannelId uint32

const (
	AntChannelId_AntExtendedDeviceNumberUpperNibble AntChannelId = 4026531840
	AntChannelId_AntTransmissionTypeLowerNibble     AntChannelId = 251658240
	AntChannelId_AntDeviceType                      AntChannelId = 16711680
	AntChannelId_AntDeviceNumber                    AntChannelId = 65535
	AntChannelId_Invalid                            AntChannelId = 0
)

var AntChannelIdNames = map[AntChannelId]string{
	AntChannelId_AntExtendedDeviceNumberUpperNibble: "AntExtendedDeviceNumberUpperNibble",
	AntChannelId_AntTransmissionTypeLowerNibble:     "AntTransmissionTypeLowerNibble",
	AntChannelId_AntDeviceType:                      "AntDeviceType",
	AntChannelId_AntDeviceNumber:                    "AntDeviceNumber",
}

func (k AntChannelId) String() string {
	if uint(k) < uint(len(AntChannelIdNames)) {
		if v, ok := AntChannelIdNames[k]; ok {
			return v
		}
	}
	return "AntChannelId(" + strconv.Itoa(int(k)) + ")"
}

var AntChannelIdValues = map[string]AntChannelId{
	"AntExtendedDeviceNumberUpperNibble": AntChannelId_AntExtendedDeviceNumberUpperNibble,
	"AntTransmissionTypeLowerNibble":     AntChannelId_AntTransmissionTypeLowerNibble,
	"AntDeviceType":                      AntChannelId_AntDeviceType,
	"AntDeviceNumber":                    AntChannelId_AntDeviceNumber,
}

func AntChannelIdFromString(s string) AntChannelId {
	if v, ok := AntChannelIdValues[s]; ok {
		return v
	}
	return AntChannelId(AntChannelId_Invalid)
}
