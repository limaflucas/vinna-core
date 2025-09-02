package frequency

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

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

//README
// How this should works: very similar to the linux cron
// Let's keep it simple first

// 0-59 | * (every minute)
var minutePattern = regexp.MustCompile(`^(?:\*|[0-9]|[1-5][0-9])$`)

// 0-23 | * (every hour)
var hourPattern *regexp.Regexp = regexp.MustCompile(`^(?:\*|[0-9]|1[0-9]|2[0-3])$`)

// 1-31 | * (every day)
var dayOfMonthPattern *regexp.Regexp = regexp.MustCompile(`^(?:\*|[1-9]|[12][0-9]|3[01])$`)

// 1-12 | * (every month)
var monthPattern *regexp.Regexp = regexp.MustCompile(`^(?:\*|[1-9]|1[0-2])$`)

// 0-6 (6 is Saturday) | * (every day)
var dayOfWeekPattern *regexp.Regexp = regexp.MustCompile(`^(?:\*|[0-6])$`)

func validatePattern(value, field string, pattern *regexp.Regexp) error {
	if !pattern.MatchString(value) {
		return fmt.Errorf("invalid %s format: %s", field, value)
	}
	return nil
}

// semantic validation: disallow impossible day/month combos
func validateDateCombination(day, month string) error {
	if day == "*" || month == "*" {
		return nil
	}
	dayInt, _ := strconv.Atoi(day)
	monthInt, _ := strconv.Atoi(month)

	// Use leap year (2024) to allow Feb 29
	testDate := time.Date(2024, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)
	if testDate.Day() != dayInt || int(testDate.Month()) != monthInt {
		return fmt.Errorf("invalid date combination: day %d, month %d", dayInt, monthInt)
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
	if err := validateDateCombination(dayOfMonth, month); err != nil {
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
