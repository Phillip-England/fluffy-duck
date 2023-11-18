

export const getById = (elementId: string) => {
	const element = document.getElementById(elementId) as Element
	if (element) {
		return element
	}
	console.error(`failed to locate element by id: ${elementId}`)
	const div = document.createElement('div') as Element
	div.innerHTML = `failed to located element by id: ${elementId}`
	return document.createElement('div') as Element
}