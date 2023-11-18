import { getByClass } from "../utility/getByClass"


export const eHookUpdateCemForm = () => {
	console.log('hooking event: eUpdateCemForm')
	const timescaleButtons = getByClass("update-cem-timescale-toggle")
	const scoreRows = getByClass('cem-update-score-row')
	let selectedTimescale = 'current month'

	// handling all timescale toggle functionality
	for (let i = 0; i < timescaleButtons.length; i++) {
		 const button = timescaleButtons[i] as HTMLInputElement
		button.addEventListener('click', (e) => {
			selectedTimescale = button.value
			console.log(selectedTimescale)
			button.classList.remove('bg-navy')
			button.classList.add('bg-sky')
			button.classList.remove('text-sky')
			button.classList.add('text-navy')
			for (let i = 0; i < timescaleButtons.length; i++) {
				let button = timescaleButtons[i] as Element
				if (button != e.target) {
					button.classList.remove('bg-sky')
					button.classList.remove('text-navy')
					button.classList.add('bg-navy')
					button.classList.add('text-sky')
				}
			}
		})
	}

	// handling all score rows
	for (let i = 0; i < scoreRows.length; i++) {
		const row = scoreRows[i] as Element
		const downButton = row.querySelector('.cem-score-update-down-arrow') as Element
		const upButton = row.querySelector('.cem-score-update-up-arrow') as Element
		const input = row.querySelector('.cem-score-update-input') as HTMLInputElement
		downButton.addEventListener('click', () => {
			const currentScore = Number(input.value) as number
			input.value = String(currentScore - 1)
		})
		upButton.addEventListener('click', () => {
			const currentScore = Number(input.value) as number
			input.value = String(currentScore + 1)
		})
	}


}