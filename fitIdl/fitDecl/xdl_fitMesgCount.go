package fitDecl

import strconv "strconv"

type MesgCount byte

const (
	MesgCount_NumPerFile     MesgCount = 0
	MesgCount_MaxPerFile     MesgCount = 1
	MesgCount_MaxPerFileType MesgCount = 2
	MesgCount_Invalid        MesgCount = 255
)

var MesgCountNames = map[MesgCount]string{
	MesgCount_NumPerFile:     "NumPerFile",
	MesgCount_MaxPerFile:     "MaxPerFile",
	MesgCount_MaxPerFileType: "MaxPerFileType",
}

func (k MesgCount) String() string {
	if uint(k) < uint(len(MesgCountNames)) {
		if v, ok := MesgCountNames[k]; ok {
			return v
		}
	}
	return "MesgCount(" + strconv.Itoa(int(k)) + ")"
}

var MesgCountValues = map[string]MesgCount{
	"NumPerFile":     MesgCount_NumPerFile,
	"MaxPerFile":     MesgCount_MaxPerFile,
	"MaxPerFileType": MesgCount_MaxPerFileType,
}

func MesgCountFromString(s string) MesgCount {
	if v, ok := MesgCountValues[s]; ok {
		return v
	}
	return MesgCount(MesgCount_Invalid)
}
