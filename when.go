package when

import (
	"time"

	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/br"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/olebedev/when/rules/nl"
	"github.com/olebedev/when/rules/ru"
)

// Parser is a struct which contains options
// rules, and middlewares to call
type Parser struct {
	options    *rules.Options
	rules      []rules.Rule
	middleware []func(string) (string, error)
}

// Result is a struct which contains parsing meta-info
type Result struct {
	// Index is a start index
	Index int
	// Text is a text found and processed
	Text string
	// Source is input string
	Source string
	// Time is an output time
	Time time.Time
}

// Parse returns Result and error if any. If have not matches it returns nil, nil.
func (p *Parser) Parse(text string, base time.Time) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// apply middlewares

// find all matches

// not found

// find a cluster

// get borders of the matches

// apply rules

// Add adds  given rules to the main chain.
func (p *Parser) Add(r ...rules.Rule) { _ = "STUB: not implemented"; return }

// Use adds give functions to middlewares.
func (p *Parser) Use(f ...func(string) (string, error)) { _ = "STUB: not implemented"; return }

// SetOptions sets options object to use.
func (p *Parser) SetOptions(o *rules.Options) {
	_ = "STUB: not implemented"

	// New returns Parser initialised with given options.
	return
}

func New(o *rules.Options) *Parser { _ = "STUB: not implemented"; return nil }

// default options for internal usage
var defaultOptions = &rules.Options{
	Distance:     5,
	MatchByOrder: true,
}

// EN is a parser for English language
var EN *Parser

// RU is a parser for Russian language
var RU *Parser

// BR is a parser for Brazilian Portuguese language
var BR *Parser

// NL is a parser for Dutch language
var NL *Parser

func init() {
	EN = New(nil)
	EN.Add(en.All...)
	EN.Add(common.All...)

	RU = New(nil)
	RU.Add(ru.All...)
	RU.Add(common.All...)

	BR = New(nil)
	BR.Add(br.All...)
	BR.Add(common.All...)

	NL = New(nil)
	NL.Add(nl.All...)
	NL.Add(common.All...)
}
