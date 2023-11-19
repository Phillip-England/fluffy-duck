package components

import (
	"fluffy-duck/pkg/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CemFormSelector(name string, score string) string {
	lowercaseName := strings.ToLower(name)
	return fmt.Sprintf(`
		<div class='cem-update-score-row flex flex-row justify-between mx-2'>
			<label class='text-sm'>%s</label>
			<div class='flex flex-row items-center justify-evenly gap-6'>
				<div class='cem-score-update-toggle-button cem-score-update-down-arrow border cursor-pointer border-lightGray rounded-full h-6 w-6 flex items-center justify-center'>
					<i class='fa-solid fa-arrow-left p-2'></i>
				</div>
				<input class='cem-score-update-input text-center w-8 text-sm bg-navy cursor-default focus:outline-none' value='%s' name='%s' readonly>
				<div  class='cem-score-update-toggle-button cem-score-update-up-arrow border cursor-pointer border-lightGray rounded-full h-6 w-6 flex items-center justify-center'>
					<i class='fa-solid fa-arrow-right p-2'></i>
				</div>
			</div>
		</div>
	`, name, score, lowercaseName)
}

func CemTimescaleButton(text string, currentTimescale string) string {
	lowercaseText := strings.ToLower(text)
	href := "/admin/cem/update?timescale=" + lowercaseText
	if (lowercaseText == currentTimescale) {
		return fmt.Sprintf(`
			<a _="on click toggle .hidden on .cem-update-timescale-loader-group" href='%s' class='update-cem-timescale-toggle bg-sky border border-sky text-navy rounded-full w-fit text-xs font-bold p-2'>%s</a>
		`, href, text)
	}
	return fmt.Sprintf(`
		<a _="on click toggle .hidden on .cem-update-timescale-loader-group" href='%s' class='update-cem-timescale-toggle bg-navy border border-sky text-sky rounded-full w-fit text-xs font-bold p-2'>%s</a>
	`, href, text)

} 

func UpdateCemForm(c *gin.Context, scores *models.CemScore) string {
	osatSelector := CemFormSelector("OSAT", strconv.Itoa(scores.Osat))
	tasteSelector := CemFormSelector("Taste", strconv.Itoa(scores.Taste))
	speedSelector := CemFormSelector("Speed", strconv.Itoa(scores.Speed))
	aceSelector := CemFormSelector("ACE", strconv.Itoa(scores.Ace))
	cleanlinessSelector := CemFormSelector("Cleanliness", strconv.Itoa(scores.Cleanliness))
	accuracySelector := CemFormSelector("Accuracy", strconv.Itoa(scores.Accuracy))
	currentMonthButton := CemTimescaleButton("Current Month", scores.Timescale)
	threeMonthRollingButton := CemTimescaleButton("Three Month Rolling", scores.Timescale)
	yearToDateButton := CemTimescaleButton("Year To Date", scores.Timescale)
	annualGoalButton := CemTimescaleButton("Annual Goal", scores.Timescale)
	formLoader := FormLoader("update-cem-form-loader")
	formAction := "/action/cem/update?timescale=" + scores.Timescale
	form := fmt.Sprintf(`
		<form id='update-cem-form' action='%s' method="POST" class='flex flex-col gap-8 px-4'>
			<h2 class='text-xl font-bold'>Which time period would you like to update?</h2>
			<div class='cem-update-timescale-loader-group hidden flex justify-center items-center'>
				<div class='h-20 w-20 bg-navy rounded-full border-t-4 border-navyLight border-t-sky animate-spin'></div>
			</div>
			<div class='cem-update-timescale-loader-group flex container mx-auto flex-wrap gap-4'>
				%s%s%s%s
			</div>
			<h2 class='text-xl font-bold'>Score Selection</h2>
			<div class='flex flex-col gap-6'>
				%s%s%s%s%s%s
			<div>
			<div class='flex justify-center items-center mt-6'>
				<input _='on click toggle .hidden on me then toggle .hidden on the next <div/>' id='update-cem-form-submit' type='submit' value="Submit" class='text-xs w-full p-2 bg-blue-600 rounded-2xl font-bold text-navy  bg-cyan focus:outline-none focus:bg-cyanLight' />
				%s
			</div>
		</form>
	`, formAction, currentMonthButton, threeMonthRollingButton, yearToDateButton, annualGoalButton, osatSelector, tasteSelector, speedSelector, aceSelector, cleanlinessSelector, accuracySelector, formLoader)
	return form
}