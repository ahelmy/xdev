package internal

import "net/url"

func EncodeURL(s string) string {
	return url.QueryEscape(s)
}

func DecodeURL(s string) (string, error) {
	return url.QueryUnescape(s)
}
