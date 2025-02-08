package sections

import (
	"errors"
	"fmt"

	anc "goweb/ancillaries"
	"goweb/db"
)

// retrieves the id of a specific section title.
// NOTE: this assumes that sections cannot have the same title.
func GetId(title string) (int, error) {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.Query("SELECT * FROM sections WHERE title=$1", title)).([]any)
  if len(rows) == 0 {
    return 0, errors.New("Section not found.")
  }
	var row = parseRow(rows[0].([]any))
	return row.Id, nil
}

// retrieves an array of sections with the passed ids.
func Get(ids []int) ([]DataModel, error) {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.Query("SELECT * FROM sections WHERE id in $1", ids)).([]any)
	var res []DataModel
	for _, row := range rows {
		res = append(res, parseRow(row.([]any)))
	}
	return res, nil
}

// retrieves an array of sections with the passed ids.
func GetAll() ([]DataModel, error) {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.Query("SELECT * FROM sections")).([]any)
	var res []DataModel
	for _, row := range rows {
		res = append(res, parseRow(row.([]any)))
	}
	return res, nil
}

// inserts a new list of sections in the database.
func Add(list []DataModel) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	query := "INSERT INTO sections (title) VALUES "
	for _, data := range list {
		query += fmt.Sprintf("('%s'),", data.Title)
	}
	query = query[0 : len(query)-1]
	anc.Must(conn.Query(query))
	return nil
}

// removes sections with the passed ids from the database.
func Delete(ids []int) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.Query("DELETE FROM sections WHERE id in $1", ids))
	return nil
}
