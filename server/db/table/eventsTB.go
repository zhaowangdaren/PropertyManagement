package table

import (
	"fmt"
	"log"
	"strings"
	"time"

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
	StreetID    string //街道
	CommunityID string //社区
	XQID        string //投诉小区名
	Status      int    //事件状态  1-居民提交 2-待审核 3-待处理 4-已处理
	EventLevel  int    //事件等级  1-特急、2-加急、3-急
	Type        int8   //事件基本类别
	Content     string //投诉内容
	Time        int64  //提交时间
	ToCourt     int8   //0-不推送至法院 1-推送至法院
	Imgs        string //图片表格，以,为分割符
}

func createEventIndex() string {
	curTime := time.Now()
	result := curTime.Format("20060102150405")
	return result + fmt.Sprint(curTime.Unix())
}

//InsertEvent 插入
func InsertEvent(db *mgo.Database, event Event) interface{} {
	c := db.C(EventTableName)
	count, err := c.Find(bson.M{"index": event.Index}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "事件已经存在，编号为" + event.Index}
	}
	event.Index = createEventIndex()
	event.Time = time.Now().Unix()
	fmt.Println("event index=", event.Index)
	err = c.Insert(&event)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": event.Index}
}

//UpdateEvent update event
func UpdateEvent(db *mgo.Database, info Event) interface{} {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": info.Index}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": info.Index}
}

//AddEventImgs add event imgs
func AddEventImgs(db *mgo.Database, index string, imgs []string) interface{} {
	c := db.C(EventTableName)
	var event Event
	err := c.Find(bson.M{"index": index}).One(&event)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}

	imgsStr := ""
	for _, v := range imgs {
		imgsStr += v + ","
	}
	event.Imgs = imgsStr
	err = c.Update(bson.M{"index": event.Index}, event)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
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
	var startTime int64
	var endTime int64
	startTime = 0
	endTime = 9223372036854775807
	hasTime := false
	for k, v := range kvs {
		if k == "StartTime" {
			startTime = int64(v.(float64))
			hasTime = true
			continue
		}
		if k == "EndTime" {
			endTime = int64(v.(float64))
			hasTime = true
			continue
		}
		query[strings.ToLower(k)] = v
	}
	c := db.C(EventTableName)
	var result []Event
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}

	var timeResult []Event
	if hasTime {
		for _, item := range result {
			if startTime <= item.Time && item.Time <= endTime {
				timeResult = append(timeResult, item)
			}
		}
		return gin.H{"error": 0, "data": timeResult}
	}
	return gin.H{"error": 0, "data": result}
}
