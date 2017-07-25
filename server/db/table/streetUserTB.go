package table

import (
	"log"

	"github.com/gin-gonic/gin"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//StreetUserTableName table name
const StreetUserTableName = "StreetUser"

//StreetUser 街道用户
type StreetUser struct {
	UserName string //用户名
	RealName string //真实姓名
	Street   string //所在街道
	Password string //
	Tel      string //
}

//StreetUsers 街道用户集
type StreetUsers struct {
	StreetUsers []StreetUser
}

//InsertStreetUser 插入
func InsertStreetUser(db *mgo.Database, info StreetUser) interface{} {
	c := db.C(StreetUserTableName)
	count, err := c.Find(bson.M{"username": info.UserName}).Count()
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "已存在用户名为" + info.UserName + ", 请重新设置街道名"}
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindStreetUsers 如果userName为""，则查询所有的街道信息,
func FindStreetUsers(db *mgo.Database, userName string, pageNo int,
	pageSize int) interface{} {
	c := db.C(StreetUserTableName)
	var result []StreetUser
	var err error
	if userName == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"username": userName}).All(&result)
	}
	if err != nil {
		log.Println(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
