package main

import (
	"fmt"
	"os"
	"poiyo-be/src/environment"
)

func main() {
	environment.Load()
	message := fmt.Sprintf("%s:%s User:%s PW:%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
	)
	fmt.Println(message)
}
