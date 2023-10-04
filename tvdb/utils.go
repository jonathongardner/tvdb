// Go program to sort the map by Keys
package tvdb

import (
	"sort"

	"github.com/manifoldco/promptui"
)

func sortedKeys(toSort []string) []string {
	sort.Strings(toSort)

	return toSort
}

// type matcher interface {
// 	Match(string) bool
// }

// func searcher(mats []matcher) func(string, int) bool {
// 	return func(input string, index int) bool {
// 		return mats[index].Match(strings.ToLower(input))
// 	}
// }

var templates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "\U0001F3A5 {{ .Value }} ",
	Inactive: "  {{ .Value }} ",
	Selected: "\U0001F3A5 {{ .Value }} ",
}
