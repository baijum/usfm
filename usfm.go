package main

import (
	"log"
	"os"
	"strings"

	"github.com/baijum/usfm/renderer"
)

func main() {
	o := renderer.Options{}
	s := ""
	html := renderer.NewHTML(o, strings.NewReader(s))
	f, _ := os.Create("out.html")
	defer f.Close()
	err := html.Render(f)
	if err != nil {
		log.Println(err)
	}
}
