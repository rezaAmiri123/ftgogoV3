// Code generated by tools/cmd/genreadfile/main.go. DO NOT EDIT.

package jwe

import (
	"io/fs"
	"os"
)

type sysFS struct{}

func (sysFS) Open(path string) (fs.File, error) {
	return os.Open(path)
}

func ReadFile(path string, options ...ReadFileOption) (*Message, error) {

	var srcFS fs.FS = sysFS{}
	for _, option := range options {
		switch option.Ident() {
		case identFS{}:
			srcFS = option.Value().(fs.FS)
		}
	}

	f, err := srcFS.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return ParseReader(f)
}
