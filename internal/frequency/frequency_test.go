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

func TestNew_Success_CommaListsAndEdges(t *testing.T) {
	tests := []struct {
		name       string
		minute     string
		hour       string
		dayOfMonth string
		month      string
		dayOfWeek  string
	}{
		{
			name:       "typical lists",
			minute:     "0,15,30,45",
			hour:       "8,12,16",
			dayOfMonth: "1,15,31",
			month:      "1,6,12",
			dayOfWeek:  "1,3,5,7",
		},
		{
			name:       "single edge values",
			minute:     "0",
			hour:       "0",
			dayOfMonth: "1",
			month:      "1",
			dayOfWeek:  "7",
		},
		{
			name:       "upper edges",
			minute:     "59",
			hour:       "23",
			dayOfMonth: "31",
			month:      "12",
			dayOfWeek:  "6",
		},
		{
			name:       "duplicates allowed (pattern permits)",
			minute:     "5,5",
			hour:       "2,2,2",
			dayOfMonth: "10,10",
			month:      "3,3",
			dayOfWeek:  "2,2",
		},
	}

	for _, tt := range tests {
		tt := tt
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
		"",     // empty
		"60",   // out of range
		"-1",   // negative
		"01",   // leading zero (pattern forbids)
		"1,",   // trailing comma
		",1",   // leading comma
		"1,,2", // empty element
		"*,1",  // mixing star with list not allowed by pattern
		"a",    // non-digit
		"1, 2", // spaces not allowed
	}
	for _, c := range cases {
		c := c
		t.Run("minute="+c, func(t *testing.T) {
			_, err := New(c, "0", "1", "1", "1")
			if err == nil {
				t.Fatalf("expected error for minute=%q", c)
			}
			if !strings.Contains(err.Error(), "invalid minute format: "+c) {
				t.Fatalf("error message mismatch: %v", err)
			}
		})
	}
}

func TestNew_Invalid_Hour(t *testing.T) {
	cases := []string{
		"", "24", "-1", "01", "1,", ",1", "1,,2", "*,1", "a", "1, 2",
	}
	for _, c := range cases {
		c := c
		t.Run("hour="+c, func(t *testing.T) {
			_, err := New("0", c, "1", "1", "1")
			if err == nil {
				t.Fatalf("expected error for hour=%q", c)
			}
			if !strings.Contains(err.Error(), "invalid hour format: "+c) {
				t.Fatalf("error message mismatch: %v", err)
			}
		})
	}
}

func TestNew_Invalid_DayOfMonth(t *testing.T) {
	cases := []string{
		"", "0", "32", "-1", "01", "1,", ",1", "1,,2", "*,1", "a", "1, 2",
	}
	for _, c := range cases {
		c := c
		t.Run("dom="+c, func(t *testing.T) {
			_, err := New("0", "0", c, "1", "1")
			if err == nil {
				t.Fatalf("expected error for day_of_month=%q", c)
			}
			if !strings.Contains(err.Error(), "invalid day_of_month format: "+c) {
				t.Fatalf("error message mismatch: %v", err)
			}
		})
	}
}

func TestNew_Invalid_Month(t *testing.T) {
	cases := []string{
		"", "0", "13", "-1", "01", "1,", ",1", "1,,2", "*,1", "a", "1, 2",
	}
	for _, c := range cases {
		c := c
		t.Run("month="+c, func(t *testing.T) {
			_, err := New("0", "0", "1", c, "1")
			if err == nil {
				t.Fatalf("expected error for month=%q", c)
			}
			if !strings.Contains(err.Error(), "invalid month format: "+c) {
				t.Fatalf("error message mismatch: %v", err)
			}
		})
	}
}

func TestNew_Invalid_DayOfWeek(t *testing.T) {
	cases := []string{
		"", "0", "8", "1,8", "0,7", "-1", "01", "*,1", "a", "1, 2",
	}
	for _, c := range cases {
		c := c
		t.Run("dow="+c, func(t *testing.T) {
			_, err := New("0", "0", "1", "1", c)
			if err == nil {
				t.Fatalf("expected error for day_of_week=%q", c)
			}
			if !strings.Contains(err.Error(), "invalid day_of_week format: "+c) {
				t.Fatalf("error message mismatch: %v", err)
			}
		})
	}
}

func TestNew_ValidationOrder_MinuteCheckedFirst(t *testing.T) {
	// Even if multiple fields are bad, only the first invalid (minute) should be reported.
	_, err := New("60", "99", "0", "13", "9")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "invalid minute format: 60") {
		t.Fatalf("expected minute error first, got: %v", err)
	}
}

func TestNew_ValidationOrder_HourThenDOMThenMonthThenDOW(t *testing.T) {
	// Good minute; bad hour should surface
	_, err := New("0", "99", "0", "13", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid hour format: 99") {
		t.Fatalf("expected hour error, got %v", err)
	}
	// Good minute & hour; bad day_of_month next
	_, err = New("0", "0", "0", "13", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid day_of_month format: 0") {
		t.Fatalf("expected day_of_month error, got %v", err)
	}
	// Then month
	_, err = New("0", "0", "1", "13", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid month format: 13") {
		t.Fatalf("expected month error, got %v", err)
	}
	// Then day_of_week
	_, err = New("0", "0", "1", "1", "9")
	if err == nil || !strings.Contains(err.Error(), "invalid day_of_week format: 9") {
		t.Fatalf("expected day_of_week error, got %v", err)
	}
}

func TestNew_Invalid_LeadingZeros(t *testing.T) {
	t.Parallel()

	type tc struct {
		name       string
		minute     string
		hour       string
		dayOfMonth string
		month      string
		dayOfWeek  string
		expect     string
	}
	cases := []tc{
		// minute 00–09 rejected
		{"minute 01", "01", "0", "1", "1", "1", "invalid minute format: 01"},
		{"minute 09", "09", "0", "1", "1", "1", "invalid minute format: 09"},
		// hour 00–09 rejected
		{"hour 01", "0", "01", "1", "1", "1", "invalid hour format: 01"},
		{"hour 09", "0", "09", "1", "1", "1", "invalid hour format: 09"},
		// dayOfMonth 01–09 rejected
		{"dom 01", "0", "0", "01", "1", "1", "invalid day_of_month format: 01"},
		{"dom 09", "0", "0", "09", "1", "1", "invalid day_of_month format: 09"},
		// month 01–09 rejected
		{"month 01", "0", "0", "1", "01", "1", "invalid month format: 01"},
		{"month 09", "0", "0", "1", "09", "1", "invalid month format: 09"},
		// dayOfWeek 01–07 rejected (pattern expects 1–7 without leading zero)
		{"dow 01", "0", "0", "1", "1", "01", "invalid day_of_week format: 01"},
		{"dow 07", "0", "0", "1", "1", "07", "invalid day_of_week format: 07"},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.minute, tt.hour, tt.dayOfMonth, tt.month, tt.dayOfWeek)
			if err == nil || !strings.Contains(err.Error(), tt.expect) {
				t.Fatalf("got err=%v; want contains %q", err, tt.expect)
			}
		})
	}
}
