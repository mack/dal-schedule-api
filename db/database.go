package db

import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "encoding/json"
  "strconv"
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
  Category string `json:"category"`
  Code string `json:"code"`
  Title string `json:"title"`
  Classes []Class  `json:"classes"`
}

func Open(code string) []byte {
  db, _ := sql.Open("mysql", "root:lalaland123@/dal")

  // seeing if I can get a working json return
  var courseQuery string = "SELECT id, category, code, title FROM courses WHERE category='" + code + "';"
  rows, _ := db.Query(courseQuery)

  var storage struct {
    Data []Course `json:"data"`
  }

  defer rows.Close()
  for rows.Next() {

    var c Course
    var cid int
    rows.Scan(&cid, &c.Category, &c.Code, &c.Title)

    var classQuery string = "SELECT classes.* FROM classes INNER JOIN course_classes ON course_classes.class=classes.id INNER JOIN courses ON course_classes.course=courses.id where course_classes.course=" + strconv.Itoa(cid) + ";"

    classRows, _ := db.Query(classQuery)
    for classRows.Next() {
      var cl Class
      classRows.Scan(&cl.Id, &cl.Crn, &cl.Section, &cl.Type, &cl.CreditHours, &cl.Days, &cl.Times, &cl.Location, &cl.Max, &cl.Current, &cl.Waitlist, &cl.Prof)
      c.Classes = append(c.Classes, cl)
    }
    storage.Data = append(storage.Data, c)
  }

  p, _ := json.Marshal(storage)

  return p;

}
