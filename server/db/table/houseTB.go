package table

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//HouseTabelName table name
const HouseTabelName = "House"

//House 建筑信息表
type House struct {
	BuildNo                string //建筑编号
	HouseOwner             string //房屋登记人
	XiaoQuName             string //房屋所属小区名
	HouseBuildNo           string //房屋楼栋号
	HouseNo                string //房屋门牌号
	HouseType              string //房屋类型
	Year                   string //建筑年代
	ZhuTiJieGouLieFeng     string //主体结构裂缝
	DiJiChenJiangBianXing  string //地基沉降变形
	ZhuTiJieGouQingXie     string //主体结构倾斜
	XuanTiaoJieGouPoHuai   string //悬挑结构破坏
	NvErQiangTuoLuo        string //女儿墙脱落
	WaiQiangMoHuiCengBoLuo string //外墙抹灰层剥落
	FangDingBianXing       string //房顶变形
	DiZhiZaiHai            string //地质灾害
	DiZhiZaiHaiZhiLi       string //地址灾害治理情况
	PaiShuiXitong          string //排水系统
	ZhuangXiuPoHuai        string //房屋内部装修主体结构变更破坏情况
	WeiGuiJiaCeng          string //违规搭建加层
	JiangDingLevel         string //房屋鉴定等级
	Img1                   string //图片
	Img2                   string
	Img3                   string
	Img4                   string
	Img5                   string //图片
	Img6                   string
	Img7                   string
	Img8                   string
	Img9                   string
}

//Houses 集合
type Houses struct {
	Houses []House
}

//InsertHouseInfoTB 插入
func InsertHouseInfoTB(db *mgo.Database, house House) string {
	c := db.C(HouseTabelName)
	count, err := c.Find(bson.M{"buildno": house.BuildNo}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "编号为" + house.BuildNo + "的建筑已经存在"
	}
	err = c.Insert(&house)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindHouses 查询小区信息集
func FindHouses(db *mgo.Database, info House, pageNo int, pageSize int) Houses {
	c := db.C(HouseTabelName)
	var result Houses
	var err error
	if info == (House{}) {
		err = c.Find(nil).All(&result.Houses)
	} else {
		err = c.Find(bson.M{"buildno": info.BuildNo}).All(&result.Houses)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
