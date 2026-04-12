package dateformat

import "time"

func Relative(val time.Time) string {
	now := time.Now()
	if val.Year() != now.Year() {
		return val.Format("02 Jan 2006 at 04:15")
	}
	if val.Month() != now.Month() {
		return val.Format("02 Jan at 04:15")
	}
	if val.Day() == now.Day() {
		return val.Format("Today at 04:15")
	}
	if val.AddDate(0, 0, 1).Day() == now.Day() {
		return val.Format("Yesterday at 04:15")
	}
	return val.Format("02 Jan at 04:15")
}
