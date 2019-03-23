package vartool_test

import (
	"errors"
	"github.com/fedir/toolkit/vartool"
	"testing"
)

func TestTypeInt(t *testing.T) {

	expectedValue := "int"
	positiveSamples := make(map[string]int)
	positiveSamples["a simple number"] = 1
	positiveSamples["negative number"] = -1
	positiveSamples["some big number"] = 1234567890

	for k, v := range positiveSamples {
		if vartool.Type(v) != expectedValue {
			t.Error(errors.New(t.Name() + " failed on " + k + " test, the detected type is" + vartool.Type(v)))
		}
	}

	negativeSamples := make(map[string]string)
	negativeSamples["an integer in a string"] = "1"
	negativeSamples["a string, which begins by an integer"] = "123soleil"

	for k, v := range negativeSamples {
		if vartool.Type(v) == expectedValue {
			t.Error(errors.New(t.Name() + " failed on " + k + " test, the detected type is" + vartool.Type(v)))
		}
	}
}

func TestTypeFloat(t *testing.T) {
	expectedValue := "float64"
	positiveSamples := make(map[string]float64)
	positiveSamples["a simple float"] = 1.1234
	positiveSamples["negative float"] = -1.1234
	positiveSamples["some big float"] = 1234567890.1234567890

	for k, v := range positiveSamples {
		if vartool.Type(v) != expectedValue {
			t.Error(errors.New(t.Name() + " failed on " + k + " test, the detected type is" + vartool.Type(v)))
		}
	}

	negativeSamples := make(map[string]string)
	negativeSamples["a float in a string"] = "1.234"
	negativeSamples["a string, which begins by a float"] = "123.123soleil"

	for k, v := range negativeSamples {
		if vartool.Type(v) == expectedValue {
			t.Error(errors.New(t.Name() + " failed on " + k + " test, the detected type is" + vartool.Type(v)))
		}
	}
}

func TestTypeString(t *testing.T) {
	expectedValue := "string"
	positiveSamples := make(map[string]string)
	positiveSamples["a simple string"] = "hello world!"
	positiveSamples["a string, which begins by a number"] = "123soleil"
	positiveSamples["some big string"] = "lipsum orum lipsum orum lipsum orum lipsum orum lipsum orum"

	for k, v := range positiveSamples {
		if vartool.Type(v) != expectedValue {
			t.Error(errors.New(t.Name() + " failed on " + k + " test, the detected type is" + vartool.Type(v)))
		}
	}

	negativeSamples := make(map[string]int)
	negativeSamples["an integer"] = 1
	negativeSamples["a big integer"] = 1234567890

	for k, v := range negativeSamples {
		if vartool.Type(v) == expectedValue {
			t.Error(errors.New(t.Name() + " failed on " + k + " test, the detected type is" + vartool.Type(v)))
		}
	}
}
