package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type Response struct {
    Email          string `json:"email"`
    CurrentDatetime string `json:"current_datetime"`
    GithubURL      string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Your registered email address
    email := "adewuyibabajide@gmail.com"

    // Current datetime as an ISO 8601 formatted timestamp
    currentDatetime := time.Now().UTC().Format(time.RFC3339)

    // GitHub URL of the project's codebase
    githubURL := "https://github.com/adewuyi22/HNG12-task"

    // Create the JSON response
    response := Response{
        Email:           email,
        CurrentDatetime: currentDatetime,
        GithubURL:       githubURL,
    }

    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/api", handler)
    http.ListenAndServe(":8080", nil)
}
