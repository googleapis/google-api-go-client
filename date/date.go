// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"fmt"
	"time"
)

// A Date represents a date (year, month, day).
//
// This type does not include location information, and therefore does not
// describe a unique 24-hour timespan.
type Date struct {
	Year  int        // Year (e.g., 2014).
	Month time.Month // Month of the year (January = 1, ...).
	Day   int        // Day of the month, starting at 1.
}

// Of returns the Date in which a time occurs in that time's location.
func Of(t time.Time) Date {
	var d Date
	d.Year, d.Month, d.Day = t.Date()
	return d
}

// Parse parses a string in RFC3339 full-date format and returns the date value it represents.
func Parse(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, err
	}
	return Of(t), nil
}

// String returns the date in RFC3339 full-date format.
func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// Midnight returns the time corresponding to the midnight beginning
// the date in the given location.
func (d Date) Midnight(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

// AddDays returns the date that is `days` days in the future.
// `days` can also be negative to go into the past.
func (d Date) AddDays(days int) Date {
	return Of(d.Midnight(time.UTC).AddDate(0, 0, days))
}

// Sub returns the signed number of days between the date and `s`, not including the end day.
// This is the inverse operation to `AddDays`.
func (d Date) Sub(s Date) (days int) {
	// We convert to Unix time so we do not have to worry about leap seconds:
	// Unix time increases by exactly 86400 seconds per day.
	deltaUnix := d.Midnight(time.UTC).Unix() - s.Midnight(time.UTC).Unix()
	return int(deltaUnix / 86400)
}

// Before reports whether d1 occurs before d2.
func (d1 Date) Before(d2 Date) bool {
	if d1.Year != d2.Year {
		return d1.Year < d2.Year
	}
	if d1.Month != d2.Month {
		return d1.Month < d2.Month
	}
	return d1.Day < d2.Day
}
