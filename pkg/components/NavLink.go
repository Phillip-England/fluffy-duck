package components

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NavLink(c *gin.Context, href string, text string) string {
	activeClass := ""
	if c.Request.URL.Path == href {
		activeClass = "text-sky border border-sky"
	}
	return fmt.Sprintf(`
		<li class='text-sm flex bg-navy underline hover:bg-navyLight %s'>
			<a class='w-full p-6' href="%s">%s</a>
		</li>
	`, activeClass, href, text)
}