package table

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const AnnounceTBName = "Announce"

type Announce struct {
	ID       bson.ObjectId `bson:"_id"`
	UserID   string        `json:"userId"`
	Classify string        `json:"classify"` // 文件分类
	MD5      string        `json:"md5"`      // 读取文件的文件名
	FileName string        `json:"fileName"` // 显示给用户文件名
	Time     int64         `json:"time"`     //文件上传时间
}

func InsertAnnounce(db *mgo.Database, announce Announce) error {
	c := db.C(AnnounceTBName)
	announce.ID = bson.NewObjectId()
	announce.Time = time.Now().Unix()
	err := c.Insert(&announce)
	return err
}

func InsertAnnounces(db *mgo.Database, announces []Announce) error {
	c := db.C(AnnounceTBName)
	var err error
	errInfo := ""
	for _, announce := range announces {
		announce.ID = bson.NewObjectId()
		announce.Time = time.Now().Unix()
		err = c.Insert(&announce)
		if err != nil {
			errInfo += "; " + err.Error()
		}
	}
	return errors.New(errInfo)
}

func DelAnnounce(db *mgo.Database, id string) error {
	c := db.C(AnnounceTBName)
	err := c.RemoveId(bson.ObjectIdHex(id))
	return err
}

func FindAnnounce(db *mgo.Database, fileName string, pageNo int, pageSize int) (ResultsPage, error) {
	c := db.C(AnnounceTBName)
	query := c.Find(bson.M{"filename": bson.M{"$regex": fileName, "$options": "$i"}})
	var err error
	var result ResultsPage
	result.Sum, err = query.Count()
	if err != nil {
		return result, err
	}
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result.List)
	} else {
		err = query.All(&result.List)
	}
	return result, err
}
