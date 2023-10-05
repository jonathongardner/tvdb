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
	if strings.Contains(strings.ToLower(s.season), input) {
		return true
	}
	for _, e := range s.episodes {
		if e.MatchTitle(input) {
			return true
		}
	}
	return false
}

func (s *Season) SelectEpisode(file string) (*Episode, error) {
	episodes := s.Episodes()
	prompt := promptui.Select{
		Label: "What episode is " + file + "?",
		Items: episodes,
		Searcher: func(input string, index int) bool {
			return episodes[index].Match(strings.ToLower(input))
		},
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Error getting episode %v", err)
	}
	return episodes[i], nil
}

// ----------------------Season------------------------
