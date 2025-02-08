package relations

type DataModel struct {
	Parent int
	Child  int
}

func parseRow(row []any) DataModel {
	return DataModel{
		Parent: int(row[0].(int32)),
		Child:  int(row[1].(int32)),
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
