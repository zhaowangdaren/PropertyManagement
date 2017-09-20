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
	Index          string // 事件编号，属于Event表中的事件编号
	XQID           string // 小区ID
	AuthorCategory int    // 提交人类别  0-用户 1-系统管理员 2-政府、3-街道、4-物业公司 5-法官
	AuthorName     string // 提交人用户名
	OpenID         string // 提交人WeChat ID
	WXNickName     string // 提交人nickname
	Time           int64  // 提交的时间
	HandleInfo     string // 处理信息
	HandleType     int    // 处理的分类 0-提交 1-street推送给物业 2-gov约谈PM 3-询问
	// 4-回复 5-物业公司处理 6-用户评价 7-审核等级 8-street请求关闭
	// 9-同意关闭 10-street推送给政府 11-推送给法官
	EventLevel int    //事件等级  1-特急、2-加急、3-急
	Imgs       string // 图片列表，以,为分隔符
}

//FindTodayHandles 获取今日新增事件处理
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
		query = c.Find(bson.M{"index": index}).Sort("-time")
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
	query := c.Find(querys).Sort("-time")
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

func FindEventHandlesByIndexSortByTime(db *mgo.Database, index string) (error, []EventHandle) {
	c := db.C(EventHandleTableName)
	var result []EventHandle
	err := c.Find(bson.M{"index": index}).Sort("time").All(&result)
	return err, result
}
