package errs

import (
	"errors"
	"fmt"
	"net/http"
)

// Error constants for 400 bad request
const (
	ErrParameterRequired  = "E400001"
	ErrRequestBodyInvalid = "E400004"
)

// Error constants for 404 bad request
const (
	ErrResourceNotFound  = "E404001"
	ErrCodeNotFound      = "E404002"
	ErrInvalidCreds      = "E404003"
	ErrUserExist         = "E404004"
	ErrUserNotExist      = "E404005"
	ErrWorkspaceExist    = "E404006"
	ErrworkspaceNotExist = "E404007"
	ErrKeyExist          = "E404008"
	ErrKeyNotExist       = "E404009"
	ErrPropertyExist     = "E404010"
	ErrPropertyNotExist  = "E404011"
)

// Error constants for 500 bad request
const (
	ErrInternalAppError = "E500002"
	ErrInternalDBError  = "E500003"
)

// Error constants for 504 Gateway timeout
const (
	ErrGatewayTimeout = "E504001"
)

// Error constants for 422 Unprocessable Entity
const (
	ErrEmptyBodyContent = "E422001"
)

// Errors - maps of error with the error code
type Errors map[int]map[string]string

// Error - return error response
type Error struct {
	Error    string `json:"error"`
	HTTPCode int    `json:"http_code,string"`
}

// ErrorResponse - Use to trow the errors to users
type ErrorResponse struct {
	Error string `json:"error"`
}

var errs *Errors
var allErrs map[string]Error

// Init function for errs package
func init() {
	errs = &Errors{
		http.StatusBadRequest: {
			ErrParameterRequired: "Parameter `%s` is a required field",
			ErrInvalidCreds:      "Invalid login credentials. Please try again",
			ErrUserExist:         "User already exist",
			ErrUserNotExist:      "User doesn't exist",
			ErrWorkspaceExist:    "workspace already exist",
			ErrworkspaceNotExist: "workspace doesn't exist",
			ErrKeyExist:          "apikey for the client already exist",
			ErrKeyNotExist:       "apikey doesn't  exist",
			ErrPropertyExist:     "property name already exist",
			ErrPropertyNotExist:  "property name doesn't  exist",
		},
		http.StatusNotFound: {
			ErrResourceNotFound: "Resource Not found",
			ErrCodeNotFound:     "Error code not found",
		},
		http.StatusInternalServerError: {
			ErrInternalAppError: "Internal Application Error, `%s`",
			ErrInternalDBError:  "Database Error, `%s`",
		},
		http.StatusGatewayTimeout: {
			ErrGatewayTimeout: "Gateway Timeout",
		},
		http.StatusUnprocessableEntity: {
			ErrEmptyBodyContent: "Cannot parse empty body",
		},
	}
	allErrs = make(map[string]Error)

	for httpcode, err := range *errs {
		for code, msg := range err {
			tmp := &Error{HTTPCode: httpcode, Error: msg}
			allErrs[code] = *tmp
		}
	}
}

// GetErrorByCode ...
func GetErrorByCode(code string) (res Error, err error) {
	var ok bool
	if res, ok = allErrs[code]; !ok {
		err = errors.New(ErrCodeNotFound)
		return
	}
	return
}

// GetErrors ...
func GetErrors() (res map[string]Error) {
	res = allErrs
	return
}

// FormateErrorResponse ...
func FormateErrorResponse(mErr Error, val ...interface{}) (res ErrorResponse) {
	if len(val) > 0 {
		mErr.Error = fmt.Sprintf(mErr.Error, val...)
	}

	errRes := &ErrorResponse{
		Error: mErr.Error,
	}
	return *errRes
}
