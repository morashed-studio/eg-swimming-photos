package users

type DataModel struct {
	Id  int
  Title  string
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:  row[0].(int),
		Title:  row[1].(string),
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
