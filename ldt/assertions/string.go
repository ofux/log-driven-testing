package assertions

import (
	"fmt"
	"strings"
)

type containedString string

func NewContainedString(str string) Assertion {
	return containedString(str)
}

func (c containedString) Assert(logLine string) (failMessage string, successful bool) {
	if !strings.Contains(logLine, string(c)) {
		failMessage = fmt.Sprintf("Expected '%s' to contain '%s'.", logLine, c)
		return failMessage, false
	}
	return "", true
}
