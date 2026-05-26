package nl

import (
	"github.com/olebedev/when/rules"
)

/*
	{"17:30", 0, "17:30", 0},
	{"17:30u", 0, "17:30u", 0},
	{"om 17:30 uur", 3, "17:30 uur", 0},
	{"om 5:59 pm", 3, "5:59 pm", 0},

	https://play.golang.org/p/hXl7C8MWNr
*/

// 1. - at?
// 2. - int
// 3. - int
// 4. - ext?
// 5. - day part?

func HourMinute(s rules.Strategy) rules.Rule { _ = "STUB: not implemented"; return *new(rules.Rule) }

// pm

// afternoon or evening

// Truncate seconds
