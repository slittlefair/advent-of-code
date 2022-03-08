package timer

import (
	"fmt"
	"time"
)

// TimeTrack prints how long a function took to run, printing the time since start. Initialise start
// before the function runs and then call Track after it. e.g.
//    t := time.Now()
//    someFunction()
//    timer.Track(t)
func Track(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
