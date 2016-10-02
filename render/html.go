package render

import (
	"fmt"
	"io"

	"github.com/baijum/usfm/parser"
)

// NewHTML renderer
func NewHTML(o Options, r io.Reader) Renderer {
	html := &HTML{}
	html.usfmParser = parser.NewParser(r)
	return html
}

// HTML renderer
type HTML struct {
	usfmParser *parser.Parser
}

// Render html
func (h *HTML) Render(w io.Writer) error {
	content, err := h.usfmParser.Parse()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(fmt.Sprintf("%#v\n", content)))
	if err != nil {
		return err
	}

	return nil
}
