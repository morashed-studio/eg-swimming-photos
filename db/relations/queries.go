package users

import (
	"fmt"

	anc "goweb/ancillaries"
	"goweb/db"
)

// retrieves an array of children of a specific parent
func GetSectionsOf(id int) ([]int, error) {
  conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.SeqQuery("SELECT child FROM relations WHERE parent=$1", id)).([]int)
	return rows, nil
}

// inserts a new (parent-child) section relation in the database.
func Add(list []DataModel) error {
  conn := anc.Must(db.GetConnection()).(*db.Connection)
  query := "INSERT INTO relations VALUES "
  for _, data := range list {
    query += fmt.Sprintf("(%d,%d),", data.Parent, data.Child)
  }
  query = query[0:len(query)-1]
  anc.Must(conn.Query(query))
	return nil
}

// removes a specific relation from the database
func Delete(data DataModel) error {
  conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.SeqQuery(
    "DELETE FROM relations WHERE parent=$1 AND child=$2", 
    data.Parent, 
    data.Child,
  ))
	return nil
}

// removes specific parent relations from the database
func DeleteAll(parent int) error {
  conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.SeqQuery( "DELETE FROM relations WHERE parent=$1", parent))
	return nil
}

// return true if there is no parent with the passed id
func isAlbum(id int) bool {
  conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.SeqQuery("SELECT child FROM relations WHERE parent=$1", id)).([]int)
	return len(rows) == 0
}
