package timer

import (
	"fmt"
	"strings"
	"time"
)

// TimeTrack prints how long a function took to run, printing the time since start. Initialise start
// before the function runs and then call Track after it. e.g.
//
//	t := time.Now()
//	someFunction()
//	timer.Track(t)
func Track(start time.Time, messages ...string) {
	elapsed := fmt.Sprintf("took %s\n", time.Since(start))
	if len(messages) == 0 {
		fmt.Print(elapsed)
		return
	}
	fmt.Printf("%s: %s", strings.Join(messages, ": "), elapsed)
}
