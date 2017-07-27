package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//StreetTableName 街道名
const StreetTableName = "Street"

//Street 街道
type Street struct {
	Name           string //街道名
	PersonInCharge string //街道负责人
	Tel            string //街道联系电话
	AuthCode       string //授权码
	Intro          string //介绍
}

//Streets 查询街道信息，返回给前端的集合
type Streets struct {
	Streets []Street
}

//InsertStreet 插入
func InsertStreet(db *mgo.Database, street Street) interface{} {
	c := db.C(StreetTableName)
	count, err := c.Find(bson.M{"name": street.Name}).Count()
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "已存在街道\"" + street.Name + "\", 请重新设置街道名"}
	}
	err = c.Insert(&street)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindStreets 如果street.Name为""，则查询所有的街道信息,
func FindStreets(db *mgo.Database, street Street, pageNo int, pageSize int) interface{} {
	c := db.C(StreetTableName)
	var result []Street
	var err error
	if street.Name == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": street.Name}).All(&result)
	}
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindStreetDistinct 筛选出key值不重复的value
func FindStreetDistinct(db *mgo.Database, key string) interface{} {
	c := db.C(StreetTableName)
	var result []string
	err := c.Find(nil).Distinct(strings.ToLower(key), &result)
	if err != nil {
		log.Print(err)
		return gin.H{"erro": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//DelStreets 删除
func DelStreets(db *mgo.Database, names []string) interface{} {
	c := db.C(StreetTableName)
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
