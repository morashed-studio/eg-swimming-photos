package photos

import (
	"fmt"

	anc "goweb/ancillaries"
	"goweb/db"
)

// retrieves an array of photos of a specific section
func GetPhotos(sectionId int) ([]DataModel, error) {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.SeqQuery("SELECT * FROM photos WHERE section_id=$1", sectionId)).([]any)
	var res []DataModel
	for _, row := range rows {
		res = append(res, parseRow(row.([]any)))
	}
	return res, nil
}

// inserts a new photo in the database
func Add(list []DataModel) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	query := "INSERT INTO photos VALUES "
	for _, data := range list {
		query += fmt.Sprintf("(%s,%s,%d),", data.Name, data.Url, data.SectionId)
	}
	query = query[0 : len(query)-1]
	anc.Must(conn.Query(query))
	return nil
}

// removes a specific photo from the database
func Delete(id int) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.SeqQuery("DELETE FROM photos WHERE id=$1", id))
	return nil
}

// removes all photos of a specific section
func DeleteAll(sectionId int) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.SeqQuery("DELETE FROM photos WHERE section_id=$1", sectionId))
	return nil
}
