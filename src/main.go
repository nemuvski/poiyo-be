package main

import (
	"fmt"
	"poiyo-be/src/environment"
)

func main() {
	dotenv := environment.Load()
	message := fmt.Sprintf("%s:%s User:%s PW:%s",
		dotenv.Host,
		dotenv.Port,
		dotenv.User,
		dotenv.Pass,
	)
	fmt.Println(message)
}
