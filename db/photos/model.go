package users

type DataModel struct {
	Id  int
  Name  string
  Url  string
  SectionId  int
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:  row[0].(int),
		Name:  row[1].(string),
		Url:  row[2].(string),
		SectionId:  row[3].(int),
	}
}

func parseModel(m *DataModel) map[string]any {
	var modelMap = make(map[string]any)
	if m.Id != 0 {
		modelMap["Id"] = m.Id
	}
	if m.Name != "" {
		modelMap["Name"] = m.Name
	}
	if m.Url != "" {
		modelMap["Url"] = m.Url
	}
	if m.SectionId != 0 {
		modelMap["SectionId"] = m.SectionId
	}
	return modelMap
}
