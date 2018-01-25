package web

import (
	"net/http"

	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/web/handler"
)

func registerHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	var registerRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decodeErr := decodeJSONRequest(r, &registerRequest)
	if decodeErr != nil {
		return requestMalformedErr("register request malformed")
	}

	user := model.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
	}
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: "",
	}

	registerErr := handler.UserRegister(user, registerRequest.Password, context)
	if registerErr != nil {
		log.Warnf("register request failed [%s]", registerErr.Error())
		return registerErr
	}

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
	return nil
}
