package pageRoutes

import (
	"fluffy-duck/pkg/middleware"
	"fluffy-duck/pkg/render"

	"github.com/gin-gonic/gin"
)


func AdminHome(c *gin.Context) {
	userType, err := middleware.Auth(c)
	if err != nil || userType != "admin" {
		c.Redirect(302, "/")
		return
	}
	c.Data(200, "text/html", []byte(render.AdminHome(c, userType)))
}