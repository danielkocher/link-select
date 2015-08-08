// Package sel provides primitives for selecting entries from a JSON file.
package sel

import (
	"os"
	"os/exec"
	"fmt"
	"time"
	"math/rand"
	"link-select/types"
)

// SelectLink selects a link from file and opens it in a bg process in browser.
// file is supposed to be a JSON file in the read, watch or book format (see
// link-select/types for further details).
// It returns nil on success and the error on failure (propagates the error).
func SelectLink(file string, browser string) error {
	fmt.Fprintf(os.Stdout, "Selecting from %s\n", file)

	var list types.RecordList
	if err := list.Read(file); err != nil {
		return err
	}

	// set new seed to get non-deterministic random number generation
	randSrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSrc)

	// select random article
	rndArticle := r.Intn(len(list))
	fmt.Fprintf(os.Stdout, "Selected \"%s\"\n", list[rndArticle].Title)
	
	// open link in browser
	cmd := exec.Command(browser, list[rndArticle].Link)
	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening %s in %s\n",
			list[rndArticle], browser)
		return err
	}

	// remove selected article from list
	list = append(list[:rndArticle], list[rndArticle + 1:]...)

	// write back to JSON file
	if err := list.Write(file); err != nil {
		return err
	}

	return nil
}