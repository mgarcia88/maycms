package DTO

type ContentRequest struct {
	Title       string `json:"Title" binding:"required,min=10,max=100"`
	ContentText string `json:"ContentText" binding:"required,min=20,max=255"`
	Status      string `json:"Status" binding:"required,min=5,max=15"`
}
