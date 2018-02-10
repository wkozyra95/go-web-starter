// Package errors error module.
package errors

import (
	"encoding/json"
	"fmt"
)

var (
	// ErrNotFound error not found.
	NotFound = fmt.Errorf("notfound")
	// ErrUnauthorized error unauthorized request.
	Unauthorized = fmt.Errorf("unauthorized")
	// ErrNotLoggedIn error not logged in.
	NotLoggedIn = fmt.Errorf("notloggedin")
	// ErrMalformed error malformed request.
	Malformed = fmt.Errorf("malformed")
	// ErrExpired token expired error.
	Expired = fmt.Errorf("expired")
	// ErrFormError form error.
	InvalidForm = fmt.Errorf("formerror")
	// ErrInternalServerError Internal Server Error.
	InternalServerError = fmt.Errorf("internal")
)

type FormError map[string]string

func NewFormError() FormError {
	return FormError{"reason": InvalidForm.Error()}
}

func (fe FormError) Error() string {
	return fmt.Sprintf("%+v", fe)
}

func (fe FormError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string(fe))
}
