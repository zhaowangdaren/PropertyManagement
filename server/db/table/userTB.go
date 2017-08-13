package table

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UserTableName = "User"

type User struct {
	ID       bson.ObjectId `bson:"_id"`
	UserName string
	Password string
	RealName string //真实姓名
	Tel      string
	StreetID string //所在街道
	Type     int    //1-admin 2-gov 3-street
}

//FindUser 查找用户信息
func FindUser(db *mgo.Database, userName string, userType int) User {
	fmt.Println("userType", userType)
	c := db.C(UserTableName)
	var result User
	fmt.Println("FindUser", userName+" "+string(userType))
	err := c.Find(bson.M{"username": userName, "type": userType}).One(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

//FindUsers 查找用户列表
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

//InsertUser 插入用户
func InsertUser(db *mgo.Database, user User) interface{} {
	c := db.C(UserTableName)
	count, err := c.Find(bson.M{"username": user.UserName, "type": user.Type}).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "已经存在用户名为" + user.UserName + "的用户，请重新设置"}
	}
	user.ID = bson.NewObjectId()
	err = nil
	err = c.Insert(&user)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//DelUsers del user by ids
func DelUsers(db *mgo.Database, ids []string) interface{} {
	c := db.C(UserTableName)
	var err error
	result := ""
	for _, v := range ids {
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(v)})
		if err != nil {
			log.Println(err.Error())
			result += err.Error()
			err = nil
		}
	}
	if result != "" {
		return gin.H{"error": 1, "data": result}
	}
	return gin.H{"error": 0, "data": Succ}
}

//UpdateUser update user by id
func UpdateUser(db *mgo.Database, update User) interface{} {
	c := db.C(UserTableName)
	err := c.Update(bson.M{"_id": update.ID}, update)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}

}

//ChangePassword 修改密码
func ChangePassword(db *mgo.Database, userType int, userName string, password string, newPassword string) interface{} {
	c := db.C(UserTableName)
	var user User
	err := c.Find(bson.M{"username": userName, "type": userType}).One(&user)
	if err != nil {
		log.Println("Error ChangePassword", err.Error())
		return gin.H{"error": 1, "data": "查询用户信息失败"}
	}

	if user == (User{}) {
		return gin.H{"error": 1, "data": "用户信息失效"}
	}

	if user.Password != password {
		return gin.H{"error": 1, "data": "密码错误"}
	}

	user.Password = newPassword
	err = c.Update(bson.M{"_id": user.ID}, user)
	if err != nil {
		log.Println("Error ChangePassword", err.Error())
		return gin.H{"error": 1, "data": "修改密码失败，请重试"}
	}
	return gin.H{"error": 0, "data": "密码修改成功"}
}
