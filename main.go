package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type      string     `json:"type"`
	Repo      Repository `json:"repo"`
	CreatedAt string     `json:"created_at"`
}

type Repository struct {
	Name string `json:"name"`
}

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
	defer resp.Body.Close()
	fmt.Println("Http status: ", resp.Status)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading the data.", err)
		return
	}
	// fmt.Println("Data: ", string(data))
	var events []Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		panic(err)

	}
	fmt.Println("Decoded Successfully: ", len(events), "events.")
	fmt.Println("Recent github activity for ", githubUsername)
	for i := range events {
		Type := events[i].Type
		Repo := events[i].Repo.Name
		CreatedAt := events[i].CreatedAt

		switch Type {
		case "PushEvent":
			fmt.Println("Pushed commits to", Repo, "on", CreatedAt)
		case "CreateEvent":
			fmt.Println("Created an event in", Repo, "on", CreatedAt)
		default:
			fmt.Println("Activity:", Type, "in", Repo)
		}

		// fmt.Println("Type: ", Type)
		// fmt.Println("Repo: ", Repo)
		// fmt.Println("CreatedAt: ", CreatedAt)
	}

}
