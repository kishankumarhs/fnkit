package datetime

import (
	"time"
)

// Parse parses a date string with the given layout (like time.Parse, but friendlier names).
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// MustParse parses a date string or panics if invalid.
func MustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

// Format formats a time.Time with the given layout.
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// AddDays returns t + n days.
func AddDays(t time.Time, n int) time.Time {
	return t.AddDate(0, 0, n)
}

// AddMonths returns t + n months.
func AddMonths(t time.Time, n int) time.Time {
	return t.AddDate(0, n, 0)
}

// AddYears returns t + n years.
func AddYears(t time.Time, n int) time.Time {
	return t.AddDate(n, 0, 0)
}

// SubtractDays returns t - n days.
func SubtractDays(t time.Time, n int) time.Time {
	return t.AddDate(0, 0, -n)
}

// SubtractMonths returns t - n months.
func SubtractMonths(t time.Time, n int) time.Time {
	return t.AddDate(0, -n, 0)
}

// SubtractYears returns t - n years.
func SubtractYears(t time.Time, n int) time.Time {
	return t.AddDate(-n, 0, 0)
}

// IsBusinessDay returns true if t is a weekday (Mon-Fri).
func IsBusinessDay(t time.Time) bool {
	w := t.Weekday()
	return w >= time.Monday && w <= time.Friday
}

// AddBusinessDays adds n business days to t (skips weekends).
func AddBusinessDays(t time.Time, n int) time.Time {
	for n > 0 {
		t = t.AddDate(0, 0, 1)
		if IsBusinessDay(t) {
			n--
		}
	}
	return t
}

// SubtractBusinessDays subtracts n business days from t (skips weekends).
func SubtractBusinessDays(t time.Time, n int) time.Time {
	for n > 0 {
		t = t.AddDate(0, 0, -1)
		if IsBusinessDay(t) {
			n--
		}
	}
	return t
}
