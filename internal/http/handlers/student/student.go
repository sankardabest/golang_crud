package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sankardabest/students-api/internal/types"
	"github.com/sankardabest/students-api/internal/utils/response"
)

func GetStudents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student API"))
	}
}

var validate = validator.New()

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a Student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.ErrorResponse(fmt.Errorf("Empty Body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ErrorResponse(err))
			return
		}

		/////request validation///
		if err := validate.Struct(student); err != nil {
			// validateErrors := err.(validator.ValidationErrors)
			var validationErrors validator.ValidationErrors

			if errors.As(err, &validationErrors) {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validationErrors))
			}
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]string{"status": "success"})
	}
}
