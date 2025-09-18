package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("🚀 TunnlrX server starting on :8080")
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
    http.ListenAndServe(":8080", nil)
}
