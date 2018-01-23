package handler

import (
	"github.com/wkozyra95/go-web-starter/model/db"
	"gopkg.in/mgo.v2/bson"
)

type ActionContext struct {
	DB     db.DB
	UserId bson.ObjectId
}
