package parser

import "errors"

var (
	ErrFileNotFound      = errors.New("file not found")
	ErrFileCannotExtract = errors.New("error extracting file")
)
