package main

import (
	"eda/algorithms"
	ix "eda/structures/invertIndex"
	s "eda/types"
	"fmt"
	"time"

	t "eda/structures/trie"
	// import bplustree "eda/structures/bplus"
	"context"
	rd "eda/reader"
	"math/rand"

	bps "eda/structures/bplus"
)

const (
	MAXFORPAGE = 200
)

type App struct {
	InvertIndex *ix.InvertIndex[s.Song]
	Trie        *t.Trie[s.Song]

	BplusYear       *bps.Bplus[int, s.Song]
	BplusPopularity *bps.Bplus[int, s.Song]
	BplusDuration   *bps.Bplus[int, s.Song]

	LastResult s.Result

	PlayList []s.Song

	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	compareTo := func(i1, i2 int) int {
		if i1 == i2 {
			return 0
		} else if i1 > i2 {
			return 1
		} else {
			return -1
		}
	}

	a.InvertIndex = ix.NewInvertIndex[s.Song](2000)
	a.Trie = t.NewTrie[s.Song]()

	a.BplusYear = bps.NewBplus[int, s.Song](257, compareTo)
	a.BplusPopularity = bps.NewBplus[int, s.Song](257, compareTo)
	a.BplusDuration = bps.NewBplus[int, s.Song](257, compareTo)

	a.LastResult = s.Result{}
	a.PlayList = make([]s.Song, 0)

	f := func(s *s.Song) {
		name := s.TrackName
		a.InvertIndex.PutMany(name, s)
		a.Trie.Add(name, s)

		a.BplusYear.Add(s.Year, s)
		a.BplusPopularity.Add(s.Popularity, s)
		a.BplusDuration.Add(s.DurationMs, s)
	}

	err := rd.ReadCSV(f, -1)
	if err != nil {
		panic(err)
	}
}

func (a *App) SearchSongInIndexInvert(name string) s.Result {
	start := time.Now()
	lists := a.InvertIndex.Search(name)

	timeLapse := float32(time.Since(start).Microseconds()) / 1000

	size := len(lists)

	if len(lists) > 50 {
		lists = lists[:50]
	}

	result := s.Result{
		TimeLapse: timeLapse,
		Songs:     lists,
		Size:      size,
	}

	a.LastResult = result

	return result
}

func (a *App) SearchSongInTrie(name string) s.Result {
	start := time.Now()

	lists := a.Trie.Suggest(name).Parse()

	timeLapse := float32(time.Since(start).Microseconds()) / 1000

	size := len(lists)

	if len(lists) > 50 {
		lists = lists[:50]
	}

	result := s.Result{
		TimeLapse: timeLapse,
		Songs:     lists,
		Size:      size,
	}

	a.LastResult = result

	return result
}

func (a *App) OrderByYear(o int, page int) s.Result {
	start := time.Now()
	res := a.BplusYear.GetArray()
	timeLapse := float32(time.Since(start).Microseconds()) / 1000
	size := len(res)

	if o == -1 {
		res = a.Invert(res)
	}

	res = getPage(res, page)

	result := s.Result{
		TimeLapse: timeLapse,
		Songs:     res,
		Size:      size,
	}

	a.LastResult = result

	return result
}

func (a *App) OrderByPopularity(o int, page int) s.Result {
	start := time.Now()
	res := a.BplusPopularity.GetArray()
	timeLapse := float32(time.Since(start).Microseconds()) / 1000
	size := len(res)

	if o == -1 {
		res = a.Invert(res)
	}

	res = getPage(res, page)

	result := s.Result{
		TimeLapse: timeLapse,
		Songs:     res,
		Size:      size,
	}

	a.LastResult = result

	return result
}

func (a *App) OrderByDuration(o int, page int) s.Result {
	start := time.Now()
	res := a.BplusDuration.GetArray()
	timeLapse := float32(time.Since(start).Microseconds()) / 1000
	size := len(res)

	if o == -1 {
		res = a.Invert(res)
	}

	res = getPage(res, page)

	result := s.Result{
		TimeLapse: timeLapse,
		Songs:     res,
		Size:      size,
	}

	a.LastResult = result

	return result
}

func (a *App) Invert(arr []s.Song) []s.Song {
	l := len(arr)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getPage(slice []s.Song, pageNumber int) []s.Song {
	start := pageNumber * MAXFORPAGE
	end := start + MAXFORPAGE

	// Manejo de los casos donde el rango excede los límites del slice
	if start > len(slice) {
		return []s.Song{}
	}
	if end > len(slice) {
		end = len(slice)
	}

	return slice[start:end]
}

func (a *App) PutSong(index int) {
	a.PlayList = append(a.PlayList, a.LastResult.Songs[index])
}

func (a *App) GetPlayList() []s.Song {
	return a.PlayList
}

func (a *App) RandomPlayList() []s.Song {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("LEN", len(a.PlayList))
	n := len(a.PlayList)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		a.PlayList[i], a.PlayList[j] = a.PlayList[j], a.PlayList[i]
	}
	return a.PlayList
}

func (a *App) RemoveSong(index int) {
	a.PlayList = append(a.PlayList[:index], a.PlayList[index+1:]...)
}

func (a *App) ClearPlayList() {
	a.PlayList = make([]s.Song, 0)
}

func (a *App) OrderPlayList(filter int, order int) {
	if filter == 0 || order == 0 {
		return
	}

	f := func(s1, s2 s.Song) int {
		if s1.Year == s2.Year {
			return 0
		} else if s1.Year > s2.Year {
			return 1
		} else {
			return -1
		}
	}

	if filter == 2 {
		f = func(s1, s2 s.Song) int {
			if s1.DurationMs == s2.DurationMs {
				return 0
			} else if s1.DurationMs > s2.DurationMs {
				return 1
			} else {
				return -1
			}
		}
	} else if filter == 3 {
		f = func(s1, s2 s.Song) int {
			if s1.Popularity == s2.Popularity {
				return 0
			} else if s1.Popularity > s2.Popularity {
				return 1
			} else {
				return -1
			}
		}
	}

	ordered := algorithms.QuickSortStart(a.PlayList, f)

	if order == -1 {
		ordered = a.Invert(ordered)
	}

	a.PlayList = ordered
}
