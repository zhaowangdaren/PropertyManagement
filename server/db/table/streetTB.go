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
	ID             bson.ObjectId `bson:"_id"`
	Name           string        //街道名
	PersonInCharge string        //街道负责人
	Tel            string        //街道联系电话
	AuthCode       string        //授权码
	Intro          string        //介绍
}

//Streets 查询街道信息，返回给前端的集合
type Streets struct {
	Streets []Street
}

//CountStreets CountStreets
func CountStreets(db *mgo.Database) interface{} {
	c := db.C(StreetTableName)
	count, err := c.Find(nil).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": count}
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
	street.ID = bson.NewObjectId()
	err = c.Insert(&street)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//UpdateStreet update street
func UpdateStreet(db *mgo.Database, street Street) interface{} {
	c := db.C(StreetTableName)
	err := c.Update(bson.M{"_id": street.ID}, street)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindStreetsByIDs Find Streets
func FindStreetsByIDs(db *mgo.Database, ids []string) interface{} {
	c := db.C(StreetTableName)
	var result []Street
	for _, id := range ids {
		if id == "" {
			continue
		}
		var street Street
		c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&street)
		if street != (Street{}) {
			result = append(result, street)
		}
	}
	return gin.H{"error": 0, "data": result}
}

//FindStreets 如果street.Name为""，则查询所有的街道信息,
func FindStreets(db *mgo.Database, pageNo int, pageSize int) interface{} {
	c := db.C(StreetTableName)
	var result []Street
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
	return gin.H{"error": 0, "data": gin.H{"streets": result, "sum": sum}}
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

func FindAllStreets(db *mgo.Database) ([]Street, error) {
	c := db.C(StreetTableName)
	var result []Street
	err := c.Find(nil).Sort("name").All(&result)
	return result, err
}

//DelStreets 删除
func DelStreets(db *mgo.Database, names []string) interface{} {
	c := db.C(StreetTableName)
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

func SearchStreetByName(db *mgo.Database, name string) interface{} {
	c := db.C(StreetTableName)
	var result []Street
	err := c.Find(bson.M{"name": bson.M{"$regex": name, "$options": "$i"}}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
