package db

import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "encoding/json"
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
  Prof string `json:"prof"`
}

type Course struct {
  Title string `json:"title"`
}

func Open() []byte {
  db, err := sql.Open("mysql", "root:lalaland123@/dal")

  // seeing if I can get a working json return
  var queryStr string = "SELECT * FROM classes"
  rows, _ := db.Query(queryStr)

  var storage struct {
    Data []Class  `json:"data"`
  }

  defer rows.Close()
  for rows.Next() {
    var c Class
    rows.Scan(&c.Id, &c.Crn, &c.Section, &c.Type, &c.CreditHours, &c.Days, &c.Times, &c.Location, &c.Max, &c.Current, &c.Waitlist, &c.Prof)
    storage.Data = append(storage.Data, c)
  }

  p, _ := json.Marshal(storage)

  return p;

}
