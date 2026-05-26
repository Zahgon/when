package zh

import (
	"github.com/olebedev/when/rules"
)

/*
	规则名称：精确到月份的日期
*/

func ExactMonthDate(s rules.Strategy) rules.Rule {
	_ = "STUB: not implemented"
	return *new(rules.Rule)
}

// can't use \W here due to Chinese characters

// the default value of month is the current month, and the default
// value of day is the first day of the month, so that we can handle
// cases like "4月" (Apr 1st) and "12号" (12th this month)
