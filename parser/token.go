package parser

// Token represents a lexical token.
type Token int

const (
	// Illegal represents an illegal/invalid character
	Illegal Token = iota

	// Whitespace represents a white space (" ", \t, \r, \n) character
	Whitespace

	// EOF represents end of file
	EOF

	// MarkerID1 represents '\id' or '\id1' marker
	MarkerID1

	// MarkerImte1 represents '\imte' or '\imte1' marker
	MarkerImte1

	// MarkerV represents '\v' marker
	MarkerV

	// Number represents a number (verse, chapter)
	Number

	// Text represents actual text
	Text
)
