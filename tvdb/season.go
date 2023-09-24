package tvdb

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/maps"
)

// ----------------------Season------------------------
type Season struct {
	show     string
	season   string
	episodes map[string]*Episode
}

func (s *Season) Value() string {
	return s.season
}

func (s *Season) Add(e Episode) {
	s.episodes[e.Episode] = &e
}

func (s *Season) Print() {
	for _, episode := range s.Episodes() {
		fmt.Printf("    %v - %v\n", episode.Episode, episode.Title)
	}
}

func (s *Season) Episodes() []*Episode {
	toReturn := make([]*Episode, 0, len(s.episodes))
	for _, e := range sortedKeys(maps.Keys(s.episodes)) {
		toReturn = append(toReturn, s.episodes[e])
	}

	return toReturn
}

func (s *Season) Match(input string) bool {
	return strings.Contains(strings.ToLower(s.season), strings.ToLower(input))
}

func (s *Season) SelectEpisode() (*Episode, error) {
	episodes := s.Episodes()
	prompt := promptui.Select{
		Label:     "Select episode from " + s.show + " " + s.season,
		Items:     episodes,
		Searcher:  searcher(episodes[0]),
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Error getting episode %v", err)
	}
	return episodes[i], nil
}

// ----------------------Season------------------------
