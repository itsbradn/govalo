package http

import (
	"fmt"
	"strings"
)

func ParseCookies(cookies []string, subs string) (string, error) {
	for _, cookie := range cookies {
		if strings.Contains(cookie, subs) {
			return cookie, nil
		}
	}
	return "", fmt.Errorf("could not find %s", subs)
}
