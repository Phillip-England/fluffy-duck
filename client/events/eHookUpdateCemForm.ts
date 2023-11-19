import { getByClass } from "../utility/getByClass"


export const eHookUpdateCemForm = async () => {
	console.log('hooking event: eUpdateCemForm')
	const scoreRows = getByClass('cem-update-score-row')
	for (let i = 0; i < scoreRows.length; i++) {
		const row = scoreRows[i] as Element
		const downButton = row.querySelector('.cem-score-update-down-arrow') as Element
		const upButton = row.querySelector('.cem-score-update-up-arrow') as Element
		const input = row.querySelector('.cem-score-update-input') as HTMLInputElement
		downButton.addEventListener('click', () => {
			const currentScore = Number(input.value) as number
			if (currentScore - 1 > -1) {
				input.value = String(currentScore - 1)
			} 
		})
		upButton.addEventListener('click', () => {
			const currentScore = Number(input.value) as number
			if (currentScore + 1 < 101) {
				input.value = String(currentScore + 1)
			}
		})
	}
}