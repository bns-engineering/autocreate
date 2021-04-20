package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	// "MambuFeeder/logging"
)

// SendHTTP - function to send HTTP requests.
func SendHTTP(reqMethod string, url string, timeout int, resp interface{}, header map[string]string, body io.Reader) (int, string, error) {
	//init logger for logging
	// logger := logging.SetVariable(map[string]interface{}{
	// 	"reqMethod": reqMethod,
	// 	"url":       url,
	// 	"timeout":   timeout,
	// 	"header":    header,
	// })

	t := &http.Transport{}
	client := http.Client{Transport: t, Timeout: time.Duration(timeout) * time.Millisecond}
	defer t.CloseIdleConnections()

	req, err := http.NewRequest(reqMethod, url, body)
	if err != nil {
		// logger.Error(err)
		fmt.Println(err)
		return 400, "", err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		// logger.Error(err)
		return 400, "", err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// logger.Error(err, "read body")
		fmt.Println(err)
		return res.StatusCode, string(responseBody), err
	}

	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		// logger.Errorf(err, "Response body: %v", string(responseBody))
		// fmt.Println(string(responseBody))

		return res.StatusCode, string(responseBody), err
	}
	return res.StatusCode, string(responseBody), nil
}
