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
func FindXiaoQus(db *mgo.Database, xiaoQu XiaoQu, pageNo int, pageSize int) interface{} {
	c := db.C(XiaoQuTableName)
	var result []XiaoQu
	var err error
	if xiaoQu == (XiaoQu{}) {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": xiaoQu.Name}).All(&result)
	}
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}

}

//FindXQDistincts 查找key对应的不同values
func FindXQDistincts(db *mgo.Database, key string) interface{} {
	c := db.C(XiaoQuTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(key, &result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindXQKVs 通过一系列的K-V来查找记录
func FindXQKVs(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(XiaoQuTableName)
	var result []XiaoQu
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

//DelXQs 删除
func DelXQs(db *mgo.Database, names []string) interface{} {
	c := db.C(XiaoQuTableName)
	var err error
	result := ""
	for _, v := range names {
		err = c.Remove(bson.M{"name": v})
		if err != nil {
			log.Println(err.Error())
			result += err.Error()
			err = nil
		}
	}
	if result != "" {
		return gin.H{"error": 1, "data": result}
	}
	return gin.H{"error": 0, "data": Succ}
}
