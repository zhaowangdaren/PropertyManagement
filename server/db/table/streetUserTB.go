package table

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//StreetUserTableName table name
const StreetUserTableName = "StreetUser"

//StreetUser 街道用户
type StreetUser struct {
	UserName string //用户名
	RealName string //真实姓名
	Steet    string //所在街道
	Password string //
	Tel      string //
}

//StreetUsers 街道用户集
type StreetUsers struct {
	StreetUsers []StreetUser
}

//InsertStreetUserTB 插入
func InsertStreetUserTB(db *mgo.Database, info StreetUser) string {
	c := db.C(StreetUserTableName)
	count, err := c.Find(bson.M{"username": info.UserName}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "已存在街道名为" + info.UserName + ", 请重新设置街道名"
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindStreetUsers 如果street.Name为""，则查询所有的街道信息,
func FindStreetUsers(db *mgo.Database, streetUser StreetUser, pageNo int, pageSize int) StreetUsers {
	c := db.C(StreetUserTableName)
	var result StreetUsers
	var err error
	if streetUser == (StreetUser{}) {
		err = c.Find(nil).All(&result.StreetUsers)
	} else {
		err = c.Find(bson.M{"username": streetUser.UserName}).All(&result.StreetUsers)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
