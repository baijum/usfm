package parser

// Content represents a part of source
// It could be a marker or text
type Content struct {
	// Type refer the type content (Verse, Text, Paragraph etc.)
	Type string

	// Value is the text (Empty for markers)
	Value string

	// Children point to the child contents (empty if no child)
	Children []*Content
}
