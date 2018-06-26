package main

const SIMPLE_HTML = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	{{if cfg "css"}}
		<style>
		{{cfg "css"}}
		</style>
	{{end}}
	{{if cfg "title"}}
		<title>{{cfg "title"}}</title>
	{{end}}
</head>
<body>
	{{if cfg "title"}}
		<h1>{{cfg "title"}}</h1>
	{{end}}
{{cfg "contents"}}
</body>
</html>
		`
