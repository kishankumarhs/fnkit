package datetime

import (
	"time"
)

// StartOf returns the time at the start of the given unit ("year", "month", "day", "hour", "minute", "second").
func StartOf(t time.Time, unit string) time.Time {
	switch unit {
	case "year":
		return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	case "month":
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	case "day":
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	case "hour":
		return t.Truncate(time.Hour)
	case "minute":
		return t.Truncate(time.Minute)
	case "second":
		return t.Truncate(time.Second)
	default:
		return t
	}
}

// EndOf returns the time at the end of the given unit ("year", "month", "day", "hour", "minute", "second").
func EndOf(t time.Time, unit string) time.Time {
	switch unit {
	case "year":
		return time.Date(t.Year()+1, 1, 1, 0, 0, 0, -1, t.Location())
	case "month":
		return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, -1, t.Location())
	case "day":
		return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, -1, t.Location())
	case "hour":
		return t.Truncate(time.Hour).Add(time.Hour - time.Nanosecond)
	case "minute":
		return t.Truncate(time.Minute).Add(time.Minute - time.Nanosecond)
	case "second":
		return t.Truncate(time.Second).Add(time.Second - time.Nanosecond)
	default:
		return t
	}
}

// IsBefore returns true if t is before u.
func IsBefore(t, u time.Time) bool {
	return t.Before(u)
}

// IsAfter returns true if t is after u.
func IsAfter(t, u time.Time) bool {
	return t.After(u)
}

// IsSame returns true if t and u are the same up to the given unit ("year", "month", "day", etc.).
func IsSame(t, u time.Time, unit string) bool {
	return StartOf(t, unit).Equal(StartOf(u, unit))
}

// Set sets a field ("year", "month", "date", "hour", "minute", "second") to a value and returns the new time.
func Set(t time.Time, field string, value int) time.Time {
	y, m, d := t.Date()
	h, min, s := t.Clock()
	ns := t.Nanosecond()
	loc := t.Location()
	switch field {
	case "year":
		y = value
	case "month":
		m = time.Month(value)
	case "date", "day":
		d = value
	case "hour":
		h = value
	case "minute":
		min = value
	case "second":
		s = value
	}
	return time.Date(y, m, d, h, min, s, ns, loc)
}

// Get returns the value of a field ("year", "month", "date", "hour", "minute", "second").
func Get(t time.Time, field string) int {
	switch field {
	case "year":
		return t.Year()
	case "month":
		return int(t.Month())
	case "date", "day":
		return t.Day()
	case "hour":
		return t.Hour()
	case "minute":
		return t.Minute()
	case "second":
		return t.Second()
	default:
		return 0
	}
}
