package frequency

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestNew_Success_AllWildcards(t *testing.T) {
	f, err := New("*", "*", "*", "*", "*")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("nil Frequency on success")
	}
	if f.ID == uuid.Nil {
		t.Errorf("expected non-nil UUID")
	}
	if f.Minute != "*" || f.Hour != "*" || f.DayOfMonth != "*" || f.Month != "*" || f.DayOfWeek != "*" {
		t.Errorf("values not preserved: %+v", f)
	}
}

func TestNew_Success_EdgeValues(t *testing.T) {
	tests := []struct {
		name       string
		minute     string
		hour       string
		dayOfMonth string
		month      string
		dayOfWeek  string
	}{
		{
			name:       "lower edges",
			minute:     "0",
			hour:       "0",
			dayOfMonth: "1",
			month:      "1",
			dayOfWeek:  "0", // Sunday
		},
		{
			name:       "upper edges",
			minute:     "59",
			hour:       "23",
			dayOfMonth: "31",
			month:      "12",
			dayOfWeek:  "6", // Saturday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := New(tt.minute, tt.hour, tt.dayOfMonth, tt.month, tt.dayOfWeek)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if f.Minute != tt.minute || f.Hour != tt.hour || f.DayOfMonth != tt.dayOfMonth || f.Month != tt.month || f.DayOfWeek != tt.dayOfWeek {
				t.Fatalf("values not preserved: %+v", f)
			}
			if f.ID == uuid.Nil {
				t.Errorf("expected non-nil UUID")
			}
		})
	}
}

func TestNew_Invalid_Minute(t *testing.T) {
	cases := []string{
		"", "60", "-1", "01", "5,10", "1-5", "*/5", " 5", "5 ", "a",
	}
	for _, c := range cases {
		t.Run("minute="+c, func(t *testing.T) {
			_, err := New(c, "0", "1", "1", "0")
			if err == nil {
				t.Fatalf("expected error for minute=%q", c)
			}
			want := "invalid minute format: " + c
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q; want contains %q", err.Error(), want)
			}
		})
	}
}

func TestNew_Invalid_Hour(t *testing.T) {
	cases := []string{
		"", "24", "-1", "09", "1,2", "1-2", "*/2", " 1", "1 ", "a",
	}
	for _, c := range cases {
		t.Run("hour="+c, func(t *testing.T) {
			_, err := New("0", c, "1", "1", "0")
			if err == nil {
				t.Fatalf("expected error for hour=%q", c)
			}
			want := "invalid hour format: " + c
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q; want contains %q", err.Error(), want)
			}
		})
	}
}

func TestNew_Invalid_DayOfMonth(t *testing.T) {
	cases := []string{
		"", "0", "32", "-1", "01", "1,2", "1-2", "*/2", " a", "a ",
	}
	for _, c := range cases {
		t.Run("dom="+c, func(t *testing.T) {
			_, err := New("0", "0", c, "1", "0")
			if err == nil {
				t.Fatalf("expected error for day_of_month=%q", c)
			}
			want := "invalid day_of_month format: " + c
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q; want contains %q", err.Error(), want)
			}
		})
	}
}

func TestNew_Invalid_Month(t *testing.T) {
	cases := []string{
		"", "0", "13", "-1", "01", "1,2", "1-2", "*/2", " a", "a ",
	}
	for _, c := range cases {
		t.Run("month="+c, func(t *testing.T) {
			_, err := New("0", "0", "1", c, "0")
			if err == nil {
				t.Fatalf("expected error for month=%q", c)
			}
			want := "invalid month format: " + c
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q; want contains %q", err.Error(), want)
			}
		})
	}
}

func TestNew_Invalid_DayOfWeek(t *testing.T) {
	cases := []string{
		"", "7", "-1", "01", "1,2", "1-2", "*/2", " a", "a ",
	}
	for _, c := range cases {
		t.Run("dow="+c, func(t *testing.T) {
			_, err := New("0", "0", "1", "1", c)
			if err == nil {
				t.Fatalf("expected error for day_of_week=%q", c)
			}
			want := "invalid day_of_week format: " + c
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q; want contains %q", err.Error(), want)
			}
		})
	}
}

func TestNew_ValidationOrder(t *testing.T) {
	// First failure should be "minute"
	_, err := New("60", "99", "0", "13", "9")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "invalid minute format: 60") {
		t.Fatalf("want minute error first, got %v", err)
	}

	// Then hour (with minute valid)
	_, err = New("0", "24", "0", "13", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid hour format: 24") {
		t.Fatalf("want hour error, got %v", err)
	}

	// Then day_of_month
	_, err = New("0", "0", "0", "13", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid day_of_month format: 0") {
		t.Fatalf("want day_of_month error, got %v", err)
	}

	// Then month
	_, err = New("0", "0", "1", "13", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid month format: 13") {
		t.Fatalf("want month error, got %v", err)
	}

	// Then day_of_week
	_, err = New("0", "0", "1", "1", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid day_of_week format: 9") {
		t.Fatalf("want day_of_week error, got %v", err)
	}
}

func TestNew_LeadingZeros_Invalid(t *testing.T) {
	type tc struct {
		name, minute, hour, dom, month, dow, want string
	}
	cases := []tc{
		{"minute 01", "01", "0", "1", "1", "0", "invalid minute format: 01"},
		{"hour 09", "0", "09", "1", "1", "0", "invalid hour format: 09"},
		{"dom 01", "0", "0", "01", "1", "0", "invalid day_of_month format: 01"},
		{"month 09", "0", "0", "1", "09", "0", "invalid month format: 09"},
		{"dow 01", "0", "0", "1", "1", "01", "invalid day_of_week format: 01"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.minute, tt.hour, tt.dom, tt.month, tt.dow)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("got %v; want contains %q", err, tt.want)
			}
		})
	}
}

func TestNew_Invalid_DateCombinations(t *testing.T) {
	cases := []struct {
		name, day, month string
	}{
		{"Feb 30", "30", "2"},
		{"Feb 31", "31", "2"},
		{"Apr 31", "31", "4"},
		{"Jun 31", "31", "6"},
		{"Sep 31", "31", "9"},
		{"Nov 31", "31", "11"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := New("0", "0", c.day, c.month, "0")
			if err == nil || !strings.Contains(err.Error(), "invalid date combination") {
				t.Fatalf("expected invalid date combination error, got %v", err)
			}
		})
	}
}

func TestNew_Valid_Feb29(t *testing.T) {
	// Allowed because we assume leap year context (2024)
	_, err := New("0", "0", "29", "2", "0")
	if err != nil {
		t.Fatalf("expected Feb 29 to be valid in leap year context, got %v", err)
	}
}
