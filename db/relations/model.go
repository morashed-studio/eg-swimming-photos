package users

type DataModel struct {
	Parent  int
  Child  int
}

func parseRow(row []any) DataModel {
	return DataModel{
		Parent:  row[0].(int),
		Child:  row[1].(int),
	}
}

func parseModel(m *DataModel) map[string]any {
	var modelMap = make(map[string]any)
	if m.Parent != 0 {
		modelMap["Parent"] = m.Parent
	}
	if m.Child != 0 {
		modelMap["Child"] = m.Child
	}
	return modelMap
}
