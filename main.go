package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Event struct {
	Type      string     `json:"type"`
	Repo      Repository `json:"repo"`
	CreatedAt string     `json:"created_at"`
}

type Repository struct {
	Name string `json:"name"`
}

func normalizeEventType(eventType string) string {
	etype := strings.ToLower(eventType)
	switch etype {
	case "push":
		eventType = "PushEvent"
	case "create":
		eventType = "CreateEvent"
	}
	return eventType

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide the github username")
		return
	}
	githubUsername := strings.TrimSpace(os.Args[1])

	if githubUsername == "" {
		fmt.Println("Please provide a valid GitHub Username.")
		return
	}
	apiFormat := "https://api.github.com/users/" + githubUsername + "/events"
	fmt.Println("GitHub Username: ", githubUsername)
	fmt.Println("API: ", apiFormat)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(apiFormat)
	if err != nil {
		fmt.Println("Error occured while doing GET request", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: GitHub API returned.", resp.Status)
		return
	}
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
	var eventType string

	if len(os.Args) >= 3 {
		eventType = normalizeEventType(os.Args[2])
	}
	eventCounter := 0
	for i := range events {
		Type := events[i].Type
		Repo := events[i].Repo.Name
		CreatedAt := events[i].CreatedAt

		parsedTime, err := time.Parse(time.RFC3339, CreatedAt)

		if err != nil {
			fmt.Println("Error: while pasrsing the time for createdAt.", err)
			return
		}

		createdTime := parsedTime.Format("02 Jan 2006, 15:04 MST")

		if eventType != "" && !strings.EqualFold(eventType, Type) {
			continue
		}
		eventCounter++

		switch Type {
		case "PushEvent":
			fmt.Println("Pushed commits to", Repo, "on", createdTime)
		case "CreateEvent":
			fmt.Println("Created an event in", Repo, "on", createdTime)
		default:
			fmt.Println("Activity:", Type, "in", Repo)
		}

		// fmt.Println("Type: ", Type)
		// fmt.Println("Repo: ", Repo)
		// fmt.Println("CreatedAt: ", CreatedAt)
	}
	if eventCounter == 0 {
		fmt.Println("No such event found.")
	}

}
