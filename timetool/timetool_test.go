package timetool_test

import (
	"errors"
	"fmt"
	"github.com/fedir/toolkit/timetool"
	"testing"
	"time"
)

func TestUnixTimestampToDate(t *testing.T) {

	timeNow := time.Now()
	timeNowString := fmt.Sprint(timeNow.Unix())
	timeNowStringFormatted := fmt.Sprintf("%d/%02d/%02d", timeNow.Year(), timeNow.Month(), timeNow.Day())
	expectedResult := timeNowStringFormatted

	positiveSamples := make(map[string]string)
	positiveSamples["unixtimestamp stringed"] = timeNowString

	for k, v := range positiveSamples {
		processedResult, _ := timetool.UnixTimestampToDate(v)
		if processedResult != expectedResult {
			t.Error(errors.New(t.Name() + " failed on " + k + " test"))
		}
	}

	negativeSamples := make(map[string]string)
	negativeSamples["day - 1"] = fmt.Sprintf("%d/%02d/%02d", timeNow.Year(), timeNow.Month(), timeNow.Day()-1)
	negativeSamples["day + 1"] = fmt.Sprintf("%d/%02d/%02d", timeNow.Year(), timeNow.Month(), timeNow.Day()+1)

	for k, v := range negativeSamples {
		processedResult, _ := timetool.UnixTimestampToDate(v)
		if processedResult == expectedResult {
			t.Error(errors.New(t.Name() + " failed on " + k + " test"))
		}
	}

}
