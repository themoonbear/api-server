package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func getClient() *http.Client {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	return &client
}

func setHeader(req *http.Request, header *map[string]string) {
	if header != nil {
		for k, v := range *header {
			req.Header.Set(k, v)
		}
	}
}

func sendRequest(url, method string, msg io.Reader, header *map[string]string) (string, error) {
	client := getClient()
	req, err := http.NewRequest(method, url, msg)
	if err != nil {
		return "", err
	}
	setHeader(req, header)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//HTTPGet Get请求
func HTTPGet(url string, header *map[string]string) (string, error) {
	return sendRequest(url, "GET", nil, header)
}
