package nl

import (
	"github.com/olebedev/when/rules"
)

// <[]string{"derde van maart", "derde", "", "maart", "", ""}>
// <[]string{"3e van march", "3e", "", "maart", "", ""}>
// <[]string{"1e van september", "1e", "", "september", "", ""}>
// <[]string{"1 sept.", "", "", "1", "sept", ""}>
// <[]string{"twintigste van december", "twintigste", "", "december", "", ""}>
// <[]string{"februari", "", "", "februari", "", ""}>
// <[]string{"oktober", "", "", "oktober", "", ""}>
// <[]string{"jul.", "", "", "jul.", "", ""}>
// <[]string{"juni", "", "", "juni", "", ""}>

// https://play.golang.org/p/Zfjl6ERNkq

// 1. - ordinal day?
// 2. - numeric day?
// 3. - month
// 4. - ordinal day?
// 5. - ordinal day?

func ExactMonthDate(s rules.Strategy) rules.Rule {
	_ = "STUB: not implemented"
	return *new(rules.Rule)
}

// skip '(?:'
