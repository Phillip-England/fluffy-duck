package components

import "fmt"

func FormLoader(id string) string {
	return fmt.Sprintf(`
		<div id="%s" class='self-center hidden  border-t-2 border-navyLight border-t-cyan w-10 h-10 rounded-full animate-spin'></div>
	`, id)
}