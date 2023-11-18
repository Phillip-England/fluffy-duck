package render

import (
	"fluffy-duck/pkg/components"
	"fluffy-duck/pkg/utility"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) string {
	fmt.Println(c.Request.URL.Path)
	banner := components.Banner("CFA Suite", utility.IfSpanishThen(c, "See under the hood of our business", "Ver bajo el capó de nuestro negocio"), false) 
	loginForm := components.LoginForm(c)
	footer := components.Footer(c)
	page := components.Root("CFA Suite - Login", fmt.Sprintf(`%s%s%s`, banner, loginForm, footer))
	return page
}