package Auth

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// structs
type UserENV struct {
	CLIENT_ID     string
	CLIENT_SECRET string
}

// error handles SPECIFIC to Auth
func errorHandler(err error) {
	if err != nil {
		log.Fatal("something went wrong with the function: ENV -> ", err)
	}
}

// Env handles the parsing of the apps env files.
func Env() *UserENV {

	err := godotenv.Load()
	errorHandler(err)

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	fmt.Printf("%s, %s", clientID, clientSecret)

	return &UserENV{
		clientID,
		clientSecret,
	}
}
