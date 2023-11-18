package components

import (
	"fluffy-duck/pkg/utility"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Footer(c *gin.Context) string {
	path := "/cookie/language/toggle?path=" + c.Request.URL.Path
	languageSelectionOption := utility.IfSpanishThen(c, "Espanol", "English")
	return fmt.Sprintf(`
		<div class='h-20 invisible w-full'></div>
		<footer class='z-30 h-20 items-center bg-navy flex flex-row justify-end w-screen pr-4 fixed bottom-0'>
			<a href="%s" class='text-xs px-4 py-1 rounded-full font-bold text-navy bg-green focus:outline-none'>%s</a>
		</footer>
	`, path, languageSelectionOption)
}