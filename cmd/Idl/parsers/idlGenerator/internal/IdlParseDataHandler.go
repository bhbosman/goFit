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

// IdlParseDataHandler
// Must implement the following interfaces:
//  1. IYYParserHandler
type IdlParseDataHandler struct {
	Writer      io.Writer
	once        sync.Once
	scopedItems map[string]ITypeSpec
	Language    ForLanguage
}

func (handler *IdlParseDataHandler) GetLanguage() ForLanguage {
	return handler.Language
}

func (handler *IdlParseDataHandler) Find(scopeName string) (ITypeSpec, error) {
	if value, ok := handler.scopedItems[scopeName]; ok {
		return value, nil
	}
	return nil, errors.New(scopeName + " not found")
}

func (handler *IdlParseDataHandler) Declare(baseDefinition IBaseDefinition) error {
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

func (handler *IdlParseDataHandler) Close() error {
	handler.once.Do(func() {
		handler.init()
	})
	return nil
}

func (handler *IdlParseDataHandler) Call(Definitions []IBaseDefinition) error {
	handler.once.Do(func() {
		handler.init()
	})

	_, _ = io.WriteString(handler.Writer, "package main\n")
	_, _ = io.WriteString(handler.Writer, "\n")
	importMap := map[string]string{}
	for _, definition := range Definitions {
		if definition == nil {
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
		_, _ = definition.BuildType(handler.Writer, BuildTypeParams{})
		_, _ = io.WriteString(handler.Writer, "\n")
	}
	return nil
}

func (handler *IdlParseDataHandler) init() {
	handler.scopedItems = map[string]ITypeSpec{}
}
