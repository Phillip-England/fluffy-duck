package pageRoutes

import "github.com/gin-gonic/gin"

func LogoutPage(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", false, true) // Delete the session cookie
	c.Redirect(301, "/") // Redirect to the root path
}