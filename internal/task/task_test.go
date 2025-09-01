package task

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"vinna.app/vinna-app/internal/frequency"
)

func TestNewBaseTask_SetsCoreFields(t *testing.T) {
	userID := uuid.New()
	start := time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)
	end := time.Date(2025, 2, 2, 3, 4, 5, 0, time.UTC)
	var freq frequency.Frequency // zero value is fine if you don't want to couple to concrete values

	before := time.Now()
	bt, err := NewBaseTask(userID, "Read Go docs", start, end, freq)
	after := time.Now()

	if err != nil {
		t.Fatalf("NewBaseTask unexpected error: %v", err)
	}

	if bt == nil {
		t.Fatal("NewBaseTask returned nil")
	}
	if bt.ID == uuid.Nil {
		t.Errorf("expected non-nil UUID")
	}
	if bt.UserID != userID {
		t.Errorf("UserID mismatch: got %s want %s", bt.UserID, userID)
	}
	if bt.Name != "Read Go docs" {
		t.Errorf("Name mismatch: got %q", bt.Name)
	}
	if !bt.StartDate.Equal(start) || !bt.EndDate.Equal(end) {
		t.Errorf("start/end mismatch: got %v..%v want %v..%v", bt.StartDate, bt.EndDate, start, end)
	}
	if bt.Frequency != freq {
		t.Errorf("Frequency mismatch: got %v want %v", bt.Frequency, freq)
	}
	if bt.CreatedAt.Before(before) || bt.CreatedAt.After(after) {
		t.Errorf("CreatedAt out of expected range [%v, %v], got %v", before, after, bt.CreatedAt)
	}
	if !bt.ModifiedAt.IsZero() {
		t.Errorf("ModifiedAt should be zero value, got %v", bt.ModifiedAt)
	}
}

func TestNewBaseTask_InvalidDateRange_ReturnsError(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC) // end before start

	_, err := NewBaseTask(uuid.New(), "Backwards dates", start, end, frequency.Frequency{})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrEndBeforeStart {
		t.Fatalf("expected ErrEndBeforeStart, got %v", err)
	}
}

func TestRegularTask_New_WiresBaseTaskAndStatus(t *testing.T) {
	userID := uuid.New()
	start := time.Date(2025, 3, 10, 12, 0, 0, 0, time.UTC)
	end := time.Date(2025, 3, 11, 12, 0, 0, 0, time.UTC)
	var freq frequency.Frequency

	before := time.Now()
	rt, err := New(userID, "Ship feature", start, end, Doing, freq)
	after := time.Now()

	if err != nil {
		t.Fatalf("RegularTask.New unexpected error: %v", err)
	}
	if rt.Task.ID == uuid.Nil {
		t.Errorf("expected Task.ID to be generated")
	}
	if rt.Task.UserID != userID {
		t.Errorf("user id mismatch: got %s want %s", rt.Task.UserID, userID)
	}
	if rt.Task.Name != "Ship feature" {
		t.Errorf("name mismatch: %q", rt.Task.Name)
	}
	if !rt.Task.StartDate.Equal(start) || !rt.Task.EndDate.Equal(end) {
		t.Errorf("date mismatch")
	}
	if rt.Status != Doing {
		t.Errorf("status mismatch: got %s want %s", rt.Status, Doing)
	}
	if rt.Task.CreatedAt.Before(before) || rt.Task.CreatedAt.After(after) {
		t.Errorf("createdAt out of range")
	}
}

func TestRegularTask_New_GeneratesUniqueIDs(t *testing.T) {
	userID := uuid.New()
	var freq frequency.Frequency
	a, err := New(userID, "t1", time.Now(), time.Now(), Todo, freq)
	if err != nil {
		t.Fatalf("unexpected error creating a: %v", err)
	}
	b, err := New(userID, "t2", time.Now(), time.Now(), Todo, freq)
	if err != nil {
		t.Fatalf("unexpected error creating a: %v", err)
	}
	if a.Task.ID == b.Task.ID {
		t.Errorf("expected unique IDs; got same %s", a.Task.ID)
	}
}

func TestStatus_AllValuesAssignable(t *testing.T) {
	cases := []Status{Recorded, Todo, Doing, Done, Canceled}
	for _, s := range cases {
		rt, err := New(uuid.New(), "status check", time.Now(), time.Now(), s, frequency.Frequency{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if rt.Status != s {
			t.Fatalf("status not retained: got %s want %s", rt.Status, s)
		}
	}
}
