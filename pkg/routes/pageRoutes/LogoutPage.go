package pageRoutes

import "github.com/gin-gonic/gin"

func LogoutPage(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", true, true) // Delete the session cookie
	c.Redirect(302, "/") // Redirect to the root path
}