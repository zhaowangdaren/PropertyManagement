package table

import (
	"log"

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
func InsertWXUser(db *mgo.Database, info WXUser) string {
	c := db.C(WXUserTableName)
	count, err := c.Find(bson.M{"openid": info.OpenID}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "微信ID：" + info.OpenID + "已绑定过"
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindWXUsers 查询小区信息集
func FindWXUsers(db *mgo.Database, info WXUser, pageNo int, pageSize int) WXUsers {
	c := db.C(WXUserTableName)
	var result WXUsers
	var err error
	if info == (WXUser{}) {
		err = c.Find(nil).All(&result.WXUsers)
	} else {
		err = c.Find(bson.M{"openid": info.OpenID}).All(&result.WXUsers)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
