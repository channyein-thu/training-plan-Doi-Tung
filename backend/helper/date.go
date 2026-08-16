package helper

import "time"

// ParseDateOnly parses a "YYYY-MM-DD" date string (as sent by the frontend date
// pickers) into a *time.Time. An empty string returns (nil, nil) so optional
// dates are handled naturally; an invalid format returns an error.
func ParseDateOnly(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
