// Package handlers in this section takes care of packaging errors.
// So far we handle the following errors: application level errors returned by the api,
// authorization errors returned by the api, http errors outside of the 200-300 range
// returned by the api.
package handlers

import (
	"errors"
	"fmt"
	"net/http"
)

// SourceError is the type of error returned by the api when there's a syntax or semantic error
type SourceError struct {
	Pointer string				`json:"pointer"`
}

type ValidationError struct {
	Detail string       `json:"detail"`
	Source *SourceError `json:"source"`
}

type ResponseErrors struct {
	Errors []*ValidationError `json:"errors"`
}

func MakeValidationError(detail, pointer string) *ValidationError {
	se := SourceError{
		Pointer: pointer,
	}

	ve := ValidationError{
		Detail: detail,
		Source: &se,
	}

	return &ve
}

type AuthenticationError struct {
	Code int
}

func MakeAuthenticationError(code int) *AuthenticationError {
	ae := &AuthenticationError{
		Code: code,
	}

	return ae
}

func (ae AuthenticationError) Error() string {
	return fmt.Sprintf("autentication error %d", ae.Code)
}

func (ae AuthenticationError) Is(target error) bool {
	ae, ok := target.(AuthenticationError)

	return ok
}

type NoSessionError struct {
	Code            int
	CodeDescription string
	PerformingWhat  string
	OriginalError   error
}

type GeneralError struct {
	Code            int
	CodeDescription string
	PerformingWhat  string
	OriginalError   error
}

func MakeGeneralError(code int, doingWhat string, err error) *GeneralError {
	codeDescription := ""
	if code > 0 {
		codeDescription = http.StatusText(code)
	}
	ge := &GeneralError{
		Code:            code,
		CodeDescription: codeDescription,
		PerformingWhat:  doingWhat,
		OriginalError:   err,
	}

	return ge
}

func (ge GeneralError) Error() string {
	return fmt.Sprintf("general error %d performing %s", ge.Code, ge.PerformingWhat)
}

func (ge GeneralError) Is(target error) bool {
	ge, ok := target.(GeneralError)

	return ok
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) bool {
	if errors.Is(err, AuthenticationError{}) {
		fmt.Printf("error %s", err.Error())
		RedirectToSessionExpired(w, r)
		return true
	}
	//ge, ok := err.(*handlers.GeneralError)
	if errors.Is(err, GeneralError{}) {
		fmt.Printf("error %s", err.Error())
		RedirectToErrorPage(w, r)
		return true
	}

	return false
}
