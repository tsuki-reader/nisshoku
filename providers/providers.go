package providers

type Provider interface {
	// Returns the results from the given query
	Search(query string) ([]ProviderResult, error)

	// Returns the chapters for the given manga/comic
	GetChapters(id string) ([]Chapter, error)

	// Returns the pages for the given chapter
	GetChapterPages(id string) ([]Page, error)

	// Headers to use when serving chapter images if needed
	ImageHeaders() map[string]string

	ProviderType() ProviderType
}

type ProviderType int64

const (
	Manga ProviderType = iota
	Comic
)

type ProviderResult struct {
	Title            string
	ID               string
	Provider         string
	AlternativeNames []string
	StartYear        int
}

type Chapter struct {
	Title          string
	ID             string
	Provider       string
	Chapter        string
	AbsoluteNumber int
}

type Page struct {
	Provider   string
	ImageURL   string
	PageNumber int
}
