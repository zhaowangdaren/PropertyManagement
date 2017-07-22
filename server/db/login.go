package db

import (
	"./table"
	"gopkg.in/mgo.v2"
)

//Login 登录
func Login(userName string, password string) string {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	result := table.FindUser(db, userName)
	if session != nil {
		session.Close()
		session = nil
	}
	if result.UserName == userName && result.Password == password {
		return "succ"
	} else if result.UserName != userName {
		return "用户名不存在"
	} else if result.Password != password {
		return "密码错误"
	}
	return "登录失败"
}
