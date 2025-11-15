package models

import "github.com/aishsanal/bookings/internal/forms"

type TemplateData struct {
	IntMap    map[string]int
	StringMap map[string]string
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Form      *forms.Form
}