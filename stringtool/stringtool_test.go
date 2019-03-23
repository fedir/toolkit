package stringtool_test

import (
	"errors"
	"github.com/fedir/toolkit/stringtool"
	"testing"
)

func TestTrimSpace(t *testing.T) {

	var expectedResult = "test"

	positiveSamples := make(map[string]string)
	positiveSamples["new line"] = "\ntest\n"
	positiveSamples["spaces"] = " test "
	positiveSamples["tabs"] = "\ttest\t"

	for k, v := range positiveSamples {
		if stringtool.TrimSpace(v) != expectedResult {
			t.Error(errors.New(t.Name() + " failed on " + k + " test"))
		}
	}

	negativeSamples := make(map[string]string)
	negativeSamples["space in the middle"] = "te st"
	negativeSamples["tab in the middle"] = "te\tst"
	negativeSamples["new line in the middle"] = "te\nst"

	for k, v := range negativeSamples {
		if stringtool.TrimSpace(v) == expectedResult {
			t.Error(errors.New(t.Name() + " failed on " + k + " test"))
		}
	}
}
