package table

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//OnlineUserTabelName table name
const OnlineUserTabelName = "OnlineUser"

//OnlineUser 在线用户
type OnlineUser struct {
	UserName    string
	Type        int   //用户类型 0-admin 2-street 3-gov
	Time        int64 //用户登录或更新时间 us
	Token       string
	ValidPeriod int //token 的有效期
}

//UpsertOnlineUser 插入或更新
func UpsertOnlineUser(db *mgo.Database, onlineUser OnlineUser) {
	c := db.C(OnlineUserTabelName)
	changeInfo, err := c.Upsert(bson.M{"username": onlineUser.UserName}, &onlineUser)
	if err != nil {
		log.Println("Error UpsertOnlineUser Failed", err.Error())
	} else {
		log.Println("UpsertOnlineUse Succ", changeInfo)
	}
}
