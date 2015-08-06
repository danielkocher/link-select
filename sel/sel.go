package sel

import (
	"os"
	"flag"
	"fmt"
	"log"
	"encoding/json"
	"link-select/types"
)

func SelectLink(arg *flag.Flag) {
	fmt.Fprintf(os.Stdout, "Selecting %s from json file\n", arg.Value)

	readFile, err := os.Open("files/read.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening read JSON file\n")
		os.Exit(-1)
	}

	var readList types.ReadList
	
	jsonParser := json.NewDecoder(readFile)
	if err = jsonParser.Decode(&readList); err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing read JSON file\n")
		log.Fatal(err)
		os.Exit(-1)
	}

	//for i, v := range articles {
	//fmt.Fprintf("title: %s, link: %s\n", article.Title, article.Link)
}