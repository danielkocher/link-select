// Package types provides types for the custom JSON configuration and the
// custom JSON read, watch and book JSON file.
package types

import (
	"os"
	"fmt"
	"encoding/json"
)

// A single record of the configuration file.
type ConfigRecord map[string]string

// A whole configuration file.
type Config map[string]ConfigRecord

// A single base record of the read, watch or book JSON file.
type BaseRecord struct {
	Title string
	Link string	
}

// A whole read, watch or book JSON file.
type RecordList []BaseRecord

// Read opens and read a given JSON file into the RecordList.
// It returns nil on success and the error on failure (propagates the error).
func (list *RecordList) Read(file string) error {
	// open JSON file
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening %s\n", file)
		return err
	}

	// decode JSON file
	jsonParser := json.NewDecoder(readFile)
	if err = jsonParser.Decode(&list); err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing %s\n", file)
		readFile.Close()
		return err
	}

	readFile.Close()
	return nil
}

// Write opens and writes the RecordList to a given JSON file.
// It returns nil on success and the error on failure (propagates the error).
func (list *RecordList) Write(file string) error {
	// write back to JSON file
	readFile, err := os.Create(file)
	jsonWriter := json.NewEncoder(readFile)
	if err = jsonWriter.Encode(&list); err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing back %s\n", file)
		readFile.Close()
		return err
	}

	readFile.Close()
	return nil
}