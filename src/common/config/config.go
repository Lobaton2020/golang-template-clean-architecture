package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	Host,
	Password,
	Username,
	Charset,
	Dbname,
	Port string
}

type ConfigApp struct {
	Port string
}
type Config struct {
	App *ConfigApp
	DB  *ConfigDB
}

const BAD_REQUEST = "BAD_REQUEST:%v"

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("Failed to load the .env file")
		panic(err)
	}
	return &Config{
		&ConfigApp{ Port: os.Getenv("HTTP_PORT") },
		&ConfigDB{
			Host: os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASSWORD"),
			Username: os.Getenv("DB_USER_NAME"),
			Charset: os.Getenv("DB_CHARSET"),
			Dbname: os.Getenv("DB_NAME"),
			Port: os.Getenv("DB_PORT"),
		},
	}
}