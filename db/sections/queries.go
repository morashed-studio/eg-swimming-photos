package users

import (
	"fmt"

	anc "goweb/ancillaries"
	"goweb/db"
)

// retrieves an array of sections with the passed ids.
func Get(ids []int) ([]DataModel, error) {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.SeqQuery("SELECT * FROM sections WHERE id in $1", ids)).([][]any)
	var res []DataModel
	for _, row := range rows {
		res = append(res, parseRow(row))
	}
	return res, nil
}

// inserts a new list of sections in the database.
func Add(list []DataModel) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	query := "INSERT INTO sections VALUES "
	for _, data := range list {
		query += fmt.Sprintf("(%s),", data.Title)
	}
	query = query[0 : len(query)-1]
	anc.Must(conn.Query(query))
	return nil
}

// removes sections with the passed ids from the database.
func Delete(ids []int) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.SeqQuery("DELETE FROM sections WHERE id in $1", ids))
	return nil
}
