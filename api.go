package itembase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/facebookgo/httpcontrol"
	"net/http"
	"net/url"
	"time"
)

// httpClient is the HTTP client used to make calls to Itembase with the default API
var httpClient = newTimeoutClient(connectTimeout, readWriteTimeout)

// itembaseAPI is the internal implementation of the Itembase API client.
type itembaseAPI struct{}

var (
	connectTimeout   = time.Duration(30 * time.Second) // timeout for http connection
	readWriteTimeout = time.Duration(30 * time.Second) // timeout for http read/write
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

	fmt.Println(path)
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

	req.Close = true

	return client.Do(req)
}

// Call invokes the appropriate HTTP method on a given Itembase URL.
func (f *itembaseAPI) Call(method, path, auth string, body interface{}, params map[string]string, dest interface{}) error {
	response, err := doItembaseRequest(httpClient, method, path, auth, "", body, params)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return err
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if response.StatusCode >= 400 {
		err := &Error{Code: response.StatusCode, Message: response.Status}
		decoder.Decode(err)
		return err
	}

	if dest != nil && response.ContentLength != 0 {
		err = decoder.Decode(dest)
		if err != nil {
			fmt.Println("error 3")
			fmt.Println(err)
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
			MaxTries:            30,
			MaxIdleConnsPerHost: 5,
		},
	}
}
