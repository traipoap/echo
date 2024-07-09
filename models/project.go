// models/project.go
package models

type Project struct {
	ID          int
	Title       string
	Year        int
	Technology  string
	Description string
	ImageURL    string
	Slug        string
}
