package tvdb

import (
	"path/filepath"
	"strings"
)

type Episode struct {
	Show     string `json:"show"`
	Year     string `json:"year"`
	Season   string `json:"season"`
	Episode  string `json:"episode"`
	Title    string `json:"title"`
	PlexName string `json:"plexName"`
}

func (e *Episode) Value() string {
	return "(" + e.Episode + ") " + e.Title
}

func (e *Episode) Match(input string) bool {
	return input == e.Episode || strings.Contains(strings.ToLower(e.Title), strings.ToLower(input))
}

func (e *Episode) showWithYear() string {
	return e.Show + " (" + e.Year + ")"
}

// /Band of Brothers (2001)/Season 01/Band of Brothers (2001) - s01e01 - Currahee.mkv
func (e *Episode) Dir() string {
	return filepath.Join(e.showWithYear(), "Season "+e.Season)
}

func (e *Episode) Filename() string {
	return e.showWithYear() + " - s" + e.Season + "e" + e.Episode + " - " + e.Title + ".mkv"
}

func (e *Episode) FullPath() string {
	return filepath.Join(e.Dir(), e.Filename())
}
