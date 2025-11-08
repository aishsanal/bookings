package models

type TemplateData struct {
	IntMap    map[string]int
	StringMap map[string]string
	FloatMap  map[string]float32
	Data      map[string]any
}