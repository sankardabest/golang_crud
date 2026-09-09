package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"ststus"`
	Error  string `json:"error"`
}

const (
	statusOk    = "Ok"
	statusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func ErrorResponse(err error) Response {
	return Response{
		Status: statusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {

	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("%s is required", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return Response{
		Status: statusError,
		Error:  strings.Join(errMsg, ", "),
	}
}
