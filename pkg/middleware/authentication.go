package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/madmuzz05/go-final-project/pkg/helper/jwt"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := auth.VerifyToken(ctx)
		if err != nil {
			sysresponse.GetResponseJson(ctx, http.StatusUnauthorized, err.Error(), err)
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
