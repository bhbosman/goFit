package common

import "io"

type OnReadIncludeFile = func(fileName string) (io.ReadCloser, string, error)

type GenerateIdl = func(inputFile string, readCloser io.ReadCloser, writer io.Writer, onReadIncludeFile OnReadIncludeFile) error
