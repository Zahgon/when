package rules

import (
	"regexp"
	"time"
)

type Strategy int

const (
	Skip Strategy = iota
	Merge
	Override
)

type Rule interface {
	Find(string) *Match
}

type Options struct {
	Afternoon, Evening, Morning, Noon int

	Distance int

	MatchByOrder bool

	// TODO
	// WeekStartsOn time.Weekday
}

type Match struct {
	Left, Right int
	Text        string
	Captures    []string
	Order       float64
	Applier     func(*Match, *Context, *Options, time.Time) (bool, error)
}

func (m Match) String() string { _ = "STUB: not implemented"; return "" }

func (m *Match) Apply(c *Context, o *Options, t time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type F struct {
	RegExp  *regexp.Regexp
	Applier func(*Match, *Context, *Options, time.Time) (bool, error)
}

func (f *F) Find(text string) *Match { _ = "STUB: not implemented"; return nil }

// check if capture was found
