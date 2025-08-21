package user

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	Birth     time.Time `json:"birth"`
	CreatedAt time.Time `json:"created_at"`
	Status    Status    `json:"status"`
}

func New(firstName, lastName, email, country string, birth time.Time, status Status) (*User, error) {
	if firstName == "" {
		return nil, errors.New("first name is required")
	}
	if lastName == "" {
		return nil, errors.New("last name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if !emailPattern.MatchString(email) {
		return nil, errors.New("invalid email format")
	}
	if country == "" {
		return nil, errors.New("country is required")
	}
	if birth.After(time.Now()) {
		return nil, errors.New("birth can not be in the future")
	}

	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Country:   country,
		Birth:     birth,
		CreatedAt: time.Now(),
		Status:    status,
	}, nil
}
