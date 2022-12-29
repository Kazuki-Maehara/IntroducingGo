package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
    <html>
      <head>
        <title>Hello, World</title>
        <link rel="stylesheet" href="/assets/test.css">
      </head>
      <body>
        Hello, World!
      </body>
    </html>`,
	)

}

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("./assets")),
		),
	)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
