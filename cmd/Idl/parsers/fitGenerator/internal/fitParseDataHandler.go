package internal

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type IOptions interface {
	GetLanguage() ForLanguage
}

// FitParseDataHandler
// Must implement the following interfaces:
//  1. idlParserHandler
type FitParseDataHandler struct {
	FileName    string
	PackageName string
	Writer      io.Writer
	once        sync.Once
	scopedItems map[string]ITypeSpec
	Language    ForLanguage
}

func (handler *FitParseDataHandler) GetLanguage() ForLanguage {
	return handler.Language
}

func (handler *FitParseDataHandler) Find(scopeName string) (ITypeSpec, error) {
	handler.once.Do(func() {
		handler.init()
	})
	exportedName := __exportName(scopeName)
	if value, ok := handler.scopedItems[exportedName]; ok {
		return value, nil
	}
	return nil, errors.New(scopeName + " not found")
}

func (handler *FitParseDataHandler) Declare(baseDefinition IBaseDefinition) error {
	handler.once.Do(func() {
		handler.init()
	})

	if baseDefinition == nil {
		return nil
	}
	switch v := baseDefinition.(type) {
	case IMultiItems:
		for idx := 0; idx < v.Count(); idx++ {
			newDefinition, name := v.Item(idx)

			handler.scopedItems[name] = newDefinition
		}
	case IDefinition:
		handler.scopedItems[v.Name()] = v
	}

	return nil
}

func (handler *FitParseDataHandler) Close() error {
	handler.once.Do(func() {
		handler.init()
	})
	return nil
}

func (handler *FitParseDataHandler) Call(Definitions []IBaseDefinition) error {
	handler.once.Do(func() {
		handler.init()
	})

	_, _ = io.WriteString(handler.Writer, fmt.Sprintf("package %v\n", handler.PackageName))
	_, _ = io.WriteString(handler.Writer, "\n")
	importMap := map[string]string{}
	for _, definition := range Definitions {
		if definition == nil {
			continue
		}
		if definition.FileName() != handler.FileName {
			continue
		}

		definition.IncludeImports(importMap)
	}

	for key, value := range importMap {
		_, _ = io.WriteString(handler.Writer, fmt.Sprintf("import %v \"%v\"\n", key, value))
	}
	_, _ = io.WriteString(handler.Writer, "\n")

	for _, definition := range Definitions {
		if definition == nil {
			continue
		}
		if definition.FileName() != handler.FileName {
			continue
		}

		_, _ = definition.BuildType(handler.Writer, BuildTypeParams{})
		_, _ = io.WriteString(handler.Writer, "\n")
	}
	return nil
}

func (handler *FitParseDataHandler) init() {
	handler.scopedItems = map[string]ITypeSpec{}
}
