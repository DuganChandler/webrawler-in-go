package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("unable to get response from url: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK http status code: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if len(contentType) < 1 {
		return "", fmt.Errorf("no content types of response")
	}

	isHTML := false
	for _, t := range strings.Split(contentType, ";") {
		if t == "text/html" {
			isHTML = true
		}
	}

	if !isHTML {
		return "", fmt.Errorf("content type does not match html")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %v", err)
	}
	return string(body), nil

}
