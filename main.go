package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
    a, _ := strconv.Atoi(r.URL.Query().Get("a"))
    b, _ := strconv.Atoi(r.URL.Query().Get("b"))
    fmt.Fprintf(w, "%d", Add(a, b))
}

func main() {
    http.HandleFunc("/add", addHandler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
