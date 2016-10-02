package main

import (
	"log"
	"os"
	"strings"

	"github.com/baijum/usfm/render"
)

func main() {
	o := render.Options{}
	s := ""
	html := render.NewHTML(o, strings.NewReader(s))
	f, _ := os.Create("out.html")
	defer f.Close()
	err := html.Render(f)
	if err != nil {
		log.Println(err)
	}
}
