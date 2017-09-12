package table

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PMKPITableName table name
const PMKPITableName = "PMKPI"

//PMKPI PMKPI 按小区进行物业考核
type PMKPI struct {
	Name        string // PM 公司名
	StreetID    string //所在街道ID
	CommunityID string //所在社区ID
	XQID        string //所在小区ID
	Year        int
	Quarter     int     //季度
	YWNo        int     //黄色警告次数
	RWNo        int     //红色警告次数
	IWNo        int     //重大警告次数
	Other       int     //其他的分数，由gov用户编辑
	Score       float32 //得分
}

//FindPMKPIsByName 通过Name查询PMKPI
func FindPMKPIsByName(db *mgo.Database, name string, pageNo int, pageSize int) interface{} {
	c := db.C(PMKPITableName)
	var result []PMKPI
	var err error
	if name == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": name}).All(&result)
	}
	if err != nil {
		log.Println(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindPMKPIByKVs 通过一系列的K-V来查找记录
func FindPMKPIByKVs(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(PMKPITableName)
	var result []PMKPI
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if result == nil {
		return gin.H{"error": 1, "data": "没有查询到结果"}
	}
	return gin.H{"error": 0, "data": result}
}

func FindPMKPIByKVsPage(db *mgo.Database, kvs map[string]interface{}, pageNo int, pageSize int) interface{} {
	querys := make(map[string]interface{})
	for k, v := range kvs {
		querys[strings.ToLower(k)] = v
	}
	c := db.C(PMKPITableName)
	query := c.Find(querys)
	sum, _ := query.Count()
	var err error
	var result []PMKPI
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"pmkpis": result, "sum": sum}}
}

//CalculatePMKPI
func CalculatePMKPI(dbc *mgo.Database, xqid string, startTime int64, endTime int64) (error, PMKPI) {
	var pmkpi PMKPI
	err, events := FindEventsByXQIDInTimeSortByType(dbc, xqid, startTime, endTime) //查询出该xq的所有事件
	if err != nil {
		return err, pmkpi
	}
	if len(events) == 0 {
		return nil, pmkpi
	}
	tempType := ""
	tempTypeCount := 0
	for _, event := range events {
		index := event.Index
		tempErr, eventHandles := FindEventHandlesByIndexSortByTime(dbc, index)
		if tempErr != nil {
			glog.Error("CalculatePMKPI " + tempErr.Error())
			continue
		}
		// 统计黄色、红色警告次数
		yellowWarn, redWarn := CalculatePMKPIWarnTimes(eventHandles)
		pmkpi.YWNo += yellowWarn
		pmkpi.RWNo += redWarn

		//统计同一类型投诉
		if tempType != event.Type {
			if tempTypeCount > 3 { //同一问题被投诉3次
				pmkpi.IWNo++
			}
			tempType = event.Type
			tempTypeCount = 0
		} else {
			tempTypeCount++
		}
	}
	return err, pmkpi
}

const ThreeDay = 3 * 24 * 60 * 60
const SevenDay = 7 * 24 * 60 * 60

//CalculatePMKPIWarnTimes 计算黄色、红色警告的次数
func CalculatePMKPIWarnTimes(eventHandles []EventHandle) (int, int) {
	var pushTime int64 //推送给物业的开始时间
	pushTime = 0
	pmHandleTime := time.Now().Unix()          //物业首次处理的时间
	for _, eventHandle := range eventHandles { // 遍历所有的handle，来记分
		if eventHandle.HandleType == 1 && pushTime == 0 { //推送给物业的时间
			pushTime = eventHandle.Time
			continue
		}
		// 物业首次处理的时间
		if eventHandle.AuthorCategory == 4 && pmHandleTime > eventHandle.Time {
			pmHandleTime = eventHandle.Time
			break
		}
	}
	dur := pmHandleTime - pushTime
	if dur >= ThreeDay && dur < SevenDay {
		return 1, 0
	}
	if dur >= SevenDay {
		return 0, 1
	}
	return 0, 0
}
