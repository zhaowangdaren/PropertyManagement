package table

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//User 管理员表格
type User struct {
	UserName string //用户名
	Password string //
	Level    string //权限等级
}

const userTableName = "User"

//InsertUser 插入user
func InsertUser(db *mgo.Database, user User) string {
	c := db.C(userTableName)
	count, err := c.Find(bson.M{"username": user.UserName}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "已存在用户名为" + user.UserName + ", 请重新设置用户名"
	}
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindUser 查找用户
func FindUser(db *mgo.Database, userName string) User {
	c := db.C(userTableName)
	query := c.Find(bson.M{"username": userName})
	count, _ := query.Count()
	result := User{}
	if count != 0 {
		query.One(&result)
	}
	return result
}
