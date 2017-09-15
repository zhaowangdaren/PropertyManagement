package table

import (
	"errors"
	"fmt"
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

//ThreeDay 3天
const ThreeDay = 3 * 24 * 60 * 60

//SevenDay 7天
const SevenDay = 7 * 24 * 60 * 60

//PMKPI PMKPI 按小区进行物业考核
type PMKPI struct {
	XQID    string //所在小区ID
	Year    int
	Quarter int     //季度 1 2 3 4
	YWNo    int     //黄色警告次数
	RWNo    int     //红色警告次数
	IWNo    int     //重大警告次数
	Other   int     //其他的分数，由gov用户编辑
	Score   float32 //得分
}

func InsertPMKPI(db *mgo.Database, pmkpi PMKPI) error {
	c := db.C(PMKPITableName)
	err := c.Insert(pmkpi)
	return err
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

func FindPMKPIByKVsPage(db *mgo.Database, kvs map[string]interface{}, pageNo int,
	pageSize int) interface{} {
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

//FindPMKPI 根据年、季度来查询物业KPI
/**
* 查询db中是否存在该记录，若不存在则重新计算
 */
func FindPMKPI(db *mgo.Database, xqid string, year int, quarter int) (error, PMKPI) {
	c := db.C(PMKPITableName)
	var result PMKPI
	var err error
	isCurQuarter := IsCurQuarter(year, quarter)
	if isCurQuarter == -1 { //查询的是往季
		fmt.Println("查询的是往季")

		query := c.Find(bson.M{"xqid": xqid, "year": year, "quarter": quarter})
		count, err := query.Count()
		if count > 0 { //数据库中存在该条记录
			err = query.One(&result)
			return err, result
		} else { //数据库中不存在该条记录，则重新计算
			err, result = CalculatePMKPIQuarter(db, xqid, year, quarter)
			if err != nil {
				glog.Error(err.Error())
			} else {
				InsertPMKPI(db, result)
			}
			return err, result
		}
	}

	if isCurQuarter == 0 { // 查询的是当前季度
		fmt.Println("查询的是当前季度")

		err, result = CalculatePMKPIQuarter(db, xqid, year, quarter)
		if err != nil {
			glog.Error(err.Error())
		} else {
			query := c.Find(bson.M{"xqid": xqid, "year": year, "quarter": quarter})
			count, _ := query.Count()
			if count > 0 { //数据库中存在该条记录，则将已有的other值取出
				var pmkpi PMKPI
				err = query.One(&pmkpi)
				if err == nil {
					result.Other = pmkpi.Other
				}
			}
			UpsertPNKPI(db, result)
		}

		return err, result
	}

	if isCurQuarter == 1 { //查询未来季节
		fmt.Println("查询未来季节")

		result.XQID = xqid
		result.Year = year
		result.Quarter = quarter
		return err, result
	}

	return err, result
}

//IsCurQuarter @return -1：过去 0：当前季度 1:未来
func IsCurQuarter(year int, quarter int) int {
	thisYear := time.Now().Year()
	thisMonth := int(time.Now().Month())
	if thisYear > year {
		return -1
	}
	if thisYear < year {
		return 1
	}
	if year == thisYear {
		if (quarter * 3) < thisMonth { //往季
			return -1
		}
		if (quarter*3-2) <= thisMonth && thisMonth <= (quarter*3) {
			return 0
		}
		if thisMonth < (quarter*3 - 2) { //未来
			return 1
		}
	}
	return 0
}

func UpsertPNKPI(db *mgo.Database, pmkpi PMKPI) (*mgo.ChangeInfo, error) {
	c := db.C(PMKPITableName)
	changeInfo, err := c.Upsert(bson.M{"xqid": pmkpi.XQID, "year": pmkpi.Year, "quarter": pmkpi.Quarter}, pmkpi)
	return changeInfo, err
}

//UpdatePMKPI 更新
func UpdatePMKPI(db *mgo.Database, pmkpi PMKPI) error {
	c := db.C(PMKPITableName)
	err := c.Update(bson.M{"xqid": pmkpi.XQID, "year": pmkpi.Year, "quarter": pmkpi.Quarter}, pmkpi)
	return err
}

func UpdatePMKPIOther(db *mgo.Database, pmkpi PMKPI) error {
	c := db.C(PMKPITableName)
	err := c.Update(bson.M{"xqid": pmkpi.XQID, "year": pmkpi.Year,
		"quarter": pmkpi.Quarter}, bson.M{"$set": bson.M{"other": pmkpi.Other}})
	return err
}

//CalculatePMKPIQuarter 按季度计算pm kpi
// @quarter 季度  如果季度为0则查询全年的
func CalculatePMKPIQuarter(dbc *mgo.Database, xqid string, year int, quarter int) (error, PMKPI) {
	if year <= 0 || quarter <= 0 {
		return errors.New("年或季度错误，year=" + string(year) + "季度：" + string(quarter)), PMKPI{}
	}
	var startTime int64
	var endTime int64
	var tempDate time.Time
	if quarter <= 0 {
		tempDate = time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		startTime = tempDate.Unix()
		endTime = tempDate.AddDate(0, 1, -1).Unix()
	} else if quarter >= 1 && quarter <= 4 {
		tempDate = time.Date(year, time.Month(quarter*3), 1, 0, 0, 0, 0, time.UTC)
		endTime = tempDate.AddDate(0, 1, -1).Unix() //
		startTime = tempDate.AddDate(0, -2, 0).Unix()
	}
	err, pmkpi := CalculatePMKPIInTime(dbc, xqid, startTime, endTime)
	pmkpi.XQID = xqid
	pmkpi.Year = year
	pmkpi.Quarter = quarter
	return err, pmkpi
}

//CalculatePMKPI
func CalculatePMKPIInTime(dbc *mgo.Database, xqid string, startTime int64, endTime int64) (error, PMKPI) {
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
		fmt.Println("CalculatePMKPIWarnTimes", eventHandles)
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

//CalculatePMKPIWarnTimes 计算黄色、红色警告的次数
func CalculatePMKPIWarnTimes(eventHandles []EventHandle) (int, int) {
	var pushTime int64 //推送给物业的开始时间
	pushTime = 0
	pmHandleTime := time.Now().Unix()          //物业首次处理的时间
	for _, eventHandle := range eventHandles { // 遍历所有的handle，来记分
		if eventHandle.HandleType == 1 && pushTime == 0 { //推送给物业的时间
			fmt.Println("推送给物业", eventHandle.HandleType)
			pushTime = eventHandle.Time
			continue
		}
		// 物业首次处理的时间
		if eventHandle.AuthorCategory == 4 && pmHandleTime > eventHandle.Time {
			pmHandleTime = eventHandle.Time
			break
		}
	}
	if pushTime == 0 { //尚未推送给物业
		return 0, 0
	}
	//已推送给物业
	dur := pmHandleTime - pushTime
	if dur >= ThreeDay && dur < SevenDay {
		return 1, 0
	}
	if dur >= SevenDay {
		return 0, 1
	}
	return 0, 0
}
