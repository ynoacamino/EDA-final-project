package reader

import (
	s "eda/types"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
)

func ReadCSV(f func(*s.Song), max int) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(pwd, "data", "data.csv")

	// path := "../data/data.csv"
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	i := 0
	for {
		record, err := reader.Read()

		if i == 0 {
			i += 1
			continue
		}
		if err == io.EOF || i == max+1 {
			break
		}
		if err != nil {
			return err
		}

		song := s.NewSong(record)
		//song := s.NewSong(record)
		// song.ToString()
		f(song)

		i += 1
	}

	return nil
}
