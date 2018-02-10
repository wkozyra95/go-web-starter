package mongo

import (
	"github.com/wkozyra95/go-web-starter/model"
	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	// ID ...
	ID         bson.ObjectId `json:"id" bson:"_id"`
	model.User `bson:",inline"`
}

// Project ...
type Project struct {
	// ID ...
	ID bson.ObjectId `json:"id" bson:"_id"`
	// UserID ...
	UserID        bson.ObjectId `json:"userId" bson:"userId"`
	model.Project `bson:",inline"`
}
