package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func NewMessageError(message string, code int) MessageError {
	return MessageError{
		Msg:  message,
		Code: code,
		Json: JsonError(map[string]SerializableError{}),
	}
}

func EmptyMessageError() MessageError {
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
