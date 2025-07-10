package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/nguyenthenguyen/docx"
	"github.com/pkg/errors"
)

type DOCXParser struct {
	path   string
	buffer bytes.Buffer
	lines  []string
}

func NewDOCXParser(path string) (ParserOutput, error) {
	parser := &DOCXParser{path: path}
	return parser, nil
}

func (p *DOCXParser) Parse() error {
	r, err := docx.ReadDocxFile(p.path)
	if err != nil {
		return errors.Wrap(ErrFileNotFound, fmt.Sprintf("failed to open DOCX file: %v", err.Error()))
	}
	defer r.Close()

	doc := r.Editable()
	content := doc.GetContent()
	lines := strings.Split(content, "\n")

	p.lines = lines
	p.buffer.Reset()
	p.buffer.WriteString(content)
	return nil
}

func (p *DOCXParser) String() string {
	content := p.buffer.String()

	// Regex untuk hapus tag XML
	re := regexp.MustCompile(`<[^>]+>`)
	textOnly := re.ReplaceAllString(content, "")

	// Opsional: Hilangkan spasi berlebih dan trim
	cleaned := strings.Join(strings.Fields(textOnly), " ")
	return cleaned
}

func (p *DOCXParser) JSON() (string, error) {
	data := map[string]interface{}{
		"type":  "docx",
		"lines": p.lines,
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (p *DOCXParser) Reader() *bytes.Buffer {
	return &p.buffer
}
