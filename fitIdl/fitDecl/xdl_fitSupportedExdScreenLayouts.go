package fitDecl

import strconv "strconv"

type SupportedExdScreenLayouts uint32

const (
	SupportedExdScreenLayouts_FullScreen                SupportedExdScreenLayouts = 1
	SupportedExdScreenLayouts_HalfVertical              SupportedExdScreenLayouts = 2
	SupportedExdScreenLayouts_HalfHorizontal            SupportedExdScreenLayouts = 4
	SupportedExdScreenLayouts_HalfVerticalRightSplit    SupportedExdScreenLayouts = 8
	SupportedExdScreenLayouts_HalfHorizontalBottomSplit SupportedExdScreenLayouts = 16
	SupportedExdScreenLayouts_FullQuarterSplit          SupportedExdScreenLayouts = 32
	SupportedExdScreenLayouts_HalfVerticalLeftSplit     SupportedExdScreenLayouts = 64
	SupportedExdScreenLayouts_HalfHorizontalTopSplit    SupportedExdScreenLayouts = 128
	SupportedExdScreenLayouts_Invalid                   SupportedExdScreenLayouts = 0
)

var SupportedExdScreenLayoutsNames = map[SupportedExdScreenLayouts]string{
	SupportedExdScreenLayouts_FullScreen:                "FullScreen",
	SupportedExdScreenLayouts_HalfVertical:              "HalfVertical",
	SupportedExdScreenLayouts_HalfHorizontal:            "HalfHorizontal",
	SupportedExdScreenLayouts_HalfVerticalRightSplit:    "HalfVerticalRightSplit",
	SupportedExdScreenLayouts_HalfHorizontalBottomSplit: "HalfHorizontalBottomSplit",
	SupportedExdScreenLayouts_FullQuarterSplit:          "FullQuarterSplit",
	SupportedExdScreenLayouts_HalfVerticalLeftSplit:     "HalfVerticalLeftSplit",
	SupportedExdScreenLayouts_HalfHorizontalTopSplit:    "HalfHorizontalTopSplit",
}

func (k SupportedExdScreenLayouts) String() string {
	if uint(k) < uint(len(SupportedExdScreenLayoutsNames)) {
		if v, ok := SupportedExdScreenLayoutsNames[k]; ok {
			return v
		}
	}
	return "SupportedExdScreenLayouts(" + strconv.Itoa(int(k)) + ")"
}

var SupportedExdScreenLayoutsValues = map[string]SupportedExdScreenLayouts{
	"FullScreen":                SupportedExdScreenLayouts_FullScreen,
	"HalfVertical":              SupportedExdScreenLayouts_HalfVertical,
	"HalfHorizontal":            SupportedExdScreenLayouts_HalfHorizontal,
	"HalfVerticalRightSplit":    SupportedExdScreenLayouts_HalfVerticalRightSplit,
	"HalfHorizontalBottomSplit": SupportedExdScreenLayouts_HalfHorizontalBottomSplit,
	"FullQuarterSplit":          SupportedExdScreenLayouts_FullQuarterSplit,
	"HalfVerticalLeftSplit":     SupportedExdScreenLayouts_HalfVerticalLeftSplit,
	"HalfHorizontalTopSplit":    SupportedExdScreenLayouts_HalfHorizontalTopSplit,
}

func SupportedExdScreenLayoutsFromString(s string) SupportedExdScreenLayouts {
	if v, ok := SupportedExdScreenLayoutsValues[s]; ok {
		return v
	}
	return SupportedExdScreenLayouts(SupportedExdScreenLayouts_Invalid)
}
