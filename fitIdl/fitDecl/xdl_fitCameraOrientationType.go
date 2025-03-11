package fitDecl

import strconv "strconv"

type CameraOrientationType byte

const (
	CameraOrientationType_CameraOrientation0   CameraOrientationType = 0
	CameraOrientationType_CameraOrientation90  CameraOrientationType = 1
	CameraOrientationType_CameraOrientation180 CameraOrientationType = 2
	CameraOrientationType_CameraOrientation270 CameraOrientationType = 3
	CameraOrientationType_Invalid              CameraOrientationType = 255
)

var CameraOrientationTypeNames = map[CameraOrientationType]string{
	CameraOrientationType_CameraOrientation0:   "CameraOrientation0",
	CameraOrientationType_CameraOrientation90:  "CameraOrientation90",
	CameraOrientationType_CameraOrientation180: "CameraOrientation180",
	CameraOrientationType_CameraOrientation270: "CameraOrientation270",
}

func (k CameraOrientationType) String() string {
	if uint(k) < uint(len(CameraOrientationTypeNames)) {
		if v, ok := CameraOrientationTypeNames[k]; ok {
			return v
		}
	}
	return "CameraOrientationType(" + strconv.Itoa(int(k)) + ")"
}

var CameraOrientationTypeValues = map[string]CameraOrientationType{
	"CameraOrientation0":   CameraOrientationType_CameraOrientation0,
	"CameraOrientation90":  CameraOrientationType_CameraOrientation90,
	"CameraOrientation180": CameraOrientationType_CameraOrientation180,
	"CameraOrientation270": CameraOrientationType_CameraOrientation270,
}

func CameraOrientationTypeFromString(s string) CameraOrientationType {
	if v, ok := CameraOrientationTypeValues[s]; ok {
		return v
	}
	return CameraOrientationType(CameraOrientationType_Invalid)
}
