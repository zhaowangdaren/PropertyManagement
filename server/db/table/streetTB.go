package table

import (
	"log"

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

//InsertStreetTB 插入
func InsertStreetTB(db *mgo.Database, street Street) string {
	c := db.C(StreetTableName)
	count, err := c.Find(bson.M{"name": street.Name}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "已存在街道名为" + street.Name + ", 请重新设置街道名"
	}
	err = c.Insert(&street)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindStreets 如果street.Name为""，则查询所有的街道信息,
func FindStreets(db *mgo.Database, street Street, pageNo int, pageSize int) []Street {
	c := db.C(StreetTableName)
	var result []Street
	var err error
	if street.Name == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": street.Name}).All(&result)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
