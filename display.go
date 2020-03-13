package glock

import (
	"fmt"
	"glock/base"
)

// Display creates a string with the table
// of all times and bases
func Display(cTime CustomTime) string {

	/*
		FORMAT TIME
	*/

	// construct base16
	h, m, s := cTime.ConvertTime(base.Base16)
	format(&h, &m, &s)
	out := fmt.Sprintf(tableRow, 16, h, m, s)

	// construct base10
	h, m, s = cTime.ConvertTime(base.Base10)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 10, h, m, s)

	// construct base8
	h, m, s = cTime.ConvertTime(base.Base8)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 8, h, m, s)

	// construct base3
	h, m, s = cTime.ConvertTime(base.Base3)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 3, h, m, s)

	// construct base2
	h, m, s = cTime.ConvertTime(base.Base2)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 2, h, m, s)

	/*
		FORMAT DATE
	*/
	var d, mo, y, yD string
	// Determine what base
	switch cTime.Seconds % 5 {
	case 0:
		d, mo, y, yD = cTime.ConvertDate(base.Base16)
		format(&d, &mo, &y, &yD)
		break
	case 1:
		d, mo, y, yD = cTime.ConvertDate(base.Base10)
		format(&d, &mo, &y, &yD)
		break
	case 2:
		d, mo, y, yD = cTime.ConvertDate(base.Base8)
		format(&d, &mo, &y, &yD)
		break
	case 3:
		d, mo, y, yD = cTime.ConvertDate(base.Base3)
		format(&d, &mo, &y, &yD)
		break
	case 4:
		d, mo, y, yD = cTime.ConvertDate(base.Base2)
		format(&d, &mo, &y, &yD)
		break
	}

	// Check the day of the year
	switch yD[len(yD)-1] {
	case '1':
		yD += "st"
		break
	case '2':
		yD += "nd"
		break
	case '3':
		yD += "rd"
		break
	default:
		yD += "th"
	}

	// Add date to out
	out += fmt.Sprintf(date, d, mo, y, yD)

	return out
}

// format the time data -> add 0 if length < 10
func format(s ...*string) {
	for _, ss := range s {

		if len(*ss) == 1 {
			*ss = "0" + *ss
		}
	}
}
