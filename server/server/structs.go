package server

type QueryBasic struct {
	Name     string
	PageNO   int
	PageSize int
}

type Values struct {
	Values []string
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
