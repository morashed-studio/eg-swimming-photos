package photos

import (
	"fmt"

	anc "goweb/ancillaries"
	"goweb/db"
)

// retrieves an array of photos of a specific section
func Get(sectionId int) ([]DataModel, error) {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	rows := anc.Must(conn.Query("SELECT * FROM photos WHERE section_id=$1", sectionId)).([]any)
	var res []DataModel
	for _, row := range rows {
		res = append(res, parseRow(row.([]any)))
	}
	return res, nil
}

// inserts a new photo in the database
func Add(list []DataModel) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	query := "INSERT INTO photos (name, url, section_id) VALUES "
	for _, data := range list {
		query += fmt.Sprintf("('%s','%s',%d),", data.Name, data.Url, data.SectionId)
	}
	query = query[0 : len(query)-1]
	anc.Must(conn.Query(query))
	return nil
}

// removes a specific photo from the database
func Delete(id int) error {
	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.Query("DELETE FROM photos WHERE id=$1", id))
	return nil
}

// removes all photos of a specific section
func DeleteAll(sectionIds []int) error {
  queryList := ""
	for _, id := range sectionIds {
    queryList += fmt.Sprintf("%d,", id)
	}
  queryList = queryList[0:len(queryList)-1]
  query := fmt.Sprintf("DELETE FROM photos WHERE section_id IN (%s)", queryList)

	conn := anc.Must(db.GetConnection()).(*db.Connection)
	anc.Must(conn.Query(query))
	return nil
}
