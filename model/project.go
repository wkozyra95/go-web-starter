package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id     bson.ObjectId
	UserId bson.ObjectId
}
