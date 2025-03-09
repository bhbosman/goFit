package main

type ConstantValue struct {
	Name    FitIdentifier
	Value   int
	Comment ConstantValueComment
}
type ConstantDeclaration struct {
	Name          FitIdentifier
	Type          FitType
	Comment       string
	ConstantValue []ConstantValue
}

type ConstantValueComment string

func (c ConstantValueComment) String() string {
	if string(c) == "" {
		return "No comment"
	}
	return string(c)
}
