package utils

import (
	"regexp"
)

func CheckColNames(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9._-]*$`).MatchString(s)
}
