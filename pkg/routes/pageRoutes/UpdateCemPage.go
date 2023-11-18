package pageRoutes

import (
	"fluffy-duck/pkg/middleware"
	"fluffy-duck/pkg/render"

	"github.com/gin-gonic/gin"
)

func UpdateCemPage(c *gin.Context) {
	userType, _ := middleware.Auth(c)
	if userType == "team" {
		c.Redirect(301, "/team")
		return
	}
	c.Data(200, "text/html", []byte(render.UpdateCemPage(c, userType)))
}