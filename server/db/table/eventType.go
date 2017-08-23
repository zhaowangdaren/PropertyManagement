package table

import (
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventType 事件类别
// 事件类别：
// 	物业管理管理费
// 	业主财产损失赔偿
// 	公用设施伤人
// 	无因管理
// 	小区公用部位侵权
// 	小区公用部位出租营利
// 	小区停车费纠纷
// 	其他类
type EventType struct {
	Type string
}

//EventTypeTableName table name
const EventTypeTableName = "EventType"

//InsertEventType insert
func InsertEventType(db *mgo.Database, eventType string) interface{} {
	c := db.C(EventTypeTableName)
	count, err := c.Find(bson.M{"type": eventType}).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}

	if count > 0 {
		return gin.H{"error": 1, "data": "该类型已经存在，请重新设置类型"}
	}

	info := EventType{eventType}
	err = c.Insert(&info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

//DelEventType DelEventType
func DelEventType(db *mgo.Database, evenTypeStr string) interface{} {
	c := db.C(EventTypeTableName)
	err := c.Remove(bson.M{"type": evenTypeStr})
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

//FindEventTypes Find All EventTypes
func FindEventTypes(db *mgo.Database) interface{} {
	c := db.C(EventTypeTableName)
	var result []EventType
	err := c.Find(nil).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
