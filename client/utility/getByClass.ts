

export const getByClass = (elementClassName: string) => {
	const elements = document.getElementsByClassName(elementClassName)
	if (elements.length > 0) {
		return elements
	}
	console.error(`failed to locate elements by class name: ${elementClassName}`)
	return `<p>failed to locate elements by class name: ${elementClassName}</p>`
}