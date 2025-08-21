package frequency

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type Frequency struct {
	ID         uuid.UUID `json:"id"`
	Minute     string    `json:"minute"`
	Hour       string    `json:"hour"`
	DayOfMonth string    `json:"day_of_month"`
	Month      string    `json:"month"`
	DayOfWeek  string    `json:"day_of_week"`
}

var minutePattern = regexp.MustCompile(`^(\*|([0-9]|[1-5][0-9])(,([0-9]|[1-5][0-9]))*)$`)
var hourPattern *regexp.Regexp = regexp.MustCompile(`^(\*|([0-9]|1[0-9]|2[0-3])(,([0-9]|1[0-9]|2[0-3]))*)$`)
var dayOfMonthPattern *regexp.Regexp = regexp.MustCompile(`^(\*|([1-9]|[12][0-9]|3[01])(,([1-9]|[12][0-9]|3[01]))*)$`)
var monthPattern *regexp.Regexp = regexp.MustCompile(`^(\*|([1-9]|1[0-2])(,([1-9]|1[0-2]))*)$`)
var dayOfWeekPattern *regexp.Regexp = regexp.MustCompile(`^(\*|[1-7](,[1-7])*)$`)

func validatePattern(value, field string, pattern *regexp.Regexp) error {
	if !pattern.MatchString(value) {
		return fmt.Errorf("invalid %s format: %s", field, value)
	}
	return nil
}

func New(minute, hour, dayOfMonth, month, dayOfWeek string) (*Frequency, error) {
	if err := validatePattern(minute, "minute", minutePattern); err != nil {
		return nil, err
	}
	if err := validatePattern(hour, "hour", hourPattern); err != nil {
		return nil, err
	}
	if err := validatePattern(dayOfMonth, "day_of_month", dayOfMonthPattern); err != nil {
		return nil, err
	}
	if err := validatePattern(month, "month", monthPattern); err != nil {
		return nil, err
	}
	if err := validatePattern(dayOfWeek, "day_of_week", dayOfWeekPattern); err != nil {
		return nil, err
	}

	return &Frequency{
		ID:         uuid.New(),
		Minute:     minute,
		Hour:       hour,
		DayOfMonth: dayOfMonth,
		Month:      month,
		DayOfWeek:  dayOfWeek,
	}, nil
}
