package render

import (
	"fluffy-duck/pkg/layouts"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminHome(c *gin.Context, userType string) string {
	fmt.Println(c.Request.URL.Path)
	page := layouts.AdminLayout(c, "CFA Suite - Admin Home", "")
	return page
}