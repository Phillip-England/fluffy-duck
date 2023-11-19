package routes

import (
	"fluffy-duck/pkg/routes/actionRoutes"
	"fluffy-duck/pkg/routes/apiRoutes"
	"fluffy-duck/pkg/routes/pageRoutes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func MountRoutes(r *gin.Engine, pool *pgxpool.Pool) {

	// static files
	r.Static("/public", "./public")

	// public pages
	r.GET("/", func(c *gin.Context) {
		pageRoutes.LoginPage(c)
	})

	// team pages
	r.GET("/team", func(c *gin.Context) {
		pageRoutes.TeamHome(c)
	})

	// admin pages
	r.GET("/admin", func(c *gin.Context) {
		pageRoutes.AdminHome(c)
	})
	r.GET("/admin/cem/update", func(c *gin.Context) {
		pageRoutes.UpdateCemPage(c, pool)
	})

	
	// cookie and redirects
	r.GET("/cookie/language/toggle", func(c *gin.Context) {
		pageRoutes.LanguageToggle(c)
	})
	r.GET("/logout", func(c *gin.Context) {
		pageRoutes.LogoutPage(c)
	})

	// api
	r.POST("/api/user/login", func(c *gin.Context) {
		apiRoutes.LoginUser(c)
	})
	r.POST("/api/cem/update", func(c *gin.Context) {
		apiRoutes.UpdateCem(c, pool)
	})
	r.GET("/api/cem", func(c *gin.Context) {
		apiRoutes.GetCems(c, pool)
	})

	// form actions
	r.POST("/action/cem/update", func(c *gin.Context) {
		actionRoutes.UpdateCem(c, pool)
	})

}