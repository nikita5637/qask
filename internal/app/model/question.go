package model

// Question model ...
type Question struct {
	Sitename string `json:"sitename"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Comment  string `json:"comment,omitempty"`
}
