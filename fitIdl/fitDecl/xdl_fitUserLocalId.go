package fitDecl

import strconv "strconv"

type UserLocalId uint16

const (
	UserLocalId_LocalMin      UserLocalId = 0
	UserLocalId_LocalMax      UserLocalId = 15
	UserLocalId_StationaryMin UserLocalId = 16
	UserLocalId_StationaryMax UserLocalId = 255
	UserLocalId_PortableMin   UserLocalId = 256
	UserLocalId_PortableMax   UserLocalId = 65534
	UserLocalId_Invalid       UserLocalId = 65535
)

var UserLocalIdNames = map[UserLocalId]string{
	UserLocalId_LocalMin:      "LocalMin",
	UserLocalId_LocalMax:      "LocalMax",
	UserLocalId_StationaryMin: "StationaryMin",
	UserLocalId_StationaryMax: "StationaryMax",
	UserLocalId_PortableMin:   "PortableMin",
	UserLocalId_PortableMax:   "PortableMax",
}

func (k UserLocalId) String() string {
	if uint(k) < uint(len(UserLocalIdNames)) {
		if v, ok := UserLocalIdNames[k]; ok {
			return v
		}
	}
	return "UserLocalId(" + strconv.Itoa(int(k)) + ")"
}

var UserLocalIdValues = map[string]UserLocalId{
	"LocalMin":      UserLocalId_LocalMin,
	"LocalMax":      UserLocalId_LocalMax,
	"StationaryMin": UserLocalId_StationaryMin,
	"StationaryMax": UserLocalId_StationaryMax,
	"PortableMin":   UserLocalId_PortableMin,
	"PortableMax":   UserLocalId_PortableMax,
}

func UserLocalIdFromString(s string) UserLocalId {
	if v, ok := UserLocalIdValues[s]; ok {
		return v
	}
	return UserLocalId(UserLocalId_Invalid)
}
