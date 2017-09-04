package table

import "gopkg.in/mgo.v2/bson"

type Company struct {
	ID           bson.ObjectId `bson:"_id"`
	Name         string
	Logo         string
	Charger      string
	Tel          string
	JoinTime     int64
	MainContent  string
	WorkSchedule string //服务时间
}
