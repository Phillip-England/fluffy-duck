import { getById } from "../utility/getById"
import { toggleClassesOnElements } from "../utility/toggleClassesOnElements"


export const eToggleNav = () => {
	console.log('hooking event: eToggleNav')
	const bannerBars = getById('banner-bars')
	const bannerX = getById('banner-x')
	const nav = getById('nav-menu')
	const overlay = getById('overlay')
	bannerBars.addEventListener('click', () => {
		toggleClassesOnElements([bannerBars, bannerX, nav, overlay], ['hidden'])
	})
	bannerX.addEventListener('click', () => {
		toggleClassesOnElements([bannerBars, bannerX, nav, overlay], ['hidden'])
	})
	overlay.addEventListener('click', () => {
		toggleClassesOnElements([bannerBars, bannerX, nav, overlay], ['hidden'])
	})
	
}