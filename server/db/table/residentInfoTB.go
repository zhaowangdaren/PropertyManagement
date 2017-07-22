package table

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ResidentInfo 小区居民信息
type ResidentInfo struct {
	WXName       string //用户微信号
	Name         string //用户姓名
	XiaoQuName   string //所属小区名
	Building     string //楼栋号
	HousNumber   string //门牌号
	Password     string //登录密码
	UserAnalysis string //用户分析
	Tel          string //联系方式
	UserLevel    int    //用户权限等级
}

//InsertResidentInfo 插入user
func InsertResidentInfo(db *mgo.Database, info ResidentInfo) string {
	c := db.C("ResidentInfo")
	count, err := c.Find(bson.M{"wxname": info.WXName}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "已存在微信号为" + info.WXName + "的用户"
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}
