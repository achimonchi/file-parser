package parser

import (
	"errors"
	"fmt"
	"testing"
)

var parserAPI ParserOutput

func init() {
	parserAPI, _ = NewParser(FormatPDF, "../sample-doc/complex.pdf")
}

func TestParse(t *testing.T) {
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
				api, _ := NewParser(FormatPDF, path)
				return api
			},
		},
		{
			name:          "success",
			path:          "../sample-doc/simple.pdf",
			expectedError: nil,
			beforeTest: func(path string) ParserOutput {
				api, _ := NewParser(FormatPDF, path)
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

func TestString(t *testing.T) {
	beforeTest()

	txt := parserAPI.String()
	if txt == "" {
		t.Error("string is empty")
	}

	fmt.Println(txt)
}

func TestJSON(t *testing.T) {
	beforeTest()

	json, err := parserAPI.JSON()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(json)
}

func TestReader(t *testing.T) {
	beforeTest()

	reader := parserAPI.Reader()
	if reader == nil {
		t.Error("reader is nil")
	}
}

func beforeTest() {
	parserAPI.Parse()
}
