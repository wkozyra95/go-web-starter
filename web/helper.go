package web

import (
	"encoding/json"
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"gopkg.in/mgo.v2/bson"
)

func helperExtractUserId(r *http.Request) bson.ObjectId {
	return bson.ObjectIdHex(r.Context().Value(contextUserId).(string))
}

func writeJSONResponse(w http.ResponseWriter, httpStatus int, body interface{}) error {
	marshaled, marshalingErr := json.Marshal(body)
	if marshalingErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return marshalingErr
	}
	w.WriteHeader(httpStatus)
	_, writeErr := w.Write(marshaled)
	return writeErr
}

func decodeJSONRequest(r *http.Request, unpackObject interface{}) error {
	err := json.NewDecoder(r.Body).Decode(unpackObject)
	if err != nil {
		return err
	}
	return nil
}

func requestMalformedErr(msg string) error {
	err := errors.NewMessageError(msg, http.StatusBadRequest)
	err.Json["request"] = errors.TextError("request malformed")
	return err
}

func handleRequestError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	serializableError, isSerializable := err.(errors.MessageError)
	if !isSerializable {
		log.Errorf("[Assert] should not happend %s", err.Error())
		return
	}

	_ = writeJSONResponse(w, serializableError.Code, serializableError)
}
