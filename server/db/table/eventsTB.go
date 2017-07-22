package table

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Event 投诉事件表
type Event struct {
	Index       string //事件编号
	Complainant string //投诉人
	XiaoQuName  string //投诉小区名
	Status      string //事件状态
	EventLevel  string //事件等级  特急、加急、急
	Type        string //事件基本类别
}

//InsertEventTB 插入
func InsertEventTB(db *mgo.Database, event Event) string {
	c := db.C("Event")
	count, err := c.Find(bson.M{"index": event.Index}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "事件已经存在，编号为" + event.Index
	}
	err = c.Insert(&event)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}
