package web

import (
	"fmt"
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func (h *handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	db := extractDBSession(r.Context())
	var registerRequest registerForm
	decodeErr := decodeJSONRequest(r, &registerRequest)
	if decodeErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	if err := registerRequest.validate(); err != nil {
		handleRequestErr(w, err)
		return
	}

	count, countErr := db.User().Find(
		bson.M{mongo.UserIDKeyUsername: registerRequest.Username},
	).Count()
	if countErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}
	if count > 0 {
		handleRequestErr(w, fmt.Errorf(
			"User with that login already exists",
		))
		return
	}

	user := registerRequest.createUser()
	insertErr := db.User().Insert(user)
	if insertErr != nil {
		log.Error(insertErr.Error())
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	_ = writeJSONResponse(w, http.StatusOK, "ok")
}

type registerForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rf registerForm) validate() error {
	if rf.Email == "" {
		return fmt.Errorf("Email can't be empty")
	}
	if rf.Username == "" {
		return fmt.Errorf("Username can't be empty")
	}
	if rf.Password == "" {
		return fmt.Errorf("Password can't be empty")
	}
	if len(rf.Password) < 8 {
		return fmt.Errorf("Password is to short, you need at least 8 characters")
	}
	return nil
}

func (rf registerForm) createUser() mongo.User {
	return mongo.User{
		ID: bson.NewObjectId(),
		User: model.User{
			Username:     rf.Username,
			Email:        rf.Email,
			PasswordHash: rf.generatePasswordHash(),
		},
	}
}

func (rf registerForm) generatePasswordHash() string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rf.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("[ASSERT] Failed to hash password")
		panic(err)
	}
	return string(hashedPassword)
}
