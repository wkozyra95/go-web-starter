package handler

import (
	"github.com/wkozyra95/go-web-starter/model"
	"gopkg.in/mgo.v2/bson"
)

func UserRegister(
	user model.User,
	password string,
	ctx ActionContext,
) error {
	user.Id = bson.NewObjectId()
	user.GeneratePasswordHash(password)
	insertErr := ctx.DB.User().Insert(user)
	if insertErr != nil {
		return insertErr
	}
	return nil
}
