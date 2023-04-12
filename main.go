package main

import (
	"fmt"
	"os/user"

	"github.com/joho/godotenv"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("user.Current error: %v\n", err)
		return
	}
	envPath := usr.HomeDir + "/.bg/.env"
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Printf("godotenv.Load error: %v\n", err)
		return
	}
}
