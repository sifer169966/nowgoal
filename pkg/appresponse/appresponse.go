package appresponse

import (
	"net/http"
)

var (
	Success             = Status{Code: http.StatusOK, Message: "Success"}
	BadRequest          = Status{Code: http.StatusBadRequest, Message: "Sorry, Not responding because of incorrect syntax"}
	Unauthorized        = Status{Code: http.StatusUnauthorized, Message: "Sorry, We are not able to process your request. Please try again"}
	Forbidden           = Status{Code: http.StatusForbidden, Message: "Sorry, Permission denied"}
	InternalServerError = Status{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
)

// ResponseBody struct
type ResponseBody struct {
	Status Status      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// Status struct
type Status struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
