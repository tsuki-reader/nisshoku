package providers

type Chapter struct {
	Title          string `json:"title"`
	ID             string `json:"id"`
	Provider       string `json:"provider"`
	AbsoluteNumber int    `json:"absoluteNumber"`
}
