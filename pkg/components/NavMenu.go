package components

import (
	"fmt"
	"strings"
)

func NavMenu(links []string) string {
	linksHtml := strings.Join(links, "")
	return fmt.Sprintf(`
		<nav id='nav-menu' class='hidden z-40 fixed left-0 w-3/5 h-full bg-navy'>
			<ul class='flex flex-col'>
				%s
			</ul>
		</nav>
	`, linksHtml)
	
}