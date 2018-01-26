package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrNotFound            = "notfound"
	ErrUnauthorized        = "unauthorized"
	ErrNotLoggedIn         = "notloggedin"
	ErrMalformed           = "malformed"
	ErrExpired             = "expired"
	ErrFormError           = "formerror"
	ErrInternalServerError = "internal"
)

type SerializableError interface {
	json.Marshaler
	Error() string
}

type JsonError map[string]SerializableError

func (je JsonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]SerializableError(je))
}

func (je JsonError) Error() string {
	return fmt.Sprintf("%+v", map[string]SerializableError(je))
}

type TextError string

func (te TextError) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(te))
}

func (te TextError) Error() string {
	return string(te)
}

type MessageError struct {
	Msg  string
	Code int
	Json JsonError
}

func New(message string, code int) MessageError {
	return MessageError{
		Msg:  message,
		Code: code,
		Json: JsonError(map[string]SerializableError{}),
	}
}

func NewSimple(message string, code int, err string) MessageError {
	return MessageError{
		Msg:  message,
		Code: code,
		Json: JsonError(map[string]SerializableError{
			"reason": TextError(err),
		}),
	}
}

func Empty() MessageError {
	return MessageError{
		Code: http.StatusBadRequest,
		Json: JsonError(map[string]SerializableError{}),
	}
}

func (me MessageError) MarshalJSON() ([]byte, error) {
	return me.Json.MarshalJSON()
}

func (me MessageError) Error() string {
	return fmt.Sprintf("%s %v", me.Msg, me.Json)
}
