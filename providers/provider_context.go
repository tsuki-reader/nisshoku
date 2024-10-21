package providers

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

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
