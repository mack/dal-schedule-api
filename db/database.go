package db

import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

type Class struct {
  Id int `json:"id"`
  Crn int `json:"crn"`
  Section string `json:"section"`
  Type string `json:"type"`
  CreditHours float32 `json:"credit_hours"`
  Days string `json:"days"` // maybe make this an array
  Times string `json:"times"` // make this a different data type easier to work with
  Location string `json:"location"`
  Max string `json:"max"` // make this int array
  Current string `json:"current"` // make this int array
  Waitlist string `json:"waitlist"`
  prof string `json:"prof"`
}

type Course struct {
  Title string `json:"title"`
}

func Open() {
  db, err := sql.Open("mysql", "root:lalaland123@/dal")


}
