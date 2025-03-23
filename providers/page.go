package providers

type Page struct {
	Provider   string `json:"provider"`
	ImageURL   string `json:"image_url"`
	PageNumber int    `json:"page_number"`
}
