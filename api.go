package main

import (
  "net/http"
  "fmt"
  "dal-api/db"
)

func main() {
  fmt.Printf("Hello, world. \n")
  http.HandleFunc("/api", welcome)
  http.HandleFunc("/api/courses", courses)
  http.HandleFunc("/api/course_list", courseList)
  http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, string(db.Open("CSCI")))
}

func courses(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Return JSON format")
}

func courseList(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Return JSON format")
}
