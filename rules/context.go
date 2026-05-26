package rules

import "time"

type Context struct {
	Text string

	// accumulator of relative values
	Duration time.Duration

	// Aboslute values
	Year, Month, Weekday, Day, Hour, Minute, Second *int

	Location *time.Location
}

func (c *Context) Time(t time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
