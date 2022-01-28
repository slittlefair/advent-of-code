package timer

import (
	"fmt"
	"time"
)

// TimeTrack prints how long a function took to run
func Track(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
