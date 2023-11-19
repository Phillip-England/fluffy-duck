package components

import "fmt"

func Root(title string, children string) string {
	root := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en" class="h-full">
		<head>
			<meta charset="UTF-8">			
			<link rel="icon" href="/public/favicon.ico" type="image/x-icon">
			<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
			<script src="https://kit.fontawesome.com/ef0709a418.js" crossorigin="anonymous"></script>
			<script src="https://unpkg.com/hyperscript.org@0.9.12"></script>
			<link rel="stylesheet" type="text/css" href="/public/output.css">
			<title>%s</title>
		</head>
		<body hx-boost='true' class='min-h-screen bg-navy text-lightGray'>
			%s
			<script src='/public/index.js'></script>
		</body>
		</html>
	`, title, children)
	return root
}
