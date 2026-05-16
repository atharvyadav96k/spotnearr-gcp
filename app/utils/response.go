package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}, err interface{},
) {

	response := APIResponse{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	WriteResponse(
		w,
		statusCode,
		true,
		message,
		data,
		nil,
	)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err interface{}) {
	WriteResponse(
		w,
		statusCode,
		false,
		message,
		nil,
		err,
	)
}

func BadRequest(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusBadRequest,
		"Bad Request",
		err,
	)
}

func Unauthorized(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusUnauthorized,
		"Unauthorized",
		err,
	)
}

func Forbidden(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusForbidden,
		"Forbidden",
		err,
	)
}

func NotFound(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusNotFound,
		"Not Found",
		err,
	)
}

func Conflict(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusConflict,
		"Conflict",
		err,
	)
}

func ValidationError(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusUnprocessableEntity,
		"Validation Error",
		err,
	)
}

func InternalServerError(w http.ResponseWriter, err interface{}) {
	ErrorResponse(
		w,
		http.StatusInternalServerError,
		"Internal Server Error",
		err,
	)
}

func OK(w http.ResponseWriter, message string, data interface{},
) {
	SuccessResponse(
		w,
		http.StatusOK,
		message,
		data,
	)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	SuccessResponse(
		w,
		http.StatusCreated,
		message,
		data,
	)
}
