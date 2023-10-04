package tvdb

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/maps"
)

//go:embed tvdb.json
var dbBytes []byte

func PrintDB() error {
	db, err := newDB()
	if err != nil {
		return err
	}
	db.Print()
	return nil
}

// ----------------------DB------------------------
type DB struct {
	shows map[string]*Show
}

func (db *DB) Add(e Episode) {
	_, ok := db.shows[e.Show]
	if !ok {
		db.shows[e.Show] = &Show{seasons: make(map[string]*Season), show: e.Show}
	}
	db.shows[e.Show].Add(e)
}

func (db *DB) Print() {
	for _, s := range db.Shows() {
		fmt.Printf("%v\n", s)
		db.shows[s.show].Print()
	}
}

func (db *DB) Shows() []*Show {
	toReturn := make([]*Show, 0, len(db.shows))
	for _, s := range sortedKeys(maps.Keys(db.shows)) {
		toReturn = append(toReturn, db.shows[s])
	}

	return toReturn
}

func (db *DB) SelectShow() (*Show, error) {
	shows := db.Shows()
	prompt := promptui.Select{
		Label: "Select show",
		Items: shows,
		Searcher: func(input string, index int) bool {
			return shows[index].Match(strings.ToLower(input))
		},
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Error getting show %v", err)
	}
	return shows[i], nil
}

func newDB() (*DB, error) {
	db := &DB{shows: make(map[string]*Show)}

	episodes := make([]Episode, 0)
	json.Unmarshal(dbBytes, &episodes)

	for _, e := range episodes {
		db.Add(e)
	}

	return db, nil
}

// ----------------------DB------------------------

// func (e Episode) PlexName() string {

// }
// func (e Episode) Key() string {
// 	return e.Show + "-" + e.Season + "-" + e.Episode
// }
