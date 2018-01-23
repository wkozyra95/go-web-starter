package web

import (
	"net/http"

	"github.com/wkozyra95/go-web-starter/web/handler"
)

func loginHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) {
	var loginRequest struct {
		Username string `json="username"`
		Password string `json="password"`
	}
	decodeErr := decodeJSONRequest(r, &loginRequest)
	if decodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: "",
	}

	token, loginErr := handler.UserLogin(loginRequest.Username, loginRequest.Password, context)
	if loginErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	_ = writeJSONResponse(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})

}
