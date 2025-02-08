package sections

import (
	"errors"
	"fmt"

	anc "goweb/ancillaries"
	"goweb/db"
	"goweb/db/photos"
	"goweb/db/relations"
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
  queryList := ""
	for _, id := range ids {
    queryList += fmt.Sprintf("%d,", id)
	}
  queryList = queryList[0:len(queryList)-1]
  query := fmt.Sprintf("SELECT * FROM sections WHERE id in (%s)", queryList)

	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.Query(query)).([]any)
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
  if len(ids) == 0 {
    return nil
  }

  anc.Must(nil, photos.DeleteAll(ids))
  anc.Must(nil, relations.DeleteAll(ids))

  queryList := ""
	for _, id := range ids {
    queryList += fmt.Sprintf("%d,", id)
	}
  queryList = queryList[0:len(queryList)-1]
  query := fmt.Sprintf("DELETE FROM sections WHERE id in (%s)", queryList)

	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.Query(query))
	return nil
}
