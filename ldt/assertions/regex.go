package assertions

import (
	"fmt"
	"regexp"
)

type regex struct {
	regexp *regexp.Regexp
}

func NewRegex(expr string) (Assertion, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	return &regex{regexp: r}, nil
}

func (r regex) Assert(logLine string) (failMessage string, successful bool) {
	if !r.regexp.MatchString(logLine) {
		failMessage = fmt.Sprintf("Line '%s' doesn't match regex '%s'.", logLine, r.regexp)
		return failMessage, false
	}
	return "", true
}
