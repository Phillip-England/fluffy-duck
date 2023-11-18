package routes

import (
	"fluffy-duck/pkg/routes/apiRoutes"
	"fluffy-duck/pkg/routes/pageRoutes"

	"github.com/gin-gonic/gin"
)

func MountRoutes(r *gin.Engine) {

	// static files
	r.Static("/public", "./public")

	// public pages
	r.GET("/", pageRoutes.LoginPage)

	// team pages
	r.GET("/team", pageRoutes.TeamHome)

	// admin pages
	r.GET("/admin", pageRoutes.AdminHome)
	r.GET("/admin/cem/update", pageRoutes.UpdateCemPage)
	
	// cookie and redirects
	r.GET("/cookie/language/toggle", pageRoutes.LanguageToggle)
	r.GET("/logout", pageRoutes.LogouPage)

	// api
	r.POST("/api/user/login", apiRoutes.LoginUser)
}