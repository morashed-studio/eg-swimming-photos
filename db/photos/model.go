package photos

type DataModel struct {
	Id        int
	Name      string
	Url       string
	SectionId int
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:        int(row[0].(int32)),
		Name:      row[1].(string),
		Url:       row[2].(string),
		SectionId: int(row[3].(int32)),
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
