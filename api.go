package main

import (
  "net/http"
  "fmt"
)

func main() {
  fmt.Printf("Hello, world. \n")
  http.HandleFunc("/api", welcome)
  http.HandleFunc("/api/courses", courses)
  http.HandleFunc("/api/course_list", courseList)
  http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome!")
}

func courses(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Return JSON format")
}

func courseList(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Return JSON format")
}
