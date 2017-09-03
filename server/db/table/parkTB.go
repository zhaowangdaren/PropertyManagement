package table

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ParkTableName = "park"

//Park park
type Park struct {
	ID      bson.ObjectId `bson:"_id"`
	XQID    string        // 小区ID
	Sum     int           // 总停车位
	FreeNum int           // 空闲车位
}

//InsertPark
func InsertPark(db *mgo.Database, info Park) interface{} {
	c := db.C(ParkTableName)
	count, err := c.Find(bson.M{"_id": info.ID}).Count()
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "该小区的停车位信息已存在"}
	}
	info.ID = bson.NewObjectId()
	err = c.Insert(&info)
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func UpdatePark(db *mgo.Database, info Park) interface{} {
	c := db.C(ParkTableName)
	err := c.Update(bson.M{"_id": info.ID}, info)
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
	if result.ID == "" {
		glog.Warning("没有找到id为" + id + "的停车信息")
		return gin.H{"error": 1, "data": "没有找到id为" + id + "的停车信息"}
	}
	result.FreeNum = freeeNum
	err = c.Update(bson.M{"_id": result.ID}, result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func FindParkByXQID(db *mgo.Database, xqid string) interface{} {
	c := db.C(ParkTableName)
	var result Park
	err := c.Find(bson.M{"xqid": xqid}).One(&result)
	fmt.Println("park", result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	if result.ID == "" {
		glog.Warning("没有找到xqid为" + xqid + "的停车信息")
		return gin.H{"error": 1, "data": "没有查询到该小区的停车信息"}
	}
	return gin.H{"error": 0, "data": result}
}
