package model

import "time"

// Date ..
type Date time.Time

// UnmarshalText ..
//
// http://grokbase.com/t/gg/golang-nuts/144q75r4kk/go-nuts-unmarshal-xml-types-time-time-bool-and-such
// https://play.golang.org/p/JBi5Qu470_
func (t *Date) UnmarshalText(b []byte) error {
	result, err := time.Parse("2006-01-02 15:04:05 -0700", string(b))
	if err != nil {
		var t2 time.Time
		err = t2.UnmarshalText(b)
		if err != nil {
			return err
		}
		*t = Date(t2)
		return nil
	}

	// Save as data
	*t = Date(result)
	return nil
}

// Time ...
func (t Date) Time() time.Time {
	return (time.Time)(t)
}

// Before reports whether the time instant t is before u.
func (t Date) Before(u Date) bool {
	return t.Time().Before(u.Time())
}

// After reports whether the time instant t is after u.
func (t Date) After(u Date) bool {
	return t.Time().After(u.Time())
}

func (t Date) String() string {
	return t.Time().String()
	// return t.Time().Format("2006-01-02")
}

// NewDate ...
func NewDate(y, m, d int, endOfDay bool) Date {
	var t time.Time
	if endOfDay {
		t = time.Date(y, (time.Month)(m), d, 23, 59, 59, 999999999, time.Local)
	} else {
		t = time.Date(y, (time.Month)(m), d, 0, 0, 0, 0, time.Local)
	}
	return (Date)(t)
}
