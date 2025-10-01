package datetime

import (
	"fmt"
	"time"
)

// Humanize returns a human-friendly relative time string (e.g., "3 days ago").
func Humanize(t, ref time.Time) string {
	delta := ref.Sub(t)
	future := delta < 0
	if future {
		delta = -delta
	}
	seconds := int(delta.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 45:
		if future {
			return "in a few seconds"
		}
		return "just now"
	case seconds < 90:
		if future {
			return "in a minute"
		}
		return "a minute ago"
	case minutes < 45:
		if future {
			return fmt.Sprintf("in %d minutes", minutes)
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	case minutes < 90:
		if future {
			return "in an hour"
		}
		return "an hour ago"
	case hours < 24:
		if future {
			return fmt.Sprintf("in %d hours", hours)
		}
		return fmt.Sprintf("%d hours ago", hours)
	case hours < 42:
		if future {
			return "in a day"
		}
		return "a day ago"
	case days < 7:
		if future {
			return fmt.Sprintf("in %d days", days)
		}
		return fmt.Sprintf("%d days ago", days)
	case weeks < 5:
		if future {
			return fmt.Sprintf("in %d weeks", weeks)
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	case months < 12:
		if future {
			return fmt.Sprintf("in %d months", months)
		}
		return fmt.Sprintf("%d months ago", months)
	default:
		if future {
			return fmt.Sprintf("in %d years", years)
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

// IsBetween checks if t is between start and end (inclusive if specified).
func IsBetween(t, start, end time.Time, inclusive bool) bool {
	if inclusive {
		return !t.Before(start) && !t.After(end)
	}
	return t.After(start) && t.Before(end)
}

// StartOfWeek returns the start of the week for t (weekStart: time.Monday, etc.).
func StartOfWeek(t time.Time, weekStart time.Weekday) time.Time {
	delta := (int(t.Weekday()) - int(weekStart) + 7) % 7
	return time.Date(t.Year(), t.Month(), t.Day()-delta, 0, 0, 0, 0, t.Location())
}

// EndOfWeek returns the end of the week for t (weekStart: time.Monday, etc.).
func EndOfWeek(t time.Time, weekStart time.Weekday) time.Time {
	start := StartOfWeek(t, weekStart)
	return start.AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// Quarter returns the quarter (1-4) for t.
func Quarter(t time.Time) int {
	return (int(t.Month())-1)/3 + 1
}

// StartOfQuarter returns the start of the quarter for t.
func StartOfQuarter(t time.Time) time.Time {
	q := Quarter(t)
	month := time.Month(3*(q-1) + 1)
	return time.Date(t.Year(), month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfQuarter returns the end of the quarter for t.
func EndOfQuarter(t time.Time) time.Time {
	return StartOfQuarter(t).AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// ISOWeek returns the ISO year and week number for t.
func ISOWeek(t time.Time) (int, int) {
	return t.ISOWeek()
}

// IsWeekend returns true if t is Saturday or Sunday.
func IsWeekend(t time.Time) bool {
	w := t.Weekday()
	return w == time.Saturday || w == time.Sunday
}

// NextWeekday returns the next occurrence of the given weekday after t.
func NextWeekday(t time.Time, weekday time.Weekday) time.Time {
	days := (int(weekday) - int(t.Weekday()) + 7) % 7
	if days == 0 {
		days = 7
	}
	return t.AddDate(0, 0, days)
}

// PreviousWeekday returns the previous occurrence of the given weekday before t.
func PreviousWeekday(t time.Time, weekday time.Weekday) time.Time {
	days := (int(t.Weekday()) - int(weekday) + 7) % 7
	if days == 0 {
		days = 7
	}
	return t.AddDate(0, 0, -days)
}

// TruncateTo truncates t to the nearest multiple of d.
func TruncateTo(t time.Time, d time.Duration) time.Time {
	return t.Truncate(d)
}

// RoundTo rounds t to the nearest multiple of d.
func RoundTo(t time.Time, d time.Duration) time.Time {
	ns := t.UnixNano()
	rem := ns % int64(d)
	if rem < int64(d)/2 {
		return t.Add(-time.Duration(rem))
	}
	return t.Add(time.Duration(int64(d) - rem))
}

// DaysInMonth returns the number of days in the month of t.
func DaysInMonth(t time.Time) int {
	first := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	next := first.AddDate(0, 1, 0)
	return int(next.Sub(first).Hours() / 24)
}

// DaysInYear returns the number of days in the year of t.
func DaysInYear(t time.Time) int {
	if IsLeapYear(t.Year()) {
		return 366
	}
	return 365
}

// IsLeapYear returns true if year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}

// ParseAny tries to parse value with any of the given layouts.
func ParseAny(layouts []string, value string) (time.Time, error) {
	var err error
	for _, layout := range layouts {
		t, e := time.Parse(layout, value)
		if e == nil {
			return t, nil
		}
		err = e
	}
	return time.Time{}, err
}

// ToUTC returns t in UTC.
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

// ToLocal returns t in local time.
func ToLocal(t time.Time) time.Time {
	return t.Local()
}

// WithLocation returns t in the given location.
func WithLocation(t time.Time, loc *time.Location) time.Time {
	return t.In(loc)
}
