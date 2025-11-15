package forms

import (
	"fmt"
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(values url.Values) *Form {
	return &Form{
		values,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(key string, r *http.Request) bool {
	if r.Form.Get(key) != "" {
		return true
	}
	f.Errors.Add(key, fmt.Sprintf("Mandatory field %s not populated", key))
	return false
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
