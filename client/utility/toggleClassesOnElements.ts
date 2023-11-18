export function toggleClassesOnElements(elements: Element[], classes: string[]): void {
	for (const element of elements) {
		for (const className of classes) {
			element.classList.toggle(className);
		}
	}
  }
