package fitDecl

import strconv "strconv"

type MessageIndex uint16

const (
	MessageIndex_Selected MessageIndex = 32768
	MessageIndex_Reserved MessageIndex = 28672
	MessageIndex_Mask     MessageIndex = 4095
	MessageIndex_Invalid  MessageIndex = 65535
)

var MessageIndexNames = map[MessageIndex]string{
	MessageIndex_Selected: "Selected",
	MessageIndex_Reserved: "Reserved",
	MessageIndex_Mask:     "Mask",
}

func (k MessageIndex) String() string {
	if uint(k) < uint(len(MessageIndexNames)) {
		if v, ok := MessageIndexNames[k]; ok {
			return v
		}
	}
	return "MessageIndex(" + strconv.Itoa(int(k)) + ")"
}

var MessageIndexValues = map[string]MessageIndex{
	"Selected": MessageIndex_Selected,
	"Reserved": MessageIndex_Reserved,
	"Mask":     MessageIndex_Mask,
}

func MessageIndexFromString(s string) MessageIndex {
	if v, ok := MessageIndexValues[s]; ok {
		return v
	}
	return MessageIndex(MessageIndex_Invalid)
}
