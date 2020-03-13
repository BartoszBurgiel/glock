package glock

import (
	"fmt"
	"glock/base"
)

// Display creates a string with the table
// of all times and bases
func Display(cTime CustomTime) string {

	// construct base16
	h, m, s := cTime.Convert(base.Base16)
	format(&h, &m, &s)
	out := fmt.Sprintf(tableRow, 16, h, m, s)

	// construct base10
	h, m, s = cTime.Convert(base.Base10)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 10, h, m, s)

	// construct base8
	h, m, s = cTime.Convert(base.Base8)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 8, h, m, s)

	// construct base3
	h, m, s = cTime.Convert(base.Base3)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 3, h, m, s)

	// construct base2
	h, m, s = cTime.Convert(base.Base2)
	format(&h, &m, &s)
	out += fmt.Sprintf(tableRow, 2, h, m, s)

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
