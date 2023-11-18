package pageRoutes

import (
	"fluffy-duck/pkg/utility"

	"github.com/gin-gonic/gin"
)

func LanguageToggle(c *gin.Context) {
	path := c.Query("path")
	language := utility.GetLanguage(c)
	if language == "" {
		c.SetCookie("language", "spanish", 3600, "/", "", false, true)
	} else if language == "english" {
		c.SetCookie("language", "spanish", 3600, "/", "", false, true)
	} else {
		c.SetCookie("language", "english", 3600, "/", "", false, true)
	}
	c.Redirect(301, path)
}