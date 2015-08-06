package types

type Config map[string][]map[string]string

type Article struct {
	Title string
	Link string
}

type Video struct {
	Title string
	Length int
	Link string
}