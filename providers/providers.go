package providers

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

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

type ProviderType string

const (
	Manga ProviderType = "MANGA"
	Comic ProviderType = "COMICS"
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

type LibraryEntry struct {
	IsDir    bool
	Fullpath string
	Library  string
	Name     string
}

func (entry *LibraryEntry) Valid() bool {
	return strings.HasPrefix(entry.Fullpath, entry.Library)
}

type ProviderContext struct {
	MangaLibraryPath string
	ComicLibraryPath string
	ProviderType     ProviderType
}

func (ctx *ProviderContext) WalkLibrary(libraryPath string) ([]LibraryEntry, error) {
	results := []LibraryEntry{}

	if !strings.HasPrefix(libraryPath, ctx.ComicLibraryPath) || !strings.HasPrefix(libraryPath, ctx.MangaLibraryPath) {
		return results, errors.New("cannot walk directory")
	}

	entries, err := os.ReadDir(libraryPath)
	if err != nil {
		return []LibraryEntry{}, err
	}

	for _, e := range entries {
		entry := LibraryEntry{
			Fullpath: filepath.Join(libraryPath, e.Name()),
			IsDir:    e.IsDir(),
			Library:  libraryPath,
			Name:     e.Name(),
		}

		results = append(results, entry)
	}

	return results, nil
}
