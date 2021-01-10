package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if os.Getenv("GO_EXEC_ENV") == "" {
		os.Setenv("GO_EXEC_ENV", "development")
	}
	envfile := fmt.Sprintf(".env.%s", os.Getenv("GO_EXEC_ENV"))
	err := godotenv.Load(envfile)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s を読み込めませんでした。", envfile))
	}
}
