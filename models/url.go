package models

type Url struct {
	Id            uint   `json:"id"`
	Full_url      string `json:"full_url"`
	Shortened_url string `json:"shortened_url"`
	Generated_by  uint   `json:"generated_by"`
}
