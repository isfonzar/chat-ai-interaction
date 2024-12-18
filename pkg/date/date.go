package date

import (
	"errors"
	"fmt"
	"regexp"
)

func Filter(date string) (*regexp.Regexp, error) {
	switch {
	case regexp.MustCompile(`^\d{4}$`).MatchString(date):
		// Matches YYYY format
		pattern := fmt.Sprintf(`^\[\d{2}/\d{2}/%s, \d{2}:\d{2}:\d{2}\] .*`, date[2:])
		return regexp.Compile(pattern)

	case regexp.MustCompile(`^\d{4}-\d{2}$`).MatchString(date):
		// Matches YYYY-MM format (e.g., 2024-06)
		year, month := date[2:4], date[5:]
		pattern := fmt.Sprintf(`^\[\d{2}/%s/%s, \d{2}:\d{2}:\d{2}\] .*`, month, year)
		return regexp.Compile(pattern)

	case regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString(date):
		// Matches YYYY-MM-DD format (e.g., 2024-06-17)
		day, month, year := date[8:], date[5:7], date[2:4]
		pattern := fmt.Sprintf(`^\[%s/%s/%s, \d{2}:\d{2}:\d{2}\] .*`, day, month, year)
		return regexp.Compile(pattern)
	default:
		return nil, errors.New("invalid date format")
	}
}
