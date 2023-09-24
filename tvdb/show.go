package tvdb

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/maps"
)

// ----------------------Show------------------------
type Show struct {
	show    string
	seasons map[string]*Season
}

func (sh *Show) Value() string {
	return sh.show
}

func (sh *Show) Add(e Episode) {
	_, ok := sh.seasons[e.Season]
	if !ok {
		sh.seasons[e.Season] = &Season{episodes: make(map[string]*Episode), show: sh.show, season: e.Season}
	}
	sh.seasons[e.Season].Add(e)
}

func (sh *Show) Print() {
	for _, s := range sh.Seasons() {
		fmt.Printf("  %v\n", s.show)
		sh.seasons[s.season].Print()
	}
}

func (sh *Show) Seasons() []*Season {
	toReturn := make([]*Season, 0, len(sh.seasons))
	for _, s := range sortedKeys(maps.Keys(sh.seasons)) {
		toReturn = append(toReturn, sh.seasons[s])
	}

	return toReturn
}

func (sh *Show) Match(input string) bool {
	return strings.Contains(strings.ToLower(sh.show), strings.ToLower(input))
}

func (sh *Show) SelectSeason() (*Season, error) {
	seasons := sh.Seasons()
	prompt := promptui.Select{
		Label:     "Select season " + sh.show,
		Items:     seasons,
		Searcher:  searcher(seasons[0]),
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Error getting season %v", err)
	}
	return seasons[i], nil
}

// ----------------------Show------------------------
