package handler

import (
	"fmt"
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/db"
	"gopkg.in/mgo.v2/bson"
)

func UserRegister(user model.User, password string, ctx ActionContext) error {
	validateErr := userRegisterValidate(user, password)
	if validateErr != nil {
		return validateErr
	}
	existingUsers := []model.User{}
	getErr := ctx.DB.User().Find(bson.M{
		db.OR: []bson.M{
			bson.M{db.UserIdKeyUsername: user.Username},
			bson.M{db.UserIdKeyEmail: user.Email},
		},
	}).All(&existingUsers)
	if getErr != nil {
		msg := fmt.Sprintf("db error %s", getErr.Error())
		return internalServerErr(msg)
	}
	for _, existingUser := range existingUsers {
		if existingUser.Email == user.Email {
			matchErr := errors.NewMessageError("register email already used", http.StatusBadRequest)
			matchErr.Json["email"] = errors.TextError("This email is already in use")
			return matchErr
		}
		if existingUser.Username == user.Username {
			matchErr := errors.NewMessageError("register username already used", http.StatusBadRequest)
			matchErr.Json["username"] = errors.TextError("This username is already in use")
			return matchErr
		}
	}

	user.Id = bson.NewObjectId()
	user.GeneratePasswordHash(password)
	insertErr := ctx.DB.User().Insert(user)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

func userRegisterValidate(user model.User, password string) error {
	formErr := errors.NewMessageError("form error", http.StatusBadRequest)

	if user.Email == "" {
		formErr.Json["email"] = errors.TextError("Email can't empty")
	}

	if user.Username == "" {
		formErr.Json["username"] = errors.TextError("Username can't empty")
	}

	if len(password) < 8 {
		formErr.Json["password"] = errors.TextError("Password is to short, you need at least 8 characters.")
	}

	if password == "" {
		formErr.Json["password"] = errors.TextError("Password can't empty")
	}

	if len(formErr.Json) > 0 {
		return formErr
	}
	return nil
}
