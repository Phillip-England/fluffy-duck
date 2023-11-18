package components

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NavLink(c *gin.Context, href string, text string) string {
	activeClass := ""
	if c.Request.URL.Path == href {
		activeClass = "underline text-cyan"
	}
	return fmt.Sprintf(`
		<li class='bg-navyLight text-sm flex rounded-full border border-lightGray rounded-full %s'>
			<a class='w-full p-6 bg-navyLight rounded-full' href="%s">%s</a>
		</li>
	`, activeClass, href, text)
}