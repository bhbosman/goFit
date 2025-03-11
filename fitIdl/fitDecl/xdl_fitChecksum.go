package fitDecl

import strconv "strconv"

type Checksum uint8

const (
	Checksum_Clear   Checksum = 0
	Checksum_Ok      Checksum = 1
	Checksum_Invalid Checksum = 255
)

var ChecksumNames = map[Checksum]string{
	Checksum_Clear: "Clear",
	Checksum_Ok:    "Ok",
}

func (k Checksum) String() string {
	if uint(k) < uint(len(ChecksumNames)) {
		if v, ok := ChecksumNames[k]; ok {
			return v
		}
	}
	return "Checksum(" + strconv.Itoa(int(k)) + ")"
}

var ChecksumValues = map[string]Checksum{
	"Clear": Checksum_Clear,
	"Ok":    Checksum_Ok,
}

func ChecksumFromString(s string) Checksum {
	if v, ok := ChecksumValues[s]; ok {
		return v
	}
	return Checksum(Checksum_Invalid)
}
