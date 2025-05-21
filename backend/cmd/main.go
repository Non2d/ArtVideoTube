package main

import (
    "log"
    "net/http"
    "os"

    "go-render-api/internal/handlers"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/items", handlers.GetItems)
    http.HandleFunc("/items/create", handlers.CreateItem)

    log.Printf("Starting server on :%s\n", port)
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}