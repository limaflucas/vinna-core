package habit

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"vinna.app/vinna-app/internal/frequency"
	"vinna.app/vinna-app/internal/task"
)

func TestNew_SetsCoreFields(t *testing.T) {
	userID := uuid.New()
	start := time.Date(2025, 4, 1, 8, 0, 0, 0, time.UTC)
	end := time.Date(2025, 5, 1, 8, 0, 0, 0, time.UTC)
	var freq frequency.Frequency

	before := time.Now()
	h, err := New(userID, "Morning run", start, end, Active, freq)
	after := time.Now()

	if err != nil {
		t.Fatalf("New unexpected error: %v", err)
	}
	if h.Task.ID == uuid.Nil {
		t.Errorf("expected Task.ID to be generated")
	}
	if h.Task.UserID != userID {
		t.Errorf("user id mismatch: got %s want %s", h.Task.UserID, userID)
	}
	if h.Task.Name != "Morning run" {
		t.Errorf("name mismatch: %q", h.Task.Name)
	}
	if !h.Task.StartDate.Equal(start) || !h.Task.EndDate.Equal(end) {
		t.Errorf("date mismatch: got %v..%v want %v..%v", h.Task.StartDate, h.Task.EndDate, start, end)
	}
	if h.Task.Frequency != freq {
		t.Errorf("frequency mismatch: got %v want %v", h.Task.Frequency, freq)
	}
	if h.Status != Active {
		t.Errorf("status mismatch: got %s want %s", h.Status, Active)
	}
	// CreatedAt is set by BaseTask; should be within the construction window
	if h.Task.CreatedAt.Before(before) || h.Task.CreatedAt.After(after) {
		t.Errorf("CreatedAt out of range [%v, %v], got %v", before, after, h.Task.CreatedAt)
	}
	// Constructor doesn't set ModifiedAt
	if !h.Task.ModifiedAt.IsZero() {
		t.Errorf("ModifiedAt should be zero value, got %v", h.Task.ModifiedAt)
	}
}

func TestNew_PropagatesInvalidDateError(t *testing.T) {
	// Requires task.NewBaseTask to return ErrEndBeforeStart when end < start (as in your RegularTask refactor)
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC) // invalid
	_, err := New(uuid.New(), "Backwards", start, end, Active, frequency.Frequency{})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != task.ErrEndBeforeStart {
		t.Fatalf("expected ErrEndBeforeStart, got %v", err)
	}
}

func TestNew_GeneratesUniqueIDs(t *testing.T) {
	userID := uuid.New()
	var freq frequency.Frequency

	a, err := New(userID, "h1", time.Now(), time.Now().Add(time.Hour), Active, freq)
	if err != nil {
		t.Fatalf("unexpected error creating a: %v", err)
	}
	b, err := New(userID, "h2", time.Now(), time.Now().Add(2*time.Hour), Active, freq)
	if err != nil {
		t.Fatalf("unexpected error creating b: %v", err)
	}
	if a.Task.ID == b.Task.ID {
		t.Errorf("expected unique IDs; got same %s", a.Task.ID)
	}
}

func TestStatus_AllValuesAssignable(t *testing.T) {
	cases := []Status{Active, Paused, Inactive, Accomplished, Abandoned}
	for _, s := range cases {
		h, err := New(uuid.New(), "status check", time.Now(), time.Now().Add(time.Hour), s, frequency.Frequency{})
		if err != nil {
			t.Fatalf("unexpected error for status %q: %v", s, err)
		}
		if h.Status != s {
			t.Fatalf("status not retained: got %s want %s", h.Status, s)
		}
	}
}

func TestNew_ZeroFrequency_Preserved(t *testing.T) {
	// Documents behavior that zero-value frequency is accepted and stored
	userID := uuid.New()
	var zero frequency.Frequency
	h, err := New(userID, "no freq", time.Now(), time.Now().Add(time.Hour), Active, zero)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.Task.Frequency != zero {
		t.Fatalf("expected zero frequency to be preserved, got %v", h.Task.Frequency)
	}
}
