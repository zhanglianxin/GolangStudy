package main

import (
	"github.com/golang-commonmark/markdown"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	md := markdown.New()
	src, _ := ioutil.ReadFile("~/cgt.md")
	tokens := md.Parse(src)
	for _, token := range tokens {
		switch token := token.(type) {
		case *markdown.HeadingOpen:
			log.Println("H" + strconv.Itoa(token.HLevel))
		case *markdown.HeadingClose:
			log.Println("H" + strconv.Itoa(token.HLevel))
		case *markdown.ParagraphOpen:
			log.Println("P" + strconv.Itoa(token.Lvl))
		case *markdown.ParagraphClose:
			log.Println("P" + strconv.Itoa(token.Lvl))
		case *markdown.Inline:
			log.Println(token.Content)
		}
	}
}
