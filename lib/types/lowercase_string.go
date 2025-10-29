package types

import "strings"

type LowercaseString string

func NewLowercaseString(s string) LowercaseString {
	return LowercaseString(strings.ToLower(s))
}
