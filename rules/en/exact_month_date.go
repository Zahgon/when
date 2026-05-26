package en

import (
	"github.com/olebedev/when/rules"
)

// <[]string{"third of march", "third", "", "march", "", ""}>
// <[]string{"march third", "", "", "march", "third", ""}>
// <[]string{"march 3rd", "", "", "march", "3rd", ""}>
// <[]string{"3rd march", "3rd", "", "march", "", ""}>
// <[]string{"march 3", "", "", "march", "", "3"}>
// <[]string{"1st of september", "1st", "", "september", "", ""}>
// <[]string{"sept. 1st", "", "", "sept.", "1st", ""}>
// <[]string{"march 7th", "", "", "march", "7th", ""}>
// <[]string{"october 21st", "", "", "october", "21st", ""}>
// <[]string{"twentieth of december", "twentieth", "", "december", "", ""}>
// <[]string{"march 10th", "", "", "march", "10th", ""}>
// <[]string{"jan. 6", "", "", "jan.", "", "6"}>
// <[]string{"february", "", "", "february", "", ""}>
// <[]string{"october", "", "", "october", "", ""}>
// <[]string{"jul.", "", "", "jul.", "", ""}>
// <[]string{"june", "", "", "june", "", ""}>

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
