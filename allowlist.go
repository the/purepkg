package main

import (
	"strings"
)

type allowList []string

func (l allowList) contains(element string) bool {
	for _, allowed := range l {
		if l.match(allowed, element) {
			return true
		}
	}
	return false
}

func (l allowList) match(allowed, element string) bool {
	wildcard := strings.HasSuffix(allowed, "*")
	if wildcard {
		allowedPrefix := allowed[:len(allowed)-1] // chop off wildcard character
		return strings.HasPrefix(element, allowedPrefix)
	}
	return allowed == element // exact match
}

// implement flag.Value

func (l *allowList) String() string {
	return strings.Join(*l, ",")
}

func (l *allowList) Set(value string) error {
	*l = strings.Split(value, ",")
	return nil
}
