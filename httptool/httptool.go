package httptool

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func Request(url string) ([]byte, int, error) {

	var httpClient = &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Cannont prepare the HTTP request", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Cannot process the HTTP request", err)
	}

	defer resp.Body.Close()
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal("Cannont dump the body of HTTP response", err)
	}
	statusProcessing(resp.StatusCode, url)
	return body, resp.StatusCode, err
}

func statusProcessing(statusCode int, url string) {
	switch statusCode {
	case http.StatusForbidden:
		log.Fatalf("Looks like the rate limit is exceeded, please try again later")
	case http.StatusAccepted:
		log.Printf("Looks like the server need some time to prepare request.")
	case http.StatusOK:
		return
	default:
		log.Fatalf("The status code of URL %s is not OK : %d", url, statusCode)
	}
}

// ReadResp :  reads response from HTTP query.
func ReadResp(fullResp []byte) ([]byte, error) {
	r := bufio.NewReader(bytes.NewReader(fullResp))
	resp, err := http.ReadResponse(r, nil)
	if err != nil {
		log.Printf("Error reading response: %v\n%s", err, fullResp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n%s", err, resp.Body)
	}
	return body, err
}
