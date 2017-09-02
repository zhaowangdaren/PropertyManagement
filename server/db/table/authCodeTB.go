package table

import (
	"gopkg.in/mgo.v2/bson"
)

type AuthCode struct {
	id     bson.ObjectId `bson:"_id"`
	target string        // park-小区停车为绑定专用
	code   string        // auth code
	used   int           //0-未使用 1-已被使用
}
