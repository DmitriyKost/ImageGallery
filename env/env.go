package env

import (
	"github.com/joho/godotenv"
)


func init() { 
    if err := godotenv.Load("creds.env", "variables.env"); err != nil {
        panic("Error loading environment")
    }
}
