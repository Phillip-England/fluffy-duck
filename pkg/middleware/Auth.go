package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) (string, error) {
	sessionToken, _ := c.Cookie("session")
	if sessionToken == os.Getenv("TEAM_SESSION_TOKEN") {
		return "team", nil
	}
	if sessionToken == os.Getenv("ADMIN_SESSION_TOKEN") {
		return "admin", nil
	}
	return "", fmt.Errorf("session token is not set")
}