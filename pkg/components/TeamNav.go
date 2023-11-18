package components

import "github.com/gin-gonic/gin"

func TeamNav(c *gin.Context) string {
	teamHomeLink := NavLink(c, "/team", "Home")
	logoutLink := NavLink(c, "/logout", "Logout")
	stringArray := []string{teamHomeLink, logoutLink}
	return NavMenu(stringArray)
}