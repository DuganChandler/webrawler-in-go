package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(baseURL string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse url: %w", err)
	}

	newURL := u.Host + u.Path

	newURL = strings.ToLower(newURL)

	newURL = strings.TrimSuffix(newURL, "/")

	return newURL, nil
}
