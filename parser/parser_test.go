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
			s: `\ide 65001 - Unicode (UTF-8)`,
			content: &parser.Content{
				Type:  "book",
				Value: "",
				Children: []*parser.Content{
					&parser.Content{
						Type:  "marker",
						Value: "\\ide",
						Children: []*parser.Content{
							&parser.Content{Type: "text", Value: "65001"},
							&parser.Content{Type: "text", Value: "-"},
							&parser.Content{Type: "text", Value: "Unicode"},
							&parser.Content{Type: "text", Value: "(UTF-8)"},
						},
					},
				},
			},
		},
		{
			s: `\c 42`,
			content: &parser.Content{
				Type:  "book",
				Value: "",
				Children: []*parser.Content{
					&parser.Content{
						Type:  "marker",
						Value: "\\c",
						Children: []*parser.Content{
							&parser.Content{Type: "chapternumber", Value: "42"},
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
		{
			s: `\id RUT T1 T2 \ide UTF-8 \c 1 \v 1 T3 200 \v 28 T4 T5`,
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
					&parser.Content{
						Type:  "marker",
						Value: "\\ide",
						Children: []*parser.Content{
							&parser.Content{Type: "text", Value: "UTF-8"},
						},
					},
					&parser.Content{
						Type:  "marker",
						Value: "\\c",
						Children: []*parser.Content{
							&parser.Content{Type: "chapternumber", Value: "1"},
						},
					},
					&parser.Content{
						Type:  "marker",
						Value: "\\v",
						Children: []*parser.Content{
							&parser.Content{Type: "versenumber", Value: "1"},
							&parser.Content{Type: "text", Value: "T3"},
							&parser.Content{Type: "text", Value: "200"},
						},
					},
					&parser.Content{
						Type:  "marker",
						Value: "\\v",
						Children: []*parser.Content{
							&parser.Content{Type: "versenumber", Value: "28"},
							&parser.Content{Type: "text", Value: "T4"},
							&parser.Content{Type: "text", Value: "T5"},
						},
					},
				},
			},
		},

		// Errors
		{s: `\id X T1 200`, err: `found "X", expected book code`},
		{s: `\v X T1 200`, err: `found "X", expected verse number`},
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
