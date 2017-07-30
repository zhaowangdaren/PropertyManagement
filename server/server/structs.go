package server

type QueryUser struct {
	UserName string
	Type     int
	PageSize int
	PageNo   int
}
type QueryBasic struct {
	Name     string
	PageNo   int
	PageSize int
}

type QueryHouse struct {
	BuildNo  string
	PageNo   int
	PageSize int
}

type QueryGov struct {
	UserName string
	PageNo   int
	PageSize int
}

type QueryEvent struct {
	Index    string //事件编号
	PageNo   int
	PageSize int
}
type Values struct {
	Values []string
}

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
