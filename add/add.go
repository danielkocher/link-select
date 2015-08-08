// Package add provides primitives for adding links to a JSON file.
package add

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"link-select/types"
)

// AddLink adds a link to file (in JSON format).
// file is supposed to be a JSON file in the read, watch or book format (see
// link-select/types for further details).
// It returns nil on success and the error on failure (propagates the error).
func AddLink(file string) error {
	fmt.Fprintf(os.Stdout, "Adding to %s\n", file)

	// input of title and link
	var newEntry types.BaseRecord
	kbInput := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Title: ")
	newEntry.Title, _ = kbInput.ReadString('\n')
	fmt.Fprintf(os.Stdout, "Link: ")
	newEntry.Link, _ = kbInput.ReadString('\n')

	// trim trailing newlines
	newEntry.Title = strings.TrimSuffix(newEntry.Title, "\n")
	newEntry.Link = strings.TrimSuffix(newEntry.Link, "\n")

	// as appending to a JSON file does not work directory we do the following:
	// (1) decode/parse JSON file
	var list types.RecordList
	if err := list.Read(file); err != nil {
		return err
	}

	// (2) append input data to decoded/parsed data
	list = append(list, newEntry)

	// (3) serialize the new data back to the JSON file
	// write back to JSON file
	if err := list.Write(file); err != nil {
		return err
	}

	return nil
}