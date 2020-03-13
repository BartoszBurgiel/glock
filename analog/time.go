package watch

import "time"

// getSeconds returns current seconds
// as an int
func getSeconds() int {
	return time.Now().Second()
}

// getMinutes returns current minutes as
// an int
func getMinutes() int {
	return time.Now().Minute()
}

// getHours returns current hours as an int
func getHours() int {
	return time.Now().Hour()
}
