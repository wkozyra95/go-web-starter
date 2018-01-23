package web

import (
	"encoding/json"
	"net/http"

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
