// pdf.go
package parser

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pkg/errors"
)

type PDFParser struct {
	path    string
	content *bytes.Buffer
}

func NewPDFParser(path string) (*PDFParser, error) {
	return &PDFParser{
		path:    path,
		content: &bytes.Buffer{},
	}, nil
}

// func (p *PDFParser) Parse() error {
// 	doc, err := pdf.Open(p.path)
// 	if err != nil {
// 		return ErrFileNotFound
// 	}

// 	prevX := 0.0
// 	threshold := 1.0 // jarak antar teks yang dianggap sebagai spasi

// 	for pageNum := 1; pageNum <= doc.NumPage(); pageNum++ {
// 		page := doc.Page(pageNum)
// 		if page.V.IsNull() {
// 			continue
// 		}
// 		content := page.Content()
// 		for _, txt := range content.Text {
// 			if prevX > 0 && txt.X-prevX > threshold {
// 				p.content.WriteString(" ")
// 			}
// 			p.content.WriteString(txt.S)
// 			prevX = txt.X
// 		}
// 		p.content.WriteString("\n")
// 	}

// 	return nil
// }

func (p *PDFParser) Parse() error {
	inFile := filepath.Join(p.path)
	// Create a context.
	ctx, err := api.ReadContextFile(inFile)
	if err != nil {
		return errors.Wrap(ErrFileNotFound, err.Error())
	}

	i := 1
	r, err := pdfcpu.ExtractPageContent(ctx, i)
	if err != nil {
		return errors.Wrap(ErrFileCannotExtract, err.Error())
	}

	bb, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	p.content.Write(bb)

	return nil
}

func (p *PDFParser) String() string {
	return p.content.String()
}

func (p *PDFParser) JSON() (string, error) {
	escaped := strings.ReplaceAll(p.content.String(), "\"", `\\"`)
	return fmt.Sprintf(`{"text": "%s"}`, escaped), nil
}

func (p *PDFParser) Reader() *bytes.Buffer {
	return p.content
}
