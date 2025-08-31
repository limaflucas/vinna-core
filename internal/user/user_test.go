package user

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNew_Success(t *testing.T) {
	// capture time bounds to validate CreatedAt
	start := time.Now()
	u, err := New("Ada", "Lovelace", "ada@example.com", "UK", time.Date(1815, time.December, 10, 0, 0, 0, 0, time.UTC), Active)
	end := time.Now()

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}
	if u == nil {
		t.Fatal("New() returned nil user without error")
	}
	if u.ID == uuid.Nil {
		t.Errorf("expected non-zero UUID")
	}
	if u.FirstName != "Ada" || u.LastName != "Lovelace" {
		t.Errorf("name not set correctly: %+v", u)
	}
	if u.Email != "ada@example.com" {
		t.Errorf("email not set correctly: %s", u.Email)
	}
	if u.Country != "UK" {
		t.Errorf("country not set correctly: %s", u.Country)
	}
	if u.CreatedAt.Before(start) || u.CreatedAt.After(end) {
		t.Errorf("CreatedAt out of expected range [%v, %v], got %v", start, end, u.CreatedAt)
	}
}

func TestNew_ValidationErrors(t *testing.T) {
	type args struct {
		first, last, email, country string
		birth                       time.Time
	}
	now := time.Now()
	tests := []struct {
		name      string
		args      args
		wantErr   string
	}{
		{
			name:    "empty first name",
			args:    args{"", "Lovelace", "ada@example.com", "UK", now.Add(-24 * time.Hour)},
			wantErr: "first name is required",
		},
		{
			name:    "empty last name",
			args:    args{"Ada", "", "ada@example.com", "UK", now.Add(-24 * time.Hour)},
			wantErr: "last name is required",
		},
		{
			name:    "empty email",
			args:    args{"Ada", "Lovelace", "", "UK", now.Add(-24 * time.Hour)},
			wantErr: "email is required",
		},
		{
			name:    "invalid email format",
			args:    args{"Ada", "Lovelace", "not-an-email", "UK", now.Add(-24 * time.Hour)},
			wantErr: "invalid email format",
		},
		{
			name:    "empty country",
			args:    args{"Ada", "Lovelace", "ada@example.com", "", now.Add(-24 * time.Hour)},
			wantErr: "country is required",
		},
		{
			name:    "birth in the future",
			args:    args{"Ada", "Lovelace", "ada@example.com", "UK", now.Add(1 * time.Hour)},
			wantErr: "birth can not be in the future",
		},
	}

	var s Status // zero value is fine for tests
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			u, err := New(tt.args.first, tt.args.last, tt.args.email, tt.args.country, tt.args.birth, s)
			if err == nil {
				t.Fatalf("expected error %q, got nil (user=%+v)", tt.wantErr, u)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error %q does not contain expected %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestNew_BirthEqualNow_IsAllowed(t *testing.T) {
	// Edge case: birth == time.Now() should be allowed (code checks After, not After or Equal)
	now := time.Now()
	u, err := New("Ada", "Lovelace", "ada@example.com", "UK", now, Active)
	if err != nil {
		t.Fatalf("expected no error when birth == now, got %v", err)
	}
	if u == nil {
		t.Fatal("user is nil")
	}
}

func TestNew_EmailRegex_ValidCases(t *testing.T) {
	valids := []string{
		"a.b+c_d%test@example.co",
		"user@sub.example.com",
		"USER123@example.io",
		"first.last@domain.travel",
	}
	for _, e := range valids {
		e := e
		t.Run(e, func(t *testing.T) {
			u, err := New("A", "B", e, "CA", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Active)
			if err != nil {
				t.Fatalf("email %q should be valid, got error: %v", e, err)
			}
			if u.Email != e {
				t.Fatalf("email not set correctly, got %q", u.Email)
			}
		})
	}
}

func TestNew_EmailRegex_InvalidCases(t *testing.T) {
	t.Parallel()
	now := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	invalids := []string{
		"plainaddress",                 // no @
		"@no-local-part.com",           // missing local part
		"user@@example.com",            // double @
		"user@example",                 // no TLD
		"user@example.c",               // TLD too short
		"user@.example.com",            // dot right after @
		"user@example..com",            // double dot in domain
		".user@example.com",            // leading dot in local
		"user.@example.com",            // trailing dot in local
		"user..name@example.com",       // consecutive dots in local
		"user@example.com.",            // trailing dot
		"user@-example.com",            // domain label starts with hyphen
		"user@example-.com",            // domain label ends with hyphen
		"user@ex_ample.com",            // underscore in domain
		"user@exam!ple.com",            // punctuation in domain
		"user name@example.com",        // space
		"user@[127.0.0.1]",             // IP literal not supported by this regex
		`user"quote"@example.com`,      // quotes not allowed by this regex
		"user@例子.测试",                   // unicode domain not allowed by this regex
		"user@123",                     // numeric “TLD”
	}

	for _, e := range invalids {
		e := e
		t.Run(e, func(t *testing.T) {
			t.Parallel()
			u, err := New("A", "B", e, "CA", now, Active)
			if err == nil {
				t.Fatalf("expected error for invalid email %q, got nil (user=%+v)", e, u)
			}
			if !strings.Contains(err.Error(), "invalid email format") {
				t.Fatalf("got error %q; want it to contain %q", err.Error(), "invalid email format")
			}
		})
	}
}
