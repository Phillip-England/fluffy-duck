export function removeClassesFromElements(elements: Element[], classes: string[]): void {
	for (const element of elements) {
		element.classList.remove(...classes);
	}
}
  
