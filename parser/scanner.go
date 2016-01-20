package parser

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if unicode.IsSpace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isBackslash(ch) {
		s.unread()
		return s.scanMarker()
	} else if isLetter(ch) {
		s.unread()
		return s.scanText()
	} else if unicode.IsDigit(ch) {
		s.unread()
		return s.scanNumber()
	}

	switch ch {
	case eof:
		return EOF, ""
	}

	return Illegal, string(ch)
}

// scanMarker consumes the current rune and read whole marker
func (s *Scanner) scanMarker() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent non-whitespace character into the buffer.
	// Whitespace character and EOF will cause the loop to exit.
	for i := 0; i <= 6; i++ {
		// FIXME: illegal?
		if ch := s.read(); ch == eof {
			break
		} else if unicode.IsSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
		// Handle largest marker like \imte1
		// anything beyond that is illegal
		if i == 6 {
			return Illegal, buf.String()
		}
	}

	switch strings.ToUpper(buf.String()) {
	case `\ID`, `\ID1`:
		return MarkerID1, buf.String()
	case `\IMTE`, `\IMTE1`:
		return MarkerImte1, buf.String()
	case `\V`:
		return MarkerV, buf.String()
	}

	return Illegal, buf.String()

}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !unicode.IsSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Whitespace, buf.String()
}

// scanText consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanText() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent runes part of scripture into the buffer.
	// Non-letter, non-digit characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return Text, buf.String()
}

// scanNumber consumes the current rune and all contiguous number runes.
func (s *Scanner) scanNumber() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !unicode.IsDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return Number, buf.String()
}

// isLetter returns true if the rune is backslash (\)
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsPunct(ch)
}

// isBackslash returns true if the rune is backslash (\)
func isBackslash(ch rune) bool { return ch == '\\' }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
