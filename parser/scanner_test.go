package parser_test

import (
	"strings"
	"testing"

	"github.com/baijum/usfm/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestScan(t *testing.T) {
	var tests = []struct {
		s   string
		tok parser.Token
		lit string
	}{
		// Special tokens (EOF, Illegal, Whitespace)
		{s: ``, tok: parser.EOF},
		{s: `#`, tok: parser.Text, lit: `#`},
		{s: `;`, tok: parser.Text, lit: `;`},
		{s: ` `, tok: parser.Whitespace, lit: " "},
		{s: "\t", tok: parser.Whitespace, lit: "\t"},
		{s: "\r", tok: parser.Whitespace, lit: "\r"},
		{s: "\n", tok: parser.Whitespace, lit: "\n"},
		{s: " \n\t\r", tok: parser.Whitespace, lit: " \n\t\r"},
		{s: `\id`, tok: parser.MarkerID1, lit: `\id`},
		{s: `\imte`, tok: parser.MarkerImte1, lit: `\imte`},
		{s: `\imte1`, tok: parser.MarkerImte1, lit: `\imte1`},
		{s: "123", tok: parser.Number, lit: "123"},
		{s: "Jesus", tok: parser.Text, lit: "Jesus"},
	}

	for i, tt := range tests {
		s := parser.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
