package handler

import (
	"github.com/wkozyra95/go-web-starter/model/db"
	"gopkg.in/mgo.v2/bson"
)

func validateReadRights(id bson.ObjectId, collection db.Collection, context ActionContext) (bool, error) {
	document := db.Document{}
	documentErr := collection.FindId(id).One(&document)
	if documentErr != nil {
		return false, documentErr
	}

	return document.UserId == context.UserId, nil
}

func validateWriteRights(id bson.ObjectId, collection db.Collection, context ActionContext) (bool, error) {
	document := db.Document{}
	documentErr := collection.FindId(id).One(&document)
	if documentErr != nil {
		return false, documentErr
	}
	return document.UserId == context.UserId, nil
}
