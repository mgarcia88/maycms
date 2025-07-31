package DTO

type CategoryRequestBody struct {
	Title       string `json:"Title" binding:"required,min=10,max=100"`
	Description string `json:"Description" binding:"required,min=20,max=255"`
}
