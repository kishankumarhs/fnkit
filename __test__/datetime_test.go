package datetime_test

import (
	"testing"
	"time"

	"github.com/kishankumarhs/fnkit/datetime"
)

func TestParseAndFormat(t *testing.T) {
	layout := time.DateOnly
	str := "2025-10-02"
	tm, err := datetime.Parse(layout, str)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if got := datetime.Format(tm, layout); got != str {
		t.Errorf("Format: got %q, want %q", got, str)
	}
}

func TestMustParse(t *testing.T) {
	layout := time.DateOnly
	str := "2025-10-02"
	tm := datetime.MustParse(layout, str)
	if tm.Format(layout) != str {
		t.Errorf("MustParse: got %q, want %q", tm.Format(layout), str)
	}
}

func TestAddSubtractDays(t *testing.T) {
	base := time.Date(2025, 10, 2, 0, 0, 0, 0, time.UTC)
	if got := datetime.AddDays(base, 5); !got.Equal(base.AddDate(0, 0, 5)) {
		t.Errorf("AddDays failed")
	}
	if got := datetime.SubtractDays(base, 2); !got.Equal(base.AddDate(0, 0, -2)) {
		t.Errorf("SubtractDays failed")
	}
}

func TestAddSubtractMonthsYears(t *testing.T) {
	base := time.Date(2025, 10, 2, 0, 0, 0, 0, time.UTC)
	if got := datetime.AddMonths(base, 2); !got.Equal(base.AddDate(0, 2, 0)) {
		t.Errorf("AddMonths failed")
	}
	if got := datetime.SubtractMonths(base, 1); !got.Equal(base.AddDate(0, -1, 0)) {
		t.Errorf("SubtractMonths failed")
	}
	if got := datetime.AddYears(base, 3); !got.Equal(base.AddDate(3, 0, 0)) {
		t.Errorf("AddYears failed")
	}
	if got := datetime.SubtractYears(base, 2); !got.Equal(base.AddDate(-2, 0, 0)) {
		t.Errorf("SubtractYears failed")
	}
}

func TestBusinessDays(t *testing.T) {
	fri := time.Date(2025, 10, 3, 0, 0, 0, 0, time.UTC) // Friday
	sat := time.Date(2025, 10, 4, 0, 0, 0, 0, time.UTC) // Saturday
	mon := time.Date(2025, 10, 6, 0, 0, 0, 0, time.UTC) // Monday
	if !datetime.IsBusinessDay(fri) {
		t.Error("Friday should be business day")
	}
	if datetime.IsBusinessDay(sat) {
		t.Error("Saturday should not be business day")
	}
	if got := datetime.AddBusinessDays(fri, 1); !got.Equal(mon) {
		t.Errorf("AddBusinessDays failed: got %v, want %v", got, mon)
	}
	if got := datetime.SubtractBusinessDays(mon, 1); !got.Equal(fri) {
		t.Errorf("SubtractBusinessDays failed: got %v, want %v", got, fri)
	}
}
