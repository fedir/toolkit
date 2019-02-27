package httptool

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func bytesToString(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func TestOKStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	testURL := ts.URL
	t.Logf("%s\n", testURL)
	_, _, err := Request(testURL)
	if err != nil {
		t.Errorf("httpRequest() returned an error: %s", err)
	}
}
