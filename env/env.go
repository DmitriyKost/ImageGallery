package env

import (
	"github.com/joho/godotenv"
)


func init() {
    InitEnv()
}

func InitEnv() {
    if err := godotenv.Load("config/creds.env", "config/variables.env"); err != nil {
        panic("Error loading environment")
    }
}
