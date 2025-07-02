package parser

import (
	"errors"
	"fmt"
	"testing"
)

var xlsxAPI ParserOutput

func init() {
	xlsxAPI, _ = NewParser(FormatXLSX, "../sample-doc/simple.xlsx")
}

func TestParseXlsx(t *testing.T) {
	type tableTest struct {
		name          string
		path          string
		expectedError error
		beforeTest    func(path string) ParserOutput
	}
	var tableTests = []tableTest{
		{
			name:          "error - file not found",
			path:          "simple.go",
			expectedError: ErrFileNotFound,
			beforeTest: func(path string) ParserOutput {
				api, _ := NewParser(FormatXLSX, path)
				return api
			},
		},
		{
			name:          "success",
			path:          "../sample-doc/simple.xlsx",
			expectedError: nil,
			beforeTest: func(path string) ParserOutput {
				api, _ := NewParser(FormatXLSX, path)
				return api
			},
		},
	}

	for _, test := range tableTests {
		t.Run(test.name, func(t *testing.T) {
			api := test.beforeTest(test.path)
			err := api.Parse()
			if !errors.Is(err, test.expectedError) {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestXlsxString(t *testing.T) {
	xlsxAPI.Parse()
	if xlsxAPI.String() == "" {
		t.Error("string is empty")
	}

	fmt.Println(xlsxAPI.String())
}

func TestXlsxJSON(t *testing.T) {
	xlsxAPI.Parse()
	json, err := xlsxAPI.JSON()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(json)
}

func TestXlsxReader(t *testing.T) {
	xlsxAPI.Parse()
	reader := xlsxAPI.Reader()
	if reader == nil {
		t.Error("reader is nil")
	}
}
