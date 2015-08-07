package types

//type Config map[string][]map[string]string
type ConfigRecord map[string]string
type Config map[string]ConfigRecord

type BaseRecord struct {
	Title string
	Link string	
}

type RecordList []BaseRecord