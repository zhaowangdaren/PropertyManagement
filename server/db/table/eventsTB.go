package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventTableName 表名
const EventTableName = "Event"

//Event 投诉事件表
type Event struct {
	Index       string //事件编号
	Complainant string //投诉人
	Street      string //街道
	Community   string //社区
	XQ          string //投诉小区名
	Status      int    //事件状态  1-居民提交 2-待审核 3-待处理 4-已处理
	EventLevel  int    //事件等级  1-特急、2-加急、3-急
	Type        string //事件基本类别
	Content     string //投诉内容
	Time        string //提交时间
	ToCourt     int8   //0-不推送至法院 1-推送至法院
	Img1        string
	Img2        string
	Img3        string
	Img4        string
}

//InsertEventTB 插入
func InsertEventTB(db *mgo.Database, event Event) string {
	c := db.C(EventTableName)
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

//FindEvents 按编号index查询事件
func FindEvents(db *mgo.Database, index string, pageNo int, pageSize int) interface{} {
	c := db.C(EventTableName)
	var result []Event
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

//FindEventKVs find key-value
func FindEventKVs(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(EventTableName)
	var result []Event
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
