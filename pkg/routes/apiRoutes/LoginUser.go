package apiRoutes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}
	username := jsonData["username"].(string)
	password := jsonData["password"].(string)
	teamUsername := os.Getenv("TEAM_USERNAME")
	teamPassword := os.Getenv("TEAM_PASSWORD")
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if username == teamUsername && password == teamPassword {
		teamSessionToken := os.Getenv("TEAM_SESSION_TOKEN")
		c.SetCookie("session", teamSessionToken, 3600, "/", "", true, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"redirect": "/team",
		})
		return
	}
	if username == adminUsername && password == adminPassword {
		adminSessionToken := os.Getenv("ADMIN_SESSION_TOKEN")
		c.SetCookie("session", adminSessionToken, 3600, "/", "", true, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"redirect": "/admin",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "invalid credentials",
	})
} 