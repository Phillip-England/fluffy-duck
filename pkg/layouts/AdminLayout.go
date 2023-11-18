package layouts

import (
	"fluffy-duck/pkg/components"
	"fluffy-duck/pkg/utility"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminLayout(c *gin.Context, title string, children string) string {
	banner := components.Banner("CFA Suite", utility.IfSpanishThen(c, "Having a positive influence", "Tener una influencia positiva"), true) 
	footer := components.Footer(c)
	mainLoader := components.MainLoader()
	nav := components.AdminNav(c)
	overlay := components.Overlay()
	layout := components.Root(title, fmt.Sprintf(`
		%s
		%s
		<main class='px-2'>%s</main>
		%s
		%s
		%s`, banner, nav, children, footer, mainLoader, overlay))
	return layout
}