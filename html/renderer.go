package html

import "io"

// Renderer render the parsed content
type Renderer interface {
	Render(w io.Writer) error
}

// Options for rendering
type Options struct {
	Title string
}
