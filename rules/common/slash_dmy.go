package common

import (
	"github.com/olebedev/when/rules"
)

/*

- DD/MM/YYYY
- 11/3/2015
- 11/3/2015
- 11/3

also with "\", gift for windows' users

https://play.golang.org/p/29LkTfe1Xr
*/

var MONTHS_DAYS = []int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func getDays(year, month int) int {
	_ = "STUB: not implemented"
	// naive leap year check
	return 0
}

func SlashDMY(s rules.Strategy) rules.Rule { _ = "STUB: not implemented"; return *new(rules.Rule) }
