package registration

import (
	"awesomeProject/fitIdl/messages"
	"fmt"
)

type IFitMessage interface {
	MsgNumber() uint16
	Read(data []byte, md *messages.MessageDefinition) error
}

var m = map[uint16]func() (IFitMessage, error){}
var globalFieldMap = map[uint16]map[uint16]DefinedFields{}

func CreateMessage(msgNumber uint16) (IFitMessage, error) {
	if onCreate, ok := m[msgNumber]; ok {
		return onCreate()
	}
	return nil, fmt.Errorf("could not find message %v", msgNumber)
}

func RegisterMessage(msgNumber uint16, fieldMap map[uint16]DefinedFields, onCreate func() (IFitMessage, error)) error {
	m[msgNumber] = onCreate
	globalFieldMap[msgNumber] = fieldMap
	return nil
}

var couldNotFindField = fmt.Errorf("could not find message/field")

func GetField(msgNumber uint16, fieldId uint16) (DefinedFields, error) {
	if message, ok := globalFieldMap[msgNumber]; ok {
		if fieldMap, ok := message[fieldId]; ok {
			return fieldMap, nil
		}
	}
	return DefinedFields{}, couldNotFindField
}

type DefinedFields struct {
	BaseType      messages.BaseType
	FieldId       int
	Name          string
	IsArray       bool
	ArrayLength   byte
	Components    string
	Scale         float64
	Offset        float64
	Units         string
	Bits          string
	Accumulate    string
	RefFieldName  string
	RefFieldValue string
	Comment       string
	Products      string
	Example       string
}
