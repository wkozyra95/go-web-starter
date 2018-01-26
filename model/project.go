package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	UserId bson.ObjectId `json:"userId" bson:"userId"`
}
