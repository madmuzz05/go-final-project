package sysresponse

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Error      error
	StatusCode int
	Message    string
}

type IError interface {
	GetMessage() string
	GetStatusCode() int
	GetError() error
}

func GetErrorMessage(err error, statusCode int, message string) IError {
	message = strings.Trim(message, " ")
	sysErr := &Error{
		Message:    message,
		StatusCode: 500,
		Error:      err,
	}

	if err != nil {
		sysErr.Error = err
	}

	if statusCode != 0 {
		sysErr.StatusCode = statusCode
	}

	return sysErr
}

func (s *Error) GetMessage() string {
	return s.Message
}

func (s *Error) GetStatusCode() int {
	if s.StatusCode == 0 {
		return 200
	}
	return s.StatusCode
}

func (s *Error) GetError() error {
	return s.Error
}

type Success struct {
	StatusCode int         `json:"status"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func GetResponseJson(ctx *gin.Context, status int, message string, data interface{}) {
	success := false
	if status < 400 {
		success = true
	}
	response := Success{
		StatusCode: status,
		Message:    message,
		Success:    success,
	}

	if success == false {
		response.Data = handleErrorResponse(data, status)
		ctx.AbortWithStatusJSON(status, response)
		return
	}

	response.Data = data
	ctx.JSON(status, response)
}

// handleErrorResponse handles the error response formatting
func handleErrorResponse(data interface{}, status int) interface{} {
	errorData, ok := data.(IError)
	if !ok {
		return data
	}
	if status >= 500 && errorData.GetError() != nil {
		return formatError(errorData)
	}

	if errorData.GetError() == nil {
		return nil
	}

	return formatError(errorData)
}

// formatError formats the error data into a proper response format
func formatError(errorData IError) interface{} {
	var jsonErr map[string]interface{}
	if json.Unmarshal([]byte(errorData.GetError().Error()), &jsonErr) == nil {
		return map[string]interface{}{
			"error": errorData.GetError(),
		}
	}
	return map[string]string{
		"error": fmt.Sprintf("%s", errorData.GetError()),
	}
}
