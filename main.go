package main

import (
	"log"
	"net/http"

	"github.com/jaehanbyun/api/app"
)

func main() {
	a := app.MakeHandler()
	if err := http.ListenAndServe(":8080", a); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
