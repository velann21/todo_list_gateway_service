package helpers

import "errors"

var (
	// ErrParamMissing required field missing error
	InvalidRequest = errors.New("Invalid Request")
	SomethingWrong = errors.New("Something went wrong")
	UnAuthorized = errors.New("UnAuthorized")
)
