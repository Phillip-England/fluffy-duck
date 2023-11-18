import { getById } from "../utility/getById"


export const eHookLoginForm = () => {
	console.log('hooking event: eSubmitLoginForm')
	const form = getById('login-form')
	const error = getById('login-form-error')
	const hiddenTeamLink = getById('login-form-hidden-team-link') as HTMLAnchorElement
	const hiddenAdminLink = getById('login-form-hidden-admin-link') as HTMLAnchorElement
	const loader = getById('login-form-loader')
	const submit= getById('login-form-submit')
	form.addEventListener('submit', async (e) => {
		e.preventDefault()
		submit.classList.add('hidden')
		loader.classList.remove('hidden')
		const data = new FormData(form as HTMLFormElement)
		const username = data.get('username')
		const password = data.get('password')
		const res = await fetch('/api/user/login', {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({"username": username, "password": password})
		})
		const json = await res.json()
		if (res.status == 200) {
			if (json.redirect == "/team") {
				hiddenTeamLink.click()
				return
			}
			if (json.redirect == "/admin") {
				hiddenAdminLink.click()
				return
			}
		}
		const errorMessage = json.message
		error.innerHTML = errorMessage
		error.classList.remove('hidden')
		submit.classList.remove('hidden')
		loader.classList.add('hidden')
	})
}