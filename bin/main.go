package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yuanyu90221/link"
)

func main() {
	filename := flag.String("file", "ex1.html", "the html file to parse")
	flag.Parse()
	fmt.Printf("Parse the html file in %s.\n", *filename)
	r, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
