import { getByClass } from "../utility/getByClass"
import { getById } from "../utility/getById"


export const eHookUpdateCemForm = async () => {
	console.log('hooking event: eUpdateCemForm')
	const cemUpdateSubmit = getById('update-cem-form-submit')
	const scoreRows = getByClass('cem-update-score-row')
	const toggleButtons = getByClass('cem-score-update-toggle-button')
	console.log(toggleButtons)
	// makes toggle button change orange when clicked
	for (let i = 0; i < toggleButtons.length; i++) {
		const currentButton = toggleButtons[i] as HTMLElement
		currentButton.addEventListener('click', (e) => {
			currentButton.classList.add('bg-green')
			currentButton.classList.add('text-navy')
			currentButton.classList.add('border-green')
			for (let i2 = 0; i2 < toggleButtons.length; i2++) {
				const anotherCurrentButton = toggleButtons[i2] as HTMLElement
				if (anotherCurrentButton != currentButton) {
					anotherCurrentButton.classList.remove('bg-green')
					anotherCurrentButton.classList.remove('text-navy')
					anotherCurrentButton.classList.remove('border-green')
				}
			}
		})
	}
	// enable score rows to toggle
	for (let i = 0; i < scoreRows.length; i++) {
		const row = scoreRows[i] as Element
		const downButton = row.querySelector('.cem-score-update-down-arrow') as Element
		const upButton = row.querySelector('.cem-score-update-up-arrow') as Element
		const input = row.querySelector('.cem-score-update-input') as HTMLInputElement
		downButton.addEventListener('click', () => {
			cemUpdateSubmit.scrollIntoView({behavior: "smooth"})
			const currentScore = Number(input.value) as number
			if (currentScore - 1 > -1) {
				input.value = String(currentScore - 1)
			} 
		})
		upButton.addEventListener('click', () => {
			cemUpdateSubmit.scrollIntoView({behavior: "auto"})
			const currentScore = Number(input.value) as number
			if (currentScore + 1 < 101) {
				input.value = String(currentScore + 1)
			}
		})
	}
}