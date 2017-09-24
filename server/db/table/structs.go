package table

//ResultsPage 分页数据返回基本格式
type ResultsPage struct {
	List []interface{} `json:"list"`
	Sum  int           `json:"sum"`
}

//EventOverview 事件总览
type EventOverview struct {
	Year       int
	Month      int
	StreetName string
	StreetID   string
	Sum        int
	Unhandle   int //未处理
	Handled    int //已处理
	Unsolved   int //未解决
	Solved     int //已解决
}

type EventOverviewGroup struct {
	ID    map[string]string `bson:"_id"`
	Count int
}
