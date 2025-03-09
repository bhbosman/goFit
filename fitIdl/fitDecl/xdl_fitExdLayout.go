package fitDecl

import strconv "strconv"

type ExdLayout byte

const (
	ExdLayout_FullScreen                ExdLayout = 0
	ExdLayout_HalfVertical              ExdLayout = 1
	ExdLayout_HalfHorizontal            ExdLayout = 2
	ExdLayout_HalfVerticalRightSplit    ExdLayout = 3
	ExdLayout_HalfHorizontalBottomSplit ExdLayout = 4
	ExdLayout_FullQuarterSplit          ExdLayout = 5
	ExdLayout_HalfVerticalLeftSplit     ExdLayout = 6
	ExdLayout_HalfHorizontalTopSplit    ExdLayout = 7
	ExdLayout_Dynamic                   ExdLayout = 8
	ExdLayout_Invalid                   ExdLayout = 255
)

var ExdLayoutNames = map[ExdLayout]string{
	ExdLayout_FullScreen:                "FullScreen",
	ExdLayout_HalfVertical:              "HalfVertical",
	ExdLayout_HalfHorizontal:            "HalfHorizontal",
	ExdLayout_HalfVerticalRightSplit:    "HalfVerticalRightSplit",
	ExdLayout_HalfHorizontalBottomSplit: "HalfHorizontalBottomSplit",
	ExdLayout_FullQuarterSplit:          "FullQuarterSplit",
	ExdLayout_HalfVerticalLeftSplit:     "HalfVerticalLeftSplit",
	ExdLayout_HalfHorizontalTopSplit:    "HalfHorizontalTopSplit",
	ExdLayout_Dynamic:                   "Dynamic",
}

func (k ExdLayout) String() string {
	if uint(k) < uint(len(ExdLayoutNames)) {
		if v, ok := ExdLayoutNames[k]; ok {
			return v
		}
	}
	return "ExdLayout(" + strconv.Itoa(int(k)) + ")"
}

var ExdLayoutValues = map[string]ExdLayout{
	"FullScreen":                ExdLayout_FullScreen,
	"HalfVertical":              ExdLayout_HalfVertical,
	"HalfHorizontal":            ExdLayout_HalfHorizontal,
	"HalfVerticalRightSplit":    ExdLayout_HalfVerticalRightSplit,
	"HalfHorizontalBottomSplit": ExdLayout_HalfHorizontalBottomSplit,
	"FullQuarterSplit":          ExdLayout_FullQuarterSplit,
	"HalfVerticalLeftSplit":     ExdLayout_HalfVerticalLeftSplit,
	"HalfHorizontalTopSplit":    ExdLayout_HalfHorizontalTopSplit,
	"Dynamic":                   ExdLayout_Dynamic,
}

func ExdLayoutFromString(s string) ExdLayout {
	if v, ok := ExdLayoutValues[s]; ok {
		return v
	}
	return ExdLayout(ExdLayout_Invalid)
}
