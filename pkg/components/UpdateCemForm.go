package components

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CemFormSelector(name string, score string) string {
	return fmt.Sprintf(`
		<div class='cem-update-score-row flex flex-row justify-between mx-2'>
			<label class='text-sm'>%s</label>
			<div class='flex flex-row items-center justify-evenly gap-6'>
				<div class='cem-score-update-down-arrow border cursor-pointer border-lightGray rounded-full h-6 w-6 flex items-center justify-center'>
					<i class='fa-solid fa-arrow-left fa-sm p-1'></i>
				</div>
				<input class='cem-score-update-input text-sm bg-navy w-5 cursor-default focus:outline-none' value='%s' readonly>
				<div class='cem-score-update-up-arrow border cursor-pointer border-lightGray rounded-full h-6 w-6 flex items-center justify-center'>
					<i class='fa-solid fa-arrow-right fa-sm p-1'></i>
				</div>
			</div>
		</div>
	`, name, score)
}

func UpdateCemForm(c *gin.Context) string {
	osatSelector := CemFormSelector("OSAT", "77")
	tasteSelector := CemFormSelector("Taste", "77")
	speedSelector := CemFormSelector("Speed", "77")
	aceSelector := CemFormSelector("ACE", "77")
	cleanlinessSelector := CemFormSelector("Cleanliness", "77")
	accuracySelector := CemFormSelector("Accuracy", "77")
	form := fmt.Sprintf(`
		<form class='flex flex-col gap-8 px-4'>
			<h2 class='text-xl font-bold'>Which time period would you like to update?</h2>
			<div class='flex container mx-auto flex-wrap gap-4'>
				<button type='button' value='current month' class='update-cem-timescale-toggle bg-sky border border-sky text-navy rounded-full w-fit text-xs font-bold p-2'>Current Month</button>
				<button type='button' value='three month rolling' class='update-cem-timescale-toggle bg-navy border border-sky text-sky rounded-full w-fit text-xs font-bold p-2'>Three Month Rolling</button>
				<button type='button' value='year to date' class='update-cem-timescale-toggle bg-navy border border-sky text-sky rounded-full w-fit text-xs font-bold p-2'>Year To Date</button>
				<button type='button' value='goal' class='update-cem-timescale-toggle bg-navy border border-sky text-sky rounded-full w-fit text-xs font-bold p-2'>Annual Goal</button>
			</div>
			<h2 class='text-xl font-bold'>Score Selection</h2>
			<div class='flex flex-col gap-6'>
				%s%s%s%s%s%s
			<div>
			<input id='update-cem-form-submit' type='submit' value="Submit" class='text-xs mt-6 w-full p-2 bg-blue-600 rounded-2xl font-bold text-navy  bg-cyan focus:outline-none focus:bg-cyanLight' />
		</form>
	`, osatSelector, tasteSelector, speedSelector, aceSelector, cleanlinessSelector, accuracySelector)
	return form
}