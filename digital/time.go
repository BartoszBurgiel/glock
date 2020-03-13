package glock

import (
	"glock/base"
	"time"
)

// CustomTime holds all information
// -> hours (24h format), minutes and seconds
// as integers
// as well as date -> days, months and year
type CustomTime struct {
	Hours        int
	Minutes      int
	Seconds      int
	Days         int
	Months       int
	Year         int
	DayOfTheYear int
}

// NewGlockTime initializes a new CustomTime
// struct with the current time
func NewGlockTime() CustomTime {
	t := time.Now()
	return CustomTime{
		Hours:        t.Hour(),
		Minutes:      t.Minute(),
		Seconds:      t.Second(),
		Days:         t.Day(),
		Months:       int(t.Month()),
		Year:         t.Year(),
		DayOfTheYear: t.YearDay(),
	}
}

// ConvertTime values to
// a given base and return these as strings
func (c CustomTime) ConvertTime(base base.Base) (h, m, s string) {
	h = base.ToBase(c.Hours, "")
	m = base.ToBase(c.Minutes, "")
	s = base.ToBase(c.Seconds, "")

	return h, m, s
}

// ConvertDate values to the given
// base and return these as strings
func (c CustomTime) ConvertDate(base base.Base) (d, mo, y, dY string) {
	d = base.ToBase(c.Days, "")
	mo = base.ToBase(c.Months, "")
	y = base.ToBase(c.Year, "")
	dY = base.ToBase(c.DayOfTheYear, "")
	return d, mo, y, dY
}
