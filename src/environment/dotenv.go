package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DotEnv struct {
	Host string
	Port string
	User string
	Pass string
}

func Load() *DotEnv {
	if os.Getenv("GO_EXEC_ENV") == "" {
		os.Setenv("GO_EXEC_ENV", "development")
	}
	envfile := fmt.Sprintf(".env.%s", os.Getenv("GO_EXEC_ENV"))
	err := godotenv.Load(envfile)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s を読み込めませんでした。", envfile))
	}
	return &DotEnv{
		Host: os.Getenv("PG_HOST"),
		Port: os.Getenv("PG_PORT"),
		User: os.Getenv("PG_USER"),
		Pass: os.Getenv("PG_PASSWORD"),
	}
}
