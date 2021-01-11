package main

import (
	"poiyo-be/src/environment"
	"poiyo-be/src/router"
)

func main() {
	environment.Load()
	r := router.Init()
	r.Logger.Fatal(r.Start(":1323"))
}
