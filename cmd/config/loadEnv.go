package config

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		_ = godotenv.Load("/go/bin/.env")
	}
}
