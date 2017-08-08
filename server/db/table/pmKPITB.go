package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PMKPITableName table name
const PMKPITableName = "PMKPI"

//PMKPI PMKPI
type PMKPI struct {
	PMID        string //PM ID
	Name        string // PM 公司名
	StreetID    string //所在街道ID
	CommunityID string //所在社区ID
	XQID        string //所在小区ID
	Year        int
	Quarter     int     //季度
	YWNo        int     //黄色警告次数
	RWNo        int     //红色警告次数
	IWNo        int     //重大警告次数
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
