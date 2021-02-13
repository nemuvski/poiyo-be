package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	EXEC_ENV_DEVELOPMENT = "development"
)

// Load dotenvファイルを読み込む.
func Load() {
	if os.Getenv("GO_EXEC_ENV") == "" {
		os.Setenv("GO_EXEC_ENV", EXEC_ENV_DEVELOPMENT)
	}

	// 開発環境でのみdotenvファイルを読み込むようにする.
	// 本番環境では管理画面で環境変数を設定すること.
	if os.Getenv("GO_EXEC_ENV") == EXEC_ENV_DEVELOPMENT {
		envfile := fmt.Sprintf(".env.%s", os.Getenv("GO_EXEC_ENV"))
		err := godotenv.Load(envfile)
		if err != nil {
			log.Fatalln(fmt.Sprintf("%s を読み込めませんでした。", envfile))
		}
	}
}
