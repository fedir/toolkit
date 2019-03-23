package iptool_test

import (
	"errors"
	"github.com/fedir/toolkit/iptool"
	"testing"
)

func TestValidateIP(t *testing.T) {

	var authorizedIP = "1.2.3.4"

	positiveSamples := make(map[string]string)
	positiveSamples["single IP"] = "1.2.3.4"
	positiveSamples["multiple IPs separated by ','"] = "1.1.1.1,1.2.3.4,255.255.255.255"
	positiveSamples["multiple IPs separated by ', '"] = "1.1.1.1, 1.2.3.4, 255.255.255.255"

	for k, v := range positiveSamples {
		if !iptool.ValidateIP(v, authorizedIP, false) {
			t.Error(errors.New(t.Name() + " failed on " + k + " test"))
		}
	}

	negativeSamples := make(map[string]string)
	positiveSamples["bad single IP"] = "1.2.3.5"
	positiveSamples["multiple bad IPs separated by ','"] = "1.1.1.1,1.2.3.5,255.255.255.255"
	positiveSamples["multiple bad IPs separated by ', '"] = "1.1.1.1, 1.2.3.5, 255.255.255.255"

	for k, v := range negativeSamples {
		if iptool.ValidateIP(v, authorizedIP, false) {
			t.Error(errors.New(t.Name() + " failed on " + k + " test"))
		}
	}
}
