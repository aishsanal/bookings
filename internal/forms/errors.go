package forms

type errors map[string][]string

func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	errorList := e[field]
	if errorList != nil {
		return errorList[0]
	}
	return ""
}
