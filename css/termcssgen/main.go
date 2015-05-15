package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var pack = flag.String("p", "", "Sets the package name")
var out = flag.String("o", "style.go", "Sets the output filename")
var meth = flag.String("m", "defaultStyle", "name of the method to generate")
var in = flag.String("i", "", "Input css filename")

func main() {
	flag.Parse()

	fIn, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	defer fIn.Close()
	data, err := ioutil.ReadAll(fIn)
	if err != nil {
		log.Fatal(err)
	}
	o, err := parse(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := generate(*out, *pack, *meth, o); err != nil {
		log.Fatal(err)
	}
}
