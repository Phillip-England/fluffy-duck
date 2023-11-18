/// <reference lib="dom" />
/// <reference lib="dom.iterable" />

import { eHookLoginForm } from "./events/eHookLoginForm"
import { eHookUpdateCemForm } from "./events/eHookUpdateCemForm"
import { eToggleNav } from "./events/eToggleNav"

const clientRouter = new Map<string, () => void>()

clientRouter.set('/', () => {
	console.log("/")
	eHookLoginForm()
})

clientRouter.set('/team', () => {
	console.log("/team")
	eToggleNav()
})

clientRouter.set('/admin', () => {
	console.log('/admin')
	eToggleNav()
})


clientRouter.set('/admin/cem/update', () => {
	console.log('/admin/cem/update')
	eToggleNav()
	eHookUpdateCemForm()
})


const clientHydrationFunc = clientRouter.get(window.location.pathname)
if (!clientHydrationFunc) {
	console.error(`no client hydration function set for route: ${window.location.pathname}`)
} else {
	clientHydrationFunc()
}

// events listed here are mounted on all routes
