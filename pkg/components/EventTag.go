package components

import "fmt"

func EventTag(events string) string {
	return fmt.Sprintf(`
		<div id='event-tag' events="%s"></div>
	`, events)
}