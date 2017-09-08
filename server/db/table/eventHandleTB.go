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

//FindTodayHandles 获取今日新增事件
func FindTodayHandles(db *mgo.Database) interface{} {
	curTime := time.Now()
	startTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, time.UTC)
	endTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 23, 59, 59, 1000000000, time.UTC)
	c := db.C(EventHandleTableName)
	var result []EventHandle
	err := c.Find(bson.M{"$and": []bson.M{bson.M{"time": bson.M{"$gte": startTime.Unix()}}, bson.M{"time": bson.M{"$lte": endTime.Unix()}}}}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//InsertEventHandle 插入
func InsertEventHandle(db *mgo.Database, handle EventHandle) (EventHandle, error) {
	c := db.C(EventHandleTableName)
	handle.Time = time.Now().Unix()
	err := c.Insert(&handle)
	return handle, err
}

//FindEventHandle 按编号index查询事件
func FindEventHandle(db *mgo.Database, index string, pageNo int, pageSize int) interface{} {
	c := db.C(EventHandleTableName)
	var query *mgo.Query
	if index == "" {
		query = c.Find(nil)
	} else {
		query = c.Find(bson.M{"index": index})
	}
	sum, _ := query.Count()
	var err error
	var result []EventHandle

	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"eventHandles": result, "sum": sum}}
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

func FindEventHandleByKVsPage(db *mgo.Database, kvs map[string]interface{}, pageNo int, pageSize int) interface{} {
	querys := make(map[string]interface{})
	for k, v := range kvs {
		querys[strings.ToLower(k)] = v
	}
	c := db.C(EventHandleTableName)
	query := c.Find(querys)
	sum, _ := query.Count()
	var err error
	var result []EventHandle
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"eventHandles": result, "sum": sum}}
}
