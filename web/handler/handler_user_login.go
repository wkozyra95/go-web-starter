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

// UserLogin ...
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
	formErr.JSON["form"] = errors.TextError("Unknown combiantion of username and password")
	formErr.JSON["reason"] = errors.TextError(errors.ErrFormError)

	user := model.User{}
	userErr := ctx.DB.User().Find(bson.M{db.UserIDKeyUsername: username}).One(&user)
	if userErr == mgo.ErrNotFound {
		formErr.Msg = fmt.Sprintf("user (%s) not found", username)
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

	token, tokenErr := generateToken(user.ID)
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
		formErr.JSON["username"] = errors.TextError("Username can't empty")
	}
	if password == "" {
		formErr.JSON["password"] = errors.TextError("Password can't empty")
	}
	if len(formErr.JSON) > 0 {
		return formErr
	}
	return nil
}
