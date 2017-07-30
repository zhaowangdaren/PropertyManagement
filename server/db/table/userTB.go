package table

import (
	"log"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UserTableName = "User"

type User struct {
	UserName string
	Password string
	RealName string //真实姓名
	Tel      string
	Street   string //所在街道
	Type     int    //1-admin 2-gov 3-street
}

//FindUser 查找用户信息
func FindUser(db *mgo.Database, userName string) User {
	c := db.C(UserTableName)
	var result User
	err := c.Find(bson.M{"username": userName}).One(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func FindUsers(db *mgo.Database, userName string, userType int, pageNo int, pageSize int) interface{} {
	c := db.C(UserTableName)
	var result []User
	var err error
	if userName == "" {
		err = c.Find(bson.M{"type": userType}).All(&result)
	} else {
		err = c.Find(bson.M{"username": userName, "type": userType}).All(&result)
	}

	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
