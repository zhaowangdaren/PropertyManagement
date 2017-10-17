package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tealeg/xlsx"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//HouseTabelName table name
const HouseTabelName = "House"

//House 建筑信息表
type House struct {
	ID                  bson.ObjectId `bson:"_id"`
	BuildNo             string        //建筑编号
	Owner               string        //房产登记人
	StreetID            string        //所属街道
	CommunityID         string        //所属社区
	XQID                string        //房屋所属小区名
	HouseBuildNo        string        //房屋楼栋号
	HouseNo             string        //房屋门牌号
	HouseType           string        //房屋类型
	Year                string        //建筑年代
	UseChange           string        //使用变更
	MainCrack           string        //主体结构裂缝
	FoundationDown      string        //地基沉降变形
	MainSlant           string        //主体结构倾斜
	CantileverCrack     string        //悬挑结构破坏
	ParapetOff          string        //女儿墙脱落
	OuterLloatedCoatOff string        //外墙抹灰层剥落
	HouseDeform         string        //房顶变形
	Disaster            string        //地质灾害
	DisasterManage      string        //地址灾害治理情况
	DrainageSsystem     string        //排水系统
	InnerChange         string        //房屋内部装修主体结构变更破坏情况
	IllegalBuild        string        //违规搭建加层
	RankAppraisal       string        //房屋鉴定等级
	Imgs                string
}

//ImportHouseIndex 建筑信息表
type ImportHouseIndex struct {
	BuildNo             int //建筑编号
	Owner               int //房产登记人
	StreetID            int //所属街道
	CommunityID         int //所属社区
	XQID                int //房屋所属小区名
	HouseBuildNo        int //房屋楼栋号
	HouseNo             int //房屋门牌号
	HouseType           int //房屋类型
	Year                int //建筑年代
	UseChange           int //使用变更
	MainCrack           int //主体结构裂缝
	FoundationDown      int //地基沉降变形
	MainSlant           int //主体结构倾斜
	CantileverCrack     int //悬挑结构破坏
	ParapetOff          int //女儿墙脱落
	OuterLloatedCoatOff int //外墙抹灰层剥落
	HouseDeform         int //房顶变形
	Disaster            int //地质灾害
	DisasterManage      int //地址灾害治理情况
	DrainageSsystem     int //排水系统
	InnerChange         int //房屋内部装修主体结构变更破坏情况
	IllegalBuild        int //违规搭建加层
	RankAppraisal       int //房屋鉴定等级
	Imgs                int
}

//Houses 集合
type Houses struct {
	Houses []House
}

//CountHouses CountHouses
func CountHouses(db *mgo.Database) interface{} {
	c := db.C(HouseTabelName)
	count, err := c.Find(nil).Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": count}
}

//ImportHousesFromXML 从xml文件中导入house信息
func ImportHousesFromXML(db *mgo.Database, path string) error {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		glog.Error(err.Error())
		return err
	}
	for _, sheet := range xlFile.Sheets {
		streetName := sheet.Name
		streets, err := SearchStreetByName(db, streetName)
		if err != nil {
			return err
		}
		if len(streets) == 0 {
			glog.Error("未查询到街道：" + streetName)
			continue
		}
		for index, row := range sheet.Rows {
			if index == 0 {
				continue
			}
			var house House
			house.ID = bson.NewObjectId()
			house.BuildNo = row.Cells[1].String()
			house.Owner = row.Cells[2].String()
			house.HouseType = row.Cells[3].String()
			house.StreetID = streets[0].ID.Hex() // 4
			house.CommunityID = ""               // 5
			xqName := row.Cells[6].String()
			xqs, err := SearchXQByName(db, xqName)
			if err == nil && len(xqs) > 0 {
				house.XQID = xqs[0].ID.Hex()
				house.CommunityID = xqs[0].CommunityID
			}
			house.HouseBuildNo = row.Cells[7].String()
			house.HouseNo = row.Cells[8].String()
			house.Year = row.Cells[9].String()
			house.UseChange = row.Cells[10].String()
			house.MainCrack = row.Cells[11].String()
			house.FoundationDown = row.Cells[12].String()
			house.MainSlant = row.Cells[13].String()
			house.CantileverCrack = row.Cells[14].String()
			house.ParapetOff = row.Cells[15].String()
			house.OuterLloatedCoatOff = row.Cells[16].String()
			house.HouseDeform = row.Cells[17].String()
			house.Disaster = row.Cells[18].String()
			house.DisasterManage = row.Cells[19].String()
			house.DrainageSsystem = row.Cells[20].String()
			house.InnerChange = row.Cells[21].String()
			house.IllegalBuild = row.Cells[22].String()
			house.RankAppraisal = row.Cells[23].String()
			InsertHouse(db, house)
		}
	}
	return err
}

// func AnalysisRowIndex(row *xlsx.Row) ImportHouseIndex {
// 	for index, cell := range row.Cells {
// 		if cell.String() == ""
// 	}
// }

//InsertHouse 插入
func InsertHouse(db *mgo.Database, house House) interface{} {
	c := db.C(HouseTabelName)
	count, err := c.Find(bson.M{"buildno": house.BuildNo}).Count()
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "编号为" + house.BuildNo + "的House已经存在"}
	}
	house.ID = bson.NewObjectId()
	err = c.Insert(&house)
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//UpdateStreet update street
func UpdateHouse(db *mgo.Database, update House) interface{} {
	c := db.C(HouseTabelName)
	err := c.Update(bson.M{"_id": update.ID}, update)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindHouses 查询小区信息集
func FindHouses(db *mgo.Database, buildNO string, pageNo int, pageSize int) interface{} {
	c := db.C(HouseTabelName)
	var result []House
	var err error
	var query *mgo.Query
	if buildNO == "" {
		query = c.Find(nil)
	} else {
		query = c.Find(bson.M{"buildno": buildNO})
	}
	sum, err := query.Count()
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"builds": result, "sum": sum}}
}

//FindHouseKVs 通过一系列的K-V来查找记录
func FindHouseKVs(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(HouseTabelName)
	var result []House
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

func FindHouseByKVsPage(db *mgo.Database, kvs map[string]interface{}, pageNo int, pageSize int) interface{} {
	querys := make(map[string]interface{})
	for k, v := range kvs {
		querys[strings.ToLower(k)] = v
	}
	c := db.C(HouseTabelName)
	query := c.Find(querys)
	sum, _ := query.Count()
	var err error
	var result []House
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"houses": result, "sum": sum}}
}

//DelHouses 删除
func DelHouses(db *mgo.Database, ids []string) interface{} {
	c := db.C(HouseTabelName)
	var err error
	result := ""
	for _, v := range ids {
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
