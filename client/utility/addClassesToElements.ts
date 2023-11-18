
export function addClassesToElements(elements: Element[], classes: string[]): void {
	for (const element of elements) {
		element.classList.add(...classes);
	}
}