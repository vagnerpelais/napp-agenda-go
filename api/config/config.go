package config

import (
	"os"
)

var (
	ConnectionString = ""
	Port             = 0
	SecretKey        []byte
)

func Load() {
	// var erro error

	// if erro = godotenv.Load(); erro != nil {
	// 	log.Fatal((erro))
	// }

	// ConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"),
	// 	os.Getenv("POSTGRES_DB"),
	// )

	ConnectionString = os.Getenv("DATABASE_URL")
}
