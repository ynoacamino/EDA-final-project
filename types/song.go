package types

import (
	"fmt"
	"strconv"
)

type Song struct {
	TrackId         string
	ArtistName      string
	TrackName       string
	Popularity      int
	Year            int
	Genre           string
	Danceability    float64
	Energy          float64
	Key             int
	Loudness        float64
	Mode            int
	Speechiness     float64
	Acousticness    float64
	Intrumentalness float64
	Liveness        float64
	Valence         float64
	Tempo           float64
	DurationMs      int
	TimeSignature   int
}

func NewSong(record []string) *Song {
	popularity, err := strconv.Atoi(record[4])
	if err != nil {
		panic(err)
	}

	year, err := strconv.Atoi(record[5])
	if err != nil {
		panic(err)
	}

	danceability, err := strconv.ParseFloat(record[7], 64)
	if err != nil {
		panic(err)
	}

	energy, err := strconv.ParseFloat(record[8], 64)
	if err != nil {
		panic(err)
	}

	key, err := strconv.Atoi(record[9])
	if err != nil {
		panic(err)
	}

	loudness, err := strconv.ParseFloat(record[10], 64)
	if err != nil {
		panic(err)
	}

	mode, err := strconv.Atoi(record[11])
	if err != nil {
		panic(err)
	}

	speechiness, err := strconv.ParseFloat(record[12], 64)
	if err != nil {
		panic(err)
	}

	acousticness, err := strconv.ParseFloat(record[13], 64)
	if err != nil {
		panic(err)
	}

	instrumentalness, err := strconv.ParseFloat(record[14], 64)
	if err != nil {
		panic(err)
	}

	liveness, err := strconv.ParseFloat(record[15], 64)
	if err != nil {
		panic(err)
	}

	valence, err := strconv.ParseFloat(record[16], 64)
	if err != nil {
		panic(err)
	}

	tempo, err := strconv.ParseFloat(record[17], 64)
	if err != nil {
		panic(err)
	}

	durationMs, err := strconv.Atoi(record[18])
	if err != nil {
		panic(err)
	}

	timeSignature, err := strconv.Atoi(record[19])
	if err != nil {
		panic(err)
	}

	return &Song{
		TrackId:         record[0],
		ArtistName:      record[1],
		TrackName:       record[2],
		Popularity:      popularity,
		Year:            year,
		Genre:           record[6],
		Danceability:    danceability,
		Key:             key,
		Energy:          energy,
		Loudness:        loudness,
		Mode:            mode,
		Speechiness:     speechiness,
		Acousticness:    acousticness,
		Intrumentalness: instrumentalness,
		Liveness:        liveness,
		Valence:         valence,
		Tempo:           tempo,
		DurationMs:      durationMs,
		TimeSignature:   timeSignature,
	}
}

func (song *Song) Print() {
	fmt.Println("Artist Name: ", song.ArtistName)
	fmt.Println("Track Name: ", song.TrackName)
	fmt.Println("Track Id: ", song.TrackId)
	fmt.Println("Popularity: ", song.Popularity)
	fmt.Println("Year: ", song.Year)
	fmt.Println("Genre: ", song.Genre)
	fmt.Println("Danceability: ", song.Danceability)
	fmt.Println("Energy: ", song.Energy)
	fmt.Println("Key: ", song.Key)
	fmt.Println("Loudness: ", song.Loudness)
	fmt.Println("Mode: ", song.Mode)
	fmt.Println("Speechiness: ", song.Speechiness)
	fmt.Println("Acousticness: ", song.Acousticness)
	fmt.Println("Instrumentalness: ", song.Intrumentalness)
	fmt.Println("Liveness: ", song.Liveness)
	fmt.Println("Valence: ", song.Valence)
	fmt.Println("Tempo: ", song.Tempo)
	fmt.Println("DurationMs: ", song.DurationMs)
	fmt.Println("Time Signature: ", song.TimeSignature)

	fmt.Println("=====================================")
}
