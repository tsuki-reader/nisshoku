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

	// Either comics or manga
	ProviderType() ProviderType
}
