package main

import (
	"net/http"

	"github.com/jaehanbyun/api/app"
)

func main() {
	a := app.MakeHandler()

	http.ListenAndServe(":8080", a)
}
