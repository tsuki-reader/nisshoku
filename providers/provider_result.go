package providers

type ProviderResult struct {
	Title            string
	ID               string
	Provider         string
	AlternativeNames []string
	StartYear        int
}

func (pr *ProviderResult) Chapters(provider Provider) ([]Chapter, error) {
	if pr == nil {
		return []Chapter{}, nil
	}
	if pr.ID == "" {
		return []Chapter{}, nil
	}
	return provider.GetChapters(pr.ID)
}
