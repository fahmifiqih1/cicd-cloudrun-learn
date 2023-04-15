package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    	fmt.Fprintf(w, "HELLO Mr.Fahmi", html.EscapeString(r.URL.Path))
    })
    
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "HI Mr.Fiqih", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}
