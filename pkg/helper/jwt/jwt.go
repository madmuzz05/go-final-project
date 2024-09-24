package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/config"
)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var secretKey = config.AppConfig.JwtSecret

func GenerateToken(id uint, username string) (string, string) {
	expirationTime := time.Now().Add(6 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time.
	claims := &Claims{
		Id:       int(id),
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := parseToken.SignedString([]byte(secretKey))

	return tokenString, string(expirationTime.Format("2006-01-02 15:04:05"))
}
func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(headerToken, "Bearer ") { // Bearer tokennyaDisini
		return nil, fmt.Errorf("Missing token")
	}

	// Bearer tokennyaDisini
	stringToken := strings.TrimPrefix(headerToken, "Bearer ")
	// tokennyaDisini
	token, errToken := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if errToken != nil || !token.Valid {
		if ve, ok := errToken.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("Token has expired")
			} else {
				return nil, fmt.Errorf("Invalid token")
			}
		} else {
			return nil, fmt.Errorf("Invalid token")
		}
	}
	fmt.Println(token.Claims.(jwt.MapClaims))

	return token.Claims.(jwt.MapClaims), nil
}
