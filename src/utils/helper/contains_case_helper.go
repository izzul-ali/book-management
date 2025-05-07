package helper

import "strings"

// checks whether substring exists in the string
func ContainsIgnoreCase(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}
