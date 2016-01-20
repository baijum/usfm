package parser

import (
	"fmt"
	"io"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a USFM formatted book content
func (p *Parser) Parse() (*Content, error) {
	book := &Content{}
	book.Type = "book"
	for {
		// Read a field.
		tok, lit := p.scanIgnoreWhitespace()
		if tok == MarkerID1 {
			marker := &Content{}
			marker.Type = "marker"
			marker.Value = lit
			book.Children = append(book.Children, marker)
			tok, lit := p.scanIgnoreWhitespace()
			if tok == Text && len([]rune(lit)) == 3 {
				child := &Content{}
				child.Type = "bookcode"
				child.Value = lit
				book.Value = lit
				marker.Children = append(marker.Children, child)
				for {
					tok, lit := p.scanIgnoreWhitespace()
					if !(tok == Text || tok == Number) {
						break
					} else {
						child := &Content{}
						child.Type = "text"
						child.Value = lit
						marker.Children = append(marker.Children, child)
					}
				}
			} else {
				return nil, fmt.Errorf("found %q, expected book code", lit)
			}
		}
		if tok == MarkerIde {
			marker := &Content{}
			marker.Type = "marker"
			marker.Value = lit
			book.Children = append(book.Children, marker)
			for {
				tok, lit := p.scanIgnoreWhitespace()
				if !(tok == Text || tok == Number) {
					break
				} else {
					child := &Content{}
					child.Type = "text"
					child.Value = lit
					marker.Children = append(marker.Children, child)
				}
			}
		}

		if tok == MarkerV {
			marker := &Content{}
			marker.Type = "marker"
			marker.Value = lit
			book.Children = append(book.Children, marker)
			tok, lit := p.scanIgnoreWhitespace()
			if tok == Number {
				child := &Content{}
				child.Type = "versenumber"
				child.Value = lit
				marker.Children = append(marker.Children, child)
				for {
					tok, lit := p.scanIgnoreWhitespace()
					if !(tok == Text || tok == Number) {
						break
					} else {
						child := &Content{}
						child.Type = "text"
						child.Value = lit
						marker.Children = append(marker.Children, child)
					}
				}

			} else {
				return nil, fmt.Errorf("found %q, expected verse number", lit)
			}
		}
		if tok == EOF {
			break
		}
	}
	// Return the successfully parsed statement.
	return book, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == Whitespace {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
