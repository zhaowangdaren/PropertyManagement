package table

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ParkTableName = "park"

type Park struct {
	id      bson.ObjectId `bson:"_id"`
	xqID    string        // 小区ID
	sum     int           // 总停车位
	freeNum int           // 空闲车位
}

//InsertPark
func InsertPark(db *mgo.Database, info Park) interface{} {
	c := db.C(ParkTableName)
	count, err := c.Find(bson.M{"xqid": info.xqID}).Count()
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "该小区的停车位信息已存在"}
	}
	err = c.Insert(&info)
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func UpdatePark(db *mgo.Database, info Park) interface{} {
	c := db.C(ParkTableName)
	err := c.Update(bson.M{"_id": info.id}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func UpdateParkFreeNum(db *mgo.Database, id string, freeeNum int) interface{} {
	c := db.C(ParkTableName)
	var result Park
	err := c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	if result.id == "" {
		glog.Warning("没有找到id为" + id + "的停车信息")
		return gin.H{"error": 1, "data": "没有找到id为" + id + "的停车信息"}
	}
	result.freeNum = freeeNum
	err = c.Update(bson.M{"_id": result.id}, result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func FindParkByXQID(db *mgo.Database, xqid string) interface{} {
	c := db.C(ParkTableName)
	var result Park
	err := c.Find(bson.M{"xqid": xqid}).One(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	if result.id == "" {
		glog.Warning("没有找到xqid为" + xqid + "的停车信息")
		return gin.H{"error": 1, "data": "没有查询到该小区的停车信息"}
	}
	return gin.H{"error": 0, "data": result}
}
