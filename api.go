package itembase

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/facebookgo/httpcontrol"
)

// httpClient is the HTTP client used to make calls to Itembase with the default API
var httpClient = newTimeoutClient(connectTimeout, readWriteTimeout)

// itembaseAPI is the internal implementation of the Itembase API client.
type itembaseAPI struct{}

var (
	connectTimeout   = time.Duration(120 * time.Second) // timeout for http connection
	readWriteTimeout = time.Duration(120 * time.Second) // timeout for http read/write
)

func doItembaseRequest(client *http.Client, method, path, auth, accept string, body interface{}, params map[string]string) (*http.Response, error) {

	qs := url.Values{}

	for k, v := range params {
		qs.Set(k, v)
	}

	if len(qs) > 0 {
		path += "?" + qs.Encode()
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	log.Println(path)
	req, err := http.NewRequest(method, path, bytes.NewReader(encodedBody))
	if err != nil {
		return nil, err
	}

	// if the client has an auth, set it as a header
	if len(auth) > 0 {
		req.Header.Add("Authorization", "Bearer "+auth)
	}

	if accept != "" {
		req.Header.Add("Accept", accept)
	}

	req.Header.Add("Accept-Encoding", "gzip")

	req.Close = true

	return client.Do(req)
}

// Call invokes the appropriate HTTP method on a given Itembase URL.
func (f *itembaseAPI) Call(method, path, auth string, body interface{}, params map[string]string, dest interface{}) error {
	response, err := doItembaseRequest(httpClient, method, path, auth, "", body, params)
	if err != nil {
		log.Println("Error when making Itembase Request", err)
		return err
	}

	defer response.Body.Close()

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			log.Println("Error when decoding gzipped response body", err)
			return err
		}
		defer reader.Close()
	default:
		reader = response.Body
	}

	decoder := json.NewDecoder(reader)
	if response.StatusCode >= 400 {
		err := &Error{Code: response.StatusCode, Message: response.Status}
		decoder.Decode(err)
		return err
	}

	if dest != nil && response.ContentLength != 0 {
		err = decoder.Decode(dest)
		if err != nil {
			log.Println("Error when decoding body", err)
			return err
		}
	}

	return nil
}

func newTimeoutClient(connectTimeout time.Duration, readWriteTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &httpcontrol.Transport{
			RequestTimeout:      readWriteTimeout,
			DialTimeout:         connectTimeout,
			MaxTries:            60,
			MaxIdleConnsPerHost: 5,
			RetryAfterTimeout:   true,
		},
	}
}
