package table

import (
	"log"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventHandleTableName 表名
const EventHandleTableName = "EventHandle"

//EventHandle 事件处理表
type EventHandle struct {
	Index          string //事件编号，属于Event表中的事件编号
	AuthorCategory string //提交人类别  住建委、社区、具名
	AuthorName     string //提交人用户名
	Time           string //提交的时间
	HandleInfo     string //处理信息
	SubmitIndex    int    //提交级次 0为最早，级次越大到则时间越靠后
}

//InsertEventHandleTB 插入
func InsertEventHandleTB(db *mgo.Database, handle EventHandle) string {
	c := db.C(EventHandleTableName)
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

//FindEventHandle 按编号index查询事件
func FindEventHandle(db *mgo.Database, index string, pageNo int, pageSize int) interface{} {
	c := db.C(EventHandleTableName)
	var result []EventHandle
	var err error
	if index == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"index": index}).All(&result)
	}
	if err != nil {
		log.Println(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
