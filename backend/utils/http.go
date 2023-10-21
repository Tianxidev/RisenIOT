package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HttpGet 发送get请求
func HttpGet(url string, headers map[string]string) (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// HttpPost 发送post请求
func HttpPost(url string, headers map[string]string, data string) (string, error) {
	body := strings.NewReader(data)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, body)
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		bytes0, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("Error:%v", string(bytes0))
	}
	var r []byte
	response.Body.Read(r)
	bytes1, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bytes1), nil
}
