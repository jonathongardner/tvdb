// Go program to sort the map by Keys
package tvdb

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type Files struct {
	show   string
	season string
}

func checkoutput(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Directory %v not found", path)
	}

	if !info.IsDir() {
		return fmt.Errorf("Path %v is not a dir", path)
	}

	return nil
}

func getfiles(path string) ([]string, error) {
	toReturn := make([]string, 0)
	fileInfo, err := os.Stat(path)
	if err != nil {
		return toReturn, fmt.Errorf("%v not found", path)
	}

	if fileInfo.IsDir() {
		err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				toReturn = append(toReturn, p)
			}
			return nil
		})
		if err != nil {
			return toReturn, fmt.Errorf("Error reading input directory %v", path)
		}
	} else {
		toReturn = append(toReturn, path)
	}
	return toReturn, nil
}

func movefile(destFolder string, source string, episode *Episode) error {
	destDir := filepath.Join(destFolder, episode.Dir())

	err := os.MkdirAll(destDir, 0777)
	if err != nil {
		return err
	}

	destination := filepath.Join(destDir, episode.Filename())
	// err = os.Rename(source, destination)
	// if err != nil {
	// 	return err
	// }

	log.Infof("Moved %v to %v\n", source, destination)
	return nil
}

func MoveFiles(output string, input string) error {
	err := checkoutput(output)
	if err != nil {
		return err
	}

	files, err := getfiles(input)
	if err != nil {
		return err
	}

	db, err := newDB()
	if err != nil {
		return err
	}

	//----------Show--------
	show, err := db.SelectShow()
	if err != nil {
		return err
	}
	//----------Show--------

	//----------Season--------
	season, err := show.SelectSeason()
	if err != nil {
		return err
	}

	//----------Episodes--------
	for _, file := range files {
		episode, err := season.SelectEpisode()
		if err != nil {
			return err
		}
		movefile(output, file, episode)
	}
	return nil
}
