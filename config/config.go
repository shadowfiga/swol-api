package config

import (
	"os"
)

type Config struct {
	Port   string
	DbUrl  string
	DbName string
	Secret string
}

func newConfig() Config {
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("MONGODB_URL")
	dbName := os.Getenv("MONGODB_DB")
	secret := os.Getenv("SECRET")

	if port == "" || dbUrl == "" || secret == "" || dbName == "" {
		panic("Missing environment variables")
	}

	return Config{
		Port:   port,
		DbUrl:  dbUrl,
		Secret: secret,
		DbName: dbName,
	}
}

var (
	Conf = newConfig()
)
