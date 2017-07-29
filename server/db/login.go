package db

import (
	"./table"
	"gopkg.in/mgo.v2"
)

//Login 登录
func Login(db *mgo.Database, userName string, password string) string {
	result := table.FindAdmin(db, userName)
	if result.UserName == userName && result.Password == password {
		return "succ"
	} else if result.UserName != userName {
		return "用户名不存在"
	} else if result.Password != password {
		return "密码错误"
	}
	return "登录失败"
}
