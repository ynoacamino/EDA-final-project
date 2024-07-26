package reader

import (
	// invertI "eda/structures/invertIndex"

	s "eda/types"
	"fmt"
	"testing"
	"time"
)

func TestReadCSV(t *testing.T) {
	f := func(s *s.Song) {

	}

	start := time.Now()

	err := ReadCSV(f, 1000)
	if err != nil {
		t.Errorf("ReadCSV() failed : %v", err)
	}

	elapsed := time.Since(start)

	fmt.Println("Elapsed", elapsed.Microseconds(), "Microseconds")
}
