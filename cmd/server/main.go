package main

import (
    "log"
    "net/http"
    "github.com/didyouwin07/task-manager/internal/router"
)

func main() {
    r := router.SetupRouter()
    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}