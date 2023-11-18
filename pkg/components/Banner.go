package components

import (
	"fmt"
)

func Banner(title string, subtext string, isTogglable bool) string {
	spacer := "<div class='h-28'></div>"
	var icons string
	if isTogglable {
		title = ""
		icons = `
			<div id='banner-toggle-icons' class='flex flex-row m-2 text-lightGray'>
				<div id='banner-bars'>
					<i class='fa-solid fa-bars fa-lg'></i>
				</div>
				<div id='banner-x' class='hidden'>
					<i class='fa-solid fa-x fa-lg'></i>
				</div>
			</div>
		`
	} else {
		icons = ""
	}
	
	return fmt.Sprintf(`
		<div class='flex top-0 z-40 flex-row fixed w-full items-center gap-4 p-4 h-28 justify-between bg-navy'>
			<img class='h-12 w-12 m-2' src="/public/logo-dark.svg">
			<h1 class='text-lightGray m-2 text-2xl font-semibold'>%s</h1>
			%s
		</div>
		%s
	`, title, icons, spacer)
}

