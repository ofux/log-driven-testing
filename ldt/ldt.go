package ldt

import (
	"bufio"
	"bytes"
	"github.com/ofux/log-driven-testing/assertions"
	"io"
	"os"
	"strings"
	"testing"
)

type TestContext struct {
	t               *testing.T
	realStdout      *os.File
	fakeStdoutWrite *os.File
	output          chan string
	assertions      []assertions.Assertion
}

func BeginLogTest(t *testing.T) *TestContext {
	old := os.Stdout // keep backup of the real Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to init LogTest with error %s", err)
		return nil
	}
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		if err != nil {
			os.Stdout = old
			t.Fatalf("Failed to close LogTest with error %s", err)
			return
		}
		outC <- buf.String()
	}()

	return &TestContext{
		t:               t,
		realStdout:      old,
		fakeStdoutWrite: w,
		output:          outC,
	}
}

func (tc *TestContext) ARegex(t *testing.T, expr string) {
	assertion, err := assertions.NewRegex(expr)
	if err != nil {
		tc.restore()
		t.Fatalf("Invalid regexp. Error: %s", err)
	}
	tc.A(assertion)
}

func (tc *TestContext) AContains(str string) {
	assertion := assertions.NewContainedString(str)
	tc.A(assertion)
}

func (tc *TestContext) A(a assertions.Assertion) {
	tc.assertions = append(tc.assertions, a)
}

func (tc *TestContext) EndLogTest() {
	// back to normal state
	err := tc.fakeStdoutWrite.Close()
	if err != nil {
		tc.restore()
		tc.t.Fatalf("Failed to close LogTest with error %s", err)
		return
	}

	tc.restore()
	out := <-tc.output

	scanner := bufio.NewScanner(strings.NewReader(out))
	for i := 0; i < len(tc.assertions) && scanner.Scan(); i++ {
		failMessage, successful := tc.assertions[i].Assert(scanner.Text())
		if !successful {
			tc.t.Error(failMessage)
		}
	}
}

func (tc *TestContext) restore() {
	os.Stdout = tc.realStdout
}
