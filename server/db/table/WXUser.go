package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//WXUser 绑定的微信用户，包括正在申请绑定
type WXUser struct {
	WXName  string //微信绑定名
	OpenID  string //OpenId
	Tel     string //电话号码
	Request int    //0-不是申请中的，1-是申请中的
	Succ    int    //0-解绑 1-未解绑
}

//WXUsers 微信用户集
type WXUsers struct {
	WXUsers []WXUser
}

//WXUserTableName table name
const WXUserTableName = "WXUser"

//InsertWXUser 插入user
func InsertWXUser(db *mgo.Database, info WXUser) interface{} {
	c := db.C(WXUserTableName)
	count, err := c.Find(bson.M{"openid": info.OpenID}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "微信号已申请了绑定"}
	}
	err = c.Insert(&info)
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

//FindWXUsers 查询小区信息集
func FindWXUsers(db *mgo.Database, wxName string, pageNo int, pageSize int) interface{} {
	c := db.C(WXUserTableName)
	var result []WXUser
	var err error
	if wxName == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"wxname": wxName}).All(&result)
	}
	if err != nil {
		log.Print(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindWXUserDistinct 筛选key对应的value
func FindWXUserDistinct(db *mgo.Database, key string) interface{} {
	c := db.C(WXUserTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(strings.ToLower(key), &result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
