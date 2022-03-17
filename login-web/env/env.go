package env

import (
	"log"

	"github.com/joho/godotenv"
)

func Read() map[string]string {
	env, err := godotenv.Read("env/.env")
	checkErr(err)

	return env
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
