package strings

import (
	"regexp"
)

func IsBlank(s string) bool {
	r := regexp.MustCompile(`[\s　]`)
	s = r.ReplaceAllString(s, "")
	return len(s) == 0
}
