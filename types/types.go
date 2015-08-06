package types

type Config map[string][]map[string]string

type ReadTuple struct {
	Title string
	Link string
}

type ReadList []ReadTuple

type WatchTriple struct {
	Title string
	Length int
	Link string
}

type WatchList []WatchTriple