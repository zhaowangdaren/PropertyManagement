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
	ID          bson.ObjectId `bson:"_id"`
	Name        string
	Address     string //地理位置
	StreetID    string //所属街道
	CommunityID string //管辖社区
	ContactName string //联系人姓名
	Tel         string //联系人号码
	Intro       string //介绍
}

//XiaoQus 小区信息集合
type XiaoQus struct {
	XiaoQus []XiaoQu
}

func CountXQs(db *mgo.Database) interface{} {
	c := db.C(XiaoQuTableName)
	count, err := c.Find(nil).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": count}
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
	info.ID = bson.NewObjectId()
	err = c.Insert(&info)
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//UpdateXQ update street
func UpdateXQ(db *mgo.Database, info XiaoQu) interface{} {
	c := db.C(XiaoQuTableName)
	err := c.Update(bson.M{"_id": info.ID}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindXiaoQus 查询小区信息集
func FindXiaoQus(db *mgo.Database, pageNo int, pageSize int) interface{} {
	c := db.C(XiaoQuTableName)
	var result []XiaoQu
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
	return gin.H{"error": 0, "data": gin.H{"xqs": result, "sum": sum}}
}

func FindXQ(db *mgo.Database, id string) XiaoQu {
	c := db.C(XiaoQuTableName)
	var result XiaoQu
	c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	return result
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
	return gin.H{"error": 0, "data": result}
}

//DelXQs 删除
func DelXQs(db *mgo.Database, names []string) interface{} {
	c := db.C(XiaoQuTableName)
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
func FindXQsByIDs(db *mgo.Database, ids []string) interface{} {
	c := db.C(XiaoQuTableName)
	var result []XiaoQu
	var err error
	for _, id := range ids {
		var info XiaoQu
		err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&info)
		if err != nil {
			result = append(result, (XiaoQu{}))
		} else {
			result = append(result, info)
		}
		err = nil
	}
	return gin.H{"error": 0, "data": result}
}

func SearchXQByName(db *mgo.Database, name string) interface{} {
	c := db.C(XiaoQuTableName)
	var result []XiaoQu
	err := c.Find(bson.M{"name": bson.M{"$regex": name, "$options": "$i"}}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
