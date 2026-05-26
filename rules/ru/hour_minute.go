package ru

import (
	"github.com/olebedev/when/rules"
)

/*
	{"5:30pm", 0, "5:30pm", 0},
	{"5:30 pm", 0, "5:30 pm", 0},
	{"7-10pm", 0, "7-10pm", 0},
	{"5-30", 0, "5-30", 0},
	{"05:30pm", 0, "05:30pm", 0},
	{"05:30 pm", 0, "05:30 pm", 0},
	{"05:30", 0, "05:30", 0},
	{"05-30", 0, "05-30", 0},
	{"7-10 pm", 0, "7-10 pm", 0},
	{"11.1pm", 0, "11.1pm", 0},
	{"11.10 pm", 0, "11.10 pm", 0},

	https://go.dev/play/p/QiSvUkrni6N
*/

// 1. - int
// 2. - int
// 3. - ext?

func HourMinute(s rules.Strategy) rules.Rule { _ = "STUB: not implemented"; return *new(rules.Rule) }

// am

// pm
