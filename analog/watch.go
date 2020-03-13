package watch

import (
	"fmt"
	"strings"
)

// Clock creates a clock face with its arm
// pointing to the value as it was a clock
func Clock(curr int, lim int) string {
	// check
	if curr > lim {
		fmt.Println("watch: WatchFace: curr[ent] cannot be bigger than lim[it],", curr, "is bigger than", lim)
	}

	// calculate the angle between curr and lim on the face
	angle := getDegree(curr, lim)

	// get according face from angle
	return angleToFace(angle)
}

// ClocksWithCurrentTime returns a string
// with 3 clocks
// each one representing a time value ->
// hours, minutes and seconds
func ClocksWithCurrentTime() (out string) {

	// create clocks
	hoursClock := Clock(getHours()/2, 12)
	minutesClock := Clock(getMinutes(), 60)
	secondsClock := Clock(getSeconds(), 60)

	// divide each to slices of strings
	hoursClockDiv := strings.Split(hoursClock, "\n")
	minutesClockDiv := strings.Split(minutesClock, "\n")
	secondsClockDiv := strings.Split(secondsClock, "\n")

	// fuse clocks together
	for i := 0; i < len(hoursClockDiv); i++ {
		out += hoursClockDiv[i] + " " + minutesClockDiv[i] + " " + secondsClockDiv[i] + "\n"
	}

	// add the footer
	out += "  hours  |"
	out += " minutes |"
	out += " seconds "

	return out
}
