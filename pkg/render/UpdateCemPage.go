package render

import (
	"fluffy-duck/pkg/components"
	"fluffy-duck/pkg/layouts"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateCemPage(c *gin.Context, userType string) string {
	fmt.Println(c.Request.URL.Path)
	updateCemForm := components.UpdateCemForm(c)
	page := layouts.AdminLayout(c, "CFA Suite - Update CEM", strings.Join([]string{updateCemForm}, ""))
	return page
}