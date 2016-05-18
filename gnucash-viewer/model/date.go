package model

import "time"

// GncDate ..
type GncDate time.Time

// UnmarshalText ..
//
// http://grokbase.com/t/gg/golang-nuts/144q75r4kk/go-nuts-unmarshal-xml-types-time-time-bool-and-such
// https://play.golang.org/p/JBi5Qu470_
func (t *GncDate) UnmarshalText(b []byte) error {
	result, err := time.Parse("2006-01-02 15:04:05 -0700", string(b))
	if err != nil {
		var t2 time.Time
		err = t2.UnmarshalText(b)
		if err != nil {
			return err
		}
		*t = GncDate(t2)
		return nil
	}

	// Save as data
	*t = GncDate(result)
	return nil
}

// Time ...
func (t GncDate) Time() time.Time {
	return (time.Time)(t)
}

func (t GncDate) String() string {
	return t.Time().String()
}
