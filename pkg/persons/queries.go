package persons

func (p Persons) SelectAllQuery() string {
	return "select * from person;"
}