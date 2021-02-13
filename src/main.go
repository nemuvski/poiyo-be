package main

import (
	"os"

	"poiyo-be/src/environment"
	"poiyo-be/src/router"
)

func main() {
	environment.Load()
	r := router.Init()

	port := os.Getenv("ECHO_LISTEN_PORT")
	r.Logger.Fatal(r.Start(":" + port))
}
