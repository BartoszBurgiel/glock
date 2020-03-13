package glock

import (
	"glock/base"
	"time"
)

// CustomTime holds all information
// -> hours (24h format), minutes and seconds
// as integers
type CustomTime struct {
	Hours   int
	Minutes int
	Seconds int
}

// NewGlockTime initializes a new CustomTime
// struct with the current time
func NewGlockTime() CustomTime {
	return CustomTime{
		Hours:   time.Now().Hour(),
		Minutes: time.Now().Minute(),
		Seconds: time.Now().Second(),
	}
}

// Convert CustomTime's values to
// a given base and return these as strings
func (c CustomTime) Convert(base base.Base) (h, m, s string) {
	h = base.ToBase(c.Hours, "")
	m = base.ToBase(c.Minutes, "")
	s = base.ToBase(c.Seconds, "")
	return h, m, s
}
