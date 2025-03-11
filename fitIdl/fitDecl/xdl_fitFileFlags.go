package fitDecl

import strconv "strconv"

type FileFlags uint8

const (
	FileFlags_Read    FileFlags = 2
	FileFlags_Write   FileFlags = 4
	FileFlags_Erase   FileFlags = 8
	FileFlags_Invalid FileFlags = 0
)

var FileFlagsNames = map[FileFlags]string{
	FileFlags_Read:  "Read",
	FileFlags_Write: "Write",
	FileFlags_Erase: "Erase",
}

func (k FileFlags) String() string {
	if uint(k) < uint(len(FileFlagsNames)) {
		if v, ok := FileFlagsNames[k]; ok {
			return v
		}
	}
	return "FileFlags(" + strconv.Itoa(int(k)) + ")"
}

var FileFlagsValues = map[string]FileFlags{
	"Read":  FileFlags_Read,
	"Write": FileFlags_Write,
	"Erase": FileFlags_Erase,
}

func FileFlagsFromString(s string) FileFlags {
	if v, ok := FileFlagsValues[s]; ok {
		return v
	}
	return FileFlags(FileFlags_Invalid)
}
