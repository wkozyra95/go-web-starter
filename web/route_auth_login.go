package web

import (
	"fmt"
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/mongo"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	db := extractDBSession(r.Context())
	var loginRequest loginForm
	decodeErr := decodeJSONRequest(r, &loginRequest)
	if decodeErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	if err := loginRequest.validate(); err != nil {
		handleRequestErr(w, err)
		return
	}

	user := mongo.User{}
	userErr := db.User().Find(
		bson.M{mongo.UserIDKeyUsername: loginRequest.Username},
	).One(&user)
	if userErr == mgo.ErrNotFound {
		handleRequestErr(w, errors.NotFound)
		return
	}
	if userErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	if err := loginRequest.validatePassword(user.PasswordHash); err != nil {
		handleRequestErr(w, err)
		return
	}

	token, tokenErr := h.jwt.generate(user.ID)
	if tokenErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	_ = writeJSONResponse(w, http.StatusOK, struct {
		Token string     `json:"token"`
		User  model.User `json:"user"`
	}{
		Token: token,
		User:  user.User,
	})
}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lf loginForm) validate() error {
	if lf.Username == "" {
		return fmt.Errorf("Username can't be empty")
	}
	if lf.Password == "" {
		return fmt.Errorf("Password can't be empty")
	}
	return nil
}

func (lf loginForm) validatePassword(hashedPassword string) error {
	log.Errorf(lf.Password)
	log.Errorf(hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(lf.Password))
	if err != nil {
		return fmt.Errorf("Unknown combination of user and password.")
	}
	return nil
}
