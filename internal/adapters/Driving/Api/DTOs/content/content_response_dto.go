package DTO

type ContentResponse struct {
	Title       string `json:"Title"`
	ContentText string `json:"ContentText"`
	Status      string `json:"Status"`
	ID          int    `json:"ID"`
}
