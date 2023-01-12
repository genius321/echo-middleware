package main

import (
	"log"

	"github.com/genius321/echo-middleware/internal/pkg/app"
)

func main() {
	a := app.New()

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
