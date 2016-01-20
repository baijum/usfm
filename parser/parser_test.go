package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/baijum/usfm/parser"
)

// Ensure the parser can parse strings into Content ASTs.
func TestParser(t *testing.T) {
	var tests = []struct {
		s       string
		content *parser.Content
		err     string
	}{
		{
			s: `\id RUT T1 T2`,
			content: &parser.Content{
				Type:  "book",
				Value: "RUT",
				Children: []*parser.Content{
					&parser.Content{
						Type:  "marker",
						Value: "\\id",
						Children: []*parser.Content{
							&parser.Content{Type: "bookcode", Value: "RUT"},
							&parser.Content{Type: "text", Value: "T1"},
							&parser.Content{Type: "text", Value: "T2"},
						},
					},
				},
			},
		},
		{
			s: `\v 1 T1 200`,
			content: &parser.Content{
				Type:  "book",
				Value: "",
				Children: []*parser.Content{
					&parser.Content{
						Type:  "marker",
						Value: "\\v",
						Children: []*parser.Content{
							&parser.Content{Type: "versenumber", Value: "1"},
							&parser.Content{Type: "text", Value: "T1"},
							&parser.Content{Type: "text", Value: "200"},
						},
					},
				},
			},
		},
	}

	for i, tt := range tests {
		content, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.content, content) {
			t.Errorf("%d. %q\n\ncontent mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.content, content)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
