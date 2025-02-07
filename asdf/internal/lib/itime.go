package lib

import (
	"encoding/json"
	"strings"
	"time"
)

// JavaScript Date format without timezone
// const jsTimeLayout string = "Mon Jan 2 2006 15:04:05"
const jsTimeLayout string = `Mon Jan 2 2006 15:04:05 GMT-0700`

type ITime time.Time

// Implement Marshaler and Unmarshaler interface
func (it *ITime) UnmarshalJSON(b []byte) error {
	// Remove quotes from string
	s := strings.Trim(string(b), "\"")

	// Remove timezone name in parentheses if present
	if idx := strings.Index(s, "("); idx != -1 {
		s = strings.TrimSpace(s[:idx])
	}

	// Parse the time string
	t, err := time.Parse(jsTimeLayout, s)
	if err != nil {
		return err
	}

	*it = ITime(t)
	return nil
}

func (it ITime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(it))
}

func (it ITime) IsZero() bool {
	t := it.Time()
	return t.Second() == 0 && t.Nanosecond() == 0
}

func (it ITime) Time() time.Time {
	return time.Time(it)
}

func (it ITime) String(s string) string {
	t := it.Time()
	return t.Format(s)
}
