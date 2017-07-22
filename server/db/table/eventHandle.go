package table

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventHandle 事件处理表
type EventHandle struct {
	Index          string //事件编号，属于Event表中的事件编号
	AuthorCategory string //提交人类别  住建委、社区、具名
	AuthorName     string //提交人用户名
	Time           string //提交的时间
	HandelInfo     string //处理信息
	SubmitIndex    int    //提交级次 0为最早，级次越大到则时间越靠后
}

//InsertEventHandleTB 插入
func InsertEventHandleTB(db *mgo.Database, handle EventHandle) string {
	c := db.C("EventHandle")
	count, err := c.Find(bson.M{"index": handle.Index, "submitindex": handle.SubmitIndex}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "存在提交级次相同的处理"
	}
	err = c.Insert(&handle)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}
