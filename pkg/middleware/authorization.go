package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
)

type AuthorizationMiddleware struct {
	gormDb *postgres.GormDB
}

func InitAuthorizationMiddleware(gormDb *postgres.GormDB) IAuthorizationMiddleware {
	return &AuthorizationMiddleware{
		gormDb: gormDb,
	}
}

type IAuthorizationMiddleware interface {
	Authorization(parameterName string, tableName string) gin.HandlerFunc
}
type user struct {
	UserId uint `json:"user_id"`
}

func (m AuthorizationMiddleware) Authorization(parameterName string, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := m.gormDb.GetDB().WithContext(c)
		id, err := strconv.Atoi(c.Param(parameterName))
		if err != nil {
			sysresponse.GetResponseJson(c, http.StatusBadRequest, "Invalid parameter", gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			c.Abort()
			return
		}

		rawDataUser := c.MustGet("userData")

		// var result map[string]interface{}
		bytes, err := json.Marshal(rawDataUser)
		if err != nil {
			sysresponse.GetResponseJson(c, http.StatusInternalServerError, "Internal Server Error.", err)
			c.Abort()
			return
		}
		var userData map[string]interface{}
		// Unmarshal the JSON into the map
		if err := json.Unmarshal([]byte(bytes), &userData); err != nil {
			sysresponse.GetResponseJson(c, http.StatusInternalServerError, "Internal Server Error.", err)
			c.Abort()
			return
		}

		// Accessing values
		userId := uint(userData["id"].(float64))

		var rawUser user
		sql := "SELECT * from " + tableName + " where id = ?"
		model := db.Raw(sql, id).Scan(&rawUser)

		if model.Error != nil {
			sysresponse.GetResponseJson(c, http.StatusInternalServerError, "Internal Server Error.", model.Error)
			c.Abort()
			return
		} else if model.RowsAffected == 0 {
			sysresponse.GetResponseJson(c, http.StatusNotFound, "Data not found.", nil)
			c.Abort()
			return
		}

		if rawUser.UserId != userId {
			sysresponse.GetResponseJson(c, http.StatusUnauthorized, "You are not allowed to access this data.", gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
