package handler

import (
	"fmt"
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/db"
	"gopkg.in/mgo.v2/bson"
)

func UserLogin(
	username string,
	password string,
	generateToken func(bson.ObjectId) (string, error),
	ctx ActionContext,
) (string, error) {
	validateErr := userLoginValidate(username, password)
	if validateErr != nil {
		return "", validateErr
	}

	formErr := errors.EmptyMessageError()
	formErr.Json["form"] = errors.TextError("Unknown combiantion of username and password")

	user := model.User{}
	userErr := ctx.DB.User().Find(bson.M{db.UserIdKeyUsername: username}).One(&user)
	if userErr != nil {
		formErr.Msg = fmt.Sprintf("user don't exist [%s]", userErr.Error())
		return "", formErr
	}

	if !user.ValidatePassword(password) {
		formErr.Msg = "invalid password"
		return "", formErr
	}

	token, tokenErr := generateToken(user.Id)
	if tokenErr != nil {
		return "", internalServerErr(fmt.Sprintf("token error %s", tokenErr.Error()))
	}
	return token, nil
}

func userLoginValidate(username, password string) error {
	formErr := errors.NewMessageError("form error", http.StatusBadRequest)

	if username == "" {
		formErr.Json["username"] = errors.TextError("Username can't empty")
	}
	if password == "" {
		formErr.Json["password"] = errors.TextError("Password can't empty")
	}
	if len(formErr.Json) > 0 {
		return formErr
	}
	return nil
}
