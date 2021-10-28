package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// response
// meta payload

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
type M map[string]interface{}

func ResponseFormatter(code int, status string, message interface{}, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}
func ResponseErrorFormatter(c echo.Context, err error) error {
	response := ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
	return c.JSON(http.StatusBadRequest, response)

}
func ErrorFormatter(err error) []string {
	var errors []string

	errors = append(errors, err.Error())

	return errors
}
