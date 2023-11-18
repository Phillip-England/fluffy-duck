package components

import "github.com/gin-gonic/gin"

func AdminNav(c *gin.Context) string {
	adminHomeLink := NavLink(c, "/admin", "Admin Home")
	updateCemLink := NavLink(c, "/admin/cem/update", "Update Cem")
	logoutLink := NavLink(c, "/logout", "Logout")
	stringArray := []string{adminHomeLink, updateCemLink, logoutLink}
	return NavMenu(stringArray)
}