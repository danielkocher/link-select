package main

import (
	"os"
	"fmt"
	"flag"
	"io"
	"log"
	"encoding/json"
	"link-select/add"
	"link-select/sel"
	"link-select/types"
)

var system []types.ConfigRecord
var files []types.ConfigRecord

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

func loadConfig() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening config JSON file\n")
		log.Fatal(err)
		os.Exit(-1)
	}

	dec := json.NewDecoder(configFile)
	for {
		var c types.Config
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error while parsing config JSON file\n")
			log.Fatal(err)
			os.Exit(-1)
		}

		system = c["system"]
		files = c["files"]
	}

	/*
	// debug
	fmt.Println("system:", system)
	fmt.Println("files:", files)
	*/
}

func processArgs(arg *flag.Flag) {
	switch arg.Name {
	case "add-link":
		add.AddLink(arg)
	case "a":
		add.AddLink(arg)
	case "sel-link":
		sel.SelectLink(arg)
	case "s":
		sel.SelectLink(arg)
	default:

	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: link-select " +
			"[--add-link=<link-to-add> |" +
			" --sel-link=<type-of-link>])\n")
		os.Exit(-1)
	}

	flag.Parse()
	if flag.Parsed() {
		loadConfig()
    	flag.Visit(processArgs)
	} else {
		fmt.Fprintf(os.Stderr, "Error while parsing command-line arguments\n")
		os.Exit(-1)
	}
}