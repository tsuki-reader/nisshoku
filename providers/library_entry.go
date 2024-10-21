package providers

import "strings"

type LibraryEntry struct {
	IsDir    bool
	Fullpath string
	Library  string
	Name     string
}

func (entry *LibraryEntry) Valid() bool {
	return strings.HasPrefix(entry.Fullpath, entry.Library)
}
