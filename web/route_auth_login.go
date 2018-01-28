package web

import (
	"net/http"

	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/web/handler"
)

func loginHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	decodeErr := decodeJSONRequest(r, &loginRequest)
	if decodeErr != nil {
		return requestMalformedErr("request login malformed")
	}
	context := handler.ActionContext{
		DB:     ctx.db,
		UserID: "",
	}

	user, token, loginErr := handler.UserLogin(
		loginRequest.Username,
		loginRequest.Password,
		ctx.server.jwt.generate,
		context,
	)
	if loginErr != nil {
		log.Warnf("request login error [%s]", loginErr.Error())
		return loginErr
	}

	_ = writeJSONResponse(w, http.StatusOK, struct {
		Token string     `json:"token"`
		User  model.User `json:"user"`
	}{
		Token: token,
		User:  user,
	})
	return nil
}
