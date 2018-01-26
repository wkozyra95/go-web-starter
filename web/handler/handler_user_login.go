package handler

import (
	"fmt"
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func UserLogin(
	username string,
	password string,
	generateToken func(bson.ObjectId) (string, error),
	ctx ActionContext,
) (model.User, string, error) {
	validateErr := userLoginValidate(username, password)
	if validateErr != nil {
		return model.User{}, "", validateErr
	}

	formErr := errors.Empty()
	formErr.Code = http.StatusBadRequest
	formErr.Json["form"] = errors.TextError("Unknown combiantion of username and password")
	formErr.Json["reason"] = errors.TextError(errors.ErrFormError)

	user := model.User{}
	userErr := ctx.DB.User().Find(bson.M{db.UserIdKeyUsername: username}).One(&user)
	if userErr == mgo.ErrNotFound {
		formErr.Msg = fmt.Sprintf("user (%s) not found", username, userErr.Error())
		return user, "", formErr
	}
	if userErr != nil {
		return user, "", internalServerErr(
			fmt.Sprintf("user (%s) find error [%s]", username, userErr.Error()),
		)
	}

	if !user.ValidatePassword(password) {
		formErr.Msg = "invalid password"
		return user, "", formErr
	}

	token, tokenErr := generateToken(user.Id)
	if tokenErr != nil {
		return user, "", internalServerErr(
			fmt.Sprintf("generate token error %s", tokenErr.Error()),
		)
	}
	return user, token, nil
}

func userLoginValidate(username, password string) error {
	formErr := errors.New("form error", http.StatusBadRequest)

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
