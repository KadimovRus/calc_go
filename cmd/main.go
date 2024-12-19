package main

import (
	"finalTask/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
