package types

type ApplyFixesResult int

const (
	// No more fixes are needed, proceed to the next tool
	Done ApplyFixesResult = iota

	// Only partial fixes have been applied, but we need to re-gather data
	// to continue.
	ProcessAgain
)
