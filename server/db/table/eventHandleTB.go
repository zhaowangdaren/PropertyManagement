package table

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventHandleTableName 表名
const EventHandleTableName = "EventHandle"

//EventHandle 事件处理表
type EventHandle struct {
	Index          string //事件编号，属于Event表中的事件编号
	AuthorCategory int    //提交人类别  住建委、社区、具名
	AuthorName     string //提交人用户名
	Time           int64  //提交的时间
	HandleInfo     string //处理信息
	Imgs           string //图片列表，以,为分隔符
}

//InsertEventHandleTB 插入
func InsertEventHandle(db *mgo.Database, handle EventHandle) interface{} {
	c := db.C(EventHandleTableName)
	handle.Time = time.Now().Unix()
	err := c.Insert(&handle)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": handle}
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

//FindEventHandlesByKV find by kvs
func FindEventHandlesByKV(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(EventHandleTableName)
	var result []EventHandle
	err := c.Find(query).All(&result)
	if err != nil {
		log.Println(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
