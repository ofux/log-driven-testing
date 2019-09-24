package assertions

type Assertion interface {
	Assert(logLine string) (failMessage string, successful bool)
}
