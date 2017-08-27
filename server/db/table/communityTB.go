package table //社区

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CommunityTableName 社区表格名称
const CommunityTableName = "Community"

//Community 社区
type Community struct {
	ID             bson.ObjectId `bson:"_id"`
	Name           string        //社区名
	PersonInCharge string        //负责人
	Tel            string        //社区联系电话
	StreetID       string        //所属街道
	Intro          string        //介绍
}

//Communities 社区集
type Communities struct {
	Communities []Community
}

func CountCommunities(db *mgo.Database) interface{} {
	c := db.C(CommunityTableName)
	count, err := c.Find(nil).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": count}
}

//InsertCommunity 插入
func InsertCommunity(db *mgo.Database, comm Community) interface{} {
	c := db.C(CommunityTableName)
	count, err := c.Find(bson.M{"name": comm.Name}).Count()
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "已存在community" + comm.Name + ", 请重新设置name"}
	}
	comm.ID = bson.NewObjectId()
	err = c.Insert(&comm)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//UpdateCommunity update street
func UpdateCommunity(db *mgo.Database, comm Community) interface{} {
	c := db.C(CommunityTableName)
	err := c.Update(bson.M{"_id": comm.ID}, comm)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindCommunities 查询社区信息
func FindCommunities(db *mgo.Database, pageNo int, pageSize int) interface{} {
	c := db.C(CommunityTableName)
	var result []Community
	var err error
	query := c.Find(nil)
	sum, _ := query.Count()
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"communities": result, "sum": sum}}
}

//FindCommunitiesKV 依据传入的kv查询
func FindCommunitiesKV(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(CommunityTableName)
	var result []Community
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindCommunityDistincts 查找key对应的不同values
func FindCommunityDistincts(db *mgo.Database, key string) interface{} {
	c := db.C(CommunityTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(strings.ToLower(key), &result)
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindCommunityByStreetName 查询社区信息
func FindCommunityByStreetName(db *mgo.Database, streetName string) interface{} {
	c := db.C(CommunityTableName)
	var result []Community
	err := c.Find(bson.M{"streetname": streetName}).All(&result)
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//DelCommunities 根据名称删除记录
func DelCommunities(db *mgo.Database, names []string) interface{} {
	c := db.C(CommunityTableName)
	var err error
	result := ""
	for _, v := range names {
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(v)})
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

//FindCommunitiesByIDs Find Streets
func FindCommunitiesByIDs(db *mgo.Database, ids []string) interface{} {
	c := db.C(CommunityTableName)
	var result []Community
	for _, id := range ids {
		if id == "" {
			continue
		}
		var info Community
		c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&info)
		if info != (Community{}) {
			result = append(result, info)
		}
	}
	return gin.H{"error": 0, "data": result}
}
