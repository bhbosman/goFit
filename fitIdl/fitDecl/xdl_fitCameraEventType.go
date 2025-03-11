package fitDecl

import strconv "strconv"

type CameraEventType byte

const (
	CameraEventType_VideoStart                  CameraEventType = 0
	CameraEventType_VideoSplit                  CameraEventType = 1
	CameraEventType_VideoEnd                    CameraEventType = 2
	CameraEventType_PhotoTaken                  CameraEventType = 3
	CameraEventType_VideoSecondStreamStart      CameraEventType = 4
	CameraEventType_VideoSecondStreamSplit      CameraEventType = 5
	CameraEventType_VideoSecondStreamEnd        CameraEventType = 6
	CameraEventType_VideoSplitStart             CameraEventType = 7
	CameraEventType_VideoSecondStreamSplitStart CameraEventType = 8
	CameraEventType_VideoPause                  CameraEventType = 11
	CameraEventType_VideoSecondStreamPause      CameraEventType = 12
	CameraEventType_VideoResume                 CameraEventType = 13
	CameraEventType_VideoSecondStreamResume     CameraEventType = 14
	CameraEventType_Invalid                     CameraEventType = 255
)

var CameraEventTypeNames = map[CameraEventType]string{
	CameraEventType_VideoStart:                  "VideoStart",
	CameraEventType_VideoSplit:                  "VideoSplit",
	CameraEventType_VideoEnd:                    "VideoEnd",
	CameraEventType_PhotoTaken:                  "PhotoTaken",
	CameraEventType_VideoSecondStreamStart:      "VideoSecondStreamStart",
	CameraEventType_VideoSecondStreamSplit:      "VideoSecondStreamSplit",
	CameraEventType_VideoSecondStreamEnd:        "VideoSecondStreamEnd",
	CameraEventType_VideoSplitStart:             "VideoSplitStart",
	CameraEventType_VideoSecondStreamSplitStart: "VideoSecondStreamSplitStart",
	CameraEventType_VideoPause:                  "VideoPause",
	CameraEventType_VideoSecondStreamPause:      "VideoSecondStreamPause",
	CameraEventType_VideoResume:                 "VideoResume",
	CameraEventType_VideoSecondStreamResume:     "VideoSecondStreamResume",
}

func (k CameraEventType) String() string {
	if uint(k) < uint(len(CameraEventTypeNames)) {
		if v, ok := CameraEventTypeNames[k]; ok {
			return v
		}
	}
	return "CameraEventType(" + strconv.Itoa(int(k)) + ")"
}

var CameraEventTypeValues = map[string]CameraEventType{
	"VideoStart":                  CameraEventType_VideoStart,
	"VideoSplit":                  CameraEventType_VideoSplit,
	"VideoEnd":                    CameraEventType_VideoEnd,
	"PhotoTaken":                  CameraEventType_PhotoTaken,
	"VideoSecondStreamStart":      CameraEventType_VideoSecondStreamStart,
	"VideoSecondStreamSplit":      CameraEventType_VideoSecondStreamSplit,
	"VideoSecondStreamEnd":        CameraEventType_VideoSecondStreamEnd,
	"VideoSplitStart":             CameraEventType_VideoSplitStart,
	"VideoSecondStreamSplitStart": CameraEventType_VideoSecondStreamSplitStart,
	"VideoPause":                  CameraEventType_VideoPause,
	"VideoSecondStreamPause":      CameraEventType_VideoSecondStreamPause,
	"VideoResume":                 CameraEventType_VideoResume,
	"VideoSecondStreamResume":     CameraEventType_VideoSecondStreamResume,
}

func CameraEventTypeFromString(s string) CameraEventType {
	if v, ok := CameraEventTypeValues[s]; ok {
		return v
	}
	return CameraEventType(CameraEventType_Invalid)
}
