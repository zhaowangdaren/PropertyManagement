package table

import (
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 社区服务的第一级
const NodeTable = "node"

//Node
type Node struct {
	ID     bson.ObjectId `bson:"_id"`
	Father string
	Lable  string
	Value  string
	Sons   string //以逗号为分隔符
}

// InsertNode InsertNode
func InsertNode(db *mgo.Database, node Node) interface{} {
	c := db.C(NodeTable)
	node.ID = bson.NewObjectId()

	if node.Father == "" { // 无父节点
		err := c.Insert(node)
		if err != nil {
			return gin.H{"error": 1, "data": err.Error()}
		}
		return gin.H{"error": 0, "data": node.ID.String()}
	}
	// 有父节点
	var father Node
	err := c.Find(bson.M{"_id": node.Father}).One(&father)
	if father.Lable == "" {
		return gin.H{"error": 1, "data": "父节点信息有误"}
	}
	err = c.Insert(node)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	// Update Father Node
	father.Sons += "," + node.ID.String()
	err = c.Update(bson.M{"_id": father.ID}, father)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": node.ID.String()}
}

func UpdateNode(db *mgo.Database, node Node) interface{} {
	c := db.C(NodeTable)
	err := c.Update(bson.M{"_id": node.ID}, node)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}
