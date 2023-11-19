package pageRoutes

import (
	"fluffy-duck/pkg/render"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", true, true) // Delete the session cookie
	c.Data(200, "text/html", []byte(render.LoginPage(c)))
}