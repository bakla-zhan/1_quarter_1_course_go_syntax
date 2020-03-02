package main

import (
	"io"
	"net/http"
	"strings"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("Name")
	if name != "" {
		builder := strings.Builder{}
		builder.WriteString(`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Hello `)
		builder.WriteString(name)
		builder.WriteString(`!</title>
			</head>
			<body>
				<p>Hello, <b><i>`)
		builder.WriteString(name)
		builder.WriteString(`!</i></b></p>
			</body>
			</html>`)
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, builder.String())
	} else {
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res,
			`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Hello World!</title>
		</head>
		<body>
			<p>Hello World!</p>
			<form method="GET">
				<label>Enter your Name:<br>
				<input type="text" name="Name" placeholder="Name">
				</label><br>
				<input type="submit" value="Отправить">
				<input type="reset" value="Очистить">
			</form>
		</body>
		</html>`)
	}
}

func main() {
	fs := http.FileServer(http.Dir("homepage"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", nil)
}
