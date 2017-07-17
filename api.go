package main

import (
  "net/http"
  "fmt"
  "dal-api/db"
  "time"
)

func main() {
  fmt.Printf("API has started running on port: 8080... \n")
  http.HandleFunc("/api/courses", courses)
  http.ListenAndServe(":8080", nil)
}

func courses(w http.ResponseWriter, r *http.Request) {
  t := time.Now()
	tn := t.Format("2006/01/02 15:04:05")
  fmt.Printf("[" + tn + "] New connection made")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  category := r.URL.Query().Get("s")
  fmt.Fprintf(w, string(db.Open(category)))
}
