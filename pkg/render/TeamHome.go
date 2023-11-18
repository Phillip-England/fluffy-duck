package render

import (
	"fluffy-duck/pkg/layouts"
	"fmt"

	"github.com/gin-gonic/gin"
)

func TeamHome(c *gin.Context, userType string) string {
	fmt.Println(c.Request.URL.Path)
	page := layouts.TeamLayout(c, "CFA Suite - Team Home", "")
	return page
}