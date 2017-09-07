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
	OpenID      string // 投诉人wx open ID
	StreetID    string //街道
	CommunityID string //社区
	XQID        string //投诉小区ID
	Status      int    //事件状态  -2-已关闭 -1-用户撤销 0-居民提交 1-已审核待处理 2-已处理待确认 3-已解决
	EventLevel  int    //事件等级  1-特急、2-加急、3-急
	Type        string `bson:"type" json:"type"` //事件基本类别
	Content     string //投诉内容
	Time        int64  //提交时间
	ToCourt     int8   //0-不推送至法院 1-推送至法院
	Imgs        string //图片表格，以,为分割符
	Tel         string //联系电话
}

//EventNum 不同事件类型的数量
type EventNum struct {
	Type string
	Num  int
}

//FindTodayEvents 获取今日新增事件
func FindTodayEvents(db *mgo.Database) interface{} {
	curTime := time.Now()
	startTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, time.UTC)
	endTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 23, 59, 59, 1000000000, time.UTC)
	c := db.C(EventTableName)
	var result []Event
	fmt.Println(startTime.Unix(), endTime.Unix())
	err := c.Find(
		bson.M{
			"$and": []bson.M{
				bson.M{"time": bson.M{"$gte": startTime.Unix()}},
				bson.M{"time": bson.M{"$lte": endTime.Unix()}},
			},
		}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

// FindEventByTime FindEventByTime
func FindEventByTime(db *mgo.Database, startTime int64, endTime int64) interface{} {
	var result []Event
	c := db.C(EventTableName)
	err := c.Find(
		bson.M{
			"$and": []bson.M{
				bson.M{"time": bson.M{"$gte": startTime}},
				bson.M{"time": bson.M{"$lte": endTime}},
			},
		}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//CheckEventProgress 查询open ID时间段内的事件
func CheckEventProgress(db *mgo.Database, openID string, startTime int64,
	endTime int64) interface{} {
	var result []Event
	c := db.C(EventTableName)
	err := c.Find(
		bson.M{
			"$and": []bson.M{
				bson.M{"time": bson.M{"$gte": startTime}},
				bson.M{"time": bson.M{"$lte": endTime}},
				bson.M{"openid": openID},
			},
		}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
func createEventIndex() string {
	curTime := time.Now()
	result := curTime.Format("20060102150405")
	return result + fmt.Sprint(curTime.Nanosecond())
}

//InsertEvent 插入
func InsertEvent(db *mgo.Database, event Event) interface{} {
	c := db.C(EventTableName)
	count, err := c.Find(bson.M{"index": event.Index}).Count()
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "事件已经存在，编号为" + event.Index}
	}
	event.Index = createEventIndex()
	event.Time = time.Now().Unix()
	fmt.Println("event time", event.Time)
	err = c.Insert(&event)
	if err != nil {
		log.Println(err)
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

func FindEvent(db *mgo.Database, index string) Event {
	c := db.C(EventTableName)
	var result Event
	c.Find(bson.M{"index": index}).One(&result)
	return result
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

//CountDiffKeyEvents 统计不同key-value的数量
func CountDiffKeyEvents(db *mgo.Database, key string) interface{} {
	c := db.C(EventTableName)
	key = strings.ToLower(key)
	pipe := c.Pipe([]bson.M{bson.M{"$group": bson.M{"_id": bson.M{key: "$" + key}, "num": bson.M{"$sum": 1}}}})
	result := []bson.M{}
	err := pipe.All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//CountDiffTypeEvents 查询每个类型的事件数量
func CountDiffTypeEvents(db *mgo.Database) interface{} {
	c := db.C(EventTypeTableName)
	var eventTypes []EventType
	var types []string
	err := c.Find(nil).All(&eventTypes)
	for _, item := range eventTypes {
		types = append(types, item.Type)
	}
	fmt.Println("types", types)
	c = db.C(EventTableName)
	// query := bson.M{"$group": bson.M{"_id": "$type", "num": bson.M{"$sum": 1}}}
	queryT := bson.M{"$match": bson.M{"type": bson.M{"$in": types}}}
	queryE := bson.M{"$group": bson.M{"_id": bson.M{"type": "$type"}, "num": bson.M{"$sum": 1}}}

	pipe := c.Pipe([]bson.M{queryT, queryE})
	result := []bson.M{}
	err = pipe.All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

func CountEvents(db *mgo.Database) interface{} {
	c := db.C(EventTableName)
	count, err := c.Find(nil).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": count}
}
