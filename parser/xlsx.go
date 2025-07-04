package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

type XLSXParser struct {
	path   string
	buffer *bytes.Buffer
	data   [][]string
}

func NewXLSXParser(path string) (ParserOutput, error) {
	return &XLSXParser{
		path:   path,
		buffer: new(bytes.Buffer),
	}, nil
}

func (p *XLSXParser) Parse() error {
	f, err := excelize.OpenFile(p.path)
	if err != nil {
		return errors.Wrap(ErrFileNotFound, err.Error())
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return fmt.Errorf("no sheets found in XLSX file")
	}

	// Baca dari sheet pertama
	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return err
	}

	p.data = rows
	for _, row := range rows {
		line := strings.Join(row, "\t")
		p.buffer.WriteString(line + "\n")
	}
	return nil
}

func (p *XLSXParser) String() string {
	return p.buffer.String()
}

func (p *XLSXParser) JSON() (string, error) {
	if p.data == nil {
		return "", fmt.Errorf("no data parsed or missing header")
	}

	header := p.data[15]
	columnCount := len(header)
	result := make([]map[string]any, columnCount)

	for i := range columnCount {
		col := map[string]any{
			"key":   header[i],
			"value": []string{},
		}
		for j := 1; j < len(p.data); j++ {
			// Ensure index safe
			if i < len(p.data[j]) {
				col["value"] = append(col["value"].([]string), p.data[j][i])
			} else {
				col["value"] = append(col["value"].([]string), "")
			}
		}
		result[i] = col
	}

	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (p *XLSXParser) Reader() *bytes.Buffer {
	return p.buffer
}
