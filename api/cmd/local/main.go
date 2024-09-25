package main

import (
	"log"

	"github.com/jasmaa/gin-lambda-example/pkg/app"
)

func main() {
	store := app.CreateLocalStore()
	r := app.CreateRouter(store)
	log.Printf("Starting server...")
	r.Run()
}
