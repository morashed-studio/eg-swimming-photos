package sections

type DataModel struct {
	Id    int
	Title string
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:    int(row[0].(int32)),
		Title: row[1].(string),
	}
}

func parseModel(m *DataModel) map[string]any {
	var modelMap = make(map[string]any)
	if m.Id != 0 {
		modelMap["Id"] = m.Id
	}
	if m.Title != "" {
		modelMap["Title"] = m.Title
	}
	return modelMap
}
