package sel

import (
	"encoding/json"
	"os"
	"flag"
	"fmt"
	"link-select/types"
)

func SelectLink(arg *flag.Flag) {
	fmt.Fprintf(os.Stdout, "Selecting %s from json file\n", arg.Value)

	readFile, err := os.Open("files/read.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening read.json")
		os.Exit(-1)
	}

	var article types.Article
	
	jsonParser := json.NewDecoder(readFile)
	if err = jsonParser.Decode(&article); err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing read.json")
		os.Exit(-1)
	}

	//for i, v := range articles {
	//fmt.Fprintf("title: %s, link: %s\n", article.Title, article.Link)
}