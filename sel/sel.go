package sel

import (
	"os"
	"os/exec"
	"flag"
	"fmt"
	"log"
	"time"
	"math/rand"
	"encoding/json"
	"link-select/types"
)

func SelectLink(arg *flag.Flag, file string, browser string) {
	fmt.Fprintf(os.Stdout, "Selecting from %s\n", file)

	// open JSON file
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening %s\n", file)
		os.Exit(-1)
	}

	// decode JSON file into custom type
	var list types.RecordList
	jsonParser := json.NewDecoder(readFile)
	if err = jsonParser.Decode(&list); err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing %s\n", file)
		log.Fatal(err)
		readFile.Close()
		os.Exit(-1)
	}

	// set new seed to get non-deterministic random number generation
	randSrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSrc)

	// select random article
	rndArticle := r.Intn(len(list))
	fmt.Fprintf(os.Stdout, "Selected \"%s\"\n", list[rndArticle].Title)
	
	// open link in browser
	cmd := exec.Command(browser, list[rndArticle].Link)
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening %s in %s\n",
			list[rndArticle], browser)
		log.Fatal(err)
		readFile.Close()
		os.Exit(-1)
	}

	// remove selected article from list
	list = append(list[:rndArticle], list[rndArticle + 1:]...)

	// write back JSON file without selected article
	readFile.Close()
	readFile, err = os.Create(file)
	jsonWriter := json.NewEncoder(readFile)
	if err = jsonWriter.Encode(&list); err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing back %s\n", file)
		log.Fatal(err)
		readFile.Close()
		os.Exit(-1)
	}

	readFile.Close()
}