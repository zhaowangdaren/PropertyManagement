package table

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SheQuMember 社区人员信息表
type SheQuMember struct {
	UserName      string //
	RealName      string //
	CommunityName string //所属社区名
	Password      string //
	OfficeNumber  string //
}

//InsertSheQuMember 插入user
func InsertSheQuMember(db *mgo.Database, info SheQuMember) string {
	c := db.C("SheQuMember")
	count, err := c.Find(bson.M{"username": info.UserName}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "已存在用户名为" + info.UserName + ", 请重新设置用户名"
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}
