package common

import (
	"fmt"
	"net/http"
)

/*
not found err
duplicate err
alredy exist err
forbidden err
InvalidArgement err
*/

// BaseErr basic error class
type BaseErr struct {
	StatusCode int
	Message    string
}

func (e *BaseErr) Error() string {
	return fmt.Sprintf("%v: %v", e.StatusCode, e.Message)
}

func NotFoundError(fmtAndArgs ...interface{}) error {
	return &BaseErr{http.StatusNotFound, Format(fmtAndArgs...)}
}

func IsNotFoundError(err error) bool {
	e, ok := err.(*BaseErr)
	return ok && e.StatusCode == http.StatusNotFound
}

func InvalidArgumentErr(fmtAndArgs ...interface{}) error {
	return &BaseErr{http.StatusBadRequest, Format(fmtAndArgs...)}
}

func IsInvalidArgumentError(err error) bool {
	e, ok := err.(*BaseErr)
	return ok && e.StatusCode == http.StatusBadRequest
}

func ForbiddenError(fmtAndArgs ...interface{}) error {
	return &BaseErr{http.StatusForbidden, Format(fmtAndArgs...)}
}

func IsForbiddenError(err error) bool {
	e, ok := err.(*BaseErr)
	return ok && e.StatusCode == http.StatusForbidden
}
