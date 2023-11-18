package utility

import "github.com/gin-gonic/gin"

func GetLanguage(c *gin.Context) string {
	language, _ := c.Cookie("language")
	return language
}