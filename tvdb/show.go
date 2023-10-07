package tvdb

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/maps"
)

// ----------------------Show------------------------
type Show struct {
	show     string
	episodes map[string]*Episode
}

func (sh *Show) Value() string {
	return sh.show
}

func (sh *Show) Add(e Episode) {
	sh.episodes[e.Value()] = &e
}

func (sh *Show) Print() {
	for _, episode := range sh.Episodes() {
		fmt.Printf("  %v - %v\n", episode.Episode, episode.Title)
	}
}

func (sh *Show) Episodes() []*Episode {
	toReturn := make([]*Episode, 0, len(sh.episodes))
	for _, s := range sortedKeys(maps.Keys(sh.episodes)) {
		toReturn = append(toReturn, sh.episodes[s])
	}

	return toReturn
}

func (sh *Show) Match(input string) bool {
	return strings.Contains(strings.ToLower(sh.show), input)
}

func (sh *Show) SelectEpisode(file string) (*Episode, error) {
	episodes := sh.Episodes()
	prompt := promptui.Select{
		Label: "Select episode for" + file,
		Items: episodes,
		Searcher: func(input string, index int) bool {
			return episodes[index].Match(strings.ToLower(input))
		},
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Error getting season %v", err)
	}
	return episodes[i], nil
}

// ----------------------Show------------------------
