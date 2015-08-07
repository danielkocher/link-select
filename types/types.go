package types

type Config map[string][]map[string]string

type ReadRecord struct {
	Title string
	Link string
}

type ReadList []ReadRecord

type WatchRecord struct {
	Title string
	Length int
	Link string
}

type WatchList []WatchRecord