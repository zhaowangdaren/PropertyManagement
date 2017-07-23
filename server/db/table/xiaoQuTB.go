package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//XiaoQuTableName table name
const XiaoQuTableName = "XiaoQu"

//XiaoQu 小区信息
type XiaoQu struct {
	Name        string
	Address     string //地理位置
	Street      string //所属街道
	Community   string //管辖社区
	ContactName string //联系人姓名
	Tel         string //联系人号码
	Intro       string //介绍
}

//XiaoQus 小区信息集合
type XiaoQus struct {
	XiaoQus []XiaoQu
}

//InsertXQ 插入user
func InsertXQ(db *mgo.Database, info XiaoQu) interface{} {
	c := db.C(XiaoQuTableName)
	count, err := c.Find(bson.M{"name": info.Name}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "已存在country：" + info.Name + ", 请重新设置name"}
	}
	err = c.Insert(&info)
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindXiaoQus 查询小区信息集
func FindXiaoQus(db *mgo.Database, xiaoQu XiaoQu, pageNo int, pageSize int) XiaoQus {
	c := db.C(XiaoQuTableName)
	var result XiaoQus
	var err error
	if xiaoQu == (XiaoQu{}) {
		err = c.Find(nil).All(&result.XiaoQus)
	} else {
		err = c.Find(bson.M{"name": xiaoQu.Name}).All(&result.XiaoQus)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//FindXQDistincts 查找key对应的不同values
func FindXQDistincts(db *mgo.Database, key string) []string {
	c := db.C(XiaoQuTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(key, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func FindXQKVs(db *mgo.Database, kvs map[string]interface{}) XiaoQus {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(XiaoQuTableName)
	var result XiaoQus
	err := c.Find(query).All(&result.XiaoQus)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
