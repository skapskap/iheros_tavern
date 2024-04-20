package utility

import "github.com/labstack/echo/v4"

type ApiResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func JSONResponse(c echo.Context, statusCode int, status, message string, data interface{}) error {
	response := ApiResponse{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	return c.JSON(statusCode, response)
}