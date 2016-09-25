package main 

import (
    "net/http"
    "html"
    "log"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", html.EscapeString(r.URL.Path[1:]))
}

func main() {
    http.HandleFunc("/", defaultHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

