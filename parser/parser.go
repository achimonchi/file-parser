package parser

import "bytes"

type Parser interface {
	Parse() error
}

type Output interface {
	String() string
	JSON() (string, error)
	Reader() *bytes.Buffer
}

type ParserOutput interface {
	Parser
	Output
}

type Format string

const (
	FormatPDF  Format = "pdf"
	FormatDOCX Format = "docx"
	FormatXLSX Format = "xlsx"
	FormatCSV  Format = "csv"
)

type ParserConfig struct {
	HeaderRows int
}

func NewParser(format Format, path string) (ParserOutput, error) {
	switch format {
	case FormatPDF:
		return NewPDFParser(path)
	case FormatXLSX:
		return NewXLSXParser(path)
	case FormatDOCX:
		return NewDOCXParser(path)
	}
	return nil, nil
}
