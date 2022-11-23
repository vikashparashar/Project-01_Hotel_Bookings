package models

// Template_Date hold's the data sent from handlers to templates
type Template_Data struct {
	StringMap map[string]string
	IntMap    map[string]int64
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
