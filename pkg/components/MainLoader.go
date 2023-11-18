package components

func MainLoader() string {
	return `
		<div id='main-loader' class="fixed hidden top-0 left-0 w-full h-full flex items-center justify-center z-50 bg-black bg-opacity-50">
			<div class="w-24 h-24 border-8 border-black border-t-red-700 rounded-full mb-80 animate-spin"></div>
		</div>
	`
}
