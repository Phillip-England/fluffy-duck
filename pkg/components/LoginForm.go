package components

import (
	"fluffy-duck/pkg/utility"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginForm(c *gin.Context) string {
	usernameText := utility.IfSpanishThen(c, "username", "nombre de usuario")
	passwordText := utility.IfSpanishThen(c, "password", "contraseña")
	loginText := utility.IfSpanishThen(c, "Login", "Inicio de sesión")
	submitValue := utility.IfSpanishThen(c, "Submit", "Enviar")
	form := fmt.Sprintf(`
		<form  id='login-form' method='POST' action='/formAction/login' class='flex flex-col px-8'>
			<h2 class='text-xl mb-4 font-semibold text-lightGray'>%s</h2>
			<p id='login-form-error' class='text-red text-xs mb-4 hidden'>default error</p>
			<div class='flex flex-row mb-4 items-center border-lightGray'>
				<i class='fa-solid fa-sm fa-user text-lightGray pr-4 absolute'></i>
				<input type='text' name='username' placeholder="%s" class='text-xs bg-navy flex-grow py-2 focus:outline-none text-lightGray w-full border-b indent-6 focus:border-b-2' />
			</div>
			<div class='flex flex-row mb-10 items-center border-lightGray'>
				<i class='fa-solid fa-sm fa-key text-lightGray pr-4 absolute'></i>
				<input type='password' name='password' placeholder="%s" class='text-xs bg-navy flex-grow py-2 focus:outline-none text-lightGray w-full border-b indent-6 focus:border-b-2' />
			</div>
			<input id='login-form-submit' type='submit' value="%s" class='text-xs p-2 bg-blue-600 rounded-2xl font-bold text-navy  bg-cyan focus:outline-none focus:bg-cyanLight' />
			<div id='login-form-loader' class='self-center hidden border-2 border-navyLight border-t-cyan w-10 h-10 rounded-full animate-spin'></div>
			<a id='login-form-hidden-team-link' href='/team'></a>
			<a id='login-form-hidden-admin-link' href='/admin'></a>
		</form>
	`, loginText, usernameText, passwordText, submitValue)
	return form
}