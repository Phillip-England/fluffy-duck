package pageRoutes

import (
	"fluffy-duck/pkg/middleware"
	"fluffy-duck/pkg/render"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	userType, _ := middleware.Auth(c)
	if userType == "team" {
		c.Redirect(301, "/team")
		return
	}
	if userType == "admin" {
		c.Redirect(301, "/admin")
		return
	}
	c.Data(200, "text/html", []byte(render.LoginPage(c)))
}