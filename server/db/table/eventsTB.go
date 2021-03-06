package table

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventTableName 表名
const EventTableName = "Event"

//Event 投诉事件表
type Event struct {
	Index         string //事件编号
	Complainant   string //投诉人
	OpenID        string // 投诉人wx open ID
	StreetID      string //街道
	CommunityID   string //社区
	XQID          string //投诉小区ID
	Status        int    //事件状态  -2-已关闭 -1-已撤销 0-居民提交 1-处理中 2-已处理待确认 3-已解决 4-未解决
	RequestClose  int    // 1-申请关闭
	NoticeGov     int    // 1-推送给政府
	NoticePM      int    // 1-已推送PM Company
	TalkAbout     int    // 1-Gov约谈PM
	EventLevel    int    //事件等级  1-特急、2-加急、3-急
	Type          string //事件基本类别
	Content       string //投诉内容
	Time          int64  //提交时间
	ToCourt       int    //0-不推送至法院 1-推送至法院
	CourtAccepted int    //0-法官未受理 1-已受理
	Imgs          string //图片表格，以,为分割符
	Tel           string //联系电话
}

//EventNum 不同事件类型的数量
type EventNum struct {
	Type string
	Num  int
}

func getEventOverivewQueryGroup(year int, month int, userType int) interface{} {
	// startTime := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	// endTime := startTime.AddDate(0, 1, -1)
	// switch userType {
	// case 2: // gov
	return []bson.M{
		bson.M{
			"$match": bson.M{
				"status": 1,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": nil,
				"Sum": bson.M{"$sum": 1},
				"Unhandle": bson.M{
					"$sum": 1,
				},
			},
		},
	}
	// }
}
func QueryEventOverviewV2(db *mgo.Database, year int, month int, userType int) error {
	startTime := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	endTime := startTime.AddDate(0, 1, -1)
	c := db.C(EventTableName)
	//查总数
	selector := []bson.M{
		bson.M{
			"$match": []bson.M{
				bson.M{
					"$and": []bson.M{
						bson.M{"time": bson.M{"$gte": startTime.Unix()}},
						bson.M{"time": bson.M{"$lte": endTime.Unix()}},
					},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{"StreetID": "$streetid"},
				"sum": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id": 1,
			},
		},
	}
	pipe := c.Pipe(selector)
	var resultGroup []EventOverviewGroup
	err := pipe.All(&resultGroup)
	if err != nil {
		return err
	}
	var result []EventOverview
	for _, g := range resultGroup {
		result = append(result, EventOverview{
			StreetID: g.ID["StreetID"],
			Sum:      g.Count,
		})
	}
	// 查未解决
	selector = []bson.M{
		bson.M{
			"$match": []bson.M{
				bson.M{
					"$and": []bson.M{
						bson.M{"time": bson.M{"$gte": startTime.Unix()}},
						bson.M{"time": bson.M{"$lte": endTime.Unix()}},
						bson.M{"status": 4},
					},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{"StreetID": "$streetid"},
				"sum": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id": 1,
			},
		},
	}
	pipe = c.Pipe(selector)
	err = pipe.All(&resultGroup)
	if err != nil {
		return err
	}
	// for index, g := range resultGroup {
	// 	result[]
	// }
	return err
}

// QueryEventOverview 按年份、月份、人员类型
//userType 1-admin 2-gov 3-street
func QueryEventOverview(db *mgo.Database, streetID string, year int, month int,
	userType int) (EventOverview, error) {
	c := db.C(EventTableName)
	startTime := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	endTime := startTime.AddDate(0, 1, -1)
	query := bson.M{
		"$and": []bson.M{
			bson.M{"streetid": streetID},
			bson.M{"time": bson.M{"$gte": startTime.Unix()}},
			bson.M{"time": bson.M{"$lte": endTime.Unix()}},
		},
	}
	sum, err := c.Find(query).Count()
	var eventOverview EventOverview
	if err != nil {
		return eventOverview, err
	}
	eventOverview.Sum = sum
	query = bson.M{
		"$and": []bson.M{
			bson.M{"streetid": streetID},
			bson.M{"time": bson.M{"$gte": startTime.Unix()}},
			bson.M{"time": bson.M{"$lte": endTime.Unix()}},
			bson.M{"status": 4},
		},
	}
	unsolved, err := c.Find(query).Count()
	if err != nil {
		return eventOverview, err
	}
	eventOverview.Unsolved = unsolved
	query = bson.M{
		"$and": []bson.M{
			bson.M{"streetid": streetID},
			bson.M{"time": bson.M{"$gte": startTime.Unix()}},
			bson.M{"time": bson.M{"$lte": endTime.Unix()}},
			bson.M{"status": 3},
		},
	}
	solved, err := c.Find(query).Count()
	if err != nil {
		return eventOverview, err
	}
	eventOverview.Solved = solved
	switch userType {
	case 2: //gov
		query = bson.M{
			"$and": []bson.M{
				bson.M{"streetid": streetID},
				bson.M{"time": bson.M{"$gte": startTime.Unix()}},
				bson.M{"time": bson.M{"$lte": endTime.Unix()}},
			},
			"$or": []bson.M{
				bson.M{"talkabout": 1},
				bson.M{"tocourt": 1},
				bson.M{"status": -2},
			},
		}
		handled, err := c.Find(query).Count()
		if err != nil {
			return eventOverview, err
		}
		eventOverview.Handled = handled

		eventOverview.Unhandle = sum - handled
		break
	case 3: // street
		query = bson.M{
			"$and": []bson.M{
				bson.M{"streetid": streetID},
				bson.M{"time": bson.M{"$gte": startTime.Unix()}},
				bson.M{"time": bson.M{"$lte": endTime.Unix()}},
			},
			"$or": []bson.M{
				bson.M{"eventlevel": bson.M{"$ne": 0}},
				bson.M{"requestclose": 1},
				bson.M{"noticegov": 1},
				bson.M{"notivepm": 1},
			},
		}
		handled, err := c.Find(query).Count()
		if err != nil {
			return eventOverview, err
		}
		eventOverview.Handled = handled
		eventOverview.Unhandle = sum - handled
		break
	}
	eventOverview.Year = year
	eventOverview.Month = month
	if month%3 > 0 {
		eventOverview.Quarter = month/3 + 1
	} else {
		eventOverview.Quarter = month / 3
	}
	return eventOverview, err
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
	eventHandle := EventHandle{
		event.Index,
		event.XQID,
		0,
		event.OpenID,
		event.OpenID,
		"",
		event.Time,
		event.Content,
		0,
		0,
		"",
	}
	_, err = InsertEventHandle(db, eventHandle)
	if err != nil {
		glog.Error(err.Error())
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

func UpdateEventStatus(db *mgo.Database, index string, status int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"status": status}})
	return err
}

func UpdateEventRequestClose(db *mgo.Database, index string, requestClose int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"requestclose": requestClose}})
	return err
}
func UpdateEventLevel(db *mgo.Database, index string, eventLevel int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"eventlevel": eventLevel}})
	return err
}

func UpdateEventTalkAbout(db *mgo.Database, index string, talkAbout int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"talkabout": talkAbout}})
	return err
}

func UpdateEventNoticeGov(db *mgo.Database, index string, noticeGov int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"noticegov": noticeGov}})
	return err
}

func UpdateEventToCourt(db *mgo.Database, index string, toCourt int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"tocourt": toCourt}})
	return err
}

func UpdateEventCourtAccepted(db *mgo.Database, index string, courtAccepted int) error {
	c := db.C(EventTableName)
	err := c.Update(bson.M{"index": index}, bson.M{"$set": bson.M{"courtaccepted": courtAccepted}})
	return err
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
	var query *mgo.Query
	if index == "" {
		query = c.Find(nil)
	} else {
		query = c.Find(bson.M{"index": index})
	}
	sum, _ := query.Count()
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"events": result, "sum": sum}}
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

func FindEventKVsPage(db *mgo.Database, kvs map[string]interface{}, pageNo int, pageSize int) interface{} {
	querys := make(map[string]interface{})
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
		querys[strings.ToLower(k)] = v
	}
	c := db.C(EventTableName)
	var result []Event
	// query := c.Find(querys)
	var query *mgo.Query
	if hasTime {
		selector := bson.M{
			"$and": []bson.M{
				bson.M{"time": bson.M{"$gte": startTime}},
				bson.M{"time": bson.M{"$lte": endTime}},
				querys,
				// bson.M{"status": bson.M{"$gte": 0}},
			},
		}
		query = c.Find(selector).Sort("-time")
	} else {
		selector := bson.M{
			"$and": []bson.M{
				querys,
				// bson.M{"status": bson.M{"$gte": 0}},
			},
		}
		query = c.Find(selector).Sort("-time")
	}
	sum, _ := query.Count()
	var err error
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}

	// err := c.Find(querys).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"events": result, "sum": sum}}
}

func FindEventKVsPageForGov(db *mgo.Database, kvs map[string]interface{}, pageNo int, pageSize int) interface{} {
	querys := make(map[string]interface{})
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
		querys[strings.ToLower(k)] = v
	}
	c := db.C(EventTableName)
	var result []Event
	// query := c.Find(querys)
	var query *mgo.Query
	if hasTime {
		selector := bson.M{
			"$and": []bson.M{
				bson.M{"time": bson.M{"$gte": startTime}},
				bson.M{"time": bson.M{"$lte": endTime}},
				querys,
				// bson.M{"status": bson.M{"$gte": 0}},
			},
			"$or": []bson.M{
				bson.M{"noticegov": 1},
				bson.M{"requestclose": 1},
			},
		}
		query = c.Find(selector).Sort("-time")
	} else {
		selector := bson.M{
			"$and": []bson.M{
				querys,
				// bson.M{"status": bson.M{"$gte": 0}},
			},
			"$or": []bson.M{
				bson.M{"noticegov": 1},
				bson.M{"requestclose": 1},
			},
		}
		query = c.Find(selector).Sort("-time")
	}
	sum, _ := query.Count()
	var err error
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}

	// err := c.Find(querys).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"events": result, "sum": sum}}
}

func FindEventKV(db *mgo.Database, key string, value string,
	pageNo int, pageSize int) interface{} {
	c := db.C(EventTableName)
	var result []Event
	var err error
	query := c.Find(bson.M{strings.ToLower(key): value})
	sum, _ := query.Count()
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"events": result, "sum": sum}}
}

func FindEventsByXQIDInTimeSortByType(db *mgo.Database, xqid string,
	startTime int64, endTime int64) (error, []Event) {
	c := db.C(EventTableName)
	var result []Event
	selector := bson.M{
		"$and": []bson.M{
			bson.M{"time": bson.M{"$gte": startTime}},
			bson.M{"time": bson.M{"$lte": endTime}},
			bson.M{"xqid": xqid},
		},
	}
	err := c.Find(selector).Sort("type").All(&result)
	return err, result
}

func DelEvent(db *mgo.Database, index string) error {
	c := db.C(EventTableName)
	err := c.Remove(bson.M{"index": index})
	return err
}

//CountDiffKeyEvents 统计不同key-value的数量
func CountDiffKeyEvents(db *mgo.Database, key string, maxRecs int) interface{} {
	c := db.C(EventTableName)
	key = strings.ToLower(key)
	pipe := c.Pipe([]bson.M{
		{
			"$group": bson.M{
				"_id": bson.M{key: "$" + key},
				"num": bson.M{"$sum": 1},
			},
		}, {
			"$sort": bson.M{
				"num": -1,
			},
		},
	})
	result := []bson.M{}
	err := pipe.All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	resultLen := len(result)
	if maxRecs != 0 && resultLen >= maxRecs {
		return gin.H{"error": 0, "data": result[0:maxRecs]}
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
