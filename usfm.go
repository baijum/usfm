package main

import (
	"log"
	"os"
	"strings"

	"github.com/baijum/usfm/html"
)

func main() {
	o := html.Options{}
	s := ""
	html := html.New(o, strings.NewReader(s))
	f, _ := os.Create("out.html")
	defer f.Close()
	err := html.Render(f)
	if err != nil {
		log.Println(err)
	}
}
