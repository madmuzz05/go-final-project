package validation

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
)

func ValidateRequest(c *gin.Context, req interface{}) (err sysresponse.IError) {
	// Validate the struct using the validator
	if err := c.ShouldBindJSON(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {

			if len(validationErrors) > 0 {
				return sysresponse.GetErrorMessage(err, http.StatusBadRequest, fmt.Sprintf("Field '%s' failed validation for tag '%s'", validationErrors[0].Field(), validationErrors[0].Tag()))
			}
		}
	}
	return nil
}
