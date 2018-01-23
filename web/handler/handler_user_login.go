package handler

import (
	"fmt"
	"github.com/wkozyra95/go-web-starter/model"
	"gopkg.in/mgo.v2/bson"
)

func UserLogin(username string, password string, ctx ActionContext) (string, error) {
	user := model.User{}
	userErr := ctx.DB.User().Find(bson.M{
		"username": username,
		"password": password,
	}).One(&user)
	if userErr != nil {
		return "", userErr
	}

	if !user.ValidatePassword(password) {
		return "", fmt.Errorf("")
	}

	return "token", nil
}
