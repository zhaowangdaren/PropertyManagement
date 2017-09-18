package table

type ResultsPage struct {
	List []interface{} `json:"list"`
	Sum  int           `json:"sum"`
}
