package DTO

type ContentUserResponse struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type ContentResponse struct {
	Title       string              `json:"Title"`
	ContentText string              `json:"ContentText"`
	Status      string              `json:"Status"`
	ID          int                 `json:"ID"`
	User        ContentUserResponse `json:"User"`
	MainImage   string              `json:"MainImage"`
}
