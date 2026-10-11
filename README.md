# GitHub Activity CLI

A simple command-line application built with **Go (Golang)** that fetches recent activity from a GitHub user's profile using the GitHub Events API.

## Features

* Fetch recent activity for any GitHub username.
* Display repository activity in a readable format.
* Filter events by type, such as `push` and `create`.
* Support case-insensitive event filtering.
* Display formatted timestamps.
* Handle HTTP errors and invalid event filters.
* Include unit tests for event type normalisation.

## Getting Started

### Prerequisites

* [Go](https://go.dev/dl/) installed on your machine.
* A terminal or command prompt.

### Run the project

Clone the repository:

git clone https://github.com/codieSam/github-user-activity.git
cd github-user-activity


Run the application:


go run . <github_username>


Filter by event type:


go run . codieSam push
go run . codieSam create


Run the tests:


go test -v


## Technologies

* Go (Golang)
* GitHub Events API
* JSON
* Go standard library and testing package

## Purpose

This project was built as a hands-on learning exercise to improve my understanding of Go, HTTP requests, JSON parsing, error handling, command-line applications, and unit testing.

## Author

**Samrat Belbase**

GitHub: [@codieSam](https://github.com/codieSam)
