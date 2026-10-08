package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the github username")
		return
	}
	githubUsername := os.Args[1]
	apiFormat := "https://api.github.com/users/" + githubUsername + "/events"
	fmt.Println("GitHub Username: ", githubUsername)
	fmt.Println("API: ", apiFormat)

	resp, err := http.Get(apiFormat)

	if err != nil {
		fmt.Println("Error occured while doing GET request", err)
		return
	}
	fmt.Println("Http status: ", resp.StatusCode, "ok")
}
