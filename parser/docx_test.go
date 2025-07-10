package parser

import (
	"errors"
	"fmt"
	"testing"
)

var docxAPI ParserOutput

func init() {
	docxAPI, _ = NewParser(FormatDOCX, "../sample-doc/sample.docx")
}

func TestParseDocx(t *testing.T) {
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
				api, _ := NewParser(FormatDOCX, path)
				return api
			},
		},
		{
			name:          "success",
			path:          "../sample-doc/sample.docx",
			expectedError: nil,
			beforeTest: func(path string) ParserOutput {
				api, _ := NewParser(FormatDOCX, path)
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

func TestStringDocx(t *testing.T) {
	beforeTestDoxc()

	txt := docxAPI.String()
	if txt == "" {
		t.Error("string is empty")
	}

	fmt.Println(txt)
}

func TestJSONDocx(t *testing.T) {
	beforeTestDoxc()

	json, err := docxAPI.JSON()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(json)
}

func TestReaderDocx(t *testing.T) {
	beforeTestDoxc()

	reader := docxAPI.Reader()
	if reader == nil {
		t.Error("reader is nil")
	}
}

func beforeTestDoxc() {
	docxAPI.Parse()
}
