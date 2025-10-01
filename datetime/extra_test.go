package datetime_test

import (
	"testing"
	"time"

	"github.com/kishankumarhs/fnkit/datetime"
)

func TestHumanize(t *testing.T) {
	now := time.Now()
	cases := []struct {
		delta  time.Duration
		expect string
	}{
		{0, "just now"},
		{time.Minute, "a minute ago"},
		{2 * time.Hour, "2 hours ago"},
		{-time.Hour, "in an hour"},
		{24 * time.Hour, "a day ago"},
		{-48 * time.Hour, "in 2 days"},
	}
	for _, c := range cases {
		got := datetime.Humanize(now.Add(-c.delta), now)
		if got != c.expect {
			t.Errorf("Humanize(%v): got %q, want %q", c.delta, got, c.expect)
		}
	}
}

func TestIsBetween(t *testing.T) {
	now := time.Now()
	start := now.Add(-time.Hour)
	end := now.Add(time.Hour)
	if !datetime.IsBetween(now, start, end, true) {
		t.Error("IsBetween inclusive failed")
	}
	if datetime.IsBetween(start, now, end, false) {
		t.Error("IsBetween exclusive failed")
	}
}

func TestWeekQuarterHelpers(t *testing.T) {
	dt := time.Date(2025, 4, 10, 12, 0, 0, 0, time.UTC) // Thursday
	startW := datetime.StartOfWeek(dt, time.Monday)
	if startW.Weekday() != time.Monday {
		t.Error("StartOfWeek not Monday")
	}
	endW := datetime.EndOfWeek(dt, time.Monday)
	if endW.Weekday() != time.Sunday {
		t.Error("EndOfWeek not Sunday")
	}
	q := datetime.Quarter(dt)
	if q != 2 {
		t.Errorf("Quarter: got %d, want 2", q)
	}
	startQ := datetime.StartOfQuarter(dt)
	if startQ.Month() != 4 {
		t.Errorf("StartOfQuarter: got %d, want 4", startQ.Month())
	}
	endQ := datetime.EndOfQuarter(dt)
	if endQ.Month() != 7 || endQ.Day() != 1 {
		t.Errorf("EndOfQuarter: got %v", endQ)
	}
}

func TestISOWeekAndWeekend(t *testing.T) {
	dt := time.Date(2025, 10, 4, 0, 0, 0, 0, time.UTC) // Saturday
	year, week := datetime.ISOWeek(dt)
	if week == 0 || year == 0 {
		t.Error("ISOWeek failed")
	}
	if !datetime.IsWeekend(dt) {
		t.Error("IsWeekend failed")
	}
}

func TestNextPreviousWeekday(t *testing.T) {
	fri := time.Date(2025, 10, 3, 0, 0, 0, 0, time.UTC) // Friday
	nextMon := datetime.NextWeekday(fri, time.Monday)
	if nextMon.Weekday() != time.Monday {
		t.Error("NextWeekday failed")
	}
	prevMon := datetime.PreviousWeekday(fri, time.Monday)
	if prevMon.Weekday() != time.Monday {
		t.Error("PreviousWeekday failed")
	}
}

func TestTruncateRoundDaysInMonthYearLeap(t *testing.T) {
	dt := time.Date(2024, 2, 29, 15, 45, 30, 0, time.UTC)
	trunc := datetime.TruncateTo(dt, time.Hour)
	if trunc.Minute() != 0 || trunc.Second() != 0 {
		t.Error("TruncateTo failed")
	}
	round := datetime.RoundTo(dt, time.Hour)
	if round.Hour() != 16 {
		t.Error("RoundTo failed")
	}
	daysMonth := datetime.DaysInMonth(dt)
	if daysMonth != 29 {
		t.Errorf("DaysInMonth: got %d, want 29", daysMonth)
	}
	daysYear := datetime.DaysInYear(dt)
	if daysYear != 366 {
		t.Errorf("DaysInYear: got %d, want 366", daysYear)
	}
	if !datetime.IsLeapYear(2024) {
		t.Error("IsLeapYear failed")
	}
}

func TestParseAnyAndTimezone(t *testing.T) {
	layouts := []string{time.RFC3339, time.RFC1123, "2006-01-02"}
	_, err := datetime.ParseAny(layouts, "2025-10-02")
	if err != nil {
		t.Error("ParseAny failed")
	}
	now := time.Now()
	if !datetime.ToUTC(now).Equal(now.UTC()) {
		t.Error("ToUTC failed")
	}
	if !datetime.ToLocal(now).Equal(now.Local()) {
		t.Error("ToLocal failed")
	}
	loc, _ := time.LoadLocation("America/New_York")
	if !datetime.WithLocation(now, loc).Equal(now.In(loc)) {
		t.Error("WithLocation failed")
	}
}
