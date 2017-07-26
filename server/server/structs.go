package server

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

type Values struct {
	Values []string
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
