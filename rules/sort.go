package rules

type MatchByIndex []*Match

func (m MatchByIndex) Len() int { _ = "STUB: not implemented"; return 0 }

func (m MatchByIndex) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (m MatchByIndex) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type MatchByOrder []*Match

func (m MatchByOrder) Len() int { _ = "STUB: not implemented"; return 0 }

func (m MatchByOrder) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (m MatchByOrder) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
