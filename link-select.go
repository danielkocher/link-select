package main

import (
	"os"
	"fmt"
	"flag"
	"link-select/add"
	"link-select/sel"
)

var addLink string
var selectLink string

func init() {
	const (
		defaultAdd = "article"
		usageAdd = "add a link (to json file)"
		defaultSelect = "article"
		usageSelect = "select a link (from json file)"
	)

	flag.StringVar(&addLink, "add-link", defaultAdd, usageAdd)
	flag.StringVar(&addLink, "a", defaultAdd, usageAdd + " (shorthand)")
	flag.StringVar(&selectLink, "sel-link", defaultSelect, usageSelect)
	flag.StringVar(&selectLink, "s", defaultSelect, usageSelect + " (shorthand)")
}

func processArgs(arg *flag.Flag) {
	switch arg.Name {
	case "add-link":
		add.AddLink()
	case "a":
		add.AddLink()
	case "sel-link":
		sel.SelectLink()
	case "s":
		sel.SelectLink()
	default:

	}
}

func main() {
	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: link-select " +
			"[--add-link=<link-to-add> |" +
			" --sel-link=<type-of-link>])\n")
		os.Exit(-1)
	}

	flag.Parse()
	if flag.Parsed() {
    	flag.Visit(processArgs)
	} else {
		fmt.Fprintf(os.Stderr, "Error while parsing command-line arguments")
		os.Exit(-1)
	}
}