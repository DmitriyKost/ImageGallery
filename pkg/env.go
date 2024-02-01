package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() { 
    if err := godotenv.Load("creds.env"); err != nil {
        panic("Error loading credentials")
    }
    InitEnv()
}


func InitEnv() {
    jwtKey = []byte(getEnv("JWT_KEY"))
    users = map[string]string{
	getEnv("ADMIN_LOGIN"): getEnv("ADMIN_PASSWORD"),
    "user2": "password2", // TODO: my personal access to website
    }
}

// Reads an environment variable or panics if cannot
func getEnv(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    } else {
        errMessage := fmt.Sprintf("%s variable not exists", key)
        panic(errMessage)
    }
}
