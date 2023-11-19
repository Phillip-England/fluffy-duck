package render

import (
	"fluffy-duck/pkg/components"
	"fluffy-duck/pkg/layouts"
	"fluffy-duck/pkg/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateCemPage(c *gin.Context, scores *models.CemScore) string {
	fmt.Println(c.Request.URL.Path)
	updateCemForm := components.UpdateCemForm(c, scores)
	page := layouts.AdminLayout(c, "CFA Suite - Update CEM", strings.Join([]string{updateCemForm}, ""))
	return page
}