package main

type FitMessage struct {
	Name    FitIdentifier
	comment string
	Fields  []*FitMessageField
}

type FitMessageField struct {
	defNumber     int
	fieldName     FitIdentifier
	fieldType     FitTypeIdentifier
	arrayType     string
	components    string
	scale         string
	offset        string
	units         string
	bits          string
	accumulate    string
	refFieldName  string
	refFieldValue string
	comment       string
	products      string
	example       string
	Fields        []*FitSubMessageField
}

type FitSubMessageField struct {
	fieldName     FitIdentifier
	fieldType     string
	arrayType     string
	components    string
	scale         string
	offset        string
	units         string
	bits          string
	accumulate    string
	refFieldName  string
	refFieldValue string
	comment       string
	products      string
	example       string
}
